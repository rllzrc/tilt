// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/resource/v1/resource.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Resource information.
type Resource struct {
	// Type identifier for the resource.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Set of labels that describe the resource.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_584700775a2fc762, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "opencensus.proto.resource.v1.Resource")
	proto.RegisterMapType((map[string]string)(nil), "opencensus.proto.resource.v1.Resource.LabelsEntry")
}

func init() {
	proto.RegisterFile("opencensus/proto/resource/v1/resource.proto", fileDescriptor_584700775a2fc762)
}

var fileDescriptor_584700775a2fc762 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xce, 0x2f, 0x48, 0xcd,
	0x4b, 0x4e, 0xcd, 0x2b, 0x2e, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4a, 0x2d,
	0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2f, 0x33, 0x84, 0xb3, 0xf5, 0xc0, 0x52, 0x42, 0x32, 0x08,
	0xc5, 0x10, 0x11, 0x3d, 0xb8, 0x82, 0x32, 0x43, 0xa5, 0xa5, 0x8c, 0x5c, 0x1c, 0x41, 0x50, 0xbe,
	0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0x2d, 0xe4, 0xc5, 0xc5, 0x96, 0x93, 0x98, 0x94, 0x9a, 0x53, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1,
	0x6d, 0x64, 0xa4, 0x87, 0xcf, 0x3c, 0x3d, 0x98, 0x59, 0x7a, 0x3e, 0x60, 0x4d, 0xae, 0x79, 0x25,
	0x45, 0x95, 0x41, 0x50, 0x13, 0xa4, 0x2c, 0xb9, 0xb8, 0x91, 0x84, 0x85, 0x04, 0xb8, 0x98, 0xb3,
	0x53, 0x2b, 0xa1, 0xb6, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12,
	0x4c, 0x60, 0x31, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0xd1, 0x69, 0x06, 0x23, 0x97, 0x7c, 0x66, 0x3e,
	0x5e, 0xbb, 0x9d, 0x78, 0x61, 0x96, 0x07, 0x80, 0xa4, 0x02, 0x18, 0xa3, 0x5c, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0xba, 0x74, 0x33, 0xf3, 0x8a, 0x4b, 0x8a,
	0x4a, 0x73, 0x53, 0xf3, 0x4a, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x11, 0x06, 0xea, 0x42, 0x42,
	0x32, 0x3d, 0x35, 0x4f, 0x37, 0x1d, 0x25, 0x40, 0x5f, 0x31, 0xc9, 0xf8, 0x17, 0xa4, 0xe6, 0x39,
	0x43, 0xac, 0x05, 0x9b, 0x8d, 0xf0, 0x66, 0x98, 0x61, 0x12, 0x1b, 0x58, 0xa3, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x32, 0xff, 0x46, 0x96, 0x01, 0x00, 0x00,
}
