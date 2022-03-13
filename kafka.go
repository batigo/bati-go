package bati

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

type kfkbroker struct {
	postman *Postman
	brokers []string
	group   string
	readers int
	writers int
	mutext  sync.RWMutex
	stoped  bool
}

func (m *kfkbroker) run() {
	for i := 0; i < m.readers; i++ {
		go m.reader()
	}

	for i := 0; i < m.writers; i++ {
		go m.writer()
	}
}

func (m *kfkbroker) stop() {
	m.mutext.Lock()
	m.stoped = true
	m.mutext.Unlock()
}

func (*kfkbroker) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*kfkbroker) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (m *kfkbroker) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		sess.MarkMessage(msg, "")

		log.Printf("[%s] kfk broker recv msg: %s", m.postman.Service, msg.Value)

		var batiMsg BatiMsg
		err := json.Unmarshal(msg.Value, &batiMsg)
		if err != nil {
			log.Printf("[%s] failed to parse mqmsg: %s - %s", m.postman.Service, msg.Value, err.Error())
			continue
		}

		err = m.postman.msghandler(batiMsg, m.postman.Service)
		if err != nil {
			log.Printf("[%s] postman failed to proc mqmsg: %s - %s", m.postman.Service, msg.Value, err.Error())
			continue
		}
	}
	return nil
}

func (m *kfkbroker) reader() {
	topic := fmt.Sprintf("bati_%s_up", m.postman.Service)
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V1_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	r, err := sarama.NewConsumerGroup(m.brokers, m.group, config)
	if err != nil {
		panic(fmt.Sprintf("[%s] failed to start kfk reader: %s", m.postman.Service, err.Error()))
	}

	log.Printf("[%s] kfk broker reader run", m.postman.Service)
	defer func() {
		log.Printf("[%s] kfk broker reader stop", m.postman.Service)
		r.Close()
	}()

	// read consumer errors
	go func() {
		tiker := time.NewTicker(time.Second)
		defer tiker.Stop()
		for {
			m.mutext.RLock()
			if m.stoped {
				m.mutext.RUnlock()
				break
			}
			m.mutext.RUnlock()

			select {
			case <-tiker.C:
				//
			case err := <-r.Errors():
				if err != nil {
					log.Printf("[%s] failed to consume kfk msg: %s", m.postman.Service, err.Error())
				}
			}
		}
	}()

	topics := []string{topic}
	for {
		m.mutext.RLock()
		if m.stoped {
			m.mutext.RUnlock()
			break
		}
		m.mutext.RUnlock()

		err := r.Consume(context.Background(), topics, m)
		if err != nil {
			log.Printf("[%s] failed to consume kafka msg: %s", m.postman.Service, err.Error())
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func (m *kfkbroker) writer() {
	topic := fmt.Sprintf("bati_%s_down", m.postman.Service)
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	w, _ := sarama.NewAsyncProducer(m.brokers, config)

	log.Printf("[%s] kfk broker writer run", m.postman.Service)
	defer func() {
		log.Printf("[%s] kfk broker writer stop", m.postman.Service)
		w.Close()
	}()

	// read producer errors
	go func() {
		tiker := time.NewTicker(time.Second)
		defer tiker.Stop()
		for {
			m.mutext.RLock()
			if m.stoped {
				m.mutext.RUnlock()
				break
			}
			m.mutext.RUnlock()

			select {
			case <-tiker.C:
				//
			case err := <-w.Errors():
				if err != nil {
					log.Printf("[%s] failed to produce kfk msg: %s", m.postman.Service, err.Error())
				}
			}
		}
	}()

	for {
		var ok bool
		var msg ServiceMsg
		select {
		case msg, ok = <-m.postman.msgchan:
		}
		if !ok {
			break
		}

		msg.Ts = getNowMillisecs()
		bs, _ := json.Marshal(msg)
		kmsg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(bs),
		}
		w.Input() <- kmsg
		log.Printf("[%s] kfk broker send msg: %s", m.postman.Service, msg.Id)
	}
}
