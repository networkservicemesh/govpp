// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~162-g1765f014b
// source: plugins/idpf.api.json

// Package idpf contains generated bindings for API file idpf.api.
//
// Contents:
// -  4 messages
package idpf

import (
	interface_types "github.com/networkservicemesh/govpp/binapi/interface_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "idpf"
	APIVersion = "1.0.0"
	VersionCrc = 0x48384ea8
)

// - client_index - opaque cookie to identify the sender
//   - pci_addr - pci address as unsigned 32bit integer:
//     0-15 domain, 16-23 bus, 24-28 slot, 29-31 function
//     ddddddddddddddddbbbbbbbbsssssfff
//   - rxq_num - number of receive queues
//   - rxq_size - receive queue size
//   - txq_size - transmit queue size
//
// IdpfCreate defines message 'idpf_create'.
type IdpfCreate struct {
	PciAddr    uint32 `binapi:"u32,name=pci_addr" json:"pci_addr,omitempty"`
	RxqSingle  uint16 `binapi:"u16,name=rxq_single" json:"rxq_single,omitempty"`
	TxqSingle  uint16 `binapi:"u16,name=txq_single" json:"txq_single,omitempty"`
	RxqNum     uint16 `binapi:"u16,name=rxq_num" json:"rxq_num,omitempty"`
	TxqNum     uint16 `binapi:"u16,name=txq_num" json:"txq_num,omitempty"`
	RxqSize    uint16 `binapi:"u16,name=rxq_size" json:"rxq_size,omitempty"`
	TxqSize    uint16 `binapi:"u16,name=txq_size" json:"txq_size,omitempty"`
	ReqVportNb uint16 `binapi:"u16,name=req_vport_nb" json:"req_vport_nb,omitempty"`
}

func (m *IdpfCreate) Reset()               { *m = IdpfCreate{} }
func (*IdpfCreate) GetMessageName() string { return "idpf_create" }
func (*IdpfCreate) GetCrcString() string   { return "2ba86d91" }
func (*IdpfCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IdpfCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.PciAddr
	size += 2 // m.RxqSingle
	size += 2 // m.TxqSingle
	size += 2 // m.RxqNum
	size += 2 // m.TxqNum
	size += 2 // m.RxqSize
	size += 2 // m.TxqSize
	size += 2 // m.ReqVportNb
	return size
}
func (m *IdpfCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PciAddr)
	buf.EncodeUint16(m.RxqSingle)
	buf.EncodeUint16(m.TxqSingle)
	buf.EncodeUint16(m.RxqNum)
	buf.EncodeUint16(m.TxqNum)
	buf.EncodeUint16(m.RxqSize)
	buf.EncodeUint16(m.TxqSize)
	buf.EncodeUint16(m.ReqVportNb)
	return buf.Bytes(), nil
}
func (m *IdpfCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PciAddr = buf.DecodeUint32()
	m.RxqSingle = buf.DecodeUint16()
	m.TxqSingle = buf.DecodeUint16()
	m.RxqNum = buf.DecodeUint16()
	m.TxqNum = buf.DecodeUint16()
	m.RxqSize = buf.DecodeUint16()
	m.TxqSize = buf.DecodeUint16()
	m.ReqVportNb = buf.DecodeUint16()
	return nil
}

// - context - sender context, to match reply w/ request
//   - retval - return value for request
//   - sw_if_index - software index for the new idpf interface
//
// IdpfCreateReply defines message 'idpf_create_reply'.
type IdpfCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IdpfCreateReply) Reset()               { *m = IdpfCreateReply{} }
func (*IdpfCreateReply) GetMessageName() string { return "idpf_create_reply" }
func (*IdpfCreateReply) GetCrcString() string   { return "5383d31f" }
func (*IdpfCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IdpfCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *IdpfCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IdpfCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// - client_index - opaque cookie to identify the sender
//   - sw_if_index - interface index
//
// IdpfDelete defines message 'idpf_delete'.
type IdpfDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IdpfDelete) Reset()               { *m = IdpfDelete{} }
func (*IdpfDelete) GetMessageName() string { return "idpf_delete" }
func (*IdpfDelete) GetCrcString() string   { return "f9e6675e" }
func (*IdpfDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IdpfDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *IdpfDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IdpfDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IdpfDeleteReply defines message 'idpf_delete_reply'.
type IdpfDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IdpfDeleteReply) Reset()               { *m = IdpfDeleteReply{} }
func (*IdpfDeleteReply) GetMessageName() string { return "idpf_delete_reply" }
func (*IdpfDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*IdpfDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IdpfDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IdpfDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IdpfDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_idpf_binapi_init() }
func file_idpf_binapi_init() {
	api.RegisterMessage((*IdpfCreate)(nil), "idpf_create_2ba86d91")
	api.RegisterMessage((*IdpfCreateReply)(nil), "idpf_create_reply_5383d31f")
	api.RegisterMessage((*IdpfDelete)(nil), "idpf_delete_f9e6675e")
	api.RegisterMessage((*IdpfDeleteReply)(nil), "idpf_delete_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IdpfCreate)(nil),
		(*IdpfCreateReply)(nil),
		(*IdpfDelete)(nil),
		(*IdpfDeleteReply)(nil),
	}
}