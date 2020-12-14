// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: narinfo.proto

package schema

import (
	proto "github.com/golang/protobuf/proto"
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

type NarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorePath   *string  `protobuf:"bytes,1,req,name=StorePath" json:"StorePath,omitempty"`
	URL         *string  `protobuf:"bytes,2,req,name=URL" json:"URL,omitempty"`
	Compression *string  `protobuf:"bytes,3,req,name=Compression" json:"Compression,omitempty"`
	FileHash    *string  `protobuf:"bytes,4,req,name=FileHash" json:"FileHash,omitempty"`
	FileSize    *uint64  `protobuf:"varint,5,req,name=FileSize" json:"FileSize,omitempty"`
	NarHash     *string  `protobuf:"bytes,6,req,name=NarHash" json:"NarHash,omitempty"`
	NarSize     *uint64  `protobuf:"varint,7,req,name=NarSize" json:"NarSize,omitempty"`
	References  []string `protobuf:"bytes,8,rep,name=References" json:"References,omitempty"`
}

func (x *NarInfo) Reset() {
	*x = NarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_narinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NarInfo) ProtoMessage() {}

func (x *NarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_narinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NarInfo.ProtoReflect.Descriptor instead.
func (*NarInfo) Descriptor() ([]byte, []int) {
	return file_narinfo_proto_rawDescGZIP(), []int{0}
}

func (x *NarInfo) GetStorePath() string {
	if x != nil && x.StorePath != nil {
		return *x.StorePath
	}
	return ""
}

func (x *NarInfo) GetURL() string {
	if x != nil && x.URL != nil {
		return *x.URL
	}
	return ""
}

func (x *NarInfo) GetCompression() string {
	if x != nil && x.Compression != nil {
		return *x.Compression
	}
	return ""
}

func (x *NarInfo) GetFileHash() string {
	if x != nil && x.FileHash != nil {
		return *x.FileHash
	}
	return ""
}

func (x *NarInfo) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *NarInfo) GetNarHash() string {
	if x != nil && x.NarHash != nil {
		return *x.NarHash
	}
	return ""
}

func (x *NarInfo) GetNarSize() uint64 {
	if x != nil && x.NarSize != nil {
		return *x.NarSize
	}
	return 0
}

func (x *NarInfo) GetReferences() []string {
	if x != nil {
		return x.References
	}
	return nil
}

var File_narinfo_proto protoreflect.FileDescriptor

var file_narinfo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x61, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe7, 0x01, 0x0a, 0x07, 0x4e, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x04, 0x52, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x61, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x06, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x61, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x4e, 0x61, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x07, 0x4e, 0x61, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x65, 0x61, 0x67, 0x2f, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x6e, 0x69,
	0x78, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
}

var (
	file_narinfo_proto_rawDescOnce sync.Once
	file_narinfo_proto_rawDescData = file_narinfo_proto_rawDesc
)

func file_narinfo_proto_rawDescGZIP() []byte {
	file_narinfo_proto_rawDescOnce.Do(func() {
		file_narinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_narinfo_proto_rawDescData)
	})
	return file_narinfo_proto_rawDescData
}

var file_narinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_narinfo_proto_goTypes = []interface{}{
	(*NarInfo)(nil), // 0: NarInfo
}
var file_narinfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_narinfo_proto_init() }
func file_narinfo_proto_init() {
	if File_narinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_narinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NarInfo); i {
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
			RawDescriptor: file_narinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_narinfo_proto_goTypes,
		DependencyIndexes: file_narinfo_proto_depIdxs,
		MessageInfos:      file_narinfo_proto_msgTypes,
	}.Build()
	File_narinfo_proto = out.File
	file_narinfo_proto_rawDesc = nil
	file_narinfo_proto_goTypes = nil
	file_narinfo_proto_depIdxs = nil
}
