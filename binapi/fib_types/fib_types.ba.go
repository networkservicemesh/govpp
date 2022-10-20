// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.06.0-1~g1e20fd388
// source: /usr/share/vpp/api/core/fib_types.api.json

// Package fib_types contains generated bindings for API file fib_types.api.
//
// Contents:
//   3 enums
//   3 structs
//
package fib_types

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// FibPathFlags defines enum 'fib_path_flags'.
type FibPathFlags uint32

const (
	FIB_API_PATH_FLAG_NONE                 FibPathFlags = 0
	FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED FibPathFlags = 1
	FIB_API_PATH_FLAG_RESOLVE_VIA_HOST     FibPathFlags = 2
	FIB_API_PATH_FLAG_POP_PW_CW            FibPathFlags = 4
)

var (
	FibPathFlags_name = map[uint32]string{
		0: "FIB_API_PATH_FLAG_NONE",
		1: "FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED",
		2: "FIB_API_PATH_FLAG_RESOLVE_VIA_HOST",
		4: "FIB_API_PATH_FLAG_POP_PW_CW",
	}
	FibPathFlags_value = map[string]uint32{
		"FIB_API_PATH_FLAG_NONE":                 0,
		"FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED": 1,
		"FIB_API_PATH_FLAG_RESOLVE_VIA_HOST":     2,
		"FIB_API_PATH_FLAG_POP_PW_CW":            4,
	}
)

func (x FibPathFlags) String() string {
	s, ok := FibPathFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := FibPathFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "FibPathFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// FibPathNhProto defines enum 'fib_path_nh_proto'.
type FibPathNhProto uint32

const (
	FIB_API_PATH_NH_PROTO_IP4      FibPathNhProto = 0
	FIB_API_PATH_NH_PROTO_IP6      FibPathNhProto = 1
	FIB_API_PATH_NH_PROTO_MPLS     FibPathNhProto = 2
	FIB_API_PATH_NH_PROTO_ETHERNET FibPathNhProto = 3
	FIB_API_PATH_NH_PROTO_BIER     FibPathNhProto = 4
)

var (
	FibPathNhProto_name = map[uint32]string{
		0: "FIB_API_PATH_NH_PROTO_IP4",
		1: "FIB_API_PATH_NH_PROTO_IP6",
		2: "FIB_API_PATH_NH_PROTO_MPLS",
		3: "FIB_API_PATH_NH_PROTO_ETHERNET",
		4: "FIB_API_PATH_NH_PROTO_BIER",
	}
	FibPathNhProto_value = map[string]uint32{
		"FIB_API_PATH_NH_PROTO_IP4":      0,
		"FIB_API_PATH_NH_PROTO_IP6":      1,
		"FIB_API_PATH_NH_PROTO_MPLS":     2,
		"FIB_API_PATH_NH_PROTO_ETHERNET": 3,
		"FIB_API_PATH_NH_PROTO_BIER":     4,
	}
)

func (x FibPathNhProto) String() string {
	s, ok := FibPathNhProto_name[uint32(x)]
	if ok {
		return s
	}
	return "FibPathNhProto(" + strconv.Itoa(int(x)) + ")"
}

// FibPathType defines enum 'fib_path_type'.
type FibPathType uint32

const (
	FIB_API_PATH_TYPE_NORMAL        FibPathType = 0
	FIB_API_PATH_TYPE_LOCAL         FibPathType = 1
	FIB_API_PATH_TYPE_DROP          FibPathType = 2
	FIB_API_PATH_TYPE_UDP_ENCAP     FibPathType = 3
	FIB_API_PATH_TYPE_BIER_IMP      FibPathType = 4
	FIB_API_PATH_TYPE_ICMP_UNREACH  FibPathType = 5
	FIB_API_PATH_TYPE_ICMP_PROHIBIT FibPathType = 6
	FIB_API_PATH_TYPE_SOURCE_LOOKUP FibPathType = 7
	FIB_API_PATH_TYPE_DVR           FibPathType = 8
	FIB_API_PATH_TYPE_INTERFACE_RX  FibPathType = 9
	FIB_API_PATH_TYPE_CLASSIFY      FibPathType = 10
)

var (
	FibPathType_name = map[uint32]string{
		0:  "FIB_API_PATH_TYPE_NORMAL",
		1:  "FIB_API_PATH_TYPE_LOCAL",
		2:  "FIB_API_PATH_TYPE_DROP",
		3:  "FIB_API_PATH_TYPE_UDP_ENCAP",
		4:  "FIB_API_PATH_TYPE_BIER_IMP",
		5:  "FIB_API_PATH_TYPE_ICMP_UNREACH",
		6:  "FIB_API_PATH_TYPE_ICMP_PROHIBIT",
		7:  "FIB_API_PATH_TYPE_SOURCE_LOOKUP",
		8:  "FIB_API_PATH_TYPE_DVR",
		9:  "FIB_API_PATH_TYPE_INTERFACE_RX",
		10: "FIB_API_PATH_TYPE_CLASSIFY",
	}
	FibPathType_value = map[string]uint32{
		"FIB_API_PATH_TYPE_NORMAL":        0,
		"FIB_API_PATH_TYPE_LOCAL":         1,
		"FIB_API_PATH_TYPE_DROP":          2,
		"FIB_API_PATH_TYPE_UDP_ENCAP":     3,
		"FIB_API_PATH_TYPE_BIER_IMP":      4,
		"FIB_API_PATH_TYPE_ICMP_UNREACH":  5,
		"FIB_API_PATH_TYPE_ICMP_PROHIBIT": 6,
		"FIB_API_PATH_TYPE_SOURCE_LOOKUP": 7,
		"FIB_API_PATH_TYPE_DVR":           8,
		"FIB_API_PATH_TYPE_INTERFACE_RX":  9,
		"FIB_API_PATH_TYPE_CLASSIFY":      10,
	}
)

func (x FibPathType) String() string {
	s, ok := FibPathType_name[uint32(x)]
	if ok {
		return s
	}
	return "FibPathType(" + strconv.Itoa(int(x)) + ")"
}

// FibMplsLabel defines type 'fib_mpls_label'.
type FibMplsLabel struct {
	IsUniform uint8  `binapi:"u8,name=is_uniform" json:"is_uniform,omitempty"`
	Label     uint32 `binapi:"u32,name=label" json:"label,omitempty"`
	TTL       uint8  `binapi:"u8,name=ttl" json:"ttl,omitempty"`
	Exp       uint8  `binapi:"u8,name=exp" json:"exp,omitempty"`
}

// FibPath defines type 'fib_path'.
type FibPath struct {
	SwIfIndex  uint32           `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	TableID    uint32           `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	RpfID      uint32           `binapi:"u32,name=rpf_id" json:"rpf_id,omitempty"`
	Weight     uint8            `binapi:"u8,name=weight" json:"weight,omitempty"`
	Preference uint8            `binapi:"u8,name=preference" json:"preference,omitempty"`
	Type       FibPathType      `binapi:"fib_path_type,name=type" json:"type,omitempty"`
	Flags      FibPathFlags     `binapi:"fib_path_flags,name=flags" json:"flags,omitempty"`
	Proto      FibPathNhProto   `binapi:"fib_path_nh_proto,name=proto" json:"proto,omitempty"`
	Nh         FibPathNh        `binapi:"fib_path_nh,name=nh" json:"nh,omitempty"`
	NLabels    uint8            `binapi:"u8,name=n_labels" json:"n_labels,omitempty"`
	LabelStack [16]FibMplsLabel `binapi:"fib_mpls_label[16],name=label_stack" json:"label_stack,omitempty"`
}

// FibPathNh defines type 'fib_path_nh'.
type FibPathNh struct {
	Address            ip_types.AddressUnion `binapi:"address_union,name=address" json:"address,omitempty"`
	ViaLabel           uint32                `binapi:"u32,name=via_label" json:"via_label,omitempty"`
	ObjID              uint32                `binapi:"u32,name=obj_id" json:"obj_id,omitempty"`
	ClassifyTableIndex uint32                `binapi:"u32,name=classify_table_index" json:"classify_table_index,omitempty"`
}
