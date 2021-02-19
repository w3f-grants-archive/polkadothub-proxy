// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: grpc/height/heightpb/height.proto

package heightpb

import (
	context "context"
	blockpb "github.com/figment-networks/polkadothub-proxy/grpc/block/blockpb"
	chainpb "github.com/figment-networks/polkadothub-proxy/grpc/chain/chainpb"
	eventpb "github.com/figment-networks/polkadothub-proxy/grpc/event/eventpb"
	stakingpb "github.com/figment-networks/polkadothub-proxy/grpc/staking/stakingpb"
	transactionpb "github.com/figment-networks/polkadothub-proxy/grpc/transaction/transactionpb"
	validatorperformancepb "github.com/figment-networks/polkadothub-proxy/grpc/validatorperformance/validatorperformancepb"
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

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_height_heightpb_height_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_height_heightpb_height_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_grpc_height_heightpb_height_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain                *chainpb.GetMetaByHeightResponse            `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Block                *blockpb.GetByHeightResponse                `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Event                *eventpb.GetByHeightResponse                `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Staking              *stakingpb.GetByHeightResponse              `protobuf:"bytes,3,opt,name=staking,proto3" json:"staking,omitempty"`
	Transaction          *transactionpb.GetAnnotatedByHeightResponse `protobuf:"bytes,5,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ValidatorPerformance *validatorperformancepb.GetByHeightResponse `protobuf:"bytes,6,opt,name=validator_performance,json=validatorPerformance,proto3" json:"validator_performance,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_height_heightpb_height_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_height_heightpb_height_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_grpc_height_heightpb_height_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetChain() *chainpb.GetMetaByHeightResponse {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *GetAllResponse) GetBlock() *blockpb.GetByHeightResponse {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *GetAllResponse) GetEvent() *eventpb.GetByHeightResponse {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *GetAllResponse) GetStaking() *stakingpb.GetByHeightResponse {
	if x != nil {
		return x.Staking
	}
	return nil
}

func (x *GetAllResponse) GetTransaction() *transactionpb.GetAnnotatedByHeightResponse {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *GetAllResponse) GetValidatorPerformance() *validatorperformancepb.GetByHeightResponse {
	if x != nil {
		return x.ValidatorPerformance
	}
	return nil
}

var File_grpc_height_heightpb_height_proto protoreflect.FileDescriptor

var file_grpc_height_heightpb_height_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x1e, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x4c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x70, 0x62, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x30, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x8f, 0x03, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x79, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x4b, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x15,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x4a, 0x0a, 0x0d,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74,
	0x68, 0x75, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_height_heightpb_height_proto_rawDescOnce sync.Once
	file_grpc_height_heightpb_height_proto_rawDescData = file_grpc_height_heightpb_height_proto_rawDesc
)

func file_grpc_height_heightpb_height_proto_rawDescGZIP() []byte {
	file_grpc_height_heightpb_height_proto_rawDescOnce.Do(func() {
		file_grpc_height_heightpb_height_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_height_heightpb_height_proto_rawDescData)
	})
	return file_grpc_height_heightpb_height_proto_rawDescData
}

var file_grpc_height_heightpb_height_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_height_heightpb_height_proto_goTypes = []interface{}{
	(*GetAllRequest)(nil),                              // 0: height.GetAllRequest
	(*GetAllResponse)(nil),                             // 1: height.GetAllResponse
	(*chainpb.GetMetaByHeightResponse)(nil),            // 2: chain.GetMetaByHeightResponse
	(*blockpb.GetByHeightResponse)(nil),                // 3: block.GetByHeightResponse
	(*eventpb.GetByHeightResponse)(nil),                // 4: event.GetByHeightResponse
	(*stakingpb.GetByHeightResponse)(nil),              // 5: staking.GetByHeightResponse
	(*transactionpb.GetAnnotatedByHeightResponse)(nil), // 6: transaction.GetAnnotatedByHeightResponse
	(*validatorperformancepb.GetByHeightResponse)(nil), // 7: validatorPerformance.GetByHeightResponse
}
var file_grpc_height_heightpb_height_proto_depIdxs = []int32{
	2, // 0: height.GetAllResponse.chain:type_name -> chain.GetMetaByHeightResponse
	3, // 1: height.GetAllResponse.block:type_name -> block.GetByHeightResponse
	4, // 2: height.GetAllResponse.event:type_name -> event.GetByHeightResponse
	5, // 3: height.GetAllResponse.staking:type_name -> staking.GetByHeightResponse
	6, // 4: height.GetAllResponse.transaction:type_name -> transaction.GetAnnotatedByHeightResponse
	7, // 5: height.GetAllResponse.validator_performance:type_name -> validatorPerformance.GetByHeightResponse
	0, // 6: height.HeightService.GetAll:input_type -> height.GetAllRequest
	1, // 7: height.HeightService.GetAll:output_type -> height.GetAllResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_grpc_height_heightpb_height_proto_init() }
func file_grpc_height_heightpb_height_proto_init() {
	if File_grpc_height_heightpb_height_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_height_heightpb_height_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_grpc_height_heightpb_height_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
			RawDescriptor: file_grpc_height_heightpb_height_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_height_heightpb_height_proto_goTypes,
		DependencyIndexes: file_grpc_height_heightpb_height_proto_depIdxs,
		MessageInfos:      file_grpc_height_heightpb_height_proto_msgTypes,
	}.Build()
	File_grpc_height_heightpb_height_proto = out.File
	file_grpc_height_heightpb_height_proto_rawDesc = nil
	file_grpc_height_heightpb_height_proto_goTypes = nil
	file_grpc_height_heightpb_height_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HeightServiceClient is the client API for HeightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeightServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type heightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeightServiceClient(cc grpc.ClientConnInterface) HeightServiceClient {
	return &heightServiceClient{cc}
}

func (c *heightServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/height.HeightService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeightServiceServer is the server API for HeightService service.
type HeightServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
}

// UnimplementedHeightServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHeightServiceServer struct {
}

func (*UnimplementedHeightServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterHeightServiceServer(s *grpc.Server, srv HeightServiceServer) {
	s.RegisterService(&_HeightService_serviceDesc, srv)
}

func _HeightService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeightServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/height.HeightService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeightServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeightService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "height.HeightService",
	HandlerType: (*HeightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _HeightService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/height/heightpb/height.proto",
}
