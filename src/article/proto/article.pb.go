// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: src/proto/article.proto

package articlepb

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

type ArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ArticleRequest) Reset() {
	*x = ArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleRequest) ProtoMessage() {}

func (x *ArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleRequest.ProtoReflect.Descriptor instead.
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_article_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ArticleResponse) Reset() {
	*x = ArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResponse) ProtoMessage() {}

func (x *ArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleResponse.ProtoReflect.Descriptor instead.
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_src_proto_article_proto protoreflect.FileDescriptor

var file_src_proto_article_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x24, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x4f, 0x0a, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x50, 0x49,
	0x12, 0x41, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x16, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_article_proto_rawDescOnce sync.Once
	file_src_proto_article_proto_rawDescData = file_src_proto_article_proto_rawDesc
)

func file_src_proto_article_proto_rawDescGZIP() []byte {
	file_src_proto_article_proto_rawDescOnce.Do(func() {
		file_src_proto_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_article_proto_rawDescData)
	})
	return file_src_proto_article_proto_rawDescData
}

var file_src_proto_article_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_proto_article_proto_goTypes = []interface{}{
	(*ArticleRequest)(nil),  // 0: domain.ArticleRequest
	(*ArticleResponse)(nil), // 1: domain.ArticleResponse
}
var file_src_proto_article_proto_depIdxs = []int32{
	0, // 0: domain.ArticleAPI.StoreArticle:input_type -> domain.ArticleRequest
	1, // 1: domain.ArticleAPI.StoreArticle:output_type -> domain.ArticleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_article_proto_init() }
func file_src_proto_article_proto_init() {
	if File_src_proto_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleRequest); i {
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
		file_src_proto_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleResponse); i {
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
			RawDescriptor: file_src_proto_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_article_proto_goTypes,
		DependencyIndexes: file_src_proto_article_proto_depIdxs,
		MessageInfos:      file_src_proto_article_proto_msgTypes,
	}.Build()
	File_src_proto_article_proto = out.File
	file_src_proto_article_proto_rawDesc = nil
	file_src_proto_article_proto_goTypes = nil
	file_src_proto_article_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArticleAPIClient is the client API for ArticleAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleAPIClient interface {
	StoreArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error)
}

type articleAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleAPIClient(cc grpc.ClientConnInterface) ArticleAPIClient {
	return &articleAPIClient{cc}
}

func (c *articleAPIClient) StoreArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/domain.ArticleAPI/StoreArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleAPIServer is the server API for ArticleAPI service.
type ArticleAPIServer interface {
	StoreArticle(context.Context, *ArticleRequest) (*ArticleResponse, error)
}

// UnimplementedArticleAPIServer can be embedded to have forward compatible implementations.
type UnimplementedArticleAPIServer struct {
}

func (*UnimplementedArticleAPIServer) StoreArticle(context.Context, *ArticleRequest) (*ArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreArticle not implemented")
}

func RegisterArticleAPIServer(s *grpc.Server, srv ArticleAPIServer) {
	s.RegisterService(&_ArticleAPI_serviceDesc, srv)
}

func _ArticleAPI_StoreArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleAPIServer).StoreArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.ArticleAPI/StoreArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleAPIServer).StoreArticle(ctx, req.(*ArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "domain.ArticleAPI",
	HandlerType: (*ArticleAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreArticle",
			Handler:    _ArticleAPI_StoreArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/article.proto",
}
