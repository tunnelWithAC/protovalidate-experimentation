// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: example/hello/v1/hello.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PurchaseDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	DeliveryDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	Price        string                 `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_hello_v1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_example_hello_v1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_example_hello_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetPurchaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PurchaseDate
	}
	return nil
}

func (x *Transaction) GetDeliveryDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveryDate
	}
	return nil
}

func (x *Transaction) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

var File_example_hello_v1_hello_proto protoreflect.FileDescriptor

var file_example_hello_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x6d, 0x79, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0x48, 0x05, 0x32, 0x03, 0x20, 0xe7, 0x07, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x12, 0xc1, 0x01, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0xaa, 0x01, 0xba, 0x48, 0xa6, 0x01, 0xba, 0x01, 0xa2, 0x01, 0x0a,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x70, 0x72, 0x69, 0x63, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x20, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x20, 0x28,
	0x24, 0x20, 0x6f, 0x72, 0x20, 0xc2, 0xa3, 0x29, 0x1a, 0x47, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x77, 0x69, 0x74, 0x68, 0x28, 0x27, 0x24, 0x27, 0x29, 0x20,
	0x6f, 0x72, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x77, 0x69,
	0x74, 0x68, 0x28, 0x27, 0xc2, 0xa3, 0x27, 0x29, 0x29, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73, 0x5b, 0x31, 0x3a, 0x5d, 0x29, 0x20, 0x3e, 0x20,
	0x30, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x3a, 0x74, 0xba, 0x48, 0x71, 0x1a, 0x6f, 0x0a,
	0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x20, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x20, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x27, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x20, 0x3e, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x8e,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x42, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x64,
	0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x50, 0x58, 0xaa, 0x02,
	0x0a, 0x4d, 0x79, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xca, 0x02, 0x0a, 0x4d, 0x79,
	0x5c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xe2, 0x02, 0x16, 0x4d, 0x79, 0x5c, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0b, 0x4d, 0x79, 0x3a, 0x3a, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_hello_v1_hello_proto_rawDescOnce sync.Once
	file_example_hello_v1_hello_proto_rawDescData = file_example_hello_v1_hello_proto_rawDesc
)

func file_example_hello_v1_hello_proto_rawDescGZIP() []byte {
	file_example_hello_v1_hello_proto_rawDescOnce.Do(func() {
		file_example_hello_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_hello_v1_hello_proto_rawDescData)
	})
	return file_example_hello_v1_hello_proto_rawDescData
}

var file_example_hello_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_hello_v1_hello_proto_goTypes = []interface{}{
	(*Transaction)(nil),           // 0: my.package.Transaction
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_example_hello_v1_hello_proto_depIdxs = []int32{
	1, // 0: my.package.Transaction.purchase_date:type_name -> google.protobuf.Timestamp
	1, // 1: my.package.Transaction.delivery_date:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_hello_v1_hello_proto_init() }
func file_example_hello_v1_hello_proto_init() {
	if File_example_hello_v1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_hello_v1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_example_hello_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_hello_v1_hello_proto_goTypes,
		DependencyIndexes: file_example_hello_v1_hello_proto_depIdxs,
		MessageInfos:      file_example_hello_v1_hello_proto_msgTypes,
	}.Build()
	File_example_hello_v1_hello_proto = out.File
	file_example_hello_v1_hello_proto_rawDesc = nil
	file_example_hello_v1_hello_proto_goTypes = nil
	file_example_hello_v1_hello_proto_depIdxs = nil
}
