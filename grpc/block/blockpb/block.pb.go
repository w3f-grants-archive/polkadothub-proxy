// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/block/blockpb/block.proto

package blockpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Header struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	BlockHash            string               `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	ParentHash           string               `protobuf:"bytes,3,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Height               int64                `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	StateRoot            string               `protobuf:"bytes,5,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	ExtrinsicsRoot       string               `protobuf:"bytes,6,opt,name=extrinsics_root,json=extrinsicsRoot,proto3" json:"extrinsics_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Header) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *Header) GetParentHash() string {
	if m != nil {
		return m.ParentHash
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *Header) GetExtrinsicsRoot() string {
	if m != nil {
		return m.ExtrinsicsRoot
	}
	return ""
}

type Extrinsic struct {
	ExtrinsicIndex       int64    `protobuf:"varint,1,opt,name=extrinsic_index,json=extrinsicIndex,proto3" json:"extrinsic_index,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	IsSignedTransaction  bool     `protobuf:"varint,4,opt,name=is_signed_transaction,json=isSignedTransaction,proto3" json:"is_signed_transaction,omitempty"`
	Signature            string   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Signer               string   `protobuf:"bytes,6,opt,name=signer,proto3" json:"signer,omitempty"`
	Nonce                int64    `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Method               string   `protobuf:"bytes,8,opt,name=method,proto3" json:"method,omitempty"`
	Section              string   `protobuf:"bytes,9,opt,name=section,proto3" json:"section,omitempty"`
	Args                 string   `protobuf:"bytes,10,opt,name=args,proto3" json:"args,omitempty"`
	IsSuccess            bool     `protobuf:"varint,11,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extrinsic) Reset()         { *m = Extrinsic{} }
func (m *Extrinsic) String() string { return proto.CompactTextString(m) }
func (*Extrinsic) ProtoMessage()    {}
func (*Extrinsic) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{1}
}

func (m *Extrinsic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extrinsic.Unmarshal(m, b)
}
func (m *Extrinsic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extrinsic.Marshal(b, m, deterministic)
}
func (m *Extrinsic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extrinsic.Merge(m, src)
}
func (m *Extrinsic) XXX_Size() int {
	return xxx_messageInfo_Extrinsic.Size(m)
}
func (m *Extrinsic) XXX_DiscardUnknown() {
	xxx_messageInfo_Extrinsic.DiscardUnknown(m)
}

var xxx_messageInfo_Extrinsic proto.InternalMessageInfo

func (m *Extrinsic) GetExtrinsicIndex() int64 {
	if m != nil {
		return m.ExtrinsicIndex
	}
	return 0
}

func (m *Extrinsic) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Extrinsic) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Extrinsic) GetIsSignedTransaction() bool {
	if m != nil {
		return m.IsSignedTransaction
	}
	return false
}

func (m *Extrinsic) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Extrinsic) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *Extrinsic) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Extrinsic) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Extrinsic) GetSection() string {
	if m != nil {
		return m.Section
	}
	return ""
}

func (m *Extrinsic) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *Extrinsic) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

type Block struct {
	Header               *Header      `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Extrinsics           []*Extrinsic `protobuf:"bytes,3,rep,name=extrinsics,proto3" json:"extrinsics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{2}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetExtrinsics() []*Extrinsic {
	if m != nil {
		return m.Extrinsics
	}
	return nil
}

type GetByHeightRequest struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByHeightRequest) Reset()         { *m = GetByHeightRequest{} }
func (m *GetByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetByHeightRequest) ProtoMessage()    {}
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{3}
}

func (m *GetByHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByHeightRequest.Unmarshal(m, b)
}
func (m *GetByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightRequest.Merge(m, src)
}
func (m *GetByHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetByHeightRequest.Size(m)
}
func (m *GetByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightRequest proto.InternalMessageInfo

func (m *GetByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetByHeightResponse struct {
	Era                  int64    `protobuf:"varint,1,opt,name=era,proto3" json:"era,omitempty"`
	Session              int64    `protobuf:"varint,2,opt,name=session,proto3" json:"session,omitempty"`
	Chain                string   `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	SpecVersion          string   `protobuf:"bytes,4,opt,name=spec_version,json=specVersion,proto3" json:"spec_version,omitempty"`
	BlockHash            string   `protobuf:"bytes,5,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Block                *Block   `protobuf:"bytes,6,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByHeightResponse) Reset()         { *m = GetByHeightResponse{} }
func (m *GetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetByHeightResponse) ProtoMessage()    {}
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{4}
}

func (m *GetByHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByHeightResponse.Unmarshal(m, b)
}
func (m *GetByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByHeightResponse.Marshal(b, m, deterministic)
}
func (m *GetByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightResponse.Merge(m, src)
}
func (m *GetByHeightResponse) XXX_Size() int {
	return xxx_messageInfo_GetByHeightResponse.Size(m)
}
func (m *GetByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightResponse proto.InternalMessageInfo

func (m *GetByHeightResponse) GetEra() int64 {
	if m != nil {
		return m.Era
	}
	return 0
}

func (m *GetByHeightResponse) GetSession() int64 {
	if m != nil {
		return m.Session
	}
	return 0
}

func (m *GetByHeightResponse) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *GetByHeightResponse) GetSpecVersion() string {
	if m != nil {
		return m.SpecVersion
	}
	return ""
}

func (m *GetByHeightResponse) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *GetByHeightResponse) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "block.Header")
	proto.RegisterType((*Extrinsic)(nil), "block.Extrinsic")
	proto.RegisterType((*Block)(nil), "block.Block")
	proto.RegisterType((*GetByHeightRequest)(nil), "block.GetByHeightRequest")
	proto.RegisterType((*GetByHeightResponse)(nil), "block.GetByHeightResponse")
}

func init() { proto.RegisterFile("grpc/block/blockpb/block.proto", fileDescriptor_6fa3e78aae19c108) }

var fileDescriptor_6fa3e78aae19c108 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0xdc, 0x34, 0x69, 0x3d, 0xee, 0x07, 0xd5, 0x16, 0x90, 0x89, 0x80, 0x06, 0x4b, 0x88,
	0x1c, 0xa8, 0x83, 0xc2, 0x1f, 0x40, 0x91, 0x80, 0x72, 0x75, 0xab, 0x1e, 0xb8, 0x84, 0x8d, 0x33,
	0xb5, 0x57, 0x6d, 0xbc, 0x66, 0x77, 0x53, 0xda, 0x1f, 0xc7, 0x81, 0x7f, 0xc1, 0xcf, 0x41, 0x3b,
	0x63, 0x37, 0x6e, 0xcb, 0x25, 0xf1, 0xbc, 0xf7, 0x66, 0xb3, 0xef, 0xcd, 0x38, 0xf0, 0xaa, 0x30,
	0x75, 0x3e, 0x59, 0x5c, 0xea, 0xfc, 0x82, 0x3f, 0xeb, 0x05, 0x7f, 0xa7, 0xb5, 0xd1, 0x4e, 0x8b,
	0x3e, 0x15, 0xc3, 0xc3, 0x42, 0xeb, 0xe2, 0x12, 0x27, 0x04, 0x2e, 0xd6, 0xe7, 0x13, 0xa7, 0x56,
	0x68, 0x9d, 0x5c, 0xd5, 0xac, 0x4b, 0xfe, 0x04, 0x30, 0x38, 0x46, 0xb9, 0x44, 0x23, 0x52, 0xd8,
	0xf6, 0x6c, 0x1c, 0x8c, 0x82, 0x71, 0x34, 0x1d, 0xa6, 0xdc, 0x9a, 0xb6, 0xad, 0xe9, 0x69, 0xdb,
	0x9a, 0x91, 0x4e, 0xbc, 0x04, 0xa0, 0x1f, 0x99, 0x97, 0xd2, 0x96, 0xf1, 0xd6, 0x28, 0x18, 0x87,
	0x59, 0x48, 0xc8, 0xb1, 0xb4, 0xa5, 0x38, 0x84, 0xa8, 0x96, 0x06, 0x2b, 0xc7, 0x7c, 0x8f, 0x78,
	0x60, 0x88, 0x04, 0xcf, 0x60, 0x50, 0xa2, 0x2a, 0x4a, 0x17, 0x6f, 0x8f, 0x82, 0x71, 0x2f, 0x6b,
	0x2a, 0x7f, 0xae, 0x75, 0xd2, 0xe1, 0xdc, 0x68, 0xed, 0xe2, 0x3e, 0x9f, 0x4b, 0x48, 0xa6, 0xb5,
	0x13, 0x6f, 0xe1, 0x31, 0x5e, 0x3b, 0xa3, 0x2a, 0xab, 0x72, 0xcb, 0x9a, 0x01, 0x69, 0x1e, 0x6d,
	0x60, 0x2f, 0x4c, 0x7e, 0x6f, 0x41, 0xf8, 0xa9, 0x85, 0xee, 0xb4, 0xcd, 0x55, 0xb5, 0xc4, 0x6b,
	0x32, 0xda, 0xeb, 0xb4, 0x7d, 0xf5, 0xa8, 0x10, 0xb0, 0xdd, 0x31, 0x44, 0xcf, 0x1e, 0xa3, 0x68,
	0xd8, 0x04, 0xdb, 0x9f, 0xc2, 0x53, 0x65, 0xe7, 0x56, 0x15, 0x15, 0x2e, 0xe7, 0xce, 0xc8, 0xca,
	0xca, 0xdc, 0x29, 0x5d, 0x91, 0x9b, 0xdd, 0xec, 0x40, 0xd9, 0x13, 0xe2, 0x4e, 0x37, 0x94, 0x78,
	0x01, 0xa1, 0x6f, 0x90, 0x6e, 0x6d, 0xf0, 0xd6, 0x59, 0x0b, 0xf8, 0x40, 0xe8, 0x38, 0xd3, 0x18,
	0x6a, 0x2a, 0xf1, 0x04, 0xfa, 0x95, 0xae, 0x72, 0x8c, 0x77, 0xe8, 0xc2, 0x5c, 0x78, 0xf5, 0x0a,
	0x5d, 0xa9, 0x97, 0xf1, 0x2e, 0xab, 0xb9, 0x12, 0x31, 0xec, 0x58, 0xe4, 0x9b, 0x84, 0x44, 0xb4,
	0xa5, 0x77, 0x21, 0x4d, 0x61, 0x63, 0x60, 0x17, 0xfe, 0xd9, 0x87, 0xed, 0x5d, 0xac, 0xf3, 0x1c,
	0xad, 0x8d, 0x23, 0xba, 0x7a, 0xa8, 0xec, 0x09, 0x03, 0xc9, 0x77, 0xe8, 0xcf, 0xfc, 0x44, 0xc5,
	0x1b, 0x3f, 0x2c, 0xbf, 0x26, 0x94, 0x4b, 0x34, 0xfd, 0x3f, 0xe5, 0x6d, 0xe3, 0xdd, 0xc9, 0x1a,
	0x52, 0xbc, 0x07, 0xd8, 0x4c, 0x21, 0xee, 0x8d, 0x7a, 0xe3, 0x68, 0xba, 0xdf, 0x48, 0x6f, 0x67,
	0x91, 0x75, 0x34, 0xc9, 0x3b, 0x10, 0x5f, 0xd0, 0xcd, 0x6e, 0x8e, 0x69, 0xf8, 0x19, 0xfe, 0x58,
	0xa3, 0x75, 0x9d, 0xdd, 0x08, 0xba, 0xbb, 0x91, 0xfc, 0x0a, 0xe0, 0xe0, 0x8e, 0xdc, 0xd6, 0xba,
	0xb2, 0x28, 0xf6, 0xa1, 0x87, 0x46, 0x36, 0x62, 0xff, 0xc8, 0x31, 0x58, 0xeb, 0x63, 0xd8, 0x22,
	0xb4, 0x2d, 0x7d, 0x9c, 0x79, 0x29, 0x55, 0xd5, 0x4c, 0x93, 0x0b, 0xf1, 0x1a, 0xf6, 0x6c, 0x8d,
	0xf9, 0xfc, 0x0a, 0x8d, 0x6d, 0xa7, 0x18, 0x66, 0x91, 0xc7, 0xce, 0x18, 0xba, 0xb7, 0xf0, 0xfd,
	0xfb, 0x0b, 0x9f, 0x00, 0xbf, 0x74, 0x34, 0xbd, 0x68, 0xba, 0xd7, 0xd8, 0xa6, 0xfc, 0x32, 0xa6,
	0xa6, 0x67, 0xb0, 0x47, 0xf5, 0x09, 0x9a, 0x2b, 0x95, 0xa3, 0xf8, 0x0c, 0x51, 0xc7, 0x8e, 0x78,
	0xde, 0xf4, 0x3c, 0x4c, 0x64, 0x38, 0xfc, 0x17, 0xc5, 0xee, 0x93, 0xff, 0x66, 0xb3, 0x6f, 0x1f,
	0x0b, 0xe5, 0xca, 0xf5, 0x22, 0xcd, 0xf5, 0x6a, 0x72, 0xae, 0x8a, 0x15, 0x56, 0xee, 0xa8, 0x42,
	0xf7, 0x53, 0x9b, 0x0b, 0x3b, 0xa9, 0xf5, 0xe5, 0x85, 0x5c, 0x6a, 0x2f, 0x38, 0xaa, 0x8d, 0xbe,
	0xbe, 0x99, 0x3c, 0xfc, 0xfb, 0x58, 0x0c, 0xe8, 0x4d, 0xff, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0x0f, 0x57, 0x42, 0x5b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockServiceClient interface {
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
}

type blockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockServiceClient(cc grpc.ClientConnInterface) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/block.BlockService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServiceServer is the server API for BlockService service.
type BlockServiceServer interface {
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
}

// UnimplementedBlockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (*UnimplementedBlockServiceServer) GetByHeight(ctx context.Context, req *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterBlockServiceServer(s *grpc.Server, srv BlockServiceServer) {
	s.RegisterService(&_BlockService_serviceDesc, srv)
}

func _BlockService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block.BlockService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "block.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _BlockService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/block/blockpb/block.proto",
}
