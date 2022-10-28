// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: conn.proto

package pb

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

// 推送事件
type ConnMQBody_Event int32

const (
	// 给某些用户推送消息
	ConnMQBody_SEND_MSG_TO_USERS ConnMQBody_Event = 0
	// 强制断开用户连接
	ConnMQBody_FORCE_DISCONNECT ConnMQBody_Event = 1
)

// Enum value maps for ConnMQBody_Event.
var (
	ConnMQBody_Event_name = map[int32]string{
		0: "SEND_MSG_TO_USERS",
		1: "FORCE_DISCONNECT",
	}
	ConnMQBody_Event_value = map[string]int32{
		"SEND_MSG_TO_USERS": 0,
		"FORCE_DISCONNECT":  1,
	}
)

func (x ConnMQBody_Event) Enum() *ConnMQBody_Event {
	p := new(ConnMQBody_Event)
	*p = x
	return p
}

func (x ConnMQBody_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnMQBody_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[0].Descriptor()
}

func (ConnMQBody_Event) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[0]
}

func (x ConnMQBody_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnMQBody_Event.Descriptor instead.
func (ConnMQBody_Event) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0, 0}
}

type MQSendMsgToUsersData_MsgType int32

const (
	MQSendMsgToUsersData_MSGDATALIST MQSendMsgToUsersData_MsgType = 0 // MsgDataList
)

// Enum value maps for MQSendMsgToUsersData_MsgType.
var (
	MQSendMsgToUsersData_MsgType_name = map[int32]string{
		0: "MSGDATALIST",
	}
	MQSendMsgToUsersData_MsgType_value = map[string]int32{
		"MSGDATALIST": 0,
	}
)

func (x MQSendMsgToUsersData_MsgType) Enum() *MQSendMsgToUsersData_MsgType {
	p := new(MQSendMsgToUsersData_MsgType)
	*p = x
	return p
}

func (x MQSendMsgToUsersData_MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MQSendMsgToUsersData_MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[1].Descriptor()
}

func (MQSendMsgToUsersData_MsgType) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[1]
}

func (x MQSendMsgToUsersData_MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MQSendMsgToUsersData_MsgType.Descriptor instead.
func (MQSendMsgToUsersData_MsgType) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2, 0}
}

// 服务器通过websocket发送给客户端的消息体
type ConnMQBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event ConnMQBody_Event `protobuf:"varint,1,opt,name=event,proto3,enum=pb.ConnMQBody_Event" json:"event"`
	Data  []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *ConnMQBody) Reset() {
	*x = ConnMQBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnMQBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnMQBody) ProtoMessage() {}

func (x *ConnMQBody) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnMQBody.ProtoReflect.Descriptor instead.
func (*ConnMQBody) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

func (x *ConnMQBody) GetEvent() ConnMQBody_Event {
	if x != nil {
		return x.Event
	}
	return ConnMQBody_SEND_MSG_TO_USERS
}

func (x *ConnMQBody) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MQForceDisconnectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds"`
	Code    int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code"`
	Msg     string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg"`
}

func (x *MQForceDisconnectData) Reset() {
	*x = MQForceDisconnectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQForceDisconnectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQForceDisconnectData) ProtoMessage() {}

func (x *MQForceDisconnectData) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQForceDisconnectData.ProtoReflect.Descriptor instead.
func (*MQForceDisconnectData) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{1}
}

func (x *MQForceDisconnectData) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *MQForceDisconnectData) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MQForceDisconnectData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type MQSendMsgToUsersData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string                  `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds"`
	Msg     *MQSendMsgToUsersData_Msg `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *MQSendMsgToUsersData) Reset() {
	*x = MQSendMsgToUsersData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQSendMsgToUsersData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQSendMsgToUsersData) ProtoMessage() {}

func (x *MQSendMsgToUsersData) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQSendMsgToUsersData.ProtoReflect.Descriptor instead.
func (*MQSendMsgToUsersData) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2}
}

func (x *MQSendMsgToUsersData) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *MQSendMsgToUsersData) GetMsg() *MQSendMsgToUsersData_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type MQSendMsgToUsersData_Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType MQSendMsgToUsersData_MsgType `protobuf:"varint,1,opt,name=msgType,proto3,enum=pb.MQSendMsgToUsersData_MsgType" json:"msgType"`
	Payload []byte                       `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload"`
}

func (x *MQSendMsgToUsersData_Msg) Reset() {
	*x = MQSendMsgToUsersData_Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQSendMsgToUsersData_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQSendMsgToUsersData_Msg) ProtoMessage() {}

func (x *MQSendMsgToUsersData_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQSendMsgToUsersData_Msg.ProtoReflect.Descriptor instead.
func (*MQSendMsgToUsersData_Msg) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MQSendMsgToUsersData_Msg) GetMsgType() MQSendMsgToUsersData_MsgType {
	if x != nil {
		return x.MsgType
	}
	return MQSendMsgToUsersData_MSGDATALIST
}

func (x *MQSendMsgToUsersData_Msg) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_conn_proto protoreflect.FileDescriptor

var file_conn_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82,
	0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x51, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x51, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x34, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x4d,
	0x53, 0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x10, 0x01, 0x22, 0x57, 0x0a, 0x15, 0x4d, 0x51, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xd9, 0x01, 0x0a,
	0x14, 0x4d, 0x51, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12,
	0x2e, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x51, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x1a,
	0x5b, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x51, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x1a, 0x0a, 0x07,
	0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x53, 0x47, 0x44, 0x41,
	0x54, 0x41, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conn_proto_rawDescOnce sync.Once
	file_conn_proto_rawDescData = file_conn_proto_rawDesc
)

func file_conn_proto_rawDescGZIP() []byte {
	file_conn_proto_rawDescOnce.Do(func() {
		file_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_conn_proto_rawDescData)
	})
	return file_conn_proto_rawDescData
}

var file_conn_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conn_proto_goTypes = []interface{}{
	(ConnMQBody_Event)(0),             // 0: pb.ConnMQBody.Event
	(MQSendMsgToUsersData_MsgType)(0), // 1: pb.MQSendMsgToUsersData.MsgType
	(*ConnMQBody)(nil),                // 2: pb.ConnMQBody
	(*MQForceDisconnectData)(nil),     // 3: pb.MQForceDisconnectData
	(*MQSendMsgToUsersData)(nil),      // 4: pb.MQSendMsgToUsersData
	(*MQSendMsgToUsersData_Msg)(nil),  // 5: pb.MQSendMsgToUsersData.Msg
}
var file_conn_proto_depIdxs = []int32{
	0, // 0: pb.ConnMQBody.event:type_name -> pb.ConnMQBody.Event
	5, // 1: pb.MQSendMsgToUsersData.msg:type_name -> pb.MQSendMsgToUsersData.Msg
	1, // 2: pb.MQSendMsgToUsersData.Msg.msgType:type_name -> pb.MQSendMsgToUsersData.MsgType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conn_proto_init() }
func file_conn_proto_init() {
	if File_conn_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnMQBody); i {
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
		file_conn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQForceDisconnectData); i {
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
		file_conn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQSendMsgToUsersData); i {
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
		file_conn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQSendMsgToUsersData_Msg); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conn_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conn_proto_goTypes,
		DependencyIndexes: file_conn_proto_depIdxs,
		EnumInfos:         file_conn_proto_enumTypes,
		MessageInfos:      file_conn_proto_msgTypes,
	}.Build()
	File_conn_proto = out.File
	file_conn_proto_rawDesc = nil
	file_conn_proto_goTypes = nil
	file_conn_proto_depIdxs = nil
}
