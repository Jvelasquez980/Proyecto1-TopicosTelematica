// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/mom.proto

package mom

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MomSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reprocessed   int32                  `protobuf:"varint,1,opt,name=reprocessed,proto3" json:"reprocessed,omitempty"`
	Failed        int32                  `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MomSummary) Reset() {
	*x = MomSummary{}
	mi := &file_proto_mom_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MomSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MomSummary) ProtoMessage() {}

func (x *MomSummary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mom_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MomSummary.ProtoReflect.Descriptor instead.
func (*MomSummary) Descriptor() ([]byte, []int) {
	return file_proto_mom_proto_rawDescGZIP(), []int{0}
}

func (x *MomSummary) GetReprocessed() int32 {
	if x != nil {
		return x.Reprocessed
	}
	return 0
}

func (x *MomSummary) GetFailed() int32 {
	if x != nil {
		return x.Failed
	}
	return 0
}

var File_proto_mom_proto protoreflect.FileDescriptor

const file_proto_mom_proto_rawDesc = "" +
	"\n" +
	"\x0fproto/mom.proto\x12\x03mom\x1a\x1bgoogle/protobuf/empty.proto\"F\n" +
	"\n" +
	"MomSummary\x12 \n" +
	"\vreprocessed\x18\x01 \x01(\x05R\vreprocessed\x12\x16\n" +
	"\x06failed\x18\x02 \x01(\x05R\x06failed2J\n" +
	"\n" +
	"MomService\x12<\n" +
	"\x11ProcessAllPending\x12\x16.google.protobuf.Empty\x1a\x0f.mom.MomSummaryB\fZ\n" +
	"/proto;momb\x06proto3"

var (
	file_proto_mom_proto_rawDescOnce sync.Once
	file_proto_mom_proto_rawDescData []byte
)

func file_proto_mom_proto_rawDescGZIP() []byte {
	file_proto_mom_proto_rawDescOnce.Do(func() {
		file_proto_mom_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_mom_proto_rawDesc), len(file_proto_mom_proto_rawDesc)))
	})
	return file_proto_mom_proto_rawDescData
}

var file_proto_mom_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_mom_proto_goTypes = []any{
	(*MomSummary)(nil),    // 0: mom.MomSummary
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_proto_mom_proto_depIdxs = []int32{
	1, // 0: mom.MomService.ProcessAllPending:input_type -> google.protobuf.Empty
	0, // 1: mom.MomService.ProcessAllPending:output_type -> mom.MomSummary
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mom_proto_init() }
func file_proto_mom_proto_init() {
	if File_proto_mom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_mom_proto_rawDesc), len(file_proto_mom_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mom_proto_goTypes,
		DependencyIndexes: file_proto_mom_proto_depIdxs,
		MessageInfos:      file_proto_mom_proto_msgTypes,
	}.Build()
	File_proto_mom_proto = out.File
	file_proto_mom_proto_goTypes = nil
	file_proto_mom_proto_depIdxs = nil
}
