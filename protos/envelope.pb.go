// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: protos/envelope.proto

package protos

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

type VerifyBitcoinDepositsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyBitcoinDepositsRequest) Reset() {
	*x = VerifyBitcoinDepositsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBitcoinDepositsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBitcoinDepositsRequest) ProtoMessage() {}

func (x *VerifyBitcoinDepositsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBitcoinDepositsRequest.ProtoReflect.Descriptor instead.
func (*VerifyBitcoinDepositsRequest) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyBitcoinDepositsRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyBitcoinDepositsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyBitcoinDepositsResponse) Reset() {
	*x = VerifyBitcoinDepositsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBitcoinDepositsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBitcoinDepositsResponse) ProtoMessage() {}

func (x *VerifyBitcoinDepositsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBitcoinDepositsResponse.ProtoReflect.Descriptor instead.
func (*VerifyBitcoinDepositsResponse) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyBitcoinDepositsResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyTokenDepositsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyTokenDepositsRequest) Reset() {
	*x = VerifyTokenDepositsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenDepositsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenDepositsRequest) ProtoMessage() {}

func (x *VerifyTokenDepositsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenDepositsRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenDepositsRequest) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyTokenDepositsRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyTokenDepositsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyTokenDepositsResponse) Reset() {
	*x = VerifyTokenDepositsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenDepositsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenDepositsResponse) ProtoMessage() {}

func (x *VerifyTokenDepositsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenDepositsResponse.ProtoReflect.Descriptor instead.
func (*VerifyTokenDepositsResponse) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyTokenDepositsResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyBitcoinWithdrawalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyBitcoinWithdrawalsRequest) Reset() {
	*x = VerifyBitcoinWithdrawalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBitcoinWithdrawalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBitcoinWithdrawalsRequest) ProtoMessage() {}

func (x *VerifyBitcoinWithdrawalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBitcoinWithdrawalsRequest.ProtoReflect.Descriptor instead.
func (*VerifyBitcoinWithdrawalsRequest) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyBitcoinWithdrawalsRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyBitcoinWithdrawalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyBitcoinWithdrawalsResponse) Reset() {
	*x = VerifyBitcoinWithdrawalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBitcoinWithdrawalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBitcoinWithdrawalsResponse) ProtoMessage() {}

func (x *VerifyBitcoinWithdrawalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBitcoinWithdrawalsResponse.ProtoReflect.Descriptor instead.
func (*VerifyBitcoinWithdrawalsResponse) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyBitcoinWithdrawalsResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyTokenWithdrawalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyTokenWithdrawalsRequest) Reset() {
	*x = VerifyTokenWithdrawalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenWithdrawalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenWithdrawalsRequest) ProtoMessage() {}

func (x *VerifyTokenWithdrawalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenWithdrawalsRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenWithdrawalsRequest) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyTokenWithdrawalsRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyTokenWithdrawalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VerifyTokenWithdrawalsResponse) Reset() {
	*x = VerifyTokenWithdrawalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_envelope_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenWithdrawalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenWithdrawalsResponse) ProtoMessage() {}

func (x *VerifyTokenWithdrawalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_envelope_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenWithdrawalsResponse.ProtoReflect.Descriptor instead.
func (*VerifyTokenWithdrawalsResponse) Descriptor() ([]byte, []int) {
	return file_protos_envelope_proto_rawDescGZIP(), []int{7}
}

func (x *VerifyTokenWithdrawalsResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_protos_envelope_proto protoreflect.FileDescriptor

var file_protos_envelope_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x22, 0x36, 0x0a, 0x1c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f,
	0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x37, 0x0a, 0x1d, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x22, 0x34, 0x0a, 0x1a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x35, 0x0a, 0x1b, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x39, 0x0a, 0x1f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3a, 0x0a, 0x20, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x37, 0x0a, 0x1d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x38, 0x0a, 0x1e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xbf, 0x03, 0x0a, 0x0f, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a,
	0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x12, 0x24,
	0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x18, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b,
	0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72, 0x61, 0x2d, 0x6e,
	0x77, 0x2f, 0x62, 0x74, 0x63, 0x2d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_envelope_proto_rawDescOnce sync.Once
	file_protos_envelope_proto_rawDescData = file_protos_envelope_proto_rawDesc
)

func file_protos_envelope_proto_rawDescGZIP() []byte {
	file_protos_envelope_proto_rawDescOnce.Do(func() {
		file_protos_envelope_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_envelope_proto_rawDescData)
	})
	return file_protos_envelope_proto_rawDescData
}

var file_protos_envelope_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_envelope_proto_goTypes = []interface{}{
	(*VerifyBitcoinDepositsRequest)(nil),     // 0: envelope.VerifyBitcoinDepositsRequest
	(*VerifyBitcoinDepositsResponse)(nil),    // 1: envelope.VerifyBitcoinDepositsResponse
	(*VerifyTokenDepositsRequest)(nil),       // 2: envelope.VerifyTokenDepositsRequest
	(*VerifyTokenDepositsResponse)(nil),      // 3: envelope.VerifyTokenDepositsResponse
	(*VerifyBitcoinWithdrawalsRequest)(nil),  // 4: envelope.VerifyBitcoinWithdrawalsRequest
	(*VerifyBitcoinWithdrawalsResponse)(nil), // 5: envelope.VerifyBitcoinWithdrawalsResponse
	(*VerifyTokenWithdrawalsRequest)(nil),    // 6: envelope.VerifyTokenWithdrawalsRequest
	(*VerifyTokenWithdrawalsResponse)(nil),   // 7: envelope.VerifyTokenWithdrawalsResponse
}
var file_protos_envelope_proto_depIdxs = []int32{
	0, // 0: envelope.EnvelopeService.VerifyBitcoinDeposits:input_type -> envelope.VerifyBitcoinDepositsRequest
	2, // 1: envelope.EnvelopeService.VerifyTokenDeposits:input_type -> envelope.VerifyTokenDepositsRequest
	4, // 2: envelope.EnvelopeService.VerifyBitcoinWithdrawals:input_type -> envelope.VerifyBitcoinWithdrawalsRequest
	6, // 3: envelope.EnvelopeService.VerifyTokenWithdrawals:input_type -> envelope.VerifyTokenWithdrawalsRequest
	1, // 4: envelope.EnvelopeService.VerifyBitcoinDeposits:output_type -> envelope.VerifyBitcoinDepositsResponse
	3, // 5: envelope.EnvelopeService.VerifyTokenDeposits:output_type -> envelope.VerifyTokenDepositsResponse
	5, // 6: envelope.EnvelopeService.VerifyBitcoinWithdrawals:output_type -> envelope.VerifyBitcoinWithdrawalsResponse
	7, // 7: envelope.EnvelopeService.VerifyTokenWithdrawals:output_type -> envelope.VerifyTokenWithdrawalsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_envelope_proto_init() }
func file_protos_envelope_proto_init() {
	if File_protos_envelope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_envelope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBitcoinDepositsRequest); i {
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
		file_protos_envelope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBitcoinDepositsResponse); i {
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
		file_protos_envelope_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenDepositsRequest); i {
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
		file_protos_envelope_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenDepositsResponse); i {
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
		file_protos_envelope_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBitcoinWithdrawalsRequest); i {
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
		file_protos_envelope_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBitcoinWithdrawalsResponse); i {
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
		file_protos_envelope_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenWithdrawalsRequest); i {
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
		file_protos_envelope_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenWithdrawalsResponse); i {
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
			RawDescriptor: file_protos_envelope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_envelope_proto_goTypes,
		DependencyIndexes: file_protos_envelope_proto_depIdxs,
		MessageInfos:      file_protos_envelope_proto_msgTypes,
	}.Build()
	File_protos_envelope_proto = out.File
	file_protos_envelope_proto_rawDesc = nil
	file_protos_envelope_proto_goTypes = nil
	file_protos_envelope_proto_depIdxs = nil
}
