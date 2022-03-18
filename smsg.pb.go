// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: smsg.proto

package bati

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatiMsg_BatiMsgType int32

const (
	BatiMsg_Unused   BatiMsg_BatiMsgType = 0
	BatiMsg_Biz      BatiMsg_BatiMsgType = 1
	BatiMsg_ConnQuit BatiMsg_BatiMsgType = 2
)

// Enum value maps for BatiMsg_BatiMsgType.
var (
	BatiMsg_BatiMsgType_name = map[int32]string{
		0: "Unused",
		1: "Biz",
		2: "ConnQuit",
	}
	BatiMsg_BatiMsgType_value = map[string]int32{
		"Unused":   0,
		"Biz":      1,
		"ConnQuit": 2,
	}
)

func (x BatiMsg_BatiMsgType) Enum() *BatiMsg_BatiMsgType {
	p := new(BatiMsg_BatiMsgType)
	*p = x
	return p
}

func (x BatiMsg_BatiMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BatiMsg_BatiMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_smsg_proto_enumTypes[0].Descriptor()
}

func (BatiMsg_BatiMsgType) Type() protoreflect.EnumType {
	return &file_smsg_proto_enumTypes[0]
}

func (x BatiMsg_BatiMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatiMsg_BatiMsgType.Descriptor instead.
func (BatiMsg_BatiMsgType) EnumDescriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{0, 0}
}

type ServiceMsg_ServiceMsgType int32

const (
	ServiceMsg_Unused   ServiceMsg_ServiceMsgType = 0
	ServiceMsg_ConnJoin ServiceMsg_ServiceMsgType = 1
	ServiceMsg_ConnQuit ServiceMsg_ServiceMsgType = 2
	ServiceMsg_Biz      ServiceMsg_ServiceMsgType = 3
)

// Enum value maps for ServiceMsg_ServiceMsgType.
var (
	ServiceMsg_ServiceMsgType_name = map[int32]string{
		0: "Unused",
		1: "ConnJoin",
		2: "ConnQuit",
		3: "Biz",
	}
	ServiceMsg_ServiceMsgType_value = map[string]int32{
		"Unused":   0,
		"ConnJoin": 1,
		"ConnQuit": 2,
		"Biz":      3,
	}
)

func (x ServiceMsg_ServiceMsgType) Enum() *ServiceMsg_ServiceMsgType {
	p := new(ServiceMsg_ServiceMsgType)
	*p = x
	return p
}

func (x ServiceMsg_ServiceMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceMsg_ServiceMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_smsg_proto_enumTypes[1].Descriptor()
}

func (ServiceMsg_ServiceMsgType) Type() protoreflect.EnumType {
	return &file_smsg_proto_enumTypes[1]
}

func (x ServiceMsg_ServiceMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceMsg_ServiceMsgType.Descriptor instead.
func (ServiceMsg_ServiceMsgType) EnumDescriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{1, 0}
}

type BizData_BizMsgType int32

const (
	BizData_Unused  BizData_BizMsgType = 0
	BizData_Users   BizData_BizMsgType = 1
	BizData_Room    BizData_BizMsgType = 2
	BizData_Service BizData_BizMsgType = 3
	BizData_All     BizData_BizMsgType = 4
)

// Enum value maps for BizData_BizMsgType.
var (
	BizData_BizMsgType_name = map[int32]string{
		0: "Unused",
		1: "Users",
		2: "Room",
		3: "Service",
		4: "All",
	}
	BizData_BizMsgType_value = map[string]int32{
		"Unused":  0,
		"Users":   1,
		"Room":    2,
		"Service": 3,
		"All":     4,
	}
)

func (x BizData_BizMsgType) Enum() *BizData_BizMsgType {
	p := new(BizData_BizMsgType)
	*p = x
	return p
}

func (x BizData_BizMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BizData_BizMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_smsg_proto_enumTypes[2].Descriptor()
}

func (BizData_BizMsgType) Type() protoreflect.EnumType {
	return &file_smsg_proto_enumTypes[2]
}

func (x BizData_BizMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BizData_BizMsgType.Descriptor instead.
func (BizData_BizMsgType) EnumDescriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{4, 0}
}

type BatiMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type BatiMsg_BatiMsgType `protobuf:"varint,2,opt,name=type,proto3,enum=smsg.BatiMsg_BatiMsgType" json:"type,omitempty"`
	Data []byte              `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Cid  string              `protobuf:"bytes,4,opt,name=cid,proto3" json:"cid,omitempty"`
	Uid  string              `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty"`
	Ip   *string             `protobuf:"bytes,6,opt,name=ip,proto3,oneof" json:"ip,omitempty"`
	Ts   uint64              `protobuf:"varint,7,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *BatiMsg) Reset() {
	*x = BatiMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatiMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatiMsg) ProtoMessage() {}

func (x *BatiMsg) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatiMsg.ProtoReflect.Descriptor instead.
func (*BatiMsg) Descriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{0}
}

func (x *BatiMsg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BatiMsg) GetType() BatiMsg_BatiMsgType {
	if x != nil {
		return x.Type
	}
	return BatiMsg_Unused
}

func (x *BatiMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BatiMsg) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *BatiMsg) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *BatiMsg) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *BatiMsg) GetTs() uint64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type ServiceMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Servcie  string                    `protobuf:"bytes,2,opt,name=servcie,proto3" json:"servcie,omitempty"`
	Type     ServiceMsg_ServiceMsgType `protobuf:"varint,3,opt,name=type,proto3,enum=smsg.ServiceMsg_ServiceMsgType" json:"type,omitempty"`
	BizData  *BizData                  `protobuf:"bytes,4,opt,name=biz_data,json=bizData,proto3,oneof" json:"biz_data,omitempty"`
	JoinData *JoinData                 `protobuf:"bytes,5,opt,name=join_data,json=joinData,proto3,oneof" json:"join_data,omitempty"`
	QuitData *QuitData                 `protobuf:"bytes,6,opt,name=quit_data,json=quitData,proto3,oneof" json:"quit_data,omitempty"`
	Ts       uint64                    `protobuf:"varint,7,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *ServiceMsg) Reset() {
	*x = ServiceMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMsg) ProtoMessage() {}

func (x *ServiceMsg) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMsg.ProtoReflect.Descriptor instead.
func (*ServiceMsg) Descriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceMsg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceMsg) GetServcie() string {
	if x != nil {
		return x.Servcie
	}
	return ""
}

func (x *ServiceMsg) GetType() ServiceMsg_ServiceMsgType {
	if x != nil {
		return x.Type
	}
	return ServiceMsg_Unused
}

func (x *ServiceMsg) GetBizData() *BizData {
	if x != nil {
		return x.BizData
	}
	return nil
}

func (x *ServiceMsg) GetJoinData() *JoinData {
	if x != nil {
		return x.JoinData
	}
	return nil
}

func (x *ServiceMsg) GetQuitData() *QuitData {
	if x != nil {
		return x.QuitData
	}
	return nil
}

func (x *ServiceMsg) GetTs() uint64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type JoinData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid         *string  `protobuf:"bytes,1,opt,name=cid,proto3,oneof" json:"cid,omitempty"`
	Uid         *string  `protobuf:"bytes,2,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	JoinService *bool    `protobuf:"varint,3,opt,name=join_service,json=joinService,proto3,oneof" json:"join_service,omitempty"`
	Rooms       []string `protobuf:"bytes,4,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *JoinData) Reset() {
	*x = JoinData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinData) ProtoMessage() {}

func (x *JoinData) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinData.ProtoReflect.Descriptor instead.
func (*JoinData) Descriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{2}
}

func (x *JoinData) GetCid() string {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return ""
}

func (x *JoinData) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *JoinData) GetJoinService() bool {
	if x != nil && x.JoinService != nil {
		return *x.JoinService
	}
	return false
}

func (x *JoinData) GetRooms() []string {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type QuitData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid         *string  `protobuf:"bytes,1,opt,name=cid,proto3,oneof" json:"cid,omitempty"`
	Uid         *string  `protobuf:"bytes,2,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	QuitService *bool    `protobuf:"varint,3,opt,name=quit_service,json=quitService,proto3,oneof" json:"quit_service,omitempty"`
	Rooms       []string `protobuf:"bytes,4,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *QuitData) Reset() {
	*x = QuitData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitData) ProtoMessage() {}

func (x *QuitData) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitData.ProtoReflect.Descriptor instead.
func (*QuitData) Descriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{3}
}

func (x *QuitData) GetCid() string {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return ""
}

func (x *QuitData) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *QuitData) GetQuitService() bool {
	if x != nil && x.QuitService != nil {
		return *x.QuitService
	}
	return false
}

func (x *QuitData) GetRooms() []string {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type BizData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           BizData_BizMsgType `protobuf:"varint,1,opt,name=type,proto3,enum=smsg.BizData_BizMsgType" json:"type,omitempty"`
	Cids           []string           `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	Uids           []string           `protobuf:"bytes,3,rep,name=uids,proto3" json:"uids,omitempty"`
	Room           *string            `protobuf:"bytes,4,opt,name=room,proto3,oneof" json:"room,omitempty"`
	BroadcastRatio *uint32            `protobuf:"varint,5,opt,name=broadcast_ratio,json=broadcastRatio,proto3,oneof" json:"broadcast_ratio,omitempty"`
	BlackUids      []string           `protobuf:"bytes,6,rep,name=black_uids,json=blackUids,proto3" json:"black_uids,omitempty"`
	WhiteUids      []string           `protobuf:"bytes,7,rep,name=white_uids,json=whiteUids,proto3" json:"white_uids,omitempty"`
	Data           []byte             `protobuf:"bytes,8,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *BizData) Reset() {
	*x = BizData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BizData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BizData) ProtoMessage() {}

func (x *BizData) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BizData.ProtoReflect.Descriptor instead.
func (*BizData) Descriptor() ([]byte, []int) {
	return file_smsg_proto_rawDescGZIP(), []int{4}
}

func (x *BizData) GetType() BizData_BizMsgType {
	if x != nil {
		return x.Type
	}
	return BizData_Unused
}

func (x *BizData) GetCids() []string {
	if x != nil {
		return x.Cids
	}
	return nil
}

func (x *BizData) GetUids() []string {
	if x != nil {
		return x.Uids
	}
	return nil
}

func (x *BizData) GetRoom() string {
	if x != nil && x.Room != nil {
		return *x.Room
	}
	return ""
}

func (x *BizData) GetBroadcastRatio() uint32 {
	if x != nil && x.BroadcastRatio != nil {
		return *x.BroadcastRatio
	}
	return 0
}

func (x *BizData) GetBlackUids() []string {
	if x != nil {
		return x.BlackUids
	}
	return nil
}

func (x *BizData) GetWhiteUids() []string {
	if x != nil {
		return x.WhiteUids
	}
	return nil
}

func (x *BizData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_smsg_proto protoreflect.FileDescriptor

var file_smsg_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x6d,
	0x73, 0x67, 0x22, 0xec, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73,
	0x6d, 0x73, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x69,
	0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02, 0x69, 0x70, 0x88, 0x01, 0x01, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x73, 0x22,
	0x30, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x69,
	0x7a, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x51, 0x75, 0x69, 0x74, 0x10,
	0x02, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x70, 0x22, 0xfa, 0x02, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x63, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x63, 0x69, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x6d, 0x73, 0x67, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x62, 0x69, 0x7a, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x6d, 0x73, 0x67, 0x2e, 0x42, 0x69, 0x7a, 0x44, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x30,
	0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6d, 0x73, 0x67, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x01, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x30, 0x0a, 0x09, 0x71, 0x75, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6d, 0x73, 0x67, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x02, 0x52, 0x08, 0x71, 0x75, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x88,
	0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x74, 0x73, 0x22, 0x41, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x4a, 0x6f, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x51, 0x75, 0x69, 0x74, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x42, 0x69, 0x7a, 0x10, 0x03, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x69, 0x7a, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x71, 0x75, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97,
	0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x63, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x03, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6a, 0x6f, 0x69,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x0b, 0x6a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x69, 0x64, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6a, 0x6f, 0x69, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x69,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x63, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x71, 0x75, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0b, 0x71, 0x75, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69,
	0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x71, 0x75, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x07, 0x42, 0x69, 0x7a, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73,
	0x6d, 0x73, 0x67, 0x2e, 0x42, 0x69, 0x7a, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x7a, 0x4d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a,
	0x0f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x69, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x68,
	0x69, 0x74, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x77, 0x68, 0x69, 0x74, 0x65, 0x55, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88,
	0x01, 0x01, 0x22, 0x43, 0x0a, 0x0a, 0x42, 0x69, 0x7a, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x10, 0x04, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x18, 0x5a,
	0x16, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x69,
	0x67, 0x6f, 0x2f, 0x73, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smsg_proto_rawDescOnce sync.Once
	file_smsg_proto_rawDescData = file_smsg_proto_rawDesc
)

func file_smsg_proto_rawDescGZIP() []byte {
	file_smsg_proto_rawDescOnce.Do(func() {
		file_smsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_smsg_proto_rawDescData)
	})
	return file_smsg_proto_rawDescData
}

var file_smsg_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_smsg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_smsg_proto_goTypes = []interface{}{
	(BatiMsg_BatiMsgType)(0),       // 0: smsg.BatiMsg.BatiMsgType
	(ServiceMsg_ServiceMsgType)(0), // 1: smsg.ServiceMsg.ServiceMsgType
	(BizData_BizMsgType)(0),        // 2: smsg.BizData.BizMsgType
	(*BatiMsg)(nil),                // 3: smsg.BatiMsg
	(*ServiceMsg)(nil),             // 4: smsg.ServiceMsg
	(*JoinData)(nil),               // 5: smsg.JoinData
	(*QuitData)(nil),               // 6: smsg.QuitData
	(*BizData)(nil),                // 7: smsg.BizData
}
var file_smsg_proto_depIdxs = []int32{
	0, // 0: smsg.BatiMsg.type:type_name -> smsg.BatiMsg.BatiMsgType
	1, // 1: smsg.ServiceMsg.type:type_name -> smsg.ServiceMsg.ServiceMsgType
	7, // 2: smsg.ServiceMsg.biz_data:type_name -> smsg.BizData
	5, // 3: smsg.ServiceMsg.join_data:type_name -> smsg.JoinData
	6, // 4: smsg.ServiceMsg.quit_data:type_name -> smsg.QuitData
	2, // 5: smsg.BizData.type:type_name -> smsg.BizData.BizMsgType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_smsg_proto_init() }
func file_smsg_proto_init() {
	if File_smsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatiMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BizData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_smsg_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_smsg_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_smsg_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_smsg_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_smsg_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_smsg_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smsg_proto_goTypes,
		DependencyIndexes: file_smsg_proto_depIdxs,
		EnumInfos:         file_smsg_proto_enumTypes,
		MessageInfos:      file_smsg_proto_msgTypes,
	}.Build()
	File_smsg_proto = out.File
	file_smsg_proto_rawDesc = nil
	file_smsg_proto_goTypes = nil
	file_smsg_proto_depIdxs = nil
}