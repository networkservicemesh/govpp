// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~94-g123ca956f
// source: plugins/ip_session_redirect.api.json

// Package ip_session_redirect contains generated bindings for API file ip_session_redirect.api.
//
// Contents:
// -  6 messages
package ip_session_redirect

import (
	fib_types "github.com/networkservicemesh/govpp/binapi/fib_types"
	_ "github.com/networkservicemesh/govpp/binapi/interface_types"
	_ "github.com/networkservicemesh/govpp/binapi/ip_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ip_session_redirect"
	APIVersion = "0.3.0"
	VersionCrc = 0xf174f8ba
)

// Add or update a session redirection
//   - table_index - classifier table index
//   - opaque_index - classifier session opaque index
//   - match_len - classifier session match length in bytes (max is 80-bytes)
//   - match - classifier session match
//   - is_punt - true = punted traffic, false = forwarded traffic
//   - n_paths - number of paths
//   - paths - the paths of the redirect
//
// IPSessionRedirectAdd defines message 'ip_session_redirect_add'.
// Deprecated: the message will be removed in the future versions
type IPSessionRedirectAdd struct {
	TableIndex  uint32              `binapi:"u32,name=table_index" json:"table_index,omitempty"`
	MatchLen    uint8               `binapi:"u8,name=match_len" json:"match_len,omitempty"`
	Match       []byte              `binapi:"u8[80],name=match" json:"match,omitempty"`
	OpaqueIndex uint32              `binapi:"u32,name=opaque_index,default=4294967295" json:"opaque_index,omitempty"`
	IsPunt      bool                `binapi:"bool,name=is_punt" json:"is_punt,omitempty"`
	NPaths      uint8               `binapi:"u8,name=n_paths" json:"-"`
	Paths       []fib_types.FibPath `binapi:"fib_path[n_paths],name=paths" json:"paths,omitempty"`
}

func (m *IPSessionRedirectAdd) Reset()               { *m = IPSessionRedirectAdd{} }
func (*IPSessionRedirectAdd) GetMessageName() string { return "ip_session_redirect_add" }
func (*IPSessionRedirectAdd) GetCrcString() string   { return "2f78ffda" }
func (*IPSessionRedirectAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPSessionRedirectAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.TableIndex
	size += 1      // m.MatchLen
	size += 1 * 80 // m.Match
	size += 4      // m.OpaqueIndex
	size += 1      // m.IsPunt
	size += 1      // m.NPaths
	for j1 := 0; j1 < len(m.Paths); j1++ {
		var s1 fib_types.FibPath
		_ = s1
		if j1 < len(m.Paths) {
			s1 = m.Paths[j1]
		}
		size += 4      // s1.SwIfIndex
		size += 4      // s1.TableID
		size += 4      // s1.RpfID
		size += 1      // s1.Weight
		size += 1      // s1.Preference
		size += 4      // s1.Type
		size += 4      // s1.Flags
		size += 4      // s1.Proto
		size += 1 * 16 // s1.Nh.Address
		size += 4      // s1.Nh.ViaLabel
		size += 4      // s1.Nh.ObjID
		size += 4      // s1.Nh.ClassifyTableIndex
		size += 1      // s1.NLabels
		for j2 := 0; j2 < 16; j2++ {
			size += 1 // s1.LabelStack[j2].IsUniform
			size += 4 // s1.LabelStack[j2].Label
			size += 1 // s1.LabelStack[j2].TTL
			size += 1 // s1.LabelStack[j2].Exp
		}
	}
	return size
}
func (m *IPSessionRedirectAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TableIndex)
	buf.EncodeUint8(m.MatchLen)
	buf.EncodeBytes(m.Match, 80)
	buf.EncodeUint32(m.OpaqueIndex)
	buf.EncodeBool(m.IsPunt)
	buf.EncodeUint8(uint8(len(m.Paths)))
	for j0 := 0; j0 < len(m.Paths); j0++ {
		var v0 fib_types.FibPath // Paths
		if j0 < len(m.Paths) {
			v0 = m.Paths[j0]
		}
		buf.EncodeUint32(v0.SwIfIndex)
		buf.EncodeUint32(v0.TableID)
		buf.EncodeUint32(v0.RpfID)
		buf.EncodeUint8(v0.Weight)
		buf.EncodeUint8(v0.Preference)
		buf.EncodeUint32(uint32(v0.Type))
		buf.EncodeUint32(uint32(v0.Flags))
		buf.EncodeUint32(uint32(v0.Proto))
		buf.EncodeBytes(v0.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v0.Nh.ViaLabel)
		buf.EncodeUint32(v0.Nh.ObjID)
		buf.EncodeUint32(v0.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v0.NLabels)
		for j1 := 0; j1 < 16; j1++ {
			buf.EncodeUint8(v0.LabelStack[j1].IsUniform)
			buf.EncodeUint32(v0.LabelStack[j1].Label)
			buf.EncodeUint8(v0.LabelStack[j1].TTL)
			buf.EncodeUint8(v0.LabelStack[j1].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *IPSessionRedirectAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TableIndex = buf.DecodeUint32()
	m.MatchLen = buf.DecodeUint8()
	m.Match = make([]byte, 80)
	copy(m.Match, buf.DecodeBytes(len(m.Match)))
	m.OpaqueIndex = buf.DecodeUint32()
	m.IsPunt = buf.DecodeBool()
	m.NPaths = buf.DecodeUint8()
	m.Paths = make([]fib_types.FibPath, m.NPaths)
	for j0 := 0; j0 < len(m.Paths); j0++ {
		m.Paths[j0].SwIfIndex = buf.DecodeUint32()
		m.Paths[j0].TableID = buf.DecodeUint32()
		m.Paths[j0].RpfID = buf.DecodeUint32()
		m.Paths[j0].Weight = buf.DecodeUint8()
		m.Paths[j0].Preference = buf.DecodeUint8()
		m.Paths[j0].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.Paths[j0].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.Paths[j0].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.Paths[j0].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Paths[j0].Nh.ViaLabel = buf.DecodeUint32()
		m.Paths[j0].Nh.ObjID = buf.DecodeUint32()
		m.Paths[j0].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.Paths[j0].NLabels = buf.DecodeUint8()
		for j1 := 0; j1 < 16; j1++ {
			m.Paths[j0].LabelStack[j1].IsUniform = buf.DecodeUint8()
			m.Paths[j0].LabelStack[j1].Label = buf.DecodeUint32()
			m.Paths[j0].LabelStack[j1].TTL = buf.DecodeUint8()
			m.Paths[j0].LabelStack[j1].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// IPSessionRedirectAddReply defines message 'ip_session_redirect_add_reply'.
// Deprecated: the message will be removed in the future versions
type IPSessionRedirectAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPSessionRedirectAddReply) Reset()               { *m = IPSessionRedirectAddReply{} }
func (*IPSessionRedirectAddReply) GetMessageName() string { return "ip_session_redirect_add_reply" }
func (*IPSessionRedirectAddReply) GetCrcString() string   { return "e8d4e804" }
func (*IPSessionRedirectAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPSessionRedirectAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPSessionRedirectAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPSessionRedirectAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Add or update a session redirection - version 2
//   - table_index - classifier table index
//   - opaque_index - classifier session opaque index
//   - proto - protocol of forwarded packets (default autodetect from path nh)
//   - is_punt - true = punted traffic, false = forwarded traffic
//   - match_len - classifier session match length in bytes (max is 80-bytes)
//   - match - classifier session match
//   - n_paths - number of paths
//   - paths - the paths of the redirect
//
// IPSessionRedirectAddV2 defines message 'ip_session_redirect_add_v2'.
// InProgress: the message form may change in the future versions
type IPSessionRedirectAddV2 struct {
	TableIndex  uint32                   `binapi:"u32,name=table_index" json:"table_index,omitempty"`
	OpaqueIndex uint32                   `binapi:"u32,name=opaque_index,default=4294967295" json:"opaque_index,omitempty"`
	Proto       fib_types.FibPathNhProto `binapi:"fib_path_nh_proto,name=proto,default=4294967295" json:"proto,omitempty"`
	IsPunt      bool                     `binapi:"bool,name=is_punt" json:"is_punt,omitempty"`
	MatchLen    uint8                    `binapi:"u8,name=match_len" json:"match_len,omitempty"`
	Match       []byte                   `binapi:"u8[80],name=match" json:"match,omitempty"`
	NPaths      uint8                    `binapi:"u8,name=n_paths" json:"-"`
	Paths       []fib_types.FibPath      `binapi:"fib_path[n_paths],name=paths" json:"paths,omitempty"`
}

func (m *IPSessionRedirectAddV2) Reset()               { *m = IPSessionRedirectAddV2{} }
func (*IPSessionRedirectAddV2) GetMessageName() string { return "ip_session_redirect_add_v2" }
func (*IPSessionRedirectAddV2) GetCrcString() string   { return "0765f51f" }
func (*IPSessionRedirectAddV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPSessionRedirectAddV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.TableIndex
	size += 4      // m.OpaqueIndex
	size += 4      // m.Proto
	size += 1      // m.IsPunt
	size += 1      // m.MatchLen
	size += 1 * 80 // m.Match
	size += 1      // m.NPaths
	for j1 := 0; j1 < len(m.Paths); j1++ {
		var s1 fib_types.FibPath
		_ = s1
		if j1 < len(m.Paths) {
			s1 = m.Paths[j1]
		}
		size += 4      // s1.SwIfIndex
		size += 4      // s1.TableID
		size += 4      // s1.RpfID
		size += 1      // s1.Weight
		size += 1      // s1.Preference
		size += 4      // s1.Type
		size += 4      // s1.Flags
		size += 4      // s1.Proto
		size += 1 * 16 // s1.Nh.Address
		size += 4      // s1.Nh.ViaLabel
		size += 4      // s1.Nh.ObjID
		size += 4      // s1.Nh.ClassifyTableIndex
		size += 1      // s1.NLabels
		for j2 := 0; j2 < 16; j2++ {
			size += 1 // s1.LabelStack[j2].IsUniform
			size += 4 // s1.LabelStack[j2].Label
			size += 1 // s1.LabelStack[j2].TTL
			size += 1 // s1.LabelStack[j2].Exp
		}
	}
	return size
}
func (m *IPSessionRedirectAddV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TableIndex)
	buf.EncodeUint32(m.OpaqueIndex)
	buf.EncodeUint32(uint32(m.Proto))
	buf.EncodeBool(m.IsPunt)
	buf.EncodeUint8(m.MatchLen)
	buf.EncodeBytes(m.Match, 80)
	buf.EncodeUint8(uint8(len(m.Paths)))
	for j0 := 0; j0 < len(m.Paths); j0++ {
		var v0 fib_types.FibPath // Paths
		if j0 < len(m.Paths) {
			v0 = m.Paths[j0]
		}
		buf.EncodeUint32(v0.SwIfIndex)
		buf.EncodeUint32(v0.TableID)
		buf.EncodeUint32(v0.RpfID)
		buf.EncodeUint8(v0.Weight)
		buf.EncodeUint8(v0.Preference)
		buf.EncodeUint32(uint32(v0.Type))
		buf.EncodeUint32(uint32(v0.Flags))
		buf.EncodeUint32(uint32(v0.Proto))
		buf.EncodeBytes(v0.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v0.Nh.ViaLabel)
		buf.EncodeUint32(v0.Nh.ObjID)
		buf.EncodeUint32(v0.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v0.NLabels)
		for j1 := 0; j1 < 16; j1++ {
			buf.EncodeUint8(v0.LabelStack[j1].IsUniform)
			buf.EncodeUint32(v0.LabelStack[j1].Label)
			buf.EncodeUint8(v0.LabelStack[j1].TTL)
			buf.EncodeUint8(v0.LabelStack[j1].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *IPSessionRedirectAddV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TableIndex = buf.DecodeUint32()
	m.OpaqueIndex = buf.DecodeUint32()
	m.Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
	m.IsPunt = buf.DecodeBool()
	m.MatchLen = buf.DecodeUint8()
	m.Match = make([]byte, 80)
	copy(m.Match, buf.DecodeBytes(len(m.Match)))
	m.NPaths = buf.DecodeUint8()
	m.Paths = make([]fib_types.FibPath, m.NPaths)
	for j0 := 0; j0 < len(m.Paths); j0++ {
		m.Paths[j0].SwIfIndex = buf.DecodeUint32()
		m.Paths[j0].TableID = buf.DecodeUint32()
		m.Paths[j0].RpfID = buf.DecodeUint32()
		m.Paths[j0].Weight = buf.DecodeUint8()
		m.Paths[j0].Preference = buf.DecodeUint8()
		m.Paths[j0].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.Paths[j0].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.Paths[j0].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.Paths[j0].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Paths[j0].Nh.ViaLabel = buf.DecodeUint32()
		m.Paths[j0].Nh.ObjID = buf.DecodeUint32()
		m.Paths[j0].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.Paths[j0].NLabels = buf.DecodeUint8()
		for j1 := 0; j1 < 16; j1++ {
			m.Paths[j0].LabelStack[j1].IsUniform = buf.DecodeUint8()
			m.Paths[j0].LabelStack[j1].Label = buf.DecodeUint32()
			m.Paths[j0].LabelStack[j1].TTL = buf.DecodeUint8()
			m.Paths[j0].LabelStack[j1].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// IPSessionRedirectAddV2Reply defines message 'ip_session_redirect_add_v2_reply'.
// InProgress: the message form may change in the future versions
type IPSessionRedirectAddV2Reply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPSessionRedirectAddV2Reply) Reset() { *m = IPSessionRedirectAddV2Reply{} }
func (*IPSessionRedirectAddV2Reply) GetMessageName() string {
	return "ip_session_redirect_add_v2_reply"
}
func (*IPSessionRedirectAddV2Reply) GetCrcString() string { return "e8d4e804" }
func (*IPSessionRedirectAddV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPSessionRedirectAddV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPSessionRedirectAddV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPSessionRedirectAddV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Delete a session redirection
//   - table_index - classifier table index
//   - match_len - classifier session match length in bytes (max is 80-bytes)
//   - match - classifier session match
//
// IPSessionRedirectDel defines message 'ip_session_redirect_del'.
// InProgress: the message form may change in the future versions
type IPSessionRedirectDel struct {
	TableIndex uint32 `binapi:"u32,name=table_index" json:"table_index,omitempty"`
	MatchLen   uint8  `binapi:"u8,name=match_len" json:"-"`
	Match      []byte `binapi:"u8[match_len],name=match" json:"match,omitempty"`
}

func (m *IPSessionRedirectDel) Reset()               { *m = IPSessionRedirectDel{} }
func (*IPSessionRedirectDel) GetMessageName() string { return "ip_session_redirect_del" }
func (*IPSessionRedirectDel) GetCrcString() string   { return "fb643388" }
func (*IPSessionRedirectDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPSessionRedirectDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                // m.TableIndex
	size += 1                // m.MatchLen
	size += 1 * len(m.Match) // m.Match
	return size
}
func (m *IPSessionRedirectDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TableIndex)
	buf.EncodeUint8(uint8(len(m.Match)))
	buf.EncodeBytes(m.Match, 0)
	return buf.Bytes(), nil
}
func (m *IPSessionRedirectDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TableIndex = buf.DecodeUint32()
	m.MatchLen = buf.DecodeUint8()
	m.Match = make([]byte, m.MatchLen)
	copy(m.Match, buf.DecodeBytes(len(m.Match)))
	return nil
}

// IPSessionRedirectDelReply defines message 'ip_session_redirect_del_reply'.
// InProgress: the message form may change in the future versions
type IPSessionRedirectDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPSessionRedirectDelReply) Reset()               { *m = IPSessionRedirectDelReply{} }
func (*IPSessionRedirectDelReply) GetMessageName() string { return "ip_session_redirect_del_reply" }
func (*IPSessionRedirectDelReply) GetCrcString() string   { return "e8d4e804" }
func (*IPSessionRedirectDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPSessionRedirectDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPSessionRedirectDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPSessionRedirectDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ip_session_redirect_binapi_init() }
func file_ip_session_redirect_binapi_init() {
	api.RegisterMessage((*IPSessionRedirectAdd)(nil), "ip_session_redirect_add_2f78ffda")
	api.RegisterMessage((*IPSessionRedirectAddReply)(nil), "ip_session_redirect_add_reply_e8d4e804")
	api.RegisterMessage((*IPSessionRedirectAddV2)(nil), "ip_session_redirect_add_v2_0765f51f")
	api.RegisterMessage((*IPSessionRedirectAddV2Reply)(nil), "ip_session_redirect_add_v2_reply_e8d4e804")
	api.RegisterMessage((*IPSessionRedirectDel)(nil), "ip_session_redirect_del_fb643388")
	api.RegisterMessage((*IPSessionRedirectDelReply)(nil), "ip_session_redirect_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IPSessionRedirectAdd)(nil),
		(*IPSessionRedirectAddReply)(nil),
		(*IPSessionRedirectAddV2)(nil),
		(*IPSessionRedirectAddV2Reply)(nil),
		(*IPSessionRedirectDel)(nil),
		(*IPSessionRedirectDelReply)(nil),
	}
}
