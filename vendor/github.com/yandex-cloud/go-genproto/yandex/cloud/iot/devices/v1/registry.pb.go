// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iot/devices/v1/registry.proto

package devices

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Registry_Status int32

const (
	Registry_STATUS_UNSPECIFIED Registry_Status = 0
	// Registry is being created.
	Registry_CREATING Registry_Status = 1
	// Registry is ready to use.
	Registry_ACTIVE Registry_Status = 2
	// Registry is being deleted.
	Registry_DELETING Registry_Status = 3
)

var Registry_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "ACTIVE",
	3: "DELETING",
}

var Registry_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"ACTIVE":             2,
	"DELETING":           3,
}

func (x Registry_Status) String() string {
	return proto.EnumName(Registry_Status_name, int32(x))
}

func (Registry_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{0, 0}
}

// A registry. For more information, see [Registry](/docs/iot-core/concepts/index#registry).
type Registry struct {
	// ID of the registry.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the registry belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the registry. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the registry. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the registry.
	Status Registry_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.iot.devices.v1.Registry_Status" json:"status,omitempty"`
	// ID of the logs group for the specified registry.
	LogGroupId           string   `protobuf:"bytes,8,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{0}
}

func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Registry) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Registry) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Registry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Registry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Registry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Registry) GetStatus() Registry_Status {
	if m != nil {
		return m.Status
	}
	return Registry_STATUS_UNSPECIFIED
}

func (m *Registry) GetLogGroupId() string {
	if m != nil {
		return m.LogGroupId
	}
	return ""
}

// A registry certificate. For more information, see [Managing registry certificates](/docs/iot-core/operations/certificates/registry-certificates).
type RegistryCertificate struct {
	// ID of the registry that the certificate belongs to.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// SHA256 hash of the certificates.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Public part of the certificate.
	CertificateData string `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	// Creation timestamp.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegistryCertificate) Reset()         { *m = RegistryCertificate{} }
func (m *RegistryCertificate) String() string { return proto.CompactTextString(m) }
func (*RegistryCertificate) ProtoMessage()    {}
func (*RegistryCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{1}
}

func (m *RegistryCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryCertificate.Unmarshal(m, b)
}
func (m *RegistryCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryCertificate.Marshal(b, m, deterministic)
}
func (m *RegistryCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryCertificate.Merge(m, src)
}
func (m *RegistryCertificate) XXX_Size() int {
	return xxx_messageInfo_RegistryCertificate.Size(m)
}
func (m *RegistryCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryCertificate proto.InternalMessageInfo

func (m *RegistryCertificate) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *RegistryCertificate) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *RegistryCertificate) GetCertificateData() string {
	if m != nil {
		return m.CertificateData
	}
	return ""
}

func (m *RegistryCertificate) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// A device topic alias.
//
// Alias is an alternate name of a device topic assigned by the user. Map alias to canonical topic name prefix, e.g. `my/custom/alias` match to `$device/abcdef/events`. For more information, see [Using topic aliases](/docs/iot-core/concepts/topic#aliases).
type DeviceAlias struct {
	// ID of the device that the alias belongs to.
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Prefix of a canonical topic name to be aliased, e.g. `$devices/abcdef`.
	TopicPrefix string `protobuf:"bytes,2,opt,name=topic_prefix,json=topicPrefix,proto3" json:"topic_prefix,omitempty"`
	// Alias of a device topic.
	Alias                string   `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceAlias) Reset()         { *m = DeviceAlias{} }
func (m *DeviceAlias) String() string { return proto.CompactTextString(m) }
func (*DeviceAlias) ProtoMessage()    {}
func (*DeviceAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{2}
}

func (m *DeviceAlias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAlias.Unmarshal(m, b)
}
func (m *DeviceAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAlias.Marshal(b, m, deterministic)
}
func (m *DeviceAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAlias.Merge(m, src)
}
func (m *DeviceAlias) XXX_Size() int {
	return xxx_messageInfo_DeviceAlias.Size(m)
}
func (m *DeviceAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAlias.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAlias proto.InternalMessageInfo

func (m *DeviceAlias) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceAlias) GetTopicPrefix() string {
	if m != nil {
		return m.TopicPrefix
	}
	return ""
}

func (m *DeviceAlias) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

// A registry password.
type RegistryPassword struct {
	// ID of the registry that the password belongs to.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// ID of the password.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Creation timestamp.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegistryPassword) Reset()         { *m = RegistryPassword{} }
func (m *RegistryPassword) String() string { return proto.CompactTextString(m) }
func (*RegistryPassword) ProtoMessage()    {}
func (*RegistryPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{3}
}

func (m *RegistryPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryPassword.Unmarshal(m, b)
}
func (m *RegistryPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryPassword.Marshal(b, m, deterministic)
}
func (m *RegistryPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryPassword.Merge(m, src)
}
func (m *RegistryPassword) XXX_Size() int {
	return xxx_messageInfo_RegistryPassword.Size(m)
}
func (m *RegistryPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryPassword.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryPassword proto.InternalMessageInfo

func (m *RegistryPassword) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *RegistryPassword) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegistryPassword) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.iot.devices.v1.Registry_Status", Registry_Status_name, Registry_Status_value)
	proto.RegisterType((*Registry)(nil), "yandex.cloud.iot.devices.v1.Registry")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.iot.devices.v1.Registry.LabelsEntry")
	proto.RegisterType((*RegistryCertificate)(nil), "yandex.cloud.iot.devices.v1.RegistryCertificate")
	proto.RegisterType((*DeviceAlias)(nil), "yandex.cloud.iot.devices.v1.DeviceAlias")
	proto.RegisterType((*RegistryPassword)(nil), "yandex.cloud.iot.devices.v1.RegistryPassword")
}

func init() {
	proto.RegisterFile("yandex/cloud/iot/devices/v1/registry.proto", fileDescriptor_39c05472a87f1ea4)
}

var fileDescriptor_39c05472a87f1ea4 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x26, 0xed, 0x56, 0xda, 0xb7, 0xd3, 0xa8, 0x0c, 0x42, 0x51, 0x27, 0xb4, 0xd0, 0x53, 0x41,
	0xcc, 0xd1, 0xc6, 0x85, 0xc1, 0xa9, 0xb4, 0x61, 0x44, 0x9a, 0xa6, 0x2a, 0xed, 0x38, 0x70, 0x89,
	0xdc, 0xd8, 0x0d, 0x86, 0x34, 0x8e, 0x1c, 0xa7, 0xac, 0x17, 0x7e, 0x1a, 0xbf, 0x85, 0x9f, 0x82,
	0x62, 0x27, 0x5b, 0xb7, 0xc3, 0xf8, 0xb8, 0x39, 0x8f, 0xdf, 0x8f, 0xc7, 0xcf, 0xfb, 0xbc, 0x81,
	0x97, 0x1b, 0x92, 0x52, 0x76, 0xe5, 0x46, 0x89, 0x28, 0xa8, 0xcb, 0x85, 0x72, 0x29, 0x5b, 0xf3,
	0x88, 0xe5, 0xee, 0xfa, 0xd8, 0x95, 0x2c, 0xe6, 0xb9, 0x92, 0x1b, 0x9c, 0x49, 0xa1, 0x04, 0x3a,
	0x30, 0xb1, 0x58, 0xc7, 0x62, 0x2e, 0x14, 0xae, 0x62, 0xf1, 0xfa, 0xb8, 0x7f, 0x18, 0x0b, 0x11,
	0x27, 0xcc, 0xd5, 0xa1, 0x8b, 0x62, 0xe9, 0x2a, 0xbe, 0x62, 0xb9, 0x22, 0xab, 0xcc, 0x64, 0xf7,
	0x9f, 0xdd, 0xea, 0xb4, 0x26, 0x09, 0xa7, 0x44, 0x71, 0x91, 0x9a, 0xeb, 0xc1, 0xaf, 0x26, 0xb4,
	0x83, 0xaa, 0x1f, 0xda, 0x87, 0x06, 0xa7, 0xb6, 0xe5, 0x58, 0xc3, 0x4e, 0xd0, 0xe0, 0x14, 0x1d,
	0x40, 0x67, 0x29, 0x12, 0xca, 0x64, 0xc8, 0xa9, 0xdd, 0xd0, 0x70, 0xdb, 0x00, 0x3e, 0x45, 0xa7,
	0x00, 0x91, 0x64, 0x44, 0x31, 0x1a, 0x12, 0x65, 0x37, 0x1d, 0x6b, 0xd8, 0x3d, 0xe9, 0x63, 0x43,
	0x07, 0xd7, 0x74, 0xf0, 0xbc, 0xa6, 0x13, 0x74, 0xaa, 0xe8, 0x91, 0x42, 0x08, 0x76, 0x52, 0xb2,
	0x62, 0xf6, 0x8e, 0x2e, 0xa9, 0xcf, 0xc8, 0x81, 0x2e, 0x65, 0x79, 0x24, 0x79, 0x56, 0xb2, 0xb3,
	0x77, 0xf5, 0xd5, 0x36, 0x84, 0x7c, 0x68, 0x25, 0x64, 0xc1, 0x92, 0xdc, 0x6e, 0x39, 0xcd, 0x61,
	0xf7, 0xe4, 0x18, 0xdf, 0x23, 0x0c, 0xae, 0x1f, 0x85, 0xcf, 0x75, 0x8e, 0x97, 0x2a, 0xb9, 0x09,
	0xaa, 0x02, 0x68, 0x02, 0xad, 0x5c, 0x11, 0x55, 0xe4, 0xf6, 0x43, 0xc7, 0x1a, 0xee, 0x9f, 0xbc,
	0xfa, 0xbb, 0x52, 0x33, 0x9d, 0x13, 0x54, 0xb9, 0xc8, 0x81, 0xbd, 0x44, 0xc4, 0x61, 0x2c, 0x45,
	0x91, 0x95, 0x0a, 0xb5, 0x35, 0x67, 0x48, 0x44, 0x7c, 0x56, 0x42, 0x3e, 0xed, 0x9f, 0x42, 0x77,
	0xab, 0x3d, 0xea, 0x41, 0xf3, 0x1b, 0xdb, 0x54, 0x02, 0x97, 0x47, 0xf4, 0x04, 0x76, 0xd7, 0x24,
	0x29, 0x58, 0xa5, 0xae, 0xf9, 0x78, 0xdb, 0x78, 0x63, 0x0d, 0x3e, 0x42, 0xcb, 0xb4, 0x43, 0x4f,
	0x01, 0xcd, 0xe6, 0xa3, 0xf9, 0xe5, 0x2c, 0xbc, 0xbc, 0x98, 0x4d, 0xbd, 0xb1, 0xff, 0xc1, 0xf7,
	0x26, 0xbd, 0x07, 0x68, 0x0f, 0xda, 0xe3, 0xc0, 0x1b, 0xcd, 0xfd, 0x8b, 0xb3, 0x9e, 0x85, 0x00,
	0x5a, 0xa3, 0xf1, 0xdc, 0xff, 0xe4, 0xf5, 0x1a, 0xe5, 0xcd, 0xc4, 0x3b, 0xf7, 0xf4, 0x4d, 0x73,
	0xf0, 0xd3, 0x82, 0xc7, 0xf5, 0x13, 0xc6, 0x4c, 0x2a, 0xbe, 0xe4, 0x11, 0x51, 0x0c, 0x1d, 0x42,
	0xb7, 0x76, 0x5a, 0x78, 0x3d, 0x76, 0xa8, 0x21, 0x9f, 0x96, 0x23, 0x59, 0xf2, 0x34, 0x66, 0x32,
	0x93, 0x3c, 0x55, 0x15, 0xc5, 0x6d, 0x08, 0xbd, 0x80, 0x5e, 0x74, 0x53, 0x31, 0xa4, 0x44, 0x11,
	0xed, 0x84, 0x4e, 0xf0, 0x68, 0x0b, 0x9f, 0x10, 0x45, 0xee, 0xd8, 0x65, 0xe7, 0x1f, 0xec, 0x32,
	0x88, 0xa0, 0x3b, 0xd1, 0xd3, 0x18, 0x25, 0x9c, 0xe4, 0xa5, 0x2b, 0xcd, 0x70, 0x6e, 0x58, 0xb7,
	0x0d, 0xe0, 0x53, 0xf4, 0x1c, 0xf6, 0x94, 0xc8, 0x78, 0x14, 0x66, 0x92, 0x2d, 0xf9, 0x55, 0x4d,
	0x5a, 0x63, 0x53, 0x0d, 0x95, 0x9a, 0x93, 0xb2, 0x50, 0xc5, 0xd4, 0x7c, 0x0c, 0x7e, 0x40, 0xaf,
	0x16, 0x69, 0x4a, 0xf2, 0xfc, 0xbb, 0x90, 0xf4, 0xcf, 0x0a, 0x99, 0x85, 0x69, 0x5c, 0x2f, 0xcc,
	0xff, 0xef, 0xc4, 0xfb, 0xaf, 0x70, 0x78, 0xcb, 0x83, 0x24, 0xe3, 0x77, 0x7c, 0xf8, 0xf9, 0x2c,
	0xe6, 0xea, 0x4b, 0xb1, 0xc0, 0x91, 0x58, 0xb9, 0x26, 0xf6, 0xc8, 0x6c, 0x75, 0x2c, 0x8e, 0x62,
	0x96, 0xea, 0xfa, 0xee, 0x3d, 0x3f, 0x96, 0x77, 0xd5, 0x71, 0xd1, 0xd2, 0xa1, 0xaf, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xae, 0x9e, 0x11, 0xb4, 0x86, 0x04, 0x00, 0x00,
}
