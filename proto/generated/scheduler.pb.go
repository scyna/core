// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: scheduler.proto

package scyna_proto

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

type StartTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context  string `protobuf:"bytes,1,opt,name=Context,proto3" json:"Context,omitempty"`
	Topic    string `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Time     int64  `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`         // Unit: second
	Interval int64  `protobuf:"varint,5,opt,name=Interval,proto3" json:"Interval,omitempty"` // In second, must be greater than 60
	Loop     uint64 `protobuf:"varint,6,opt,name=Loop,proto3" json:"Loop,omitempty"`
}

func (x *StartTaskRequest) Reset() {
	*x = StartTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTaskRequest) ProtoMessage() {}

func (x *StartTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTaskRequest.ProtoReflect.Descriptor instead.
func (*StartTaskRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *StartTaskRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *StartTaskRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *StartTaskRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StartTaskRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *StartTaskRequest) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *StartTaskRequest) GetLoop() uint64 {
	if x != nil {
		return x.Loop
	}
	return 0
}

type StartTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *StartTaskResponse) Reset() {
	*x = StartTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTaskResponse) ProtoMessage() {}

func (x *StartTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTaskResponse.ProtoReflect.Descriptor instead.
func (*StartTaskResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *StartTaskResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StopTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *StopTaskRequest) Reset() {
	*x = StopTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTaskRequest) ProtoMessage() {}

func (x *StopTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTaskRequest.ProtoReflect.Descriptor instead.
func (*StopTaskRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *StopTaskRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x4c, 0x6f, 0x6f, 0x70, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x74,
	0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x42, 0x32, 0x0a,
	0x0e, 0x69, 0x6f, 0x2e, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x48,
	0x02, 0x50, 0x01, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x0b, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_proto_rawDescData = file_scheduler_proto_rawDesc
)

func file_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_proto_rawDescData)
	})
	return file_scheduler_proto_rawDescData
}

var file_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scheduler_proto_goTypes = []interface{}{
	(*StartTaskRequest)(nil),  // 0: scyna.StartTaskRequest
	(*StartTaskResponse)(nil), // 1: scyna.StartTaskResponse
	(*StopTaskRequest)(nil),   // 2: scyna.StopTaskRequest
}
var file_scheduler_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scheduler_proto_init() }
func file_scheduler_proto_init() {
	if File_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTaskRequest); i {
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
		file_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTaskResponse); i {
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
		file_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTaskRequest); i {
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
			RawDescriptor: file_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_proto = out.File
	file_scheduler_proto_rawDesc = nil
	file_scheduler_proto_goTypes = nil
	file_scheduler_proto_depIdxs = nil
}