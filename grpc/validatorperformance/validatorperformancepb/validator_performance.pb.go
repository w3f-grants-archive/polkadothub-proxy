// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/validatorperformance/validatorperformancepb/validator_performance.proto

package validatorperformancepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Validator struct {
	StashAccount         string   `protobuf:"bytes,1,opt,name=stash_account,json=stashAccount,proto3" json:"stash_account,omitempty"`
	Online               bool     `protobuf:"varint,2,opt,name=online,proto3" json:"online,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_3278576864b3645d, []int{0}
}

func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetStashAccount() string {
	if m != nil {
		return m.StashAccount
	}
	return ""
}

func (m *Validator) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
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
	return fileDescriptor_3278576864b3645d, []int{1}
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
	Validators           []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetByHeightResponse) Reset()         { *m = GetByHeightResponse{} }
func (m *GetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetByHeightResponse) ProtoMessage()    {}
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3278576864b3645d, []int{2}
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

func (m *GetByHeightResponse) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func init() {
	proto.RegisterType((*Validator)(nil), "validatorPerformance.Validator")
	proto.RegisterType((*GetByHeightRequest)(nil), "validatorPerformance.GetByHeightRequest")
	proto.RegisterType((*GetByHeightResponse)(nil), "validatorPerformance.GetByHeightResponse")
}

func init() {
	proto.RegisterFile("grpc/validatorperformance/validatorperformancepb/validator_performance.proto", fileDescriptor_3278576864b3645d)
}

var fileDescriptor_3278576864b3645d = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x49, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0x48, 0x2d, 0x4a, 0xcb, 0x2f,
	0xca, 0x4d, 0xcc, 0x4b, 0x4e, 0xc5, 0x2a, 0x58, 0x90, 0x84, 0x10, 0x8e, 0x47, 0x12, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0x4b, 0x06, 0x20, 0xe4, 0x94, 0x3c, 0xb8, 0x38, 0xc3,
	0x60, 0xe2, 0x42, 0xca, 0x5c, 0xbc, 0xc5, 0x25, 0x89, 0xc5, 0x19, 0xf1, 0x89, 0xc9, 0xc9, 0xf9,
	0xa5, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x3c, 0x60, 0x41, 0x47, 0x88, 0x98,
	0x90, 0x18, 0x17, 0x5b, 0x7e, 0x5e, 0x4e, 0x66, 0x5e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x94, 0xa7, 0xa4, 0xc3, 0x25, 0xe4, 0x9e, 0x5a, 0xe2, 0x54, 0xe9, 0x91, 0x9a, 0x99, 0x9e,
	0x51, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x0c, 0x56, 0x9d, 0x01, 0x16, 0x00, 0x9b, 0xc5, 0x1c,
	0x04, 0xe5, 0x29, 0x85, 0x71, 0x09, 0xa3, 0xa8, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2,
	0xe7, 0xe2, 0x82, 0x3b, 0xb3, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5e, 0x0f, 0x9b,
	0xcb, 0xf5, 0xe0, 0xce, 0x0e, 0x42, 0xd2, 0x62, 0xd4, 0xcc, 0xc8, 0x25, 0x1d, 0x86, 0x45, 0x79,
	0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x50, 0x0a, 0x17, 0x37, 0x92, 0xbd, 0x42, 0x1a, 0xd8,
	0xcd, 0xc6, 0xf4, 0x88, 0x94, 0x26, 0x11, 0x2a, 0x21, 0x9e, 0x50, 0x62, 0x70, 0x32, 0x8a, 0x32,
	0x20, 0x35, 0xee, 0x92, 0xd8, 0xc0, 0xd1, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xde, 0x8f,
	0x8e, 0xb0, 0xf6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ValidatorPerformanceServiceClient is the client API for ValidatorPerformanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidatorPerformanceServiceClient interface {
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
}

type validatorPerformanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorPerformanceServiceClient(cc grpc.ClientConnInterface) ValidatorPerformanceServiceClient {
	return &validatorPerformanceServiceClient{cc}
}

func (c *validatorPerformanceServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/validatorPerformance.ValidatorPerformanceService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorPerformanceServiceServer is the server API for ValidatorPerformanceService service.
type ValidatorPerformanceServiceServer interface {
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
}

// UnimplementedValidatorPerformanceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValidatorPerformanceServiceServer struct {
}

func (*UnimplementedValidatorPerformanceServiceServer) GetByHeight(ctx context.Context, req *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterValidatorPerformanceServiceServer(s *grpc.Server, srv ValidatorPerformanceServiceServer) {
	s.RegisterService(&_ValidatorPerformanceService_serviceDesc, srv)
}

func _ValidatorPerformanceService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorPerformanceServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validatorPerformance.ValidatorPerformanceService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorPerformanceServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValidatorPerformanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "validatorPerformance.ValidatorPerformanceService",
	HandlerType: (*ValidatorPerformanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _ValidatorPerformanceService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/validatorperformance/validatorperformancepb/validator_performance.proto",
}