// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: shortener.proto

package shortener

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

type GetShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetShortUrlRequest) Reset() {
	*x = GetShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortUrlRequest) ProtoMessage() {}

func (x *GetShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortUrlRequest.ProtoReflect.Descriptor instead.
func (*GetShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *GetShortUrlRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type GetShortUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetShortUrlResponse) Reset() {
	*x = GetShortUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortUrlResponse) ProtoMessage() {}

func (x *GetShortUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortUrlResponse.ProtoReflect.Descriptor instead.
func (*GetShortUrlResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetShortUrlResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOriginalUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetOriginalUrlRequest) Reset() {
	*x = GetOriginalUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalUrlRequest) ProtoMessage() {}

func (x *GetOriginalUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalUrlRequest.ProtoReflect.Descriptor instead.
func (*GetOriginalUrlRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetOriginalUrlRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOriginalUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetOriginalUrlResponse) Reset() {
	*x = GetOriginalUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalUrlResponse) ProtoMessage() {}

func (x *GetOriginalUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalUrlResponse.ProtoReflect.Descriptor instead.
func (*GetOriginalUrlResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetOriginalUrlResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_shortener_proto protoreflect.FileDescriptor

var file_shortener_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x22, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x34, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c,
	0x32, 0xc0, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x54,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x6d, 0x69, 0x72, 0x61, 0x2d, 0x67, 0x61, 0x6c, 0x65, 0x65, 0x76, 0x61,
	0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_proto_rawDescOnce sync.Once
	file_shortener_proto_rawDescData = file_shortener_proto_rawDesc
)

func file_shortener_proto_rawDescGZIP() []byte {
	file_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_proto_rawDescData)
	})
	return file_shortener_proto_rawDescData
}

var file_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortener_proto_goTypes = []interface{}{
	(*GetShortUrlRequest)(nil),     // 0: protos.shortener.GetShortUrlRequest
	(*GetShortUrlResponse)(nil),    // 1: protos.shortener.GetShortUrlResponse
	(*GetOriginalUrlRequest)(nil),  // 2: protos.shortener.GetOriginalUrlRequest
	(*GetOriginalUrlResponse)(nil), // 3: protos.shortener.GetOriginalUrlResponse
}
var file_shortener_proto_depIdxs = []int32{
	0, // 0: protos.shortener.Shortener.GetShortUrl:input_type -> protos.shortener.GetShortUrlRequest
	2, // 1: protos.shortener.Shortener.GetOriginalUrl:input_type -> protos.shortener.GetOriginalUrlRequest
	1, // 2: protos.shortener.Shortener.GetShortUrl:output_type -> protos.shortener.GetShortUrlResponse
	3, // 3: protos.shortener.Shortener.GetOriginalUrl:output_type -> protos.shortener.GetOriginalUrlResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortener_proto_init() }
func file_shortener_proto_init() {
	if File_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortUrlRequest); i {
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
		file_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortUrlResponse); i {
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
		file_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalUrlRequest); i {
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
		file_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalUrlResponse); i {
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
			RawDescriptor: file_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_proto_msgTypes,
	}.Build()
	File_shortener_proto = out.File
	file_shortener_proto_rawDesc = nil
	file_shortener_proto_goTypes = nil
	file_shortener_proto_depIdxs = nil
}