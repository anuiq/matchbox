// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	SelectGroupRequest
	SelectGroupResponse
	SelectProfileRequest
	SelectProfileResponse
	GroupPutRequest
	GroupPutResponse
	GroupGetRequest
	GroupListRequest
	GroupGetResponse
	GroupListResponse
	ProfilePutRequest
	ProfilePutResponse
	ProfileGetRequest
	ProfileGetResponse
	ProfileListRequest
	ProfileListResponse
	IgnitionPutRequest
	IgnitionPutResponse
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import storagepb "github.com/coreos/matchbox/matchbox/storage/storagepb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SelectGroupRequest struct {
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SelectGroupRequest) Reset()                    { *m = SelectGroupRequest{} }
func (m *SelectGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*SelectGroupRequest) ProtoMessage()               {}
func (*SelectGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SelectGroupRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SelectGroupResponse struct {
	Group *storagepb.Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *SelectGroupResponse) Reset()                    { *m = SelectGroupResponse{} }
func (m *SelectGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*SelectGroupResponse) ProtoMessage()               {}
func (*SelectGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SelectGroupResponse) GetGroup() *storagepb.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type SelectProfileRequest struct {
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SelectProfileRequest) Reset()                    { *m = SelectProfileRequest{} }
func (m *SelectProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*SelectProfileRequest) ProtoMessage()               {}
func (*SelectProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SelectProfileRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SelectProfileResponse struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *SelectProfileResponse) Reset()                    { *m = SelectProfileResponse{} }
func (m *SelectProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*SelectProfileResponse) ProtoMessage()               {}
func (*SelectProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SelectProfileResponse) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GroupPutRequest struct {
	Group *storagepb.Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *GroupPutRequest) Reset()                    { *m = GroupPutRequest{} }
func (m *GroupPutRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupPutRequest) ProtoMessage()               {}
func (*GroupPutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GroupPutRequest) GetGroup() *storagepb.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GroupPutResponse struct {
}

func (m *GroupPutResponse) Reset()                    { *m = GroupPutResponse{} }
func (m *GroupPutResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupPutResponse) ProtoMessage()               {}
func (*GroupPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GroupGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GroupGetRequest) Reset()                    { *m = GroupGetRequest{} }
func (m *GroupGetRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupGetRequest) ProtoMessage()               {}
func (*GroupGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GroupListRequest struct {
}

func (m *GroupListRequest) Reset()                    { *m = GroupListRequest{} }
func (m *GroupListRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupListRequest) ProtoMessage()               {}
func (*GroupListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GroupGetResponse struct {
	Group *storagepb.Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *GroupGetResponse) Reset()                    { *m = GroupGetResponse{} }
func (m *GroupGetResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupGetResponse) ProtoMessage()               {}
func (*GroupGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GroupGetResponse) GetGroup() *storagepb.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GroupListResponse struct {
	Groups []*storagepb.Group `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
}

func (m *GroupListResponse) Reset()                    { *m = GroupListResponse{} }
func (m *GroupListResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupListResponse) ProtoMessage()               {}
func (*GroupListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GroupListResponse) GetGroups() []*storagepb.Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ProfilePutRequest struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *ProfilePutRequest) Reset()                    { *m = ProfilePutRequest{} }
func (m *ProfilePutRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfilePutRequest) ProtoMessage()               {}
func (*ProfilePutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ProfilePutRequest) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ProfilePutResponse struct {
}

func (m *ProfilePutResponse) Reset()                    { *m = ProfilePutResponse{} }
func (m *ProfilePutResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfilePutResponse) ProtoMessage()               {}
func (*ProfilePutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type ProfileGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ProfileGetRequest) Reset()                    { *m = ProfileGetRequest{} }
func (m *ProfileGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileGetRequest) ProtoMessage()               {}
func (*ProfileGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type ProfileGetResponse struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *ProfileGetResponse) Reset()                    { *m = ProfileGetResponse{} }
func (m *ProfileGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileGetResponse) ProtoMessage()               {}
func (*ProfileGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ProfileGetResponse) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ProfileListRequest struct {
}

func (m *ProfileListRequest) Reset()                    { *m = ProfileListRequest{} }
func (m *ProfileListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileListRequest) ProtoMessage()               {}
func (*ProfileListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type ProfileListResponse struct {
	Profiles []*storagepb.Profile `protobuf:"bytes,1,rep,name=profiles" json:"profiles,omitempty"`
}

func (m *ProfileListResponse) Reset()                    { *m = ProfileListResponse{} }
func (m *ProfileListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileListResponse) ProtoMessage()               {}
func (*ProfileListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ProfileListResponse) GetProfiles() []*storagepb.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type IgnitionPutRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Config []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *IgnitionPutRequest) Reset()                    { *m = IgnitionPutRequest{} }
func (m *IgnitionPutRequest) String() string            { return proto.CompactTextString(m) }
func (*IgnitionPutRequest) ProtoMessage()               {}
func (*IgnitionPutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type IgnitionPutResponse struct {
}

func (m *IgnitionPutResponse) Reset()                    { *m = IgnitionPutResponse{} }
func (m *IgnitionPutResponse) String() string            { return proto.CompactTextString(m) }
func (*IgnitionPutResponse) ProtoMessage()               {}
func (*IgnitionPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func init() {
	proto.RegisterType((*SelectGroupRequest)(nil), "serverpb.SelectGroupRequest")
	proto.RegisterType((*SelectGroupResponse)(nil), "serverpb.SelectGroupResponse")
	proto.RegisterType((*SelectProfileRequest)(nil), "serverpb.SelectProfileRequest")
	proto.RegisterType((*SelectProfileResponse)(nil), "serverpb.SelectProfileResponse")
	proto.RegisterType((*GroupPutRequest)(nil), "serverpb.GroupPutRequest")
	proto.RegisterType((*GroupPutResponse)(nil), "serverpb.GroupPutResponse")
	proto.RegisterType((*GroupGetRequest)(nil), "serverpb.GroupGetRequest")
	proto.RegisterType((*GroupListRequest)(nil), "serverpb.GroupListRequest")
	proto.RegisterType((*GroupGetResponse)(nil), "serverpb.GroupGetResponse")
	proto.RegisterType((*GroupListResponse)(nil), "serverpb.GroupListResponse")
	proto.RegisterType((*ProfilePutRequest)(nil), "serverpb.ProfilePutRequest")
	proto.RegisterType((*ProfilePutResponse)(nil), "serverpb.ProfilePutResponse")
	proto.RegisterType((*ProfileGetRequest)(nil), "serverpb.ProfileGetRequest")
	proto.RegisterType((*ProfileGetResponse)(nil), "serverpb.ProfileGetResponse")
	proto.RegisterType((*ProfileListRequest)(nil), "serverpb.ProfileListRequest")
	proto.RegisterType((*ProfileListResponse)(nil), "serverpb.ProfileListResponse")
	proto.RegisterType((*IgnitionPutRequest)(nil), "serverpb.IgnitionPutRequest")
	proto.RegisterType((*IgnitionPutResponse)(nil), "serverpb.IgnitionPutResponse")
}

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x18, 0x25, 0xa9, 0x8d, 0xf5, 0x5b, 0xa9, 0xdb, 0xd9, 0xad, 0x2c, 0xbd, 0xaa, 0x23, 0x48, 0x10,
	0x9d, 0x85, 0x7a, 0x63, 0x0b, 0x85, 0x5a, 0x58, 0x8a, 0xd2, 0x8b, 0x12, 0x9f, 0x20, 0x49, 0xbf,
	0x8d, 0x83, 0x49, 0x26, 0xce, 0x4c, 0x0a, 0x7d, 0x0c, 0x2f, 0x7c, 0x5f, 0xe9, 0xfc, 0xc4, 0xc9,
	0x56, 0xc4, 0x8a, 0x57, 0x33, 0x73, 0xe6, 0x9c, 0xf3, 0xed, 0x39, 0xb3, 0x04, 0x76, 0x1b, 0x54,
	0x2a, 0xaf, 0x50, 0xb1, 0x4e, 0x0a, 0x2d, 0xc8, 0x8e, 0x42, 0x79, 0x83, 0xb2, 0x2b, 0x0e, 0x3e,
	0x55, 0x5c, 0x7f, 0xe9, 0x0b, 0x56, 0x8a, 0x66, 0x59, 0x0a, 0x89, 0x42, 0xb9, 0xe5, 0x6d, 0x91,
	0x4b, 0x6c, 0x50, 0xe7, 0xf5, 0xb2, 0x10, 0x42, 0x97, 0xeb, 0x6a, 0xa9, 0xb4, 0x90, 0x79, 0x85,
	0x7e, 0xed, 0x0a, 0xbf, 0xb3, 0xae, 0xf4, 0x7b, 0x04, 0xe4, 0x33, 0xd6, 0x58, 0xea, 0x0b, 0x29,
	0xfa, 0x2e, 0xc3, 0x6f, 0x3d, 0x2a, 0x4d, 0xce, 0x20, 0xa9, 0xf3, 0x02, 0x6b, 0xb5, 0x88, 0x0e,
	0xb7, 0xd2, 0xc9, 0x51, 0xca, 0xfc, 0x74, 0x76, 0x9f, 0xcd, 0x2e, 0x0d, 0x75, 0xd5, 0x6a, 0x79,
	0x9b, 0x39, 0xdd, 0xc1, 0x31, 0x4c, 0x02, 0x98, 0x4c, 0x61, 0xeb, 0x2b, 0xde, 0x2e, 0xa2, 0xc3,
	0x28, 0x7d, 0x92, 0xdd, 0x6d, 0xc9, 0x1c, 0xb6, 0x6f, 0xf2, 0xba, 0xc7, 0x45, 0x6c, 0x30, 0x7b,
	0x38, 0x89, 0xdf, 0x47, 0xf4, 0x14, 0x66, 0xa3, 0x21, 0xaa, 0x13, 0xad, 0x42, 0xf2, 0x0a, 0xb6,
	0xab, 0x3b, 0xc0, 0x98, 0x4c, 0x8e, 0xa6, 0x6c, 0xc8, 0xc4, 0x2c, 0xd1, 0x5e, 0xd3, 0x1f, 0x11,
	0xcc, 0xad, 0xfe, 0x4a, 0x8a, 0x35, 0xaf, 0xd1, 0x87, 0x3a, 0xdf, 0x08, 0xf5, 0x7a, 0x33, 0xd4,
	0x98, 0xff, 0xbf, 0x63, 0xad, 0x60, 0x7f, 0x63, 0x8c, 0x0b, 0xf6, 0x06, 0x1e, 0x77, 0x16, 0x72,
	0xd1, 0x48, 0x10, 0xcd, 0x93, 0x3d, 0x85, 0x1e, 0xc3, 0x33, 0x13, 0xf7, 0xaa, 0xd7, 0x3e, 0xd8,
	0xdf, 0x36, 0x43, 0x60, 0xfa, 0x4b, 0x6a, 0x87, 0xd3, 0x17, 0xce, 0xee, 0x02, 0x07, 0xbb, 0x5d,
	0x88, 0xf9, 0xb5, 0xcb, 0x14, 0xf3, 0xeb, 0x41, 0x76, 0xc9, 0x95, 0xe7, 0xd0, 0x13, 0x87, 0x19,
	0xd9, 0x03, 0x1f, 0xe8, 0x14, 0xf6, 0x02, 0x3f, 0x27, 0x4e, 0x21, 0x31, 0xb7, 0xfe, 0x71, 0xee,
	0xab, 0xdd, 0x3d, 0xfd, 0x00, 0x7b, 0xae, 0x94, 0xa0, 0x82, 0x87, 0x75, 0x38, 0x07, 0x12, 0x5a,
	0xb8, 0x2a, 0x5e, 0x0e, 0xc6, 0x7f, 0x28, 0xe3, 0x7c, 0x90, 0x86, 0xd1, 0xff, 0x75, 0x7c, 0x58,
	0xe9, 0x0a, 0x66, 0x23, 0xd4, 0x59, 0x33, 0xd8, 0x71, 0x3a, 0x5f, 0xcd, 0xef, 0xbc, 0x07, 0x0e,
	0x3d, 0x03, 0xf2, 0xb1, 0x6a, 0xb9, 0xe6, 0xa2, 0x0d, 0xfa, 0x21, 0xf0, 0xa8, 0xcd, 0x1b, 0x74,
	0x41, 0xcc, 0x9e, 0x3c, 0x87, 0xa4, 0x14, 0xed, 0x9a, 0x57, 0xe6, 0xbf, 0xfa, 0x34, 0x73, 0x27,
	0xba, 0x0f, 0xb3, 0x91, 0x83, 0xfd, 0x21, 0x45, 0x62, 0xbe, 0x18, 0xef, 0x7e, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x19, 0xef, 0xe3, 0x5b, 0x99, 0x04, 0x00, 0x00,
}
