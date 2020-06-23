// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/clickhouse/v1/resource_preset.proto

package clickhouse

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
	// Number of CPU cores for a ClickHouse host created with the preset.
	Cores int64 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	// RAM volume for a ClickHouse host created with the preset, in bytes.
	Memory               int64    `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcePreset) Reset()         { *m = ResourcePreset{} }
func (m *ResourcePreset) String() string { return proto.CompactTextString(m) }
func (*ResourcePreset) ProtoMessage()    {}
func (*ResourcePreset) Descriptor() ([]byte, []int) {
	return fileDescriptor_24eaf27369654811, []int{0}
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
	proto.RegisterType((*ResourcePreset)(nil), "yandex.cloud.mdb.clickhouse.v1.ResourcePreset")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/clickhouse/v1/resource_preset.proto", fileDescriptor_24eaf27369654811)
}

var fileDescriptor_24eaf27369654811 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0xb9, 0x3b, 0xad, 0x36, 0x43, 0x87, 0x20, 0x12, 0x17, 0x39, 0x3a, 0xdd, 0xd2, 0x84,
	0xa2, 0x9b, 0x9b, 0x9b, 0x83, 0x20, 0x19, 0x5d, 0x4a, 0x93, 0xf7, 0xb8, 0x06, 0x9b, 0x7b, 0x47,
	0x72, 0x29, 0xd6, 0xbf, 0x5e, 0x4c, 0x0a, 0x45, 0x04, 0xc7, 0xdf, 0xe3, 0xfd, 0xe0, 0xfb, 0x3e,
	0xf6, 0x78, 0xdc, 0x0e, 0x80, 0x9f, 0xca, 0xee, 0x29, 0x81, 0xf2, 0x60, 0x94, 0xdd, 0x3b, 0xfb,
	0xb1, 0xa3, 0x14, 0x51, 0x1d, 0xd6, 0x2a, 0x60, 0xa4, 0x14, 0x2c, 0x6e, 0xc6, 0x80, 0x11, 0x27,
	0x39, 0x06, 0x9a, 0x88, 0xdf, 0x17, 0x4b, 0x66, 0x4b, 0x7a, 0x30, 0xf2, 0x6c, 0xc9, 0xc3, 0x7a,
	0xe9, 0xd8, 0x42, 0x9f, 0xc4, 0xb7, 0xec, 0xf1, 0x05, 0xab, 0x1d, 0x88, 0xaa, 0xad, 0xba, 0xb9,
	0xae, 0x1d, 0xf0, 0x3b, 0x76, 0xfd, 0x45, 0x03, 0x6e, 0x1c, 0x44, 0x51, 0xb7, 0x4d, 0x37, 0xd7,
	0x57, 0x3f, 0xfc, 0x02, 0x91, 0xdf, 0xb0, 0x4b, 0x4b, 0x01, 0xa3, 0x68, 0xda, 0xaa, 0x6b, 0x74,
	0x01, 0x7e, 0xcb, 0x66, 0x1e, 0x3d, 0x85, 0xa3, 0xb8, 0xc8, 0xe7, 0x13, 0x3d, 0x47, 0xb6, 0xfc,
	0x15, 0x66, 0x3b, 0xba, 0xbf, 0x81, 0xde, 0x5f, 0x7b, 0x37, 0xed, 0x92, 0x91, 0x96, 0xbc, 0x2a,
	0xef, 0xab, 0xd2, 0xb8, 0xa7, 0x55, 0x8f, 0x43, 0x6e, 0xa5, 0xfe, 0x9f, 0xe2, 0xe9, 0x4c, 0x66,
	0x96, 0x85, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x2a, 0x7f, 0x4c, 0x3e, 0x01, 0x00,
	0x00,
}
