// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: grpc/event/eventpb/event.proto

package eventpb

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

type EventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EventData) Reset() {
	*x = EventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_eventpb_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventData) ProtoMessage() {}

func (x *EventData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_eventpb_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventData.ProtoReflect.Descriptor instead.
func (*EventData) Descriptor() ([]byte, []int) {
	return file_grpc_event_eventpb_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EventData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index          int64        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Data           []*EventData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Phase          string       `protobuf:"bytes,3,opt,name=phase,proto3" json:"phase,omitempty"`
	Method         string       `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Section        string       `protobuf:"bytes,5,opt,name=section,proto3" json:"section,omitempty"`
	Description    string       `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	ExtrinsicIndex int64        `protobuf:"varint,7,opt,name=extrinsicIndex,proto3" json:"extrinsicIndex,omitempty"`
	Error          string       `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
	Raw            string       `protobuf:"bytes,9,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_eventpb_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_eventpb_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_grpc_event_eventpb_event_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Event) GetData() []*EventData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *Event) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Event) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetExtrinsicIndex() int64 {
	if x != nil {
		return x.ExtrinsicIndex
	}
	return 0
}

func (x *Event) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Event) GetRaw() string {
	if x != nil {
		return x.Raw
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
		mi := &file_grpc_event_eventpb_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHeightRequest) ProtoMessage() {}

func (x *GetByHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_eventpb_event_proto_msgTypes[2]
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
	return file_grpc_event_eventpb_event_proto_rawDescGZIP(), []int{2}
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

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetByHeightResponse) Reset() {
	*x = GetByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_eventpb_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHeightResponse) ProtoMessage() {}

func (x *GetByHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_eventpb_event_proto_msgTypes[3]
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
	return file_grpc_event_eventpb_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetByHeightResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_grpc_event_eventpb_event_proto protoreflect.FileDescriptor

var file_grpc_event_eventpb_event_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xfd,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x61, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x22, 0x2c,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3b, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x56, 0x0a, 0x0c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x2f, 0x70, 0x6f, 0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74, 0x68, 0x75, 0x62, 0x2d, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_event_eventpb_event_proto_rawDescOnce sync.Once
	file_grpc_event_eventpb_event_proto_rawDescData = file_grpc_event_eventpb_event_proto_rawDesc
)

func file_grpc_event_eventpb_event_proto_rawDescGZIP() []byte {
	file_grpc_event_eventpb_event_proto_rawDescOnce.Do(func() {
		file_grpc_event_eventpb_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_event_eventpb_event_proto_rawDescData)
	})
	return file_grpc_event_eventpb_event_proto_rawDescData
}

var file_grpc_event_eventpb_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_event_eventpb_event_proto_goTypes = []interface{}{
	(*EventData)(nil),           // 0: event.EventData
	(*Event)(nil),               // 1: event.Event
	(*GetByHeightRequest)(nil),  // 2: event.GetByHeightRequest
	(*GetByHeightResponse)(nil), // 3: event.GetByHeightResponse
}
var file_grpc_event_eventpb_event_proto_depIdxs = []int32{
	0, // 0: event.Event.data:type_name -> event.EventData
	1, // 1: event.GetByHeightResponse.events:type_name -> event.Event
	2, // 2: event.EventService.GetByHeight:input_type -> event.GetByHeightRequest
	3, // 3: event.EventService.GetByHeight:output_type -> event.GetByHeightResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_event_eventpb_event_proto_init() }
func file_grpc_event_eventpb_event_proto_init() {
	if File_grpc_event_eventpb_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_event_eventpb_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventData); i {
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
		file_grpc_event_eventpb_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_grpc_event_eventpb_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_event_eventpb_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_event_eventpb_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_event_eventpb_event_proto_goTypes,
		DependencyIndexes: file_grpc_event_eventpb_event_proto_depIdxs,
		MessageInfos:      file_grpc_event_eventpb_event_proto_msgTypes,
	}.Build()
	File_grpc_event_eventpb_event_proto = out.File
	file_grpc_event_eventpb_event_proto_rawDesc = nil
	file_grpc_event_eventpb_event_proto_goTypes = nil
	file_grpc_event_eventpb_event_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/event.EventService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _EventService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/event/eventpb/event.proto",
}
