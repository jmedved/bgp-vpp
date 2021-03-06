// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/customer_manager_link.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents customer-manager link relationship.
type CustomerManagerLink struct {
	// Name of the resource.
	// CustomerManagerLink resource names have the form:
	//
	// `customers/{customer_id}/customerManagerLinks/{manager_customer_id}_{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The manager customer linked to the customer.
	ManagerCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=manager_customer,json=managerCustomer,proto3" json:"manager_customer,omitempty"`
	// ID of the customer-manager link. This field is read only.
	ManagerLinkId *wrappers.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// Status of the link between the customer and the manager.
	Status               enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CustomerManagerLink) Reset()         { *m = CustomerManagerLink{} }
func (m *CustomerManagerLink) String() string { return proto.CompactTextString(m) }
func (*CustomerManagerLink) ProtoMessage()    {}
func (*CustomerManagerLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_manager_link_688ed550f2e6f71c, []int{0}
}
func (m *CustomerManagerLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerManagerLink.Unmarshal(m, b)
}
func (m *CustomerManagerLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerManagerLink.Marshal(b, m, deterministic)
}
func (dst *CustomerManagerLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerManagerLink.Merge(dst, src)
}
func (m *CustomerManagerLink) XXX_Size() int {
	return xxx_messageInfo_CustomerManagerLink.Size(m)
}
func (m *CustomerManagerLink) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerManagerLink.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerManagerLink proto.InternalMessageInfo

func (m *CustomerManagerLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerManagerLink) GetManagerCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ManagerCustomer
	}
	return nil
}

func (m *CustomerManagerLink) GetManagerLinkId() *wrappers.Int64Value {
	if m != nil {
		return m.ManagerLinkId
	}
	return nil
}

func (m *CustomerManagerLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CustomerManagerLink)(nil), "google.ads.googleads.v0.resources.CustomerManagerLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/customer_manager_link.proto", fileDescriptor_customer_manager_link_688ed550f2e6f71c)
}

var fileDescriptor_customer_manager_link_688ed550f2e6f71c = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6a, 0xdb, 0x30,
	0x1c, 0xc6, 0xce, 0x16, 0x98, 0xb7, 0x2c, 0xc3, 0xbb, 0x98, 0x6c, 0x8c, 0x64, 0x23, 0x90, 0x93,
	0x6c, 0xb2, 0xb1, 0x81, 0xc6, 0x0e, 0x4e, 0xd8, 0x42, 0xc6, 0x36, 0x82, 0x03, 0x3e, 0x0c, 0x53,
	0xa3, 0xc4, 0xaa, 0x30, 0xb1, 0x24, 0x23, 0xd9, 0xe9, 0x13, 0xf4, 0x45, 0x7a, 0x2a, 0x7d, 0x94,
	0x3e, 0x4a, 0x9f, 0xa2, 0xc4, 0xb2, 0xdc, 0x94, 0x34, 0xed, 0xed, 0xb3, 0xf4, 0xfd, 0xf9, 0x7d,
	0x3f, 0xd9, 0xfa, 0x41, 0x38, 0x27, 0x19, 0x76, 0x51, 0x22, 0x5d, 0x05, 0x77, 0x68, 0xeb, 0xb9,
	0x02, 0x4b, 0x5e, 0x8a, 0x35, 0x96, 0xee, 0xba, 0x94, 0x05, 0xa7, 0x58, 0xc4, 0x14, 0x31, 0x44,
	0xb0, 0x88, 0xb3, 0x94, 0x6d, 0x40, 0x2e, 0x78, 0xc1, 0xed, 0x81, 0xd2, 0x00, 0x94, 0x48, 0xd0,
	0xc8, 0xc1, 0xd6, 0x03, 0x8d, 0xbc, 0xf7, 0xed, 0x58, 0x02, 0x66, 0x25, 0x95, 0xee, 0xbe, 0x69,
	0x2c, 0x0b, 0x54, 0x94, 0x52, 0x79, 0xf7, 0x3e, 0xd4, 0xc2, 0xea, 0x6b, 0x55, 0x9e, 0xba, 0x67,
	0x02, 0xe5, 0x39, 0x16, 0xf5, 0xfd, 0xc7, 0x4b, 0xd3, 0x7a, 0x3b, 0xad, 0x67, 0xfb, 0xab, 0x5c,
	0xfe, 0xa4, 0x6c, 0x63, 0x7f, 0xb2, 0x3a, 0x3a, 0x3d, 0x66, 0x88, 0x62, 0xc7, 0xe8, 0x1b, 0xa3,
	0x17, 0xc1, 0x2b, 0x7d, 0xf8, 0x0f, 0x51, 0x6c, 0xcf, 0xac, 0x37, 0x3a, 0x59, 0xf7, 0x73, 0x5a,
	0x7d, 0x63, 0xf4, 0x72, 0xfc, 0xbe, 0x2e, 0x02, 0x74, 0x2e, 0x58, 0x16, 0x22, 0x65, 0x24, 0x44,
	0x59, 0x89, 0x83, 0x6e, 0xad, 0xd2, 0xc1, 0xf6, 0xd4, 0xea, 0xde, 0xab, 0x90, 0x26, 0xce, 0xb3,
	0xca, 0xe7, 0xdd, 0x81, 0xcf, 0x9c, 0x15, 0x5f, 0xbf, 0x28, 0x9b, 0x0e, 0xbd, 0x1b, 0x78, 0x9e,
	0xd8, 0x27, 0x56, 0x5b, 0x55, 0x77, 0x9e, 0xf7, 0x8d, 0xd1, 0xeb, 0xf1, 0x2f, 0x70, 0x6c, 0xaf,
	0xd5, 0xd2, 0xc0, 0x5e, 0xdd, 0x65, 0xa5, 0xfb, 0xc9, 0x4a, 0x7a, 0x78, 0x1a, 0xd4, 0xae, 0x93,
	0x73, 0xd3, 0x1a, 0xae, 0x39, 0x05, 0x4f, 0xbe, 0xd6, 0xc4, 0x79, 0x60, 0xa3, 0x8b, 0x5d, 0x81,
	0x85, 0xf1, 0xff, 0x77, 0x2d, 0x27, 0x3c, 0x43, 0x8c, 0x00, 0x2e, 0x88, 0x4b, 0x30, 0xab, 0xea,
	0xe9, 0x97, 0xcd, 0x53, 0xf9, 0xc8, 0xaf, 0xf4, 0xbd, 0x41, 0x17, 0x66, 0x6b, 0xe6, 0xfb, 0x57,
	0xe6, 0x60, 0xa6, 0x2c, 0xfd, 0x44, 0x02, 0x05, 0x77, 0x28, 0xf4, 0x40, 0xa0, 0x99, 0xd7, 0x9a,
	0x13, 0xf9, 0x89, 0x8c, 0x1a, 0x4e, 0x14, 0x7a, 0x51, 0xc3, 0xb9, 0x31, 0x87, 0xea, 0x02, 0x42,
	0x3f, 0x91, 0x10, 0x36, 0x2c, 0x08, 0x43, 0x0f, 0xc2, 0x86, 0xb7, 0x6a, 0x57, 0xc3, 0x7e, 0xbe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x05, 0x7b, 0x54, 0xf6, 0x02, 0x00, 0x00,
}
