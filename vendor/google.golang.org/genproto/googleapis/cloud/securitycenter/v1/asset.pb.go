// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/asset.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Cloud Security Command Center's (Cloud SCC) representation of a Google Cloud
// Platform (GCP) resource.
//
// The Asset is a Cloud SCC resource that captures information about a single
// GCP resource. All modifications to an Asset are only within the context of
// Cloud SCC and don't affect the referenced GCP resource.
type Asset struct {
	// The relative resource name of this asset. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/assets/456".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cloud SCC managed properties. These properties are managed by
	// Cloud SCC and cannot be modified by the user.
	SecurityCenterProperties *Asset_SecurityCenterProperties `protobuf:"bytes,2,opt,name=security_center_properties,json=securityCenterProperties,proto3" json:"security_center_properties,omitempty"`
	// Resource managed properties. These properties are managed and defined by
	// the GCP resource and cannot be modified by the user.
	ResourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=resource_properties,json=resourceProperties,proto3" json:"resource_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified security marks. These marks are entirely managed by the user
	// and come from the SecurityMarks resource that belongs to the asset.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the asset was created in Cloud SCC.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the asset was last updated, added, or deleted in Cloud
	// SCC.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// IAM Policy information associated with the GCP resource described by the
	// Cloud SCC asset. This information is managed and defined by the GCP
	// resource and cannot be modified by the user.
	IamPolicy            *Asset_IamPolicy `protobuf:"bytes,11,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_e3a1e16638a04b0e, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (dst *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(dst, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetSecurityCenterProperties() *Asset_SecurityCenterProperties {
	if m != nil {
		return m.SecurityCenterProperties
	}
	return nil
}

func (m *Asset) GetResourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.ResourceProperties
	}
	return nil
}

func (m *Asset) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Asset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Asset) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Asset) GetIamPolicy() *Asset_IamPolicy {
	if m != nil {
		return m.IamPolicy
	}
	return nil
}

// Cloud SCC managed properties. These properties are managed by Cloud SCC and
// cannot be modified by the user.
type Asset_SecurityCenterProperties struct {
	// The full resource name of the GCP resource this asset
	// represents. This field is immutable after create time. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The type of the GCP resource. Examples include: APPLICATION,
	// PROJECT, and ORGANIZATION. This is a case insensitive field defined by
	// Cloud SCC and/or the producer of the resource and is immutable
	// after create time.
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The full resource name of the immediate parent of the resource. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceParent string `protobuf:"bytes,3,opt,name=resource_parent,json=resourceParent,proto3" json:"resource_parent,omitempty"`
	// The full resource name of the project the resource belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceProject string `protobuf:"bytes,4,opt,name=resource_project,json=resourceProject,proto3" json:"resource_project,omitempty"`
	// Owners of the Google Cloud resource.
	ResourceOwners       []string `protobuf:"bytes,5,rep,name=resource_owners,json=resourceOwners,proto3" json:"resource_owners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset_SecurityCenterProperties) Reset()         { *m = Asset_SecurityCenterProperties{} }
func (m *Asset_SecurityCenterProperties) String() string { return proto.CompactTextString(m) }
func (*Asset_SecurityCenterProperties) ProtoMessage()    {}
func (*Asset_SecurityCenterProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_e3a1e16638a04b0e, []int{0, 0}
}
func (m *Asset_SecurityCenterProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Unmarshal(m, b)
}
func (m *Asset_SecurityCenterProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Marshal(b, m, deterministic)
}
func (dst *Asset_SecurityCenterProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_SecurityCenterProperties.Merge(dst, src)
}
func (m *Asset_SecurityCenterProperties) XXX_Size() int {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Size(m)
}
func (m *Asset_SecurityCenterProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_SecurityCenterProperties.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_SecurityCenterProperties proto.InternalMessageInfo

func (m *Asset_SecurityCenterProperties) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceParent() string {
	if m != nil {
		return m.ResourceParent
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceProject() string {
	if m != nil {
		return m.ResourceProject
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceOwners() []string {
	if m != nil {
		return m.ResourceOwners
	}
	return nil
}

// IAM Policy information associated with the GCP resource described by the
// Cloud SCC asset. This information is managed and defined by the GCP
// resource and cannot be modified by the user.
type Asset_IamPolicy struct {
	// The JSON representation of the Policy associated with the asset.
	// See https://cloud.google.com/iam/reference/rest/v1/Policy for format
	// details.
	PolicyBlob           string   `protobuf:"bytes,1,opt,name=policy_blob,json=policyBlob,proto3" json:"policy_blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset_IamPolicy) Reset()         { *m = Asset_IamPolicy{} }
func (m *Asset_IamPolicy) String() string { return proto.CompactTextString(m) }
func (*Asset_IamPolicy) ProtoMessage()    {}
func (*Asset_IamPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_e3a1e16638a04b0e, []int{0, 1}
}
func (m *Asset_IamPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_IamPolicy.Unmarshal(m, b)
}
func (m *Asset_IamPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_IamPolicy.Marshal(b, m, deterministic)
}
func (dst *Asset_IamPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_IamPolicy.Merge(dst, src)
}
func (m *Asset_IamPolicy) XXX_Size() int {
	return xxx_messageInfo_Asset_IamPolicy.Size(m)
}
func (m *Asset_IamPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_IamPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_IamPolicy proto.InternalMessageInfo

func (m *Asset_IamPolicy) GetPolicyBlob() string {
	if m != nil {
		return m.PolicyBlob
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.cloud.securitycenter.v1.Asset")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1.Asset.ResourcePropertiesEntry")
	proto.RegisterType((*Asset_SecurityCenterProperties)(nil), "google.cloud.securitycenter.v1.Asset.SecurityCenterProperties")
	proto.RegisterType((*Asset_IamPolicy)(nil), "google.cloud.securitycenter.v1.Asset.IamPolicy")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/asset.proto", fileDescriptor_asset_e3a1e16638a04b0e)
}

var fileDescriptor_asset_e3a1e16638a04b0e = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0xdd, 0xae, 0x75, 0xef, 0xda, 0x0f, 0x46, 0xd0, 0x21, 0x88, 0x5d, 0xea, 0x83, 0xab,
	0xd4, 0x84, 0xb6, 0x2f, 0x62, 0x51, 0xb0, 0xe2, 0x83, 0xa0, 0x75, 0x89, 0x8b, 0x0f, 0x82, 0x2c,
	0xb3, 0xe9, 0x35, 0xc4, 0x26, 0x33, 0xc3, 0xcc, 0x64, 0x25, 0xe0, 0x2f, 0xf1, 0xdf, 0xf9, 0x4f,
	0x24, 0x93, 0x49, 0x9a, 0xb5, 0x6c, 0xb7, 0x6f, 0x93, 0x73, 0xcf, 0x39, 0x37, 0x77, 0xce, 0x4d,
	0xe0, 0x79, 0x22, 0x44, 0x92, 0x61, 0x18, 0x67, 0xa2, 0xb8, 0x08, 0x35, 0xc6, 0x85, 0x4a, 0x4d,
	0x19, 0x23, 0x37, 0xa8, 0xc2, 0xc5, 0x51, 0xc8, 0xb4, 0x46, 0x13, 0x48, 0x25, 0x8c, 0x20, 0x8f,
	0x6b, 0x6e, 0x60, 0xb9, 0xc1, 0x32, 0x37, 0x58, 0x1c, 0xf9, 0x8f, 0x9c, 0x97, 0x65, 0xcf, 0x8b,
	0x1f, 0xa1, 0x36, 0xaa, 0x88, 0x9d, 0xda, 0xdf, 0xff, 0xbf, 0x6a, 0xd2, 0x1c, 0xb5, 0x61, 0xb9,
	0x74, 0x84, 0x93, 0x35, 0xaf, 0xd2, 0x20, 0xb3, 0x9c, 0xa9, 0x4b, 0xed, 0x44, 0x4d, 0x4f, 0x26,
	0xd3, 0x90, 0x71, 0x2e, 0x0c, 0x33, 0xa9, 0xe0, 0xae, 0x7a, 0xf0, 0x67, 0x0b, 0xfa, 0x6f, 0xab,
	0x09, 0x08, 0x81, 0x4d, 0xce, 0x72, 0xa4, 0xde, 0xc8, 0x1b, 0x0f, 0x22, 0x7b, 0x26, 0xbf, 0xc1,
	0x6f, 0x3d, 0xeb, 0x36, 0x33, 0xa9, 0x84, 0x44, 0x65, 0x52, 0xd4, 0x74, 0x63, 0xe4, 0x8d, 0x87,
	0xc7, 0x6f, 0x82, 0x9b, 0x87, 0x0e, 0xac, 0x7d, 0xf0, 0xc5, 0xe1, 0xef, 0x2c, 0x3e, 0x69, 0x5d,
	0x22, 0xaa, 0x57, 0x54, 0x08, 0x87, 0xfb, 0x0a, 0xb5, 0x28, 0x54, 0x8c, 0xdd, 0xb6, 0x5b, 0xa3,
	0xde, 0x78, 0x78, 0xfc, 0xfa, 0x76, 0x6d, 0x23, 0x67, 0x70, 0x65, 0xfb, 0x9e, 0x1b, 0x55, 0x46,
	0x44, 0x5d, 0x2b, 0x90, 0x29, 0xec, 0x2c, 0xdf, 0x20, 0xbd, 0x6b, 0x27, 0x7c, 0xb1, 0xae, 0x55,
	0x33, 0xdb, 0xa7, 0x4a, 0x14, 0x6d, 0xeb, 0xee, 0x23, 0x39, 0x85, 0x61, 0xac, 0x90, 0x19, 0x9c,
	0x55, 0x71, 0xd2, 0x81, 0xb5, 0xf4, 0x1b, 0xcb, 0x26, 0xeb, 0x60, 0xda, 0x64, 0x1d, 0x41, 0x4d,
	0xaf, 0x80, 0x4a, 0x5c, 0xc8, 0x8b, 0x56, 0x0c, 0xeb, 0xc5, 0x35, 0xdd, 0x8a, 0xcf, 0x01, 0x52,
	0x96, 0xcf, 0xa4, 0xc8, 0xd2, 0xb8, 0xa4, 0x43, 0xab, 0x0d, 0x6f, 0x77, 0x6d, 0x1f, 0x58, 0x3e,
	0xb1, 0xb2, 0x68, 0x90, 0x36, 0x47, 0xff, 0xaf, 0x07, 0x74, 0x55, 0x8c, 0xe4, 0x09, 0x6c, 0xb7,
	0x61, 0x75, 0xf6, 0xe8, 0x5e, 0x03, 0x9e, 0x57, 0xfb, 0xd4, 0x25, 0x99, 0x52, 0xa2, 0x5d, 0xa1,
	0x0e, 0x69, 0x5a, 0x4a, 0x24, 0x4f, 0x61, 0xf7, 0x2a, 0x76, 0xa6, 0x90, 0x1b, 0xda, 0xb3, 0xb4,
	0x9d, 0x36, 0x33, 0x8b, 0x92, 0x67, 0xb0, 0xd7, 0xdd, 0x8f, 0x9f, 0x18, 0x1b, 0xba, 0x69, 0x99,
	0xbb, 0x9d, 0x74, 0x2b, 0x78, 0xc9, 0x53, 0xfc, 0xe2, 0xa8, 0x34, 0xed, 0x8f, 0x7a, 0x5d, 0xcf,
	0xcf, 0x16, 0xf5, 0x0f, 0x61, 0xd0, 0xce, 0x4e, 0xf6, 0x61, 0x58, 0x5f, 0xde, 0x6c, 0x9e, 0x89,
	0xb9, 0x9b, 0x08, 0x6a, 0xe8, 0x2c, 0x13, 0x73, 0xff, 0x3b, 0x3c, 0x5c, 0xb1, 0x60, 0x64, 0x0f,
	0x7a, 0x97, 0x58, 0x3a, 0x4d, 0x75, 0x24, 0x87, 0xd0, 0x5f, 0xb0, 0xac, 0x40, 0xf7, 0xdd, 0x3c,
	0xb8, 0x96, 0xe2, 0xd7, 0xaa, 0x1a, 0xd5, 0xa4, 0x57, 0x1b, 0x2f, 0xbd, 0x33, 0x03, 0x07, 0xb1,
	0xc8, 0xd7, 0x24, 0x36, 0xf1, 0xbe, 0x7d, 0x74, 0x8c, 0x44, 0x64, 0x8c, 0x27, 0x81, 0x50, 0x49,
	0x98, 0x20, 0xb7, 0xbe, 0x61, 0x5d, 0x62, 0x32, 0xd5, 0xab, 0x7e, 0x1b, 0xa7, 0xcb, 0xc8, 0xfc,
	0x8e, 0x15, 0x9e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x46, 0xa7, 0xee, 0xa3, 0xf9, 0x04, 0x00,
	0x00,
}
