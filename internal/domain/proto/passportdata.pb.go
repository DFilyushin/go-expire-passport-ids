// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: internal/proto/passportdata.proto

package proto

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

type PassportDataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CsvHeader      string        `protobuf:"bytes,1,opt,name=csv_header,json=csvHeader,proto3" json:"csv_header,omitempty"`
	NumbersOnlyMap []*NumbersMap `protobuf:"bytes,2,rep,name=numbers_only_map,json=numbersOnlyMap,proto3" json:"numbers_only_map,omitempty"`
	OtherLines     []string      `protobuf:"bytes,3,rep,name=other_lines,json=otherLines,proto3" json:"other_lines,omitempty"`
}

func (x *PassportDataMessage) Reset() {
	*x = PassportDataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_passportdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassportDataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassportDataMessage) ProtoMessage() {}

func (x *PassportDataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_passportdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassportDataMessage.ProtoReflect.Descriptor instead.
func (*PassportDataMessage) Descriptor() ([]byte, []int) {
	return file_internal_proto_passportdata_proto_rawDescGZIP(), []int{0}
}

func (x *PassportDataMessage) GetCsvHeader() string {
	if x != nil {
		return x.CsvHeader
	}
	return ""
}

func (x *PassportDataMessage) GetNumbersOnlyMap() []*NumbersMap {
	if x != nil {
		return x.NumbersOnlyMap
	}
	return nil
}

func (x *PassportDataMessage) GetOtherLines() []string {
	if x != nil {
		return x.OtherLines
	}
	return nil
}

type NumbersMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SevenDigitsKey       int32  `protobuf:"varint,1,opt,name=seven_digits_key,json=sevenDigitsKey,proto3" json:"seven_digits_key,omitempty"`
	ThreeDigitsBitsValue []byte `protobuf:"bytes,2,opt,name=three_digits_bits_value,json=threeDigitsBitsValue,proto3" json:"three_digits_bits_value,omitempty"`
}

func (x *NumbersMap) Reset() {
	*x = NumbersMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_passportdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumbersMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumbersMap) ProtoMessage() {}

func (x *NumbersMap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_passportdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumbersMap.ProtoReflect.Descriptor instead.
func (*NumbersMap) Descriptor() ([]byte, []int) {
	return file_internal_proto_passportdata_proto_rawDescGZIP(), []int{1}
}

func (x *NumbersMap) GetSevenDigitsKey() int32 {
	if x != nil {
		return x.SevenDigitsKey
	}
	return 0
}

func (x *NumbersMap) GetThreeDigitsBitsValue() []byte {
	if x != nil {
		return x.ThreeDigitsBitsValue
	}
	return nil
}

var File_internal_proto_passportdata_proto protoreflect.FileDescriptor

var file_internal_proto_passportdata_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x95, 0x01,
	0x0a, 0x13, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x73, 0x76, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x73, 0x76, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x4d, 0x61, 0x70, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c,
	0x79, 0x4d, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x4c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x4d, 0x61, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x5f, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73,
	0x65, 0x76, 0x65, 0x6e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x35, 0x0a,
	0x17, 0x74, 0x68, 0x72, 0x65, 0x65, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x5f, 0x62, 0x69,
	0x74, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14,
	0x74, 0x68, 0x72, 0x65, 0x65, 0x44, 0x69, 0x67, 0x69, 0x74, 0x73, 0x42, 0x69, 0x74, 0x73, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_passportdata_proto_rawDescOnce sync.Once
	file_internal_proto_passportdata_proto_rawDescData = file_internal_proto_passportdata_proto_rawDesc
)

func file_internal_proto_passportdata_proto_rawDescGZIP() []byte {
	file_internal_proto_passportdata_proto_rawDescOnce.Do(func() {
		file_internal_proto_passportdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_passportdata_proto_rawDescData)
	})
	return file_internal_proto_passportdata_proto_rawDescData
}

var file_internal_proto_passportdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_passportdata_proto_goTypes = []interface{}{
	(*PassportDataMessage)(nil), // 0: internal.PassportDataMessage
	(*NumbersMap)(nil),          // 1: internal.NumbersMap
}
var file_internal_proto_passportdata_proto_depIdxs = []int32{
	1, // 0: internal.PassportDataMessage.numbers_only_map:type_name -> internal.NumbersMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_passportdata_proto_init() }
func file_internal_proto_passportdata_proto_init() {
	if File_internal_proto_passportdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_passportdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassportDataMessage); i {
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
		file_internal_proto_passportdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumbersMap); i {
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
			RawDescriptor: file_internal_proto_passportdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_passportdata_proto_goTypes,
		DependencyIndexes: file_internal_proto_passportdata_proto_depIdxs,
		MessageInfos:      file_internal_proto_passportdata_proto_msgTypes,
	}.Build()
	File_internal_proto_passportdata_proto = out.File
	file_internal_proto_passportdata_proto_rawDesc = nil
	file_internal_proto_passportdata_proto_goTypes = nil
	file_internal_proto_passportdata_proto_depIdxs = nil
}