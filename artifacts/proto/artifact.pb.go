// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifact.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldDescriptor struct {
	FriendlyName string   `protobuf:"bytes,1,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Repeated     bool     `protobuf:"varint,3,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Type         string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Doc          string   `protobuf:"bytes,5,opt,name=doc,proto3" json:"doc,omitempty"`
	Labels       []string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	Default      string   `protobuf:"bytes,7,opt,name=default,proto3" json:"default,omitempty"`
	// Same as doc field.
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldDescriptor) Reset()         { *m = FieldDescriptor{} }
func (m *FieldDescriptor) String() string { return proto.CompactTextString(m) }
func (*FieldDescriptor) ProtoMessage()    {}
func (*FieldDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{0}
}
func (m *FieldDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldDescriptor.Unmarshal(m, b)
}
func (m *FieldDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldDescriptor.Marshal(b, m, deterministic)
}
func (dst *FieldDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldDescriptor.Merge(dst, src)
}
func (m *FieldDescriptor) XXX_Size() int {
	return xxx_messageInfo_FieldDescriptor.Size(m)
}
func (m *FieldDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FieldDescriptor proto.InternalMessageInfo

func (m *FieldDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *FieldDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldDescriptor) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *FieldDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FieldDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *FieldDescriptor) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *FieldDescriptor) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *FieldDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type EnumValue struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumValue) Reset()         { *m = EnumValue{} }
func (m *EnumValue) String() string { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()    {}
func (*EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{1}
}
func (m *EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValue.Unmarshal(m, b)
}
func (m *EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValue.Marshal(b, m, deterministic)
}
func (dst *EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValue.Merge(dst, src)
}
func (m *EnumValue) XXX_Size() int {
	return xxx_messageInfo_EnumValue.Size(m)
}
func (m *EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValue proto.InternalMessageInfo

func (m *EnumValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *EnumValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TypeDescriptor struct {
	Doc string `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
	// Same as doc field.
	Description  string             `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Fields       []*FieldDescriptor `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name         string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FriendlyName string             `protobuf:"bytes,7,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Kind         string             `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	// The fields are all part of a single one of structure. NOTE:
	// Currently we only support an all or nothing approach to oneof -
	// there can be at most a single oneof group within the proto and
	// it implies that all the fields belong to it.
	Oneof                bool         `protobuf:"varint,5,opt,name=oneof,proto3" json:"oneof,omitempty"`
	Default              string       `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	AllowedValues        []*EnumValue `protobuf:"bytes,8,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TypeDescriptor) Reset()         { *m = TypeDescriptor{} }
func (m *TypeDescriptor) String() string { return proto.CompactTextString(m) }
func (*TypeDescriptor) ProtoMessage()    {}
func (*TypeDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{2}
}
func (m *TypeDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeDescriptor.Unmarshal(m, b)
}
func (m *TypeDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeDescriptor.Marshal(b, m, deterministic)
}
func (dst *TypeDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeDescriptor.Merge(dst, src)
}
func (m *TypeDescriptor) XXX_Size() int {
	return xxx_messageInfo_TypeDescriptor.Size(m)
}
func (m *TypeDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TypeDescriptor proto.InternalMessageInfo

func (m *TypeDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *TypeDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TypeDescriptor) GetFields() []*FieldDescriptor {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TypeDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypeDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *TypeDescriptor) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *TypeDescriptor) GetOneof() bool {
	if m != nil {
		return m.Oneof
	}
	return false
}

func (m *TypeDescriptor) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *TypeDescriptor) GetAllowedValues() []*EnumValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

type Types struct {
	Items                []*TypeDescriptor `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{3}
}
func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (dst *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(dst, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetItems() []*TypeDescriptor {
	if m != nil {
		return m.Items
	}
	return nil
}

type ArtifactParameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default              string   `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	FriendlyName         string   `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactParameter) Reset()         { *m = ArtifactParameter{} }
func (m *ArtifactParameter) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameter) ProtoMessage()    {}
func (*ArtifactParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{4}
}
func (m *ArtifactParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameter.Unmarshal(m, b)
}
func (m *ArtifactParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameter.Marshal(b, m, deterministic)
}
func (dst *ArtifactParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameter.Merge(dst, src)
}
func (m *ArtifactParameter) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameter.Size(m)
}
func (m *ArtifactParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameter proto.InternalMessageInfo

func (m *ArtifactParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactParameter) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *ArtifactParameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactParameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ArtifactParameter) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

type ArtifactSource struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Precondition         string   `protobuf:"bytes,1,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Queries              []string `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactSource) Reset()         { *m = ArtifactSource{} }
func (m *ArtifactSource) String() string { return proto.CompactTextString(m) }
func (*ArtifactSource) ProtoMessage()    {}
func (*ArtifactSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{5}
}
func (m *ArtifactSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactSource.Unmarshal(m, b)
}
func (m *ArtifactSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactSource.Marshal(b, m, deterministic)
}
func (dst *ArtifactSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactSource.Merge(dst, src)
}
func (m *ArtifactSource) XXX_Size() int {
	return xxx_messageInfo_ArtifactSource.Size(m)
}
func (m *ArtifactSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactSource.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactSource proto.InternalMessageInfo

func (m *ArtifactSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactSource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactSource) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *ArtifactSource) GetQueries() []string {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Artifact struct {
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Reference   []string `protobuf:"bytes,5,rep,name=reference,proto3" json:"reference,omitempty"`
	// If set here this applies to all sources.
	Precondition string               `protobuf:"bytes,8,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Parameters   []*ArtifactParameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	Sources      []*ArtifactSource    `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	// Internal use only
	Path string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	// Internal use only
	Raw                  string   `protobuf:"bytes,7,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{6}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (dst *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(dst, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Artifact) GetReference() []string {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Artifact) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *Artifact) GetParameters() []*ArtifactParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Artifact) GetSources() []*ArtifactSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Artifact) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Artifact) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type ArtifactDescriptors struct {
	Items                []*Artifact `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArtifactDescriptors) Reset()         { *m = ArtifactDescriptors{} }
func (m *ArtifactDescriptors) String() string { return proto.CompactTextString(m) }
func (*ArtifactDescriptors) ProtoMessage()    {}
func (*ArtifactDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifact_89fda6e8d5324ffc, []int{7}
}
func (m *ArtifactDescriptors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactDescriptors.Unmarshal(m, b)
}
func (m *ArtifactDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactDescriptors.Marshal(b, m, deterministic)
}
func (dst *ArtifactDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactDescriptors.Merge(dst, src)
}
func (m *ArtifactDescriptors) XXX_Size() int {
	return xxx_messageInfo_ArtifactDescriptors.Size(m)
}
func (m *ArtifactDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactDescriptors proto.InternalMessageInfo

func (m *ArtifactDescriptors) GetItems() []*Artifact {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*FieldDescriptor)(nil), "proto.FieldDescriptor")
	proto.RegisterType((*EnumValue)(nil), "proto.EnumValue")
	proto.RegisterType((*TypeDescriptor)(nil), "proto.TypeDescriptor")
	proto.RegisterType((*Types)(nil), "proto.Types")
	proto.RegisterType((*ArtifactParameter)(nil), "proto.ArtifactParameter")
	proto.RegisterType((*ArtifactSource)(nil), "proto.ArtifactSource")
	proto.RegisterType((*Artifact)(nil), "proto.Artifact")
	proto.RegisterType((*ArtifactDescriptors)(nil), "proto.ArtifactDescriptors")
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor_artifact_89fda6e8d5324ffc) }

var fileDescriptor_artifact_89fda6e8d5324ffc = []byte{
	// 1280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x8e, 0xdc, 0x34,
	0x14, 0x56, 0x76, 0xf6, 0xd7, 0x6d, 0xb7, 0xc5, 0x40, 0x15, 0xf5, 0x02, 0x1d, 0xa6, 0x02, 0x76,
	0xa1, 0xcd, 0xa2, 0x42, 0x45, 0xb5, 0xaa, 0x90, 0x66, 0x68, 0x11, 0xa0, 0x65, 0xb7, 0x4d, 0x57,
	0xad, 0xe8, 0x4d, 0xe5, 0x4d, 0x4e, 0x26, 0x66, 0x1d, 0x3b, 0xb5, 0x9d, 0xce, 0x0e, 0x37, 0x3c,
	0x01, 0x12, 0xe2, 0xe7, 0x06, 0xf1, 0x0e, 0x88, 0x77, 0xe0, 0x19, 0x90, 0x40, 0x42, 0x2a, 0xf0,
	0x1a, 0x48, 0x20, 0x3b, 0xc9, 0x64, 0x66, 0xa7, 0x82, 0x2b, 0xae, 0x12, 0xdb, 0xc7, 0xc7, 0x9f,
	0xbf, 0xef, 0xf3, 0x39, 0x64, 0x93, 0x69, 0xcb, 0x33, 0x96, 0xd8, 0xa8, 0xd4, 0xca, 0x2a, 0xba,
	0xe2, 0x3f, 0x97, 0x76, 0xc7, 0xe3, 0x71, 0xf4, 0x04, 0x85, 0x4a, 0x78, 0x8a, 0x27, 0x51, 0xa2,
	0x8a, 0x9d, 0x91, 0x12, 0x4c, 0x8e, 0x76, 0xea, 0x49, 0xcd, 0x4a, 0xab, 0xf4, 0x8e, 0x0f, 0xde,
	0x31, 0x58, 0x30, 0x69, 0x79, 0x52, 0xa7, 0xe8, 0x3f, 0x0d, 0xc8, 0xf9, 0xf7, 0x39, 0x8a, 0xf4,
	0x16, 0x9a, 0x44, 0x73, 0x17, 0x48, 0x2f, 0x93, 0x73, 0x99, 0xe6, 0x28, 0x53, 0x31, 0x79, 0x24,
	0x59, 0x81, 0x61, 0x00, 0xc1, 0xd6, 0x46, 0x7c, 0xb6, 0x9d, 0xdc, 0x67, 0x05, 0x52, 0x4a, 0x96,
	0xfd, 0xda, 0x92, 0x5f, 0xf3, 0xff, 0xf4, 0x12, 0x59, 0xd7, 0x58, 0x22, 0xb3, 0x98, 0x86, 0x3d,
	0x08, 0xb6, 0xd6, 0xe3, 0xe9, 0xd8, 0xc5, 0xdb, 0x49, 0x89, 0xe1, 0x72, 0x1d, 0xef, 0xfe, 0xe9,
	0x05, 0xd2, 0x4b, 0x55, 0x12, 0xae, 0xf8, 0x29, 0xf7, 0x4b, 0x2f, 0x92, 0x55, 0xc1, 0x8e, 0x50,
	0x98, 0x70, 0x15, 0x7a, 0x5b, 0x1b, 0x71, 0x33, 0xa2, 0x21, 0x59, 0x4b, 0x31, 0x63, 0x95, 0xb0,
	0xe1, 0x9a, 0x8f, 0x6e, 0x87, 0x14, 0xc8, 0x99, 0xb4, 0x81, 0xce, 0x95, 0x0c, 0xd7, 0xfd, 0xea,
	0xec, 0x54, 0xff, 0x3a, 0xd9, 0xb8, 0x2d, 0xab, 0xe2, 0x3e, 0x13, 0x15, 0xd2, 0x17, 0xc8, 0xca,
	0x13, 0xf7, 0xe3, 0xef, 0xb4, 0x12, 0xd7, 0x83, 0x67, 0x5d, 0xa6, 0xff, 0xfb, 0x12, 0xd9, 0x3c,
	0x9c, 0x94, 0x38, 0x43, 0x4c, 0x83, 0x37, 0xe8, 0xf0, 0x9e, 0x3a, 0x7d, 0x63, 0xe1, 0x74, 0x1a,
	0x91, 0xd5, 0xcc, 0xf1, 0x6b, 0xc2, 0x25, 0xe8, 0x6d, 0x9d, 0xb9, 0x76, 0xb1, 0x26, 0x3e, 0x3a,
	0x45, 0x7a, 0xdc, 0x44, 0x4d, 0xa1, 0xf4, 0x66, 0x78, 0x5d, 0x10, 0x64, 0xed, 0xd9, 0x82, 0x1c,
	0x73, 0x99, 0xb6, 0x04, 0xbb, 0x7f, 0x77, 0x5b, 0x25, 0x51, 0x65, 0x9e, 0xe2, 0xf5, 0xb8, 0x1e,
	0xd0, 0x83, 0x8e, 0xcc, 0x55, 0x17, 0x3c, 0xbc, 0xfe, 0xc7, 0x5f, 0x7f, 0xfe, 0x14, 0xec, 0xd0,
	0xab, 0x87, 0x39, 0xc2, 0xa7, 0x46, 0x49, 0x40, 0x99, 0xa8, 0x14, 0x53, 0x68, 0xe2, 0xc0, 0xd3,
	0x04, 0x99, 0xd2, 0x60, 0x73, 0x6e, 0xc0, 0xa9, 0x17, 0x75, 0x1a, 0xbc, 0x43, 0x36, 0x99, 0x10,
	0x6a, 0x8c, 0xe9, 0x23, 0x1f, 0x68, 0xc2, 0x75, 0x7f, 0xd7, 0x0b, 0xcd, 0x5d, 0xa7, 0xf4, 0xc7,
	0xe7, 0x9a, 0x38, 0x3f, 0x32, 0xfd, 0xb7, 0xc9, 0x8a, 0xa3, 0xd8, 0xd0, 0x37, 0xc8, 0x0a, 0xb7,
	0x58, 0x98, 0x30, 0xf0, 0x1b, 0x5f, 0x6c, 0x36, 0xce, 0xf3, 0x1f, 0xd7, 0x31, 0xfd, 0xbf, 0x97,
	0xc8, 0x73, 0x83, 0xe6, 0x25, 0xdc, 0x61, 0x9a, 0x15, 0x68, 0x51, 0x4f, 0x89, 0x0b, 0x66, 0x88,
	0x9b, 0xb1, 0xcd, 0xd2, 0xbf, 0xda, 0xa6, 0xb7, 0x28, 0xdc, 0xb3, 0x0c, 0xbb, 0x20, 0xc4, 0xca,
	0xa2, 0x10, 0xbb, 0x4f, 0x83, 0x5f, 0x1d, 0x9d, 0xbf, 0x04, 0xe4, 0xe7, 0xa0, 0x85, 0x69, 0xa0,
	0x60, 0x13, 0x60, 0x49, 0x82, 0xa5, 0x85, 0xb2, 0xc5, 0x6c, 0x60, 0x9c, 0xf3, 0x24, 0x07, 0xa6,
	0x11, 0x58, 0xea, 0xf8, 0xb6, 0x0a, 0x6c, 0x8e, 0x60, 0x12, 0x55, 0x22, 0x94, 0x9a, 0x3b, 0xc2,
	0x15, 0xe0, 0x09, 0x26, 0x95, 0x03, 0x15, 0xc1, 0xfe, 0xc1, 0xe1, 0xed, 0x5d, 0x60, 0x42, 0xcc,
	0x66, 0x71, 0xfb, 0x8d, 0xd5, 0x5c, 0x8e, 0x0c, 0x5c, 0x05, 0x9e, 0xc1, 0x44, 0x55, 0x20, 0x11,
	0x53, 0x30, 0xaa, 0x40, 0x9b, 0x73, 0x39, 0x02, 0x14, 0x06, 0x7d, 0xee, 0xc7, 0x15, 0xea, 0x09,
	0x24, 0x4c, 0x42, 0x25, 0x4b, 0x96, 0x1c, 0x03, 0x46, 0xa3, 0x08, 0x32, 0xad, 0x0a, 0xf8, 0xe8,
	0xde, 0xc1, 0x3e, 0x54, 0xc6, 0x85, 0xbb, 0x48, 0x37, 0xbc, 0xc3, 0xb4, 0xc1, 0xad, 0x6d, 0xb8,
	0x7f, 0x77, 0x0f, 0xb2, 0x4a, 0x26, 0x1e, 0x45, 0xff, 0xb7, 0x65, 0xb2, 0xd9, 0x5e, 0xed, 0x9e,
	0xaa, 0x74, 0x82, 0xf4, 0x87, 0x60, 0xd6, 0xb8, 0xc3, 0xef, 0x02, 0xef, 0xa9, 0xaf, 0x03, 0xfa,
	0x65, 0xe0, 0x5c, 0xe5, 0x96, 0x40, 0x65, 0xb5, 0x7d, 0xda, 0x32, 0x06, 0xc6, 0xef, 0x8d, 0xe0,
	0xc3, 0x0c, 0xa4, 0xb2, 0x60, 0xd0, 0xc2, 0x18, 0xa1, 0x6a, 0x90, 0x9a, 0xe9, 0x1e, 0xec, 0xb6,
	0x70, 0x6b, 0x50, 0x64, 0x11, 0x1c, 0xce, 0x4e, 0x26, 0xaa, 0x28, 0xb9, 0x40, 0x0d, 0x63, 0x2e,
	0x04, 0x8c, 0x50, 0xa2, 0x66, 0x16, 0x81, 0x35, 0x97, 0x1d, 0x73, 0x9b, 0xd7, 0x27, 0x3b, 0x18,
	0x51, 0x63, 0x8e, 0x2f, 0x82, 0x79, 0x0f, 0x78, 0xa1, 0x87, 0xc7, 0x1e, 0x37, 0xd2, 0x64, 0x00,
	0x33, 0x8b, 0x0d, 0xbd, 0xdd, 0x0b, 0x68, 0x91, 0xef, 0x2b, 0x8b, 0xc0, 0xad, 0xe7, 0xf3, 0x08,
	0x81, 0x4b, 0x8b, 0xba, 0x54, 0xc2, 0x95, 0xbc, 0xfa, 0x58, 0x65, 0x73, 0xd4, 0x1d, 0xd2, 0x99,
	0x9c, 0x26, 0x9a, 0x37, 0x5c, 0x4e, 0xce, 0x96, 0x1a, 0x13, 0x25, 0x53, 0xee, 0xf1, 0x78, 0x23,
	0x0f, 0x6f, 0x79, 0x3c, 0xef, 0xd2, 0x9b, 0x03, 0xaf, 0x01, 0x9e, 0x94, 0x1a, 0x8d, 0x71, 0x90,
	0xac, 0x72, 0x27, 0xa2, 0x7b, 0x71, 0xfe, 0xb8, 0xa9, 0x5d, 0x5a, 0x05, 0x3b, 0x98, 0xf1, 0x5c,
	0x66, 0x6a, 0xc9, 0x9a, 0x23, 0x86, 0x63, 0x5d, 0x94, 0x36, 0x86, 0x0f, 0xfd, 0x21, 0x87, 0x34,
	0xbe, 0x5b, 0x4f, 0x83, 0xcd, 0x99, 0xad, 0xc9, 0xd4, 0x95, 0x04, 0x2e, 0x41, 0xe9, 0x14, 0x75,
	0x04, 0x07, 0x52, 0x4c, 0x40, 0x55, 0xb6, 0xac, 0x6c, 0x6d, 0x19, 0xa7, 0x8d, 0x60, 0xc6, 0x4e,
	0x99, 0x16, 0xc2, 0xa1, 0x4a, 0x94, 0x10, 0x98, 0x58, 0x4c, 0xa3, 0xb8, 0x3d, 0x6a, 0x77, 0xdb,
	0x3f, 0x8b, 0xcb, 0xe4, 0xe5, 0x07, 0x39, 0x6a, 0x9c, 0x57, 0x75, 0x84, 0xd6, 0x38, 0x69, 0x21,
	0x65, 0x96, 0x45, 0xfd, 0x6f, 0xd6, 0xc8, 0x7a, 0xeb, 0x2f, 0xfa, 0x63, 0x30, 0xfb, 0xb2, 0x87,
	0xdf, 0xd7, 0xce, 0xfa, 0x36, 0xa0, 0x5f, 0x9d, 0x72, 0x56, 0x97, 0x2f, 0x82, 0x7b, 0xb9, 0xaa,
	0x44, 0xea, 0xa0, 0x54, 0x92, 0x3f, 0xae, 0x10, 0x98, 0x4c, 0xfd, 0x1b, 0x4c, 0x94, 0xb4, 0x8c,
	0x4b, 0x48, 0x95, 0x35, 0x11, 0x0c, 0x9c, 0xd9, 0xb2, 0x4a, 0x80, 0x49, 0x72, 0x2c, 0x10, 0x5c,
	0x6d, 0x53, 0x70, 0xa4, 0x91, 0x1d, 0x43, 0xc2, 0x2c, 0x8e, 0x94, 0xa7, 0xc3, 0x2b, 0x99, 0x2a,
	0x5b, 0xbf, 0x95, 0x3d, 0x2e, 0xab, 0x93, 0x68, 0xa8, 0xd5, 0xd8, 0xa0, 0x36, 0xd1, 0x7b, 0xb9,
	0x56, 0x05, 0x7e, 0xc0, 0x8d, 0x55, 0x7a, 0xd2, 0x78, 0xeb, 0xee, 0xbc, 0xb5, 0x7c, 0xf1, 0x19,
	0xee, 0x78, 0xe0, 0xdb, 0xf4, 0xb5, 0x07, 0x8e, 0xdd, 0x79, 0x57, 0x1b, 0xb0, 0x7a, 0xe2, 0x65,
	0x53, 0x2d, 0x79, 0xa7, 0xec, 0x71, 0x8b, 0x6c, 0x68, 0xcc, 0x50, 0xa3, 0x4c, 0x5c, 0xdd, 0x71,
	0xb2, 0xbd, 0xea, 0x13, 0x02, 0x7d, 0x69, 0x00, 0xd3, 0xa5, 0xce, 0xa2, 0x53, 0x32, 0xe2, 0x6e,
	0xe3, 0x82, 0xc9, 0xd6, 0xff, 0x37, 0x93, 0x09, 0x42, 0xba, 0x12, 0x15, 0xf6, 0x7c, 0x5d, 0x0f,
	0x9b, 0xba, 0xbe, 0x50, 0xbd, 0x87, 0xd7, 0x3c, 0x82, 0x2b, 0xf4, 0xf5, 0x3b, 0x5d, 0x59, 0xab,
	0xcf, 0x2e, 0xb5, 0x7a, 0xc2, 0x67, 0x6a, 0x63, 0x77, 0xad, 0x99, 0xfc, 0xf4, 0x21, 0x59, 0xab,
	0x61, 0x98, 0x70, 0x79, 0xae, 0x85, 0xcc, 0x97, 0xa9, 0xe1, 0xb6, 0x3f, 0xe7, 0x32, 0xfd, 0x6f,
	0x13, 0xc6, 0x6d, 0x42, 0xba, 0x47, 0x96, 0x4b, 0x66, 0xf3, 0xa6, 0x59, 0xde, 0xf0, 0x19, 0xae,
	0xd1, 0x37, 0x9d, 0xf9, 0xdc, 0xfc, 0x62, 0x59, 0xe3, 0x72, 0x3e, 0xb1, 0x73, 0x86, 0x2b, 0x3b,
	0x2e, 0x9a, 0xde, 0x20, 0x3d, 0xcd, 0xc6, 0x75, 0x0b, 0xef, 0x14, 0x74, 0xc9, 0x34, 0x1b, 0xc3,
	0x27, 0x83, 0x8f, 0xf7, 0x16, 0x12, 0x46, 0xb1, 0xdb, 0xb2, 0xfb, 0xb9, 0x7f, 0x40, 0x13, 0x32,
	0x1e, 0xc8, 0x2e, 0xf7, 0x58, 0xb3, 0xd2, 0x00, 0xf3, 0xba, 0xd5, 0x6f, 0x90, 0x4b, 0xd0, 0x58,
	0x19, 0x76, 0x24, 0xf0, 0x0a, 0xa4, 0x2a, 0xa9, 0x0a, 0x94, 0xbe, 0x1c, 0xb1, 0x49, 0xd4, 0xf5,
	0x23, 0xdf, 0x73, 0x84, 0x00, 0x76, 0xa4, 0x2a, 0xdb, 0x5a, 0xae, 0x91, 0xd4, 0xf5, 0x11, 0x57,
	0x8f, 0x99, 0x64, 0x62, 0xf2, 0x59, 0xd3, 0x0e, 0x8a, 0xa8, 0x7f, 0x93, 0x3c, 0xdf, 0x26, 0xe8,
	0xba, 0xb2, 0xa1, 0xaf, 0xcc, 0x37, 0xef, 0xf3, 0xa7, 0x98, 0x6f, 0xda, 0xf6, 0xd1, 0xaa, 0x9f,
	0x7e, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xb5, 0x8d, 0x89, 0xc6, 0x0a, 0x00, 0x00,
}
