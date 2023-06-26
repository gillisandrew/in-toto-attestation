// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: in_toto_attestation/predicates/link/v0/link.proto

package v0

import (
	v1 "github.com/in-toto/attestation/go/v1"
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

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command     []string                 `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	Materials   []*v1.ResourceDescriptor `protobuf:"bytes,3,rep,name=materials,proto3" json:"materials,omitempty"`
	Byproducts  *structpb.Struct         `protobuf:"bytes,4,opt,name=byproducts,proto3" json:"byproducts,omitempty"`
	Environment *structpb.Struct         `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_toto_attestation_predicates_link_v0_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_link_v0_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_link_v0_link_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *Link) GetMaterials() []*v1.ResourceDescriptor {
	if x != nil {
		return x.Materials
	}
	return nil
}

func (x *Link) GetByproducts() *structpb.Struct {
	if x != nil {
		return x.Byproducts
	}
	return nil
}

func (x *Link) GetEnvironment() *structpb.Struct {
	if x != nil {
		return x.Environment
	}
	return nil
}

var File_in_toto_attestation_predicates_link_v0_link_proto protoreflect.FileDescriptor

var file_in_toto_attestation_predicates_link_v0_link_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x30, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x26, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x30, 0x1a, 0x30, 0x69, 0x6e, 0x5f,
	0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x04,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x0a,
	0x62, 0x79, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x62, 0x79, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x67, 0x0a, 0x2f, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6e,
	0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x30, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_in_toto_attestation_predicates_link_v0_link_proto_rawDescOnce sync.Once
	file_in_toto_attestation_predicates_link_v0_link_proto_rawDescData = file_in_toto_attestation_predicates_link_v0_link_proto_rawDesc
)

func file_in_toto_attestation_predicates_link_v0_link_proto_rawDescGZIP() []byte {
	file_in_toto_attestation_predicates_link_v0_link_proto_rawDescOnce.Do(func() {
		file_in_toto_attestation_predicates_link_v0_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_in_toto_attestation_predicates_link_v0_link_proto_rawDescData)
	})
	return file_in_toto_attestation_predicates_link_v0_link_proto_rawDescData
}

var file_in_toto_attestation_predicates_link_v0_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_in_toto_attestation_predicates_link_v0_link_proto_goTypes = []interface{}{
	(*Link)(nil),                  // 0: in_toto_attestation.predicates.link.v0.Link
	(*v1.ResourceDescriptor)(nil), // 1: in_toto_attestation.v1.ResourceDescriptor
	(*structpb.Struct)(nil),       // 2: google.protobuf.Struct
}
var file_in_toto_attestation_predicates_link_v0_link_proto_depIdxs = []int32{
	1, // 0: in_toto_attestation.predicates.link.v0.Link.materials:type_name -> in_toto_attestation.v1.ResourceDescriptor
	2, // 1: in_toto_attestation.predicates.link.v0.Link.byproducts:type_name -> google.protobuf.Struct
	2, // 2: in_toto_attestation.predicates.link.v0.Link.environment:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_in_toto_attestation_predicates_link_v0_link_proto_init() }
func file_in_toto_attestation_predicates_link_v0_link_proto_init() {
	if File_in_toto_attestation_predicates_link_v0_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_in_toto_attestation_predicates_link_v0_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_in_toto_attestation_predicates_link_v0_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_toto_attestation_predicates_link_v0_link_proto_goTypes,
		DependencyIndexes: file_in_toto_attestation_predicates_link_v0_link_proto_depIdxs,
		MessageInfos:      file_in_toto_attestation_predicates_link_v0_link_proto_msgTypes,
	}.Build()
	File_in_toto_attestation_predicates_link_v0_link_proto = out.File
	file_in_toto_attestation_predicates_link_v0_link_proto_rawDesc = nil
	file_in_toto_attestation_predicates_link_v0_link_proto_goTypes = nil
	file_in_toto_attestation_predicates_link_v0_link_proto_depIdxs = nil
}
