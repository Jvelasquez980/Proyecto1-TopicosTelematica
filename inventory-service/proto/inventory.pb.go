// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/inventory.proto

package inventory

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_proto_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type InventoryCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Item                `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryCheckRequest) Reset() {
	*x = InventoryCheckRequest{}
	mi := &file_proto_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryCheckRequest) ProtoMessage() {}

func (x *InventoryCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryCheckRequest.ProtoReflect.Descriptor instead.
func (*InventoryCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *InventoryCheckRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type InventoryCheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Available     bool                   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryCheckResponse) Reset() {
	*x = InventoryCheckResponse{}
	mi := &file_proto_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryCheckResponse) ProtoMessage() {}

func (x *InventoryCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryCheckResponse.ProtoReflect.Descriptor instead.
func (*InventoryCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *InventoryCheckResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

var File_proto_inventory_proto protoreflect.FileDescriptor

const file_proto_inventory_proto_rawDesc = "" +
	"\n" +
	"\x15proto/inventory.proto\x12\tinventory\"A\n" +
	"\x04Item\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\">\n" +
	"\x15InventoryCheckRequest\x12%\n" +
	"\x05items\x18\x01 \x03(\v2\x0f.inventory.ItemR\x05items\"6\n" +
	"\x16InventoryCheckResponse\x12\x1c\n" +
	"\tavailable\x18\x01 \x01(\bR\tavailable2i\n" +
	"\x10InventoryService\x12U\n" +
	"\x0eCheckInventory\x12 .inventory.InventoryCheckRequest\x1a!.inventory.InventoryCheckResponseB\x12Z\x10/proto;inventoryb\x06proto3"

var (
	file_proto_inventory_proto_rawDescOnce sync.Once
	file_proto_inventory_proto_rawDescData []byte
)

func file_proto_inventory_proto_rawDescGZIP() []byte {
	file_proto_inventory_proto_rawDescOnce.Do(func() {
		file_proto_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)))
	})
	return file_proto_inventory_proto_rawDescData
}

var file_proto_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_inventory_proto_goTypes = []any{
	(*Item)(nil),                   // 0: inventory.Item
	(*InventoryCheckRequest)(nil),  // 1: inventory.InventoryCheckRequest
	(*InventoryCheckResponse)(nil), // 2: inventory.InventoryCheckResponse
}
var file_proto_inventory_proto_depIdxs = []int32{
	0, // 0: inventory.InventoryCheckRequest.items:type_name -> inventory.Item
	1, // 1: inventory.InventoryService.CheckInventory:input_type -> inventory.InventoryCheckRequest
	2, // 2: inventory.InventoryService.CheckInventory:output_type -> inventory.InventoryCheckResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_inventory_proto_init() }
func file_proto_inventory_proto_init() {
	if File_proto_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inventory_proto_goTypes,
		DependencyIndexes: file_proto_inventory_proto_depIdxs,
		MessageInfos:      file_proto_inventory_proto_msgTypes,
	}.Build()
	File_proto_inventory_proto = out.File
	file_proto_inventory_proto_goTypes = nil
	file_proto_inventory_proto_depIdxs = nil
}
