// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~162-g1765f014b
// source: plugins/bpf_trace_filter.api.json

// Package bpf_trace_filter contains generated bindings for API file bpf_trace_filter.api.
//
// Contents:
// -  2 messages
package bpf_trace_filter

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "bpf_trace_filter"
	APIVersion = "0.1.0"
	VersionCrc = 0x79160184
)

// /*
//   - bpf_trace_filter.api - BPF Trace filter API
//     *
//   - Copyright (c) 2023 Cisco and/or its affiliates
//   - Licensed under the Apache License, Version 2.0 (the "License");
//   - you may not use this file except in compliance with the License.
//   - You may obtain a copy of the License at:
//     *
//   - http://www.apache.org/licenses/LICENSE-2.0
//     *
//   - Unless required by applicable law or agreed to in writing, software
//   - distributed under the License is distributed on an "AS IS" BASIS,
//   - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   - See the License for the specific language governing permissions and
//   - limitations under the License.
//
// BpfTraceFilterSet defines message 'bpf_trace_filter_set'.
type BpfTraceFilterSet struct {
	IsAdd  bool   `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	Filter string `binapi:"string[],name=filter" json:"filter,omitempty"`
}

func (m *BpfTraceFilterSet) Reset()               { *m = BpfTraceFilterSet{} }
func (*BpfTraceFilterSet) GetMessageName() string { return "bpf_trace_filter_set" }
func (*BpfTraceFilterSet) GetCrcString() string   { return "3171346e" }
func (*BpfTraceFilterSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *BpfTraceFilterSet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1                 // m.IsAdd
	size += 4 + len(m.Filter) // m.Filter
	return size
}
func (m *BpfTraceFilterSet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeString(m.Filter, 0)
	return buf.Bytes(), nil
}
func (m *BpfTraceFilterSet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Filter = buf.DecodeString(0)
	return nil
}

// BpfTraceFilterSetReply defines message 'bpf_trace_filter_set_reply'.
type BpfTraceFilterSetReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *BpfTraceFilterSetReply) Reset()               { *m = BpfTraceFilterSetReply{} }
func (*BpfTraceFilterSetReply) GetMessageName() string { return "bpf_trace_filter_set_reply" }
func (*BpfTraceFilterSetReply) GetCrcString() string   { return "e8d4e804" }
func (*BpfTraceFilterSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *BpfTraceFilterSetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *BpfTraceFilterSetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *BpfTraceFilterSetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_bpf_trace_filter_binapi_init() }
func file_bpf_trace_filter_binapi_init() {
	api.RegisterMessage((*BpfTraceFilterSet)(nil), "bpf_trace_filter_set_3171346e")
	api.RegisterMessage((*BpfTraceFilterSetReply)(nil), "bpf_trace_filter_set_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BpfTraceFilterSet)(nil),
		(*BpfTraceFilterSetReply)(nil),
	}
}
