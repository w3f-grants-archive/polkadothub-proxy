// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: grpc/transaction/transactionpb/transaction.proto

package transactionpb

import (
	context "context"
	eventpb "github.com/figment-networks/polkadothub-proxy/grpc/event/eventpb"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CallArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Section string `protobuf:"bytes,2,opt,name=section,proto3" json:"section,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CallArg) Reset() {
	*x = CallArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallArg) ProtoMessage() {}

func (x *CallArg) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallArg.ProtoReflect.Descriptor instead.
func (*CallArg) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *CallArg) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CallArg) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *CallArg) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtrinsicIndex      int64      `protobuf:"varint,1,opt,name=extrinsic_index,json=extrinsicIndex,proto3" json:"extrinsic_index,omitempty"`
	Hash                string     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Time                string     `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Signature           string     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Signer              string     `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
	Nonce               int64      `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Method              string     `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	Section             string     `protobuf:"bytes,8,opt,name=section,proto3" json:"section,omitempty"`
	Args                string     `protobuf:"bytes,9,opt,name=args,proto3" json:"args,omitempty"`
	IsSuccess           bool       `protobuf:"varint,10,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	PartialFee          string     `protobuf:"bytes,11,opt,name=partial_fee,json=partialFee,proto3" json:"partial_fee,omitempty"`
	Tip                 string     `protobuf:"bytes,12,opt,name=tip,proto3" json:"tip,omitempty"`
	Raw                 string     `protobuf:"bytes,13,opt,name=raw,proto3" json:"raw,omitempty"`
	IsSignedTransaction bool       `protobuf:"varint,14,opt,name=is_signed_transaction,json=isSignedTransaction,proto3" json:"is_signed_transaction,omitempty"`
	CallArgs            []*CallArg `protobuf:"bytes,15,rep,name=callArgs,proto3" json:"callArgs,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetExtrinsicIndex() int64 {
	if x != nil {
		return x.ExtrinsicIndex
	}
	return 0
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Transaction) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Transaction) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *Transaction) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Transaction) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Transaction) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *Transaction) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *Transaction) GetPartialFee() string {
	if x != nil {
		return x.PartialFee
	}
	return ""
}

func (x *Transaction) GetTip() string {
	if x != nil {
		return x.Tip
	}
	return ""
}

func (x *Transaction) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *Transaction) GetIsSignedTransaction() bool {
	if x != nil {
		return x.IsSignedTransaction
	}
	return false
}

func (x *Transaction) GetCallArgs() []*CallArg {
	if x != nil {
		return x.CallArgs
	}
	return nil
}

type Annotated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtrinsicIndex int64  `protobuf:"varint,1,opt,name=extrinsic_index,json=extrinsicIndex,proto3" json:"extrinsic_index,omitempty"`
	Hash           string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Method         string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Section        string `protobuf:"bytes,4,opt,name=section,proto3" json:"section,omitempty"`
	IsSigned       bool   `protobuf:"varint,5,opt,name=is_signed,json=isSigned,proto3" json:"is_signed,omitempty"`
	Args           string `protobuf:"bytes,6,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *Annotated) Reset() {
	*x = Annotated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Annotated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Annotated) ProtoMessage() {}

func (x *Annotated) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Annotated.ProtoReflect.Descriptor instead.
func (*Annotated) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *Annotated) GetExtrinsicIndex() int64 {
	if x != nil {
		return x.ExtrinsicIndex
	}
	return 0
}

func (x *Annotated) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Annotated) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Annotated) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Annotated) GetIsSigned() bool {
	if x != nil {
		return x.IsSigned
	}
	return false
}

func (x *Annotated) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

type GetByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetByHeightRequest) Reset() {
	*x = GetByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHeightRequest) ProtoMessage() {}

func (x *GetByHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByHeightRequest.ProtoReflect.Descriptor instead.
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *GetByHeightRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetByHeightResponse) Reset() {
	*x = GetByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHeightResponse) ProtoMessage() {}

func (x *GetByHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByHeightResponse.ProtoReflect.Descriptor instead.
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *GetByHeightResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type GetAnnotatedByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetAnnotatedByHeightRequest) Reset() {
	*x = GetAnnotatedByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnotatedByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnotatedByHeightRequest) ProtoMessage() {}

func (x *GetAnnotatedByHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnotatedByHeightRequest.ProtoReflect.Descriptor instead.
func (*GetAnnotatedByHeightRequest) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *GetAnnotatedByHeightRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetAnnotatedByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Annotated `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetAnnotatedByHeightResponse) Reset() {
	*x = GetAnnotatedByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnotatedByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnotatedByHeightResponse) ProtoMessage() {}

func (x *GetAnnotatedByHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnotatedByHeightResponse.ProtoReflect.Descriptor instead.
func (*GetAnnotatedByHeightResponse) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *GetAnnotatedByHeightResponse) GetTransactions() []*Annotated {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_grpc_transaction_transactionpb_transaction_proto protoreflect.FileDescriptor

var file_grpc_transaction_transactionpb_transaction_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x51, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xba, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x74,
	0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x73, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x73, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x41, 0x72, 0x67, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x22,
	0xab, 0x01, 0x0a, 0x09, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x2c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x35, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5a, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0xd7, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4e, 0x5a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x6f, 0x6c,
	0x6b, 0x61, 0x64, 0x6f, 0x74, 0x68, 0x75, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_transaction_transactionpb_transaction_proto_rawDescOnce sync.Once
	file_grpc_transaction_transactionpb_transaction_proto_rawDescData = file_grpc_transaction_transactionpb_transaction_proto_rawDesc
)

func file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP() []byte {
	file_grpc_transaction_transactionpb_transaction_proto_rawDescOnce.Do(func() {
		file_grpc_transaction_transactionpb_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_transaction_transactionpb_transaction_proto_rawDescData)
	})
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescData
}

var file_grpc_transaction_transactionpb_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_transaction_transactionpb_transaction_proto_goTypes = []interface{}{
	(*CallArg)(nil),                      // 0: transaction.CallArg
	(*Transaction)(nil),                  // 1: transaction.Transaction
	(*Annotated)(nil),                    // 2: transaction.Annotated
	(*GetByHeightRequest)(nil),           // 3: transaction.GetByHeightRequest
	(*GetByHeightResponse)(nil),          // 4: transaction.GetByHeightResponse
	(*GetAnnotatedByHeightRequest)(nil),  // 5: transaction.GetAnnotatedByHeightRequest
	(*GetAnnotatedByHeightResponse)(nil), // 6: transaction.GetAnnotatedByHeightResponse
}
var file_grpc_transaction_transactionpb_transaction_proto_depIdxs = []int32{
	0, // 0: transaction.Transaction.callArgs:type_name -> transaction.CallArg
	1, // 1: transaction.GetByHeightResponse.transactions:type_name -> transaction.Transaction
	2, // 2: transaction.GetAnnotatedByHeightResponse.transactions:type_name -> transaction.Annotated
	3, // 3: transaction.TransactionService.GetByHeight:input_type -> transaction.GetByHeightRequest
	5, // 4: transaction.TransactionService.GetAnnotatedByHeight:input_type -> transaction.GetAnnotatedByHeightRequest
	4, // 5: transaction.TransactionService.GetByHeight:output_type -> transaction.GetByHeightResponse
	6, // 6: transaction.TransactionService.GetAnnotatedByHeight:output_type -> transaction.GetAnnotatedByHeightResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_transaction_transactionpb_transaction_proto_init() }
func file_grpc_transaction_transactionpb_transaction_proto_init() {
	if File_grpc_transaction_transactionpb_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallArg); i {
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
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Annotated); i {
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
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByHeightRequest); i {
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
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByHeightResponse); i {
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
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnotatedByHeightRequest); i {
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
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnotatedByHeightResponse); i {
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
			RawDescriptor: file_grpc_transaction_transactionpb_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_transaction_transactionpb_transaction_proto_goTypes,
		DependencyIndexes: file_grpc_transaction_transactionpb_transaction_proto_depIdxs,
		MessageInfos:      file_grpc_transaction_transactionpb_transaction_proto_msgTypes,
	}.Build()
	File_grpc_transaction_transactionpb_transaction_proto = out.File
	file_grpc_transaction_transactionpb_transaction_proto_rawDesc = nil
	file_grpc_transaction_transactionpb_transaction_proto_goTypes = nil
	file_grpc_transaction_transactionpb_transaction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
	GetAnnotatedByHeight(ctx context.Context, in *GetAnnotatedByHeightRequest, opts ...grpc.CallOption) (*GetAnnotatedByHeightResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetAnnotatedByHeight(ctx context.Context, in *GetAnnotatedByHeightRequest, opts ...grpc.CallOption) (*GetAnnotatedByHeightResponse, error) {
	out := new(GetAnnotatedByHeightResponse)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/GetAnnotatedByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
	GetAnnotatedByHeight(context.Context, *GetAnnotatedByHeightRequest) (*GetAnnotatedByHeightResponse, error)
}

// UnimplementedTransactionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (*UnimplementedTransactionServiceServer) GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}
func (*UnimplementedTransactionServiceServer) GetAnnotatedByHeight(context.Context, *GetAnnotatedByHeightRequest) (*GetAnnotatedByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnotatedByHeight not implemented")
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetAnnotatedByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnotatedByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetAnnotatedByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/GetAnnotatedByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetAnnotatedByHeight(ctx, req.(*GetAnnotatedByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _TransactionService_GetByHeight_Handler,
		},
		{
			MethodName: "GetAnnotatedByHeight",
			Handler:    _TransactionService_GetAnnotatedByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/transaction/transactionpb/transaction.proto",
}
