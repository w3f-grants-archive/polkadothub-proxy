// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: grpc/validator/validatorpb/validator.proto

package validatorpb

import (
	context "context"
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

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance    string `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Commission string `protobuf:"bytes,3,opt,name=commission,proto3" json:"commission,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_validator_validatorpb_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_validator_validatorpb_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_grpc_validator_validatorpb_validator_proto_rawDescGZIP(), []int{0}
}

func (x *Validator) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Validator) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *Validator) GetCommission() string {
	if x != nil {
		return x.Commission
	}
	return ""
}

type GetAllByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetAllByHeightRequest) Reset() {
	*x = GetAllByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_validator_validatorpb_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByHeightRequest) ProtoMessage() {}

func (x *GetAllByHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_validator_validatorpb_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByHeightRequest.ProtoReflect.Descriptor instead.
func (*GetAllByHeightRequest) Descriptor() ([]byte, []int) {
	return file_grpc_validator_validatorpb_validator_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllByHeightRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetAllByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *GetAllByHeightResponse) Reset() {
	*x = GetAllByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_validator_validatorpb_validator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByHeightResponse) ProtoMessage() {}

func (x *GetAllByHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_validator_validatorpb_validator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByHeightResponse.ProtoReflect.Descriptor instead.
func (*GetAllByHeightResponse) Descriptor() ([]byte, []int) {
	return file_grpc_validator_validatorpb_validator_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllByHeightResponse) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

var File_grpc_validator_validatorpb_validator_proto protoreflect.FileDescriptor

var file_grpc_validator_validatorpb_validator_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x32, 0x6b, 0x0a, 0x10, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x20, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74, 0x68, 0x75,
	0x62, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_validator_validatorpb_validator_proto_rawDescOnce sync.Once
	file_grpc_validator_validatorpb_validator_proto_rawDescData = file_grpc_validator_validatorpb_validator_proto_rawDesc
)

func file_grpc_validator_validatorpb_validator_proto_rawDescGZIP() []byte {
	file_grpc_validator_validatorpb_validator_proto_rawDescOnce.Do(func() {
		file_grpc_validator_validatorpb_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_validator_validatorpb_validator_proto_rawDescData)
	})
	return file_grpc_validator_validatorpb_validator_proto_rawDescData
}

var file_grpc_validator_validatorpb_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_validator_validatorpb_validator_proto_goTypes = []interface{}{
	(*Validator)(nil),              // 0: validator.Validator
	(*GetAllByHeightRequest)(nil),  // 1: validator.GetAllByHeightRequest
	(*GetAllByHeightResponse)(nil), // 2: validator.GetAllByHeightResponse
}
var file_grpc_validator_validatorpb_validator_proto_depIdxs = []int32{
	0, // 0: validator.GetAllByHeightResponse.validators:type_name -> validator.Validator
	1, // 1: validator.ValidatorService.GetAllByHeight:input_type -> validator.GetAllByHeightRequest
	2, // 2: validator.ValidatorService.GetAllByHeight:output_type -> validator.GetAllByHeightResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_validator_validatorpb_validator_proto_init() }
func file_grpc_validator_validatorpb_validator_proto_init() {
	if File_grpc_validator_validatorpb_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_validator_validatorpb_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
		file_grpc_validator_validatorpb_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByHeightRequest); i {
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
		file_grpc_validator_validatorpb_validator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByHeightResponse); i {
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
			RawDescriptor: file_grpc_validator_validatorpb_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_validator_validatorpb_validator_proto_goTypes,
		DependencyIndexes: file_grpc_validator_validatorpb_validator_proto_depIdxs,
		MessageInfos:      file_grpc_validator_validatorpb_validator_proto_msgTypes,
	}.Build()
	File_grpc_validator_validatorpb_validator_proto = out.File
	file_grpc_validator_validatorpb_validator_proto_rawDesc = nil
	file_grpc_validator_validatorpb_validator_proto_goTypes = nil
	file_grpc_validator_validatorpb_validator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ValidatorServiceClient is the client API for ValidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidatorServiceClient interface {
	GetAllByHeight(ctx context.Context, in *GetAllByHeightRequest, opts ...grpc.CallOption) (*GetAllByHeightResponse, error)
}

type validatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorServiceClient(cc grpc.ClientConnInterface) ValidatorServiceClient {
	return &validatorServiceClient{cc}
}

func (c *validatorServiceClient) GetAllByHeight(ctx context.Context, in *GetAllByHeightRequest, opts ...grpc.CallOption) (*GetAllByHeightResponse, error) {
	out := new(GetAllByHeightResponse)
	err := c.cc.Invoke(ctx, "/validator.ValidatorService/GetAllByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorServiceServer is the server API for ValidatorService service.
type ValidatorServiceServer interface {
	GetAllByHeight(context.Context, *GetAllByHeightRequest) (*GetAllByHeightResponse, error)
}

// UnimplementedValidatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValidatorServiceServer struct {
}

func (*UnimplementedValidatorServiceServer) GetAllByHeight(context.Context, *GetAllByHeightRequest) (*GetAllByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByHeight not implemented")
}

func RegisterValidatorServiceServer(s *grpc.Server, srv ValidatorServiceServer) {
	s.RegisterService(&_ValidatorService_serviceDesc, srv)
}

func _ValidatorService_GetAllByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).GetAllByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validator.ValidatorService/GetAllByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).GetAllByHeight(ctx, req.(*GetAllByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValidatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "validator.ValidatorService",
	HandlerType: (*ValidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllByHeight",
			Handler:    _ValidatorService_GetAllByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/validator/validatorpb/validator.proto",
}
