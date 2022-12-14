// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: friend.proto

package proto

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//ADD FRIEND
type AddFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Friend string `protobuf:"bytes,2,opt,name=friend,proto3" json:"friend,omitempty"` //email of friend
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendRequest.ProtoReflect.Descriptor instead.
func (*AddFriendRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{1}
}

func (x *AddFriendRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddFriendRequest) GetFriend() string {
	if x != nil {
		return x.Friend
	}
	return ""
}

//LIST
type ListFriendByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListFriendByIDRequest) Reset() {
	*x = ListFriendByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendByIDRequest) ProtoMessage() {}

func (x *ListFriendByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendByIDRequest.ProtoReflect.Descriptor instead.
func (*ListFriendByIDRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{2}
}

func (x *ListFriendByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ListFriendRequest) Reset() {
	*x = ListFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendRequest) ProtoMessage() {}

func (x *ListFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendRequest.ProtoReflect.Descriptor instead.
func (*ListFriendRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{3}
}

func (x *ListFriendRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ListFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Account `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListFriendResponse) Reset() {
	*x = ListFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendResponse) ProtoMessage() {}

func (x *ListFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendResponse.ProtoReflect.Descriptor instead.
func (*ListFriendResponse) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{4}
}

func (x *ListFriendResponse) GetItems() []*Account {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_friend_proto protoreflect.FileDescriptor

var file_friend_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x34, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42,
	0x13, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_proto_rawDescOnce sync.Once
	file_friend_proto_rawDescData = file_friend_proto_rawDesc
)

func file_friend_proto_rawDescGZIP() []byte {
	file_friend_proto_rawDescOnce.Do(func() {
		file_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_proto_rawDescData)
	})
	return file_friend_proto_rawDescData
}

var file_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_friend_proto_goTypes = []interface{}{
	(*Account)(nil),               // 0: Account
	(*AddFriendRequest)(nil),      // 1: AddFriendRequest
	(*ListFriendByIDRequest)(nil), // 2: ListFriendByIDRequest
	(*ListFriendRequest)(nil),     // 3: ListFriendRequest
	(*ListFriendResponse)(nil),    // 4: ListFriendResponse
}
var file_friend_proto_depIdxs = []int32{
	0, // 0: ListFriendResponse.items:type_name -> Account
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_proto_init() }
func file_friend_proto_init() {
	if File_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendRequest); i {
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
		file_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendByIDRequest); i {
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
		file_friend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendRequest); i {
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
		file_friend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendResponse); i {
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
			RawDescriptor: file_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_friend_proto_goTypes,
		DependencyIndexes: file_friend_proto_depIdxs,
		MessageInfos:      file_friend_proto_msgTypes,
	}.Build()
	File_friend_proto = out.File
	file_friend_proto_rawDesc = nil
	file_friend_proto_goTypes = nil
	file_friend_proto_depIdxs = nil
}
