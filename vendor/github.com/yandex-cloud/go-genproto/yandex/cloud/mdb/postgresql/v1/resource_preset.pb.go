// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/postgresql/v1/resource_preset.proto

package postgresql

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

// A ResourcePreset resource for describing hardware configuration presets.
type ResourcePreset struct {
	// ID of the ResourcePreset resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IDs of availability zones where the resource preset is available.
	ZoneIds []string `protobuf:"bytes,2,rep,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
	// Number of CPU cores for a PostgreSQL host created with the preset.
	Cores int64 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	// RAM volume for a PostgreSQL host created with the preset, in bytes.
	Memory               int64    `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcePreset) Reset()         { *m = ResourcePreset{} }
func (m *ResourcePreset) String() string { return proto.CompactTextString(m) }
func (*ResourcePreset) ProtoMessage()    {}
func (*ResourcePreset) Descriptor() ([]byte, []int) {
	return fileDescriptor_b92450db83d64039, []int{0}
}

func (m *ResourcePreset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourcePreset.Unmarshal(m, b)
}
func (m *ResourcePreset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourcePreset.Marshal(b, m, deterministic)
}
func (m *ResourcePreset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcePreset.Merge(m, src)
}
func (m *ResourcePreset) XXX_Size() int {
	return xxx_messageInfo_ResourcePreset.Size(m)
}
func (m *ResourcePreset) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcePreset.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcePreset proto.InternalMessageInfo

func (m *ResourcePreset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourcePreset) GetZoneIds() []string {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *ResourcePreset) GetCores() int64 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *ResourcePreset) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourcePreset)(nil), "yandex.cloud.mdb.postgresql.v1.ResourcePreset")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/postgresql/v1/resource_preset.proto", fileDescriptor_b92450db83d64039)
}

var fileDescriptor_b92450db83d64039 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0xb9, 0x3b, 0xad, 0x36, 0x43, 0x87, 0x20, 0x72, 0x2e, 0x72, 0x74, 0xba, 0xa5, 0x09,
	0x45, 0x37, 0x37, 0x37, 0x07, 0x41, 0x6e, 0x74, 0x29, 0xbd, 0x7b, 0x8f, 0x18, 0x68, 0xee, 0xc5,
	0xbc, 0x5c, 0xb1, 0xfe, 0xf5, 0x62, 0x52, 0x28, 0x22, 0x74, 0xfc, 0x3d, 0xde, 0x0f, 0xbe, 0xef,
	0x13, 0x8f, 0x87, 0xed, 0x08, 0xf8, 0xa5, 0x87, 0x1d, 0x4d, 0xa0, 0x1d, 0xf4, 0xda, 0x13, 0x47,
	0x13, 0x90, 0x3f, 0x77, 0x7a, 0xbf, 0xd6, 0x01, 0x99, 0xa6, 0x30, 0xe0, 0xc6, 0x07, 0x64, 0x8c,
	0xca, 0x07, 0x8a, 0x24, 0xef, 0xb3, 0xa5, 0x92, 0xa5, 0x1c, 0xf4, 0xea, 0x64, 0xa9, 0xfd, 0x7a,
	0x69, 0xc5, 0xa2, 0x3b, 0x8a, 0x6f, 0xc9, 0x93, 0x0b, 0x51, 0x5a, 0xa8, 0x8b, 0xa6, 0x68, 0xe7,
	0x5d, 0x69, 0x41, 0xde, 0x89, 0xeb, 0x6f, 0x1a, 0x71, 0x63, 0x81, 0xeb, 0xb2, 0xa9, 0xda, 0x79,
	0x77, 0xf5, 0xcb, 0x2f, 0xc0, 0xf2, 0x46, 0x5c, 0x0e, 0x14, 0x90, 0xeb, 0xaa, 0x29, 0xda, 0xaa,
	0xcb, 0x20, 0x6f, 0xc5, 0xcc, 0xa1, 0xa3, 0x70, 0xa8, 0x2f, 0xd2, 0xf9, 0x48, 0xcf, 0x2c, 0x96,
	0x7f, 0xc2, 0x6c, 0xbd, 0xfd, 0x1f, 0xe8, 0xfd, 0xd5, 0xd8, 0xf8, 0x31, 0xf5, 0x6a, 0x20, 0xa7,
	0xf3, 0xfb, 0x2a, 0x37, 0x36, 0xb4, 0x32, 0x38, 0xa6, 0x56, 0xfa, 0xfc, 0x14, 0x4f, 0x27, 0xea,
	0x67, 0x49, 0x78, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0xff, 0x3b, 0x03, 0x58, 0x3e, 0x01, 0x00,
	0x00,
}
