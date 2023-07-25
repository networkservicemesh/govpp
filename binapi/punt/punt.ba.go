// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~162-g1765f014b
// source: core/punt.api.json

// Package punt contains generated bindings for API file punt.api.
//
// Contents:
// -  1 enum
// -  5 structs
// -  1 union
// - 10 messages
package punt

import (
	"strconv"

	ip_types "github.com/networkservicemesh/govpp/binapi/ip_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "punt"
	APIVersion = "2.2.1"
	VersionCrc = 0x24d11934
)

// PuntType defines enum 'punt_type'.
type PuntType uint32

const (
	PUNT_API_TYPE_L4        PuntType = 0
	PUNT_API_TYPE_IP_PROTO  PuntType = 1
	PUNT_API_TYPE_EXCEPTION PuntType = 2
)

var (
	PuntType_name = map[uint32]string{
		0: "PUNT_API_TYPE_L4",
		1: "PUNT_API_TYPE_IP_PROTO",
		2: "PUNT_API_TYPE_EXCEPTION",
	}
	PuntType_value = map[string]uint32{
		"PUNT_API_TYPE_L4":        0,
		"PUNT_API_TYPE_IP_PROTO":  1,
		"PUNT_API_TYPE_EXCEPTION": 2,
	}
)

func (x PuntType) String() string {
	s, ok := PuntType_name[uint32(x)]
	if ok {
		return s
	}
	return "PuntType(" + strconv.Itoa(int(x)) + ")"
}

// Punt defines type 'punt'.
type Punt struct {
	Type PuntType  `binapi:"punt_type,name=type" json:"type,omitempty"`
	Punt PuntUnion `binapi:"punt_union,name=punt" json:"punt,omitempty"`
}

// PuntException defines type 'punt_exception'.
type PuntException struct {
	ID uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

// PuntIPProto defines type 'punt_ip_proto'.
type PuntIPProto struct {
	Af       ip_types.AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	Protocol ip_types.IPProto       `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
}

// PuntL4 defines type 'punt_l4'.
type PuntL4 struct {
	Af       ip_types.AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	Protocol ip_types.IPProto       `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
	Port     uint16                 `binapi:"u16,name=port" json:"port,omitempty"`
}

// PuntReason defines type 'punt_reason'.
type PuntReason struct {
	ID   uint32 `binapi:"u32,name=id" json:"id,omitempty"`
	Name string `binapi:"string[],name=name" json:"name,omitempty"`
}

// PuntUnion defines union 'punt_union'.
type PuntUnion struct {
	// PuntUnion can be one of:
	// - Exception *PuntException
	// - L4 *PuntL4
	// - IPProto *PuntIPProto
	XXX_UnionData [4]byte
}

func PuntUnionException(a PuntException) (u PuntUnion) {
	u.SetException(a)
	return
}
func (u *PuntUnion) SetException(a PuntException) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint32(a.ID)
}
func (u *PuntUnion) GetException() (a PuntException) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint32()
	return
}

func PuntUnionL4(a PuntL4) (u PuntUnion) {
	u.SetL4(a)
	return
}
func (u *PuntUnion) SetL4(a PuntL4) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint8(uint8(a.Af))
	buf.EncodeUint8(uint8(a.Protocol))
	buf.EncodeUint16(a.Port)
}
func (u *PuntUnion) GetL4() (a PuntL4) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Af = ip_types.AddressFamily(buf.DecodeUint8())
	a.Protocol = ip_types.IPProto(buf.DecodeUint8())
	a.Port = buf.DecodeUint16()
	return
}

func PuntUnionIPProto(a PuntIPProto) (u PuntUnion) {
	u.SetIPProto(a)
	return
}
func (u *PuntUnion) SetIPProto(a PuntIPProto) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint8(uint8(a.Af))
	buf.EncodeUint8(uint8(a.Protocol))
}
func (u *PuntUnion) GetIPProto() (a PuntIPProto) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Af = ip_types.AddressFamily(buf.DecodeUint8())
	a.Protocol = ip_types.IPProto(buf.DecodeUint8())
	return
}

// PuntReasonDetails defines message 'punt_reason_details'.
type PuntReasonDetails struct {
	Reason PuntReason `binapi:"punt_reason,name=reason" json:"reason,omitempty"`
}

func (m *PuntReasonDetails) Reset()               { *m = PuntReasonDetails{} }
func (*PuntReasonDetails) GetMessageName() string { return "punt_reason_details" }
func (*PuntReasonDetails) GetCrcString() string   { return "2c9d4a40" }
func (*PuntReasonDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PuntReasonDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.Reason.ID
	size += 4 + len(m.Reason.Name) // m.Reason.Name
	return size
}
func (m *PuntReasonDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Reason.ID)
	buf.EncodeString(m.Reason.Name, 0)
	return buf.Bytes(), nil
}
func (m *PuntReasonDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Reason.ID = buf.DecodeUint32()
	m.Reason.Name = buf.DecodeString(0)
	return nil
}

// Dump all or one of the exception punt reasons
// *   - - If the string is not set punt dump all reasons
// *            else dump only the one specified
// PuntReasonDump defines message 'punt_reason_dump'.
type PuntReasonDump struct {
	Reason PuntReason `binapi:"punt_reason,name=reason" json:"reason,omitempty"`
}

func (m *PuntReasonDump) Reset()               { *m = PuntReasonDump{} }
func (*PuntReasonDump) GetMessageName() string { return "punt_reason_dump" }
func (*PuntReasonDump) GetCrcString() string   { return "5c0dd4fe" }
func (*PuntReasonDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PuntReasonDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.Reason.ID
	size += 4 + len(m.Reason.Name) // m.Reason.Name
	return size
}
func (m *PuntReasonDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Reason.ID)
	buf.EncodeString(m.Reason.Name, 0)
	return buf.Bytes(), nil
}
func (m *PuntReasonDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Reason.ID = buf.DecodeUint32()
	m.Reason.Name = buf.DecodeString(0)
	return nil
}

// PuntSocketDeregister defines message 'punt_socket_deregister'.
type PuntSocketDeregister struct {
	Punt Punt `binapi:"punt,name=punt" json:"punt,omitempty"`
}

func (m *PuntSocketDeregister) Reset()               { *m = PuntSocketDeregister{} }
func (*PuntSocketDeregister) GetMessageName() string { return "punt_socket_deregister" }
func (*PuntSocketDeregister) GetCrcString() string   { return "75afa766" }
func (*PuntSocketDeregister) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PuntSocketDeregister) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Punt.Type
	size += 1 * 4 // m.Punt.Punt
	return size
}
func (m *PuntSocketDeregister) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Punt.Type))
	buf.EncodeBytes(m.Punt.Punt.XXX_UnionData[:], 4)
	return buf.Bytes(), nil
}
func (m *PuntSocketDeregister) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Punt.Type = PuntType(buf.DecodeUint32())
	copy(m.Punt.Punt.XXX_UnionData[:], buf.DecodeBytes(4))
	return nil
}

// PuntSocketDeregisterReply defines message 'punt_socket_deregister_reply'.
type PuntSocketDeregisterReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PuntSocketDeregisterReply) Reset()               { *m = PuntSocketDeregisterReply{} }
func (*PuntSocketDeregisterReply) GetMessageName() string { return "punt_socket_deregister_reply" }
func (*PuntSocketDeregisterReply) GetCrcString() string   { return "e8d4e804" }
func (*PuntSocketDeregisterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PuntSocketDeregisterReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PuntSocketDeregisterReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PuntSocketDeregisterReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PuntSocketDetails defines message 'punt_socket_details'.
type PuntSocketDetails struct {
	Punt     Punt   `binapi:"punt,name=punt" json:"punt,omitempty"`
	Pathname string `binapi:"string[108],name=pathname" json:"pathname,omitempty"`
}

func (m *PuntSocketDetails) Reset()               { *m = PuntSocketDetails{} }
func (*PuntSocketDetails) GetMessageName() string { return "punt_socket_details" }
func (*PuntSocketDetails) GetCrcString() string   { return "330466e4" }
func (*PuntSocketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PuntSocketDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Punt.Type
	size += 1 * 4 // m.Punt.Punt
	size += 108   // m.Pathname
	return size
}
func (m *PuntSocketDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Punt.Type))
	buf.EncodeBytes(m.Punt.Punt.XXX_UnionData[:], 4)
	buf.EncodeString(m.Pathname, 108)
	return buf.Bytes(), nil
}
func (m *PuntSocketDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Punt.Type = PuntType(buf.DecodeUint32())
	copy(m.Punt.Punt.XXX_UnionData[:], buf.DecodeBytes(4))
	m.Pathname = buf.DecodeString(108)
	return nil
}

// PuntSocketDump defines message 'punt_socket_dump'.
type PuntSocketDump struct {
	Type PuntType `binapi:"punt_type,name=type" json:"type,omitempty"`
}

func (m *PuntSocketDump) Reset()               { *m = PuntSocketDump{} }
func (*PuntSocketDump) GetMessageName() string { return "punt_socket_dump" }
func (*PuntSocketDump) GetCrcString() string   { return "916fb004" }
func (*PuntSocketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PuntSocketDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Type
	return size
}
func (m *PuntSocketDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Type))
	return buf.Bytes(), nil
}
func (m *PuntSocketDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Type = PuntType(buf.DecodeUint32())
	return nil
}

// Punt traffic to the host via socket
//   - header_version - expected meta data header version (currently 1)
//   - punt - punt definition
//
// PuntSocketRegister defines message 'punt_socket_register'.
type PuntSocketRegister struct {
	HeaderVersion uint32 `binapi:"u32,name=header_version" json:"header_version,omitempty"`
	Punt          Punt   `binapi:"punt,name=punt" json:"punt,omitempty"`
	Pathname      string `binapi:"string[108],name=pathname" json:"pathname,omitempty"`
}

func (m *PuntSocketRegister) Reset()               { *m = PuntSocketRegister{} }
func (*PuntSocketRegister) GetMessageName() string { return "punt_socket_register" }
func (*PuntSocketRegister) GetCrcString() string   { return "7875badb" }
func (*PuntSocketRegister) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PuntSocketRegister) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.HeaderVersion
	size += 4     // m.Punt.Type
	size += 1 * 4 // m.Punt.Punt
	size += 108   // m.Pathname
	return size
}
func (m *PuntSocketRegister) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.HeaderVersion)
	buf.EncodeUint32(uint32(m.Punt.Type))
	buf.EncodeBytes(m.Punt.Punt.XXX_UnionData[:], 4)
	buf.EncodeString(m.Pathname, 108)
	return buf.Bytes(), nil
}
func (m *PuntSocketRegister) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HeaderVersion = buf.DecodeUint32()
	m.Punt.Type = PuntType(buf.DecodeUint32())
	copy(m.Punt.Punt.XXX_UnionData[:], buf.DecodeBytes(4))
	m.Pathname = buf.DecodeString(108)
	return nil
}

// PuntSocketRegisterReply defines message 'punt_socket_register_reply'.
type PuntSocketRegisterReply struct {
	Retval   int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Pathname string `binapi:"string[108],name=pathname" json:"pathname,omitempty"`
}

func (m *PuntSocketRegisterReply) Reset()               { *m = PuntSocketRegisterReply{} }
func (*PuntSocketRegisterReply) GetMessageName() string { return "punt_socket_register_reply" }
func (*PuntSocketRegisterReply) GetCrcString() string   { return "bd30ae90" }
func (*PuntSocketRegisterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PuntSocketRegisterReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.Retval
	size += 108 // m.Pathname
	return size
}
func (m *PuntSocketRegisterReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.Pathname, 108)
	return buf.Bytes(), nil
}
func (m *PuntSocketRegisterReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Pathname = buf.DecodeString(108)
	return nil
}

// Punt traffic to the host
//   - is_add - add punt if non-zero, else delete
//   - punt - punt definition, only UDP (0x11) is supported
//
// SetPunt defines message 'set_punt'.
type SetPunt struct {
	IsAdd bool `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Punt  Punt `binapi:"punt,name=punt" json:"punt,omitempty"`
}

func (m *SetPunt) Reset()               { *m = SetPunt{} }
func (*SetPunt) GetMessageName() string { return "set_punt" }
func (*SetPunt) GetCrcString() string   { return "47d0e347" }
func (*SetPunt) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SetPunt) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsAdd
	size += 4     // m.Punt.Type
	size += 1 * 4 // m.Punt.Punt
	return size
}
func (m *SetPunt) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(uint32(m.Punt.Type))
	buf.EncodeBytes(m.Punt.Punt.XXX_UnionData[:], 4)
	return buf.Bytes(), nil
}
func (m *SetPunt) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Punt.Type = PuntType(buf.DecodeUint32())
	copy(m.Punt.Punt.XXX_UnionData[:], buf.DecodeBytes(4))
	return nil
}

// SetPuntReply defines message 'set_punt_reply'.
type SetPuntReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SetPuntReply) Reset()               { *m = SetPuntReply{} }
func (*SetPuntReply) GetMessageName() string { return "set_punt_reply" }
func (*SetPuntReply) GetCrcString() string   { return "e8d4e804" }
func (*SetPuntReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SetPuntReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SetPuntReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SetPuntReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_punt_binapi_init() }
func file_punt_binapi_init() {
	api.RegisterMessage((*PuntReasonDetails)(nil), "punt_reason_details_2c9d4a40")
	api.RegisterMessage((*PuntReasonDump)(nil), "punt_reason_dump_5c0dd4fe")
	api.RegisterMessage((*PuntSocketDeregister)(nil), "punt_socket_deregister_75afa766")
	api.RegisterMessage((*PuntSocketDeregisterReply)(nil), "punt_socket_deregister_reply_e8d4e804")
	api.RegisterMessage((*PuntSocketDetails)(nil), "punt_socket_details_330466e4")
	api.RegisterMessage((*PuntSocketDump)(nil), "punt_socket_dump_916fb004")
	api.RegisterMessage((*PuntSocketRegister)(nil), "punt_socket_register_7875badb")
	api.RegisterMessage((*PuntSocketRegisterReply)(nil), "punt_socket_register_reply_bd30ae90")
	api.RegisterMessage((*SetPunt)(nil), "set_punt_47d0e347")
	api.RegisterMessage((*SetPuntReply)(nil), "set_punt_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PuntReasonDetails)(nil),
		(*PuntReasonDump)(nil),
		(*PuntSocketDeregister)(nil),
		(*PuntSocketDeregisterReply)(nil),
		(*PuntSocketDetails)(nil),
		(*PuntSocketDump)(nil),
		(*PuntSocketRegister)(nil),
		(*PuntSocketRegisterReply)(nil),
		(*SetPunt)(nil),
		(*SetPuntReply)(nil),
	}
}
