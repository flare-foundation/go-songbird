// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: gkeystoreproto/keystore.proto

package gkeystoreproto

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

type GetDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetDatabaseRequest) Reset() {
	*x = GetDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkeystoreproto_keystore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseRequest) ProtoMessage() {}

func (x *GetDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gkeystoreproto_keystore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_gkeystoreproto_keystore_proto_rawDescGZIP(), []int{0}
}

func (x *GetDatabaseRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetDatabaseRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbServer uint32 `protobuf:"varint,1,opt,name=db_server,json=dbServer,proto3" json:"db_server,omitempty"`
}

func (x *GetDatabaseResponse) Reset() {
	*x = GetDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkeystoreproto_keystore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseResponse) ProtoMessage() {}

func (x *GetDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gkeystoreproto_keystore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseResponse.ProtoReflect.Descriptor instead.
func (*GetDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_gkeystoreproto_keystore_proto_rawDescGZIP(), []int{1}
}

func (x *GetDatabaseResponse) GetDbServer() uint32 {
	if x != nil {
		return x.DbServer
	}
	return 0
}

var File_gkeystoreproto_keystore_proto protoreflect.FileDescriptor

var file_gkeystoreproto_keystore_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x32, 0x62, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x56, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x22, 0x2e, 0x67,
	0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x76, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x67, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gkeystoreproto_keystore_proto_rawDescOnce sync.Once
	file_gkeystoreproto_keystore_proto_rawDescData = file_gkeystoreproto_keystore_proto_rawDesc
)

func file_gkeystoreproto_keystore_proto_rawDescGZIP() []byte {
	file_gkeystoreproto_keystore_proto_rawDescOnce.Do(func() {
		file_gkeystoreproto_keystore_proto_rawDescData = protoimpl.X.CompressGZIP(file_gkeystoreproto_keystore_proto_rawDescData)
	})
	return file_gkeystoreproto_keystore_proto_rawDescData
}

var file_gkeystoreproto_keystore_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gkeystoreproto_keystore_proto_goTypes = []interface{}{
	(*GetDatabaseRequest)(nil),  // 0: gkeystoreproto.GetDatabaseRequest
	(*GetDatabaseResponse)(nil), // 1: gkeystoreproto.GetDatabaseResponse
}
var file_gkeystoreproto_keystore_proto_depIdxs = []int32{
	0, // 0: gkeystoreproto.Keystore.GetDatabase:input_type -> gkeystoreproto.GetDatabaseRequest
	1, // 1: gkeystoreproto.Keystore.GetDatabase:output_type -> gkeystoreproto.GetDatabaseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gkeystoreproto_keystore_proto_init() }
func file_gkeystoreproto_keystore_proto_init() {
	if File_gkeystoreproto_keystore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gkeystoreproto_keystore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseRequest); i {
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
		file_gkeystoreproto_keystore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseResponse); i {
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
			RawDescriptor: file_gkeystoreproto_keystore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gkeystoreproto_keystore_proto_goTypes,
		DependencyIndexes: file_gkeystoreproto_keystore_proto_depIdxs,
		MessageInfos:      file_gkeystoreproto_keystore_proto_msgTypes,
	}.Build()
	File_gkeystoreproto_keystore_proto = out.File
	file_gkeystoreproto_keystore_proto_rawDesc = nil
	file_gkeystoreproto_keystore_proto_goTypes = nil
	file_gkeystoreproto_keystore_proto_depIdxs = nil
}
