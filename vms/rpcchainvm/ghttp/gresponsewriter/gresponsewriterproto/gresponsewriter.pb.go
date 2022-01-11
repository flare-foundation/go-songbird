// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gresponsewriter.proto

package gresponsewriterproto

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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []*Header `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Payload []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{1}
}

func (x *WriteRequest) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *WriteRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Written int32 `protobuf:"varint,1,opt,name=written,proto3" json:"written,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{2}
}

func (x *WriteResponse) GetWritten() int32 {
	if x != nil {
		return x.Written
	}
	return 0
}

type WriteHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers    []*Header `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	StatusCode int32     `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *WriteHeaderRequest) Reset() {
	*x = WriteHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteHeaderRequest) ProtoMessage() {}

func (x *WriteHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteHeaderRequest.ProtoReflect.Descriptor instead.
func (*WriteHeaderRequest) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{3}
}

func (x *WriteHeaderRequest) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *WriteHeaderRequest) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type WriteHeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteHeaderResponse) Reset() {
	*x = WriteHeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteHeaderResponse) ProtoMessage() {}

func (x *WriteHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteHeaderResponse.ProtoReflect.Descriptor instead.
func (*WriteHeaderResponse) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{4}
}

type FlushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushRequest) Reset() {
	*x = FlushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushRequest) ProtoMessage() {}

func (x *FlushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushRequest.ProtoReflect.Descriptor instead.
func (*FlushRequest) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{5}
}

type FlushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushResponse) Reset() {
	*x = FlushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushResponse) ProtoMessage() {}

func (x *FlushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushResponse.ProtoReflect.Descriptor instead.
func (*FlushResponse) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{6}
}

type HijackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HijackRequest) Reset() {
	*x = HijackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HijackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HijackRequest) ProtoMessage() {}

func (x *HijackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HijackRequest.ProtoReflect.Descriptor instead.
func (*HijackRequest) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{7}
}

type HijackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnServer    uint32 `protobuf:"varint,1,opt,name=connServer,proto3" json:"connServer,omitempty"`
	LocalNetwork  string `protobuf:"bytes,2,opt,name=localNetwork,proto3" json:"localNetwork,omitempty"`
	LocalString   string `protobuf:"bytes,3,opt,name=localString,proto3" json:"localString,omitempty"`
	RemoteNetwork string `protobuf:"bytes,4,opt,name=remoteNetwork,proto3" json:"remoteNetwork,omitempty"`
	RemoteString  string `protobuf:"bytes,5,opt,name=remoteString,proto3" json:"remoteString,omitempty"`
	ReaderServer  uint32 `protobuf:"varint,6,opt,name=readerServer,proto3" json:"readerServer,omitempty"`
	WriterServer  uint32 `protobuf:"varint,7,opt,name=writerServer,proto3" json:"writerServer,omitempty"`
}

func (x *HijackResponse) Reset() {
	*x = HijackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gresponsewriter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HijackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HijackResponse) ProtoMessage() {}

func (x *HijackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gresponsewriter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HijackResponse.ProtoReflect.Descriptor instead.
func (*HijackResponse) Descriptor() ([]byte, []int) {
	return file_gresponsewriter_proto_rawDescGZIP(), []int{8}
}

func (x *HijackResponse) GetConnServer() uint32 {
	if x != nil {
		return x.ConnServer
	}
	return 0
}

func (x *HijackResponse) GetLocalNetwork() string {
	if x != nil {
		return x.LocalNetwork
	}
	return ""
}

func (x *HijackResponse) GetLocalString() string {
	if x != nil {
		return x.LocalString
	}
	return ""
}

func (x *HijackResponse) GetRemoteNetwork() string {
	if x != nil {
		return x.RemoteNetwork
	}
	return ""
}

func (x *HijackResponse) GetRemoteString() string {
	if x != nil {
		return x.RemoteString
	}
	return ""
}

func (x *HijackResponse) GetReaderServer() uint32 {
	if x != nil {
		return x.ReaderServer
	}
	return 0
}

func (x *HijackResponse) GetWriterServer() uint32 {
	if x != nil {
		return x.WriterServer
	}
	return 0
}

var File_gresponsewriter_proto protoreflect.FileDescriptor

var file_gresponsewriter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x60, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x22, 0x6c,
	0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x69, 0x6a, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x88, 0x02, 0x0a, 0x0e, 0x48, 0x69, 0x6a, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x24,
	0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x32, 0xe5, 0x02, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x05, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x67,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x05, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x48, 0x69, 0x6a, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x2e,
	0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x69, 0x6a, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x69, 0x6a, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2d, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x6d, 0x2f, 0x67, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gresponsewriter_proto_rawDescOnce sync.Once
	file_gresponsewriter_proto_rawDescData = file_gresponsewriter_proto_rawDesc
)

func file_gresponsewriter_proto_rawDescGZIP() []byte {
	file_gresponsewriter_proto_rawDescOnce.Do(func() {
		file_gresponsewriter_proto_rawDescData = protoimpl.X.CompressGZIP(file_gresponsewriter_proto_rawDescData)
	})
	return file_gresponsewriter_proto_rawDescData
}

var file_gresponsewriter_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gresponsewriter_proto_goTypes = []interface{}{
	(*Header)(nil),              // 0: gresponsewriterproto.Header
	(*WriteRequest)(nil),        // 1: gresponsewriterproto.WriteRequest
	(*WriteResponse)(nil),       // 2: gresponsewriterproto.WriteResponse
	(*WriteHeaderRequest)(nil),  // 3: gresponsewriterproto.WriteHeaderRequest
	(*WriteHeaderResponse)(nil), // 4: gresponsewriterproto.WriteHeaderResponse
	(*FlushRequest)(nil),        // 5: gresponsewriterproto.FlushRequest
	(*FlushResponse)(nil),       // 6: gresponsewriterproto.FlushResponse
	(*HijackRequest)(nil),       // 7: gresponsewriterproto.HijackRequest
	(*HijackResponse)(nil),      // 8: gresponsewriterproto.HijackResponse
}
var file_gresponsewriter_proto_depIdxs = []int32{
	0, // 0: gresponsewriterproto.WriteRequest.headers:type_name -> gresponsewriterproto.Header
	0, // 1: gresponsewriterproto.WriteHeaderRequest.headers:type_name -> gresponsewriterproto.Header
	1, // 2: gresponsewriterproto.Writer.Write:input_type -> gresponsewriterproto.WriteRequest
	3, // 3: gresponsewriterproto.Writer.WriteHeader:input_type -> gresponsewriterproto.WriteHeaderRequest
	5, // 4: gresponsewriterproto.Writer.Flush:input_type -> gresponsewriterproto.FlushRequest
	7, // 5: gresponsewriterproto.Writer.Hijack:input_type -> gresponsewriterproto.HijackRequest
	2, // 6: gresponsewriterproto.Writer.Write:output_type -> gresponsewriterproto.WriteResponse
	4, // 7: gresponsewriterproto.Writer.WriteHeader:output_type -> gresponsewriterproto.WriteHeaderResponse
	6, // 8: gresponsewriterproto.Writer.Flush:output_type -> gresponsewriterproto.FlushResponse
	8, // 9: gresponsewriterproto.Writer.Hijack:output_type -> gresponsewriterproto.HijackResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gresponsewriter_proto_init() }
func file_gresponsewriter_proto_init() {
	if File_gresponsewriter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gresponsewriter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_gresponsewriter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_gresponsewriter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_gresponsewriter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteHeaderRequest); i {
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
		file_gresponsewriter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteHeaderResponse); i {
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
		file_gresponsewriter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushRequest); i {
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
		file_gresponsewriter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushResponse); i {
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
		file_gresponsewriter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HijackRequest); i {
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
		file_gresponsewriter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HijackResponse); i {
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
			RawDescriptor: file_gresponsewriter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gresponsewriter_proto_goTypes,
		DependencyIndexes: file_gresponsewriter_proto_depIdxs,
		MessageInfos:      file_gresponsewriter_proto_msgTypes,
	}.Build()
	File_gresponsewriter_proto = out.File
	file_gresponsewriter_proto_rawDesc = nil
	file_gresponsewriter_proto_goTypes = nil
	file_gresponsewriter_proto_depIdxs = nil
}
