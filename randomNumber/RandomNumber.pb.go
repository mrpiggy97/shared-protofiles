// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protofiles/RandomNumber.proto

package randomNumber

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

type RandomNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *RandomNumberRequest) Reset() {
	*x = RandomNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_RandomNumber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNumberRequest) ProtoMessage() {}

func (x *RandomNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_RandomNumber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNumberRequest.ProtoReflect.Descriptor instead.
func (*RandomNumberRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_RandomNumber_proto_rawDescGZIP(), []int{0}
}

func (x *RandomNumberRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type RandomNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalNumber int64 `protobuf:"varint,1,opt,name=originalNumber,proto3" json:"originalNumber,omitempty"`
	TotalNumber    int64 `protobuf:"varint,2,opt,name=totalNumber,proto3" json:"totalNumber,omitempty"`
}

func (x *RandomNumberResponse) Reset() {
	*x = RandomNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_RandomNumber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNumberResponse) ProtoMessage() {}

func (x *RandomNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_RandomNumber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNumberResponse.ProtoReflect.Descriptor instead.
func (*RandomNumberResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_RandomNumber_proto_rawDescGZIP(), []int{1}
}

func (x *RandomNumberResponse) GetOriginalNumber() int64 {
	if x != nil {
		return x.OriginalNumber
	}
	return 0
}

func (x *RandomNumberResponse) GetTotalNumber() int64 {
	if x != nil {
		return x.TotalNumber
	}
	return 0
}

var File_protofiles_RandomNumber_proto protoreflect.FileDescriptor

var file_protofiles_RandomNumber_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x14, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xb9, 0x01, 0x0a, 0x0d, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x56, 0x0a,
	0x15, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_RandomNumber_proto_rawDescOnce sync.Once
	file_protofiles_RandomNumber_proto_rawDescData = file_protofiles_RandomNumber_proto_rawDesc
)

func file_protofiles_RandomNumber_proto_rawDescGZIP() []byte {
	file_protofiles_RandomNumber_proto_rawDescOnce.Do(func() {
		file_protofiles_RandomNumber_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_RandomNumber_proto_rawDescData)
	})
	return file_protofiles_RandomNumber_proto_rawDescData
}

var file_protofiles_RandomNumber_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protofiles_RandomNumber_proto_goTypes = []interface{}{
	(*RandomNumberRequest)(nil),  // 0: numbers.RandomNumberRequest
	(*RandomNumberResponse)(nil), // 1: numbers.RandomNumberResponse
}
var file_protofiles_RandomNumber_proto_depIdxs = []int32{
	0, // 0: numbers.RandomService.AddRandomNumber:input_type -> numbers.RandomNumberRequest
	0, // 1: numbers.RandomService.SubstractRandomNumber:input_type -> numbers.RandomNumberRequest
	1, // 2: numbers.RandomService.AddRandomNumber:output_type -> numbers.RandomNumberResponse
	1, // 3: numbers.RandomService.SubstractRandomNumber:output_type -> numbers.RandomNumberResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofiles_RandomNumber_proto_init() }
func file_protofiles_RandomNumber_proto_init() {
	if File_protofiles_RandomNumber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_RandomNumber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNumberRequest); i {
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
		file_protofiles_RandomNumber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNumberResponse); i {
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
			RawDescriptor: file_protofiles_RandomNumber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_RandomNumber_proto_goTypes,
		DependencyIndexes: file_protofiles_RandomNumber_proto_depIdxs,
		MessageInfos:      file_protofiles_RandomNumber_proto_msgTypes,
	}.Build()
	File_protofiles_RandomNumber_proto = out.File
	file_protofiles_RandomNumber_proto_rawDesc = nil
	file_protofiles_RandomNumber_proto_goTypes = nil
	file_protofiles_RandomNumber_proto_depIdxs = nil
}
