// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: google_value.proto

package protoc_gen_connect_openapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GoogleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arg        *structpb.Value     `protobuf:"bytes,1,opt,name=arg,proto3" json:"arg,omitempty"`
	SomeList   *structpb.ListValue `protobuf:"bytes,2,opt,name=some_list,json=someList,proto3" json:"some_list,omitempty"`
	SomeStruct *structpb.Struct    `protobuf:"bytes,3,opt,name=some_struct,json=someStruct,proto3" json:"some_struct,omitempty"`
}

func (x *GoogleValue) Reset() {
	*x = GoogleValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleValue) ProtoMessage() {}

func (x *GoogleValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleValue.ProtoReflect.Descriptor instead.
func (*GoogleValue) Descriptor() ([]byte, []int) {
	return file_google_value_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleValue) GetArg() *structpb.Value {
	if x != nil {
		return x.Arg
	}
	return nil
}

func (x *GoogleValue) GetSomeList() *structpb.ListValue {
	if x != nil {
		return x.SomeList
	}
	return nil
}

func (x *GoogleValue) GetSomeStruct() *structpb.Struct {
	if x != nil {
		return x.SomeStruct
	}
	return nil
}

var File_google_value_proto protoreflect.FileDescriptor

var file_google_value_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x61,
	0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x03, 0x61, 0x72, 0x67, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x0b, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x6f,
	0x6d, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_value_proto_rawDescOnce sync.Once
	file_google_value_proto_rawDescData = file_google_value_proto_rawDesc
)

func file_google_value_proto_rawDescGZIP() []byte {
	file_google_value_proto_rawDescOnce.Do(func() {
		file_google_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_value_proto_rawDescData)
	})
	return file_google_value_proto_rawDescData
}

var file_google_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_value_proto_goTypes = []interface{}{
	(*GoogleValue)(nil),        // 0: samples.GoogleValue
	(*structpb.Value)(nil),     // 1: google.protobuf.Value
	(*structpb.ListValue)(nil), // 2: google.protobuf.ListValue
	(*structpb.Struct)(nil),    // 3: google.protobuf.Struct
}
var file_google_value_proto_depIdxs = []int32{
	1, // 0: samples.GoogleValue.arg:type_name -> google.protobuf.Value
	2, // 1: samples.GoogleValue.some_list:type_name -> google.protobuf.ListValue
	3, // 2: samples.GoogleValue.some_struct:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_value_proto_init() }
func file_google_value_proto_init() {
	if File_google_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleValue); i {
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
			RawDescriptor: file_google_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_value_proto_goTypes,
		DependencyIndexes: file_google_value_proto_depIdxs,
		MessageInfos:      file_google_value_proto_msgTypes,
	}.Build()
	File_google_value_proto = out.File
	file_google_value_proto_rawDesc = nil
	file_google_value_proto_goTypes = nil
	file_google_value_proto_depIdxs = nil
}