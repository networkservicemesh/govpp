// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~94-g123ca956f
// source: core/ipip.api.json

// Package ipip contains generated bindings for API file ipip.api.
//
// Contents:
// -  1 struct
// - 10 messages
package ipip

import (
	interface_types "github.com/networkservicemesh/govpp/binapi/interface_types"
	ip_types "github.com/networkservicemesh/govpp/binapi/ip_types"
	tunnel_types "github.com/networkservicemesh/govpp/binapi/tunnel_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ipip"
	APIVersion = "2.0.2"
	VersionCrc = 0x4609caba
)

// IpipTunnel defines type 'ipip_tunnel'.
type IpipTunnel struct {
	Instance  uint32                             `binapi:"u32,name=instance" json:"instance,omitempty"`
	Src       ip_types.Address                   `binapi:"address,name=src" json:"src,omitempty"`
	Dst       ip_types.Address                   `binapi:"address,name=dst" json:"dst,omitempty"`
	SwIfIndex interface_types.InterfaceIndex     `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	TableID   uint32                             `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	Flags     tunnel_types.TunnelEncapDecapFlags `binapi:"tunnel_encap_decap_flags,name=flags" json:"flags,omitempty"`
	Mode      tunnel_types.TunnelMode            `binapi:"tunnel_mode,name=mode" json:"mode,omitempty"`
	Dscp      ip_types.IPDscp                    `binapi:"ip_dscp,name=dscp" json:"dscp,omitempty"`
}

// * Create an IPv4 over IPv6 automatic tunnel (6RD)
// Ipip6rdAddTunnel defines message 'ipip_6rd_add_tunnel'.
type Ipip6rdAddTunnel struct {
	IP6TableID    uint32              `binapi:"u32,name=ip6_table_id" json:"ip6_table_id,omitempty"`
	IP4TableID    uint32              `binapi:"u32,name=ip4_table_id" json:"ip4_table_id,omitempty"`
	IP6Prefix     ip_types.IP6Prefix  `binapi:"ip6_prefix,name=ip6_prefix" json:"ip6_prefix,omitempty"`
	IP4Prefix     ip_types.IP4Prefix  `binapi:"ip4_prefix,name=ip4_prefix" json:"ip4_prefix,omitempty"`
	IP4Src        ip_types.IP4Address `binapi:"ip4_address,name=ip4_src" json:"ip4_src,omitempty"`
	SecurityCheck bool                `binapi:"bool,name=security_check" json:"security_check,omitempty"`
	TcTos         uint8               `binapi:"u8,name=tc_tos" json:"tc_tos,omitempty"`
}

func (m *Ipip6rdAddTunnel) Reset()               { *m = Ipip6rdAddTunnel{} }
func (*Ipip6rdAddTunnel) GetMessageName() string { return "ipip_6rd_add_tunnel" }
func (*Ipip6rdAddTunnel) GetCrcString() string   { return "b9ec1863" }
func (*Ipip6rdAddTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Ipip6rdAddTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.IP6TableID
	size += 4      // m.IP4TableID
	size += 1 * 16 // m.IP6Prefix.Address
	size += 1      // m.IP6Prefix.Len
	size += 1 * 4  // m.IP4Prefix.Address
	size += 1      // m.IP4Prefix.Len
	size += 1 * 4  // m.IP4Src
	size += 1      // m.SecurityCheck
	size += 1      // m.TcTos
	return size
}
func (m *Ipip6rdAddTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.IP6TableID)
	buf.EncodeUint32(m.IP4TableID)
	buf.EncodeBytes(m.IP6Prefix.Address[:], 16)
	buf.EncodeUint8(m.IP6Prefix.Len)
	buf.EncodeBytes(m.IP4Prefix.Address[:], 4)
	buf.EncodeUint8(m.IP4Prefix.Len)
	buf.EncodeBytes(m.IP4Src[:], 4)
	buf.EncodeBool(m.SecurityCheck)
	buf.EncodeUint8(m.TcTos)
	return buf.Bytes(), nil
}
func (m *Ipip6rdAddTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IP6TableID = buf.DecodeUint32()
	m.IP4TableID = buf.DecodeUint32()
	copy(m.IP6Prefix.Address[:], buf.DecodeBytes(16))
	m.IP6Prefix.Len = buf.DecodeUint8()
	copy(m.IP4Prefix.Address[:], buf.DecodeBytes(4))
	m.IP4Prefix.Len = buf.DecodeUint8()
	copy(m.IP4Src[:], buf.DecodeBytes(4))
	m.SecurityCheck = buf.DecodeBool()
	m.TcTos = buf.DecodeUint8()
	return nil
}

// Ipip6rdAddTunnelReply defines message 'ipip_6rd_add_tunnel_reply'.
type Ipip6rdAddTunnelReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *Ipip6rdAddTunnelReply) Reset()               { *m = Ipip6rdAddTunnelReply{} }
func (*Ipip6rdAddTunnelReply) GetMessageName() string { return "ipip_6rd_add_tunnel_reply" }
func (*Ipip6rdAddTunnelReply) GetCrcString() string   { return "5383d31f" }
func (*Ipip6rdAddTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Ipip6rdAddTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *Ipip6rdAddTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *Ipip6rdAddTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// * Delete an IPv4 over IPv6 automatic tunnel (6RD)
// Ipip6rdDelTunnel defines message 'ipip_6rd_del_tunnel'.
type Ipip6rdDelTunnel struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *Ipip6rdDelTunnel) Reset()               { *m = Ipip6rdDelTunnel{} }
func (*Ipip6rdDelTunnel) GetMessageName() string { return "ipip_6rd_del_tunnel" }
func (*Ipip6rdDelTunnel) GetCrcString() string   { return "f9e6675e" }
func (*Ipip6rdDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Ipip6rdDelTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *Ipip6rdDelTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *Ipip6rdDelTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Ipip6rdDelTunnelReply defines message 'ipip_6rd_del_tunnel_reply'.
type Ipip6rdDelTunnelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *Ipip6rdDelTunnelReply) Reset()               { *m = Ipip6rdDelTunnelReply{} }
func (*Ipip6rdDelTunnelReply) GetMessageName() string { return "ipip_6rd_del_tunnel_reply" }
func (*Ipip6rdDelTunnelReply) GetCrcString() string   { return "e8d4e804" }
func (*Ipip6rdDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Ipip6rdDelTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *Ipip6rdDelTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *Ipip6rdDelTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// * Create an IP{v4,v6} over IP{v4,v6} tunnel.
// IpipAddTunnel defines message 'ipip_add_tunnel'.
type IpipAddTunnel struct {
	Tunnel IpipTunnel `binapi:"ipip_tunnel,name=tunnel" json:"tunnel,omitempty"`
}

func (m *IpipAddTunnel) Reset()               { *m = IpipAddTunnel{} }
func (*IpipAddTunnel) GetMessageName() string { return "ipip_add_tunnel" }
func (*IpipAddTunnel) GetCrcString() string   { return "2ac399f5" }
func (*IpipAddTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpipAddTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Tunnel.Instance
	size += 1      // m.Tunnel.Src.Af
	size += 1 * 16 // m.Tunnel.Src.Un
	size += 1      // m.Tunnel.Dst.Af
	size += 1 * 16 // m.Tunnel.Dst.Un
	size += 4      // m.Tunnel.SwIfIndex
	size += 4      // m.Tunnel.TableID
	size += 1      // m.Tunnel.Flags
	size += 1      // m.Tunnel.Mode
	size += 1      // m.Tunnel.Dscp
	return size
}
func (m *IpipAddTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Tunnel.Instance)
	buf.EncodeUint8(uint8(m.Tunnel.Src.Af))
	buf.EncodeBytes(m.Tunnel.Src.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Tunnel.Dst.Af))
	buf.EncodeBytes(m.Tunnel.Dst.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.Tunnel.SwIfIndex))
	buf.EncodeUint32(m.Tunnel.TableID)
	buf.EncodeUint8(uint8(m.Tunnel.Flags))
	buf.EncodeUint8(uint8(m.Tunnel.Mode))
	buf.EncodeUint8(uint8(m.Tunnel.Dscp))
	return buf.Bytes(), nil
}
func (m *IpipAddTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Tunnel.Instance = buf.DecodeUint32()
	m.Tunnel.Src.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Src.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.Dst.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Dst.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Tunnel.TableID = buf.DecodeUint32()
	m.Tunnel.Flags = tunnel_types.TunnelEncapDecapFlags(buf.DecodeUint8())
	m.Tunnel.Mode = tunnel_types.TunnelMode(buf.DecodeUint8())
	m.Tunnel.Dscp = ip_types.IPDscp(buf.DecodeUint8())
	return nil
}

// IpipAddTunnelReply defines message 'ipip_add_tunnel_reply'.
type IpipAddTunnelReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IpipAddTunnelReply) Reset()               { *m = IpipAddTunnelReply{} }
func (*IpipAddTunnelReply) GetMessageName() string { return "ipip_add_tunnel_reply" }
func (*IpipAddTunnelReply) GetCrcString() string   { return "5383d31f" }
func (*IpipAddTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpipAddTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *IpipAddTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IpipAddTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// * Delete an IP{v4,v6} over IP{v4,v6} tunnel.
// IpipDelTunnel defines message 'ipip_del_tunnel'.
type IpipDelTunnel struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IpipDelTunnel) Reset()               { *m = IpipDelTunnel{} }
func (*IpipDelTunnel) GetMessageName() string { return "ipip_del_tunnel" }
func (*IpipDelTunnel) GetCrcString() string   { return "f9e6675e" }
func (*IpipDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpipDelTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *IpipDelTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IpipDelTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IpipDelTunnelReply defines message 'ipip_del_tunnel_reply'.
type IpipDelTunnelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IpipDelTunnelReply) Reset()               { *m = IpipDelTunnelReply{} }
func (*IpipDelTunnelReply) GetMessageName() string { return "ipip_del_tunnel_reply" }
func (*IpipDelTunnelReply) GetCrcString() string   { return "e8d4e804" }
func (*IpipDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpipDelTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IpipDelTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IpipDelTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IpipTunnelDetails defines message 'ipip_tunnel_details'.
type IpipTunnelDetails struct {
	Tunnel IpipTunnel `binapi:"ipip_tunnel,name=tunnel" json:"tunnel,omitempty"`
}

func (m *IpipTunnelDetails) Reset()               { *m = IpipTunnelDetails{} }
func (*IpipTunnelDetails) GetMessageName() string { return "ipip_tunnel_details" }
func (*IpipTunnelDetails) GetCrcString() string   { return "d31cb34e" }
func (*IpipTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpipTunnelDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Tunnel.Instance
	size += 1      // m.Tunnel.Src.Af
	size += 1 * 16 // m.Tunnel.Src.Un
	size += 1      // m.Tunnel.Dst.Af
	size += 1 * 16 // m.Tunnel.Dst.Un
	size += 4      // m.Tunnel.SwIfIndex
	size += 4      // m.Tunnel.TableID
	size += 1      // m.Tunnel.Flags
	size += 1      // m.Tunnel.Mode
	size += 1      // m.Tunnel.Dscp
	return size
}
func (m *IpipTunnelDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Tunnel.Instance)
	buf.EncodeUint8(uint8(m.Tunnel.Src.Af))
	buf.EncodeBytes(m.Tunnel.Src.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Tunnel.Dst.Af))
	buf.EncodeBytes(m.Tunnel.Dst.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.Tunnel.SwIfIndex))
	buf.EncodeUint32(m.Tunnel.TableID)
	buf.EncodeUint8(uint8(m.Tunnel.Flags))
	buf.EncodeUint8(uint8(m.Tunnel.Mode))
	buf.EncodeUint8(uint8(m.Tunnel.Dscp))
	return buf.Bytes(), nil
}
func (m *IpipTunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Tunnel.Instance = buf.DecodeUint32()
	m.Tunnel.Src.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Src.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.Dst.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Dst.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Tunnel.TableID = buf.DecodeUint32()
	m.Tunnel.Flags = tunnel_types.TunnelEncapDecapFlags(buf.DecodeUint8())
	m.Tunnel.Mode = tunnel_types.TunnelMode(buf.DecodeUint8())
	m.Tunnel.Dscp = ip_types.IPDscp(buf.DecodeUint8())
	return nil
}

// * List all IPIP tunnels
// IpipTunnelDump defines message 'ipip_tunnel_dump'.
type IpipTunnelDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IpipTunnelDump) Reset()               { *m = IpipTunnelDump{} }
func (*IpipTunnelDump) GetMessageName() string { return "ipip_tunnel_dump" }
func (*IpipTunnelDump) GetCrcString() string   { return "f9e6675e" }
func (*IpipTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpipTunnelDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *IpipTunnelDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IpipTunnelDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_ipip_binapi_init() }
func file_ipip_binapi_init() {
	api.RegisterMessage((*Ipip6rdAddTunnel)(nil), "ipip_6rd_add_tunnel_b9ec1863")
	api.RegisterMessage((*Ipip6rdAddTunnelReply)(nil), "ipip_6rd_add_tunnel_reply_5383d31f")
	api.RegisterMessage((*Ipip6rdDelTunnel)(nil), "ipip_6rd_del_tunnel_f9e6675e")
	api.RegisterMessage((*Ipip6rdDelTunnelReply)(nil), "ipip_6rd_del_tunnel_reply_e8d4e804")
	api.RegisterMessage((*IpipAddTunnel)(nil), "ipip_add_tunnel_2ac399f5")
	api.RegisterMessage((*IpipAddTunnelReply)(nil), "ipip_add_tunnel_reply_5383d31f")
	api.RegisterMessage((*IpipDelTunnel)(nil), "ipip_del_tunnel_f9e6675e")
	api.RegisterMessage((*IpipDelTunnelReply)(nil), "ipip_del_tunnel_reply_e8d4e804")
	api.RegisterMessage((*IpipTunnelDetails)(nil), "ipip_tunnel_details_d31cb34e")
	api.RegisterMessage((*IpipTunnelDump)(nil), "ipip_tunnel_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Ipip6rdAddTunnel)(nil),
		(*Ipip6rdAddTunnelReply)(nil),
		(*Ipip6rdDelTunnel)(nil),
		(*Ipip6rdDelTunnelReply)(nil),
		(*IpipAddTunnel)(nil),
		(*IpipAddTunnelReply)(nil),
		(*IpipDelTunnel)(nil),
		(*IpipDelTunnelReply)(nil),
		(*IpipTunnelDetails)(nil),
		(*IpipTunnelDump)(nil),
	}
}
