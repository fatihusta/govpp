// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.11.0
//  VPP:              24.06-release
// source: core/dev.api.json

// Package dev contains generated bindings for API file dev.api.
//
// Contents:
// -  2 enums
// -  8 messages
package dev

import (
	"strconv"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "dev"
	APIVersion = "0.0.1"
	VersionCrc = 0x36dfff5d
)

// DevFlags defines enum 'dev_flags'.
type DevFlags uint32

const (
	VL_API_DEV_FLAG_NO_STATS DevFlags = 1
)

var (
	DevFlags_name = map[uint32]string{
		1: "VL_API_DEV_FLAG_NO_STATS",
	}
	DevFlags_value = map[string]uint32{
		"VL_API_DEV_FLAG_NO_STATS": 1,
	}
)

func (x DevFlags) String() string {
	s, ok := DevFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := DevFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "DevFlags(" + strconv.Itoa(int(n)) + ")"
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

// DevPortFlags defines enum 'dev_port_flags'.
type DevPortFlags uint32

const (
	VL_API_DEV_PORT_FLAG_INTERRUPT_MODE DevPortFlags = 1
)

var (
	DevPortFlags_name = map[uint32]string{
		1: "VL_API_DEV_PORT_FLAG_INTERRUPT_MODE",
	}
	DevPortFlags_value = map[string]uint32{
		"VL_API_DEV_PORT_FLAG_INTERRUPT_MODE": 1,
	}
)

func (x DevPortFlags) String() string {
	s, ok := DevPortFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := DevPortFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "DevPortFlags(" + strconv.Itoa(int(n)) + ")"
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

// /* SPDX-License-Identifier: Apache-2.0
//   - Copyright(c) 2022 Cisco Systems, Inc.
//
// DevAttach defines message 'dev_attach'.
type DevAttach struct {
	DeviceID   string   `binapi:"string[48],name=device_id" json:"device_id,omitempty"`
	DriverName string   `binapi:"string[16],name=driver_name" json:"driver_name,omitempty"`
	Flags      DevFlags `binapi:"dev_flags,name=flags" json:"flags,omitempty"`
	Args       string   `binapi:"string[],name=args" json:"args,omitempty"`
}

func (m *DevAttach) Reset()               { *m = DevAttach{} }
func (*DevAttach) GetMessageName() string { return "dev_attach" }
func (*DevAttach) GetCrcString() string   { return "44b725fc" }
func (*DevAttach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DevAttach) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 48              // m.DeviceID
	size += 16              // m.DriverName
	size += 4               // m.Flags
	size += 4 + len(m.Args) // m.Args
	return size
}
func (m *DevAttach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.DeviceID, 48)
	buf.EncodeString(m.DriverName, 16)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeString(m.Args, 0)
	return buf.Bytes(), nil
}
func (m *DevAttach) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DeviceID = buf.DecodeString(48)
	m.DriverName = buf.DecodeString(16)
	m.Flags = DevFlags(buf.DecodeUint32())
	m.Args = buf.DecodeString(0)
	return nil
}

// DevAttachReply defines message 'dev_attach_reply'.
type DevAttachReply struct {
	DevIndex    uint32 `binapi:"u32,name=dev_index" json:"dev_index,omitempty"`
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ErrorString string `binapi:"string[],name=error_string" json:"error_string,omitempty"`
}

func (m *DevAttachReply) Reset()               { *m = DevAttachReply{} }
func (*DevAttachReply) GetMessageName() string { return "dev_attach_reply" }
func (*DevAttachReply) GetCrcString() string   { return "6082b181" }
func (*DevAttachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DevAttachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.DevIndex
	size += 4                      // m.Retval
	size += 4 + len(m.ErrorString) // m.ErrorString
	return size
}
func (m *DevAttachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.DevIndex)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.ErrorString, 0)
	return buf.Bytes(), nil
}
func (m *DevAttachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DevIndex = buf.DecodeUint32()
	m.Retval = buf.DecodeInt32()
	m.ErrorString = buf.DecodeString(0)
	return nil
}

// DevCreatePortIf defines message 'dev_create_port_if'.
type DevCreatePortIf struct {
	DevIndex    uint32       `binapi:"u32,name=dev_index" json:"dev_index,omitempty"`
	IntfName    string       `binapi:"string[32],name=intf_name" json:"intf_name,omitempty"`
	NumRxQueues uint16       `binapi:"u16,name=num_rx_queues" json:"num_rx_queues,omitempty"`
	NumTxQueues uint16       `binapi:"u16,name=num_tx_queues" json:"num_tx_queues,omitempty"`
	RxQueueSize uint16       `binapi:"u16,name=rx_queue_size" json:"rx_queue_size,omitempty"`
	TxQueueSize uint16       `binapi:"u16,name=tx_queue_size" json:"tx_queue_size,omitempty"`
	PortID      uint16       `binapi:"u16,name=port_id" json:"port_id,omitempty"`
	Flags       DevPortFlags `binapi:"dev_port_flags,name=flags" json:"flags,omitempty"`
	Args        string       `binapi:"string[],name=args" json:"args,omitempty"`
}

func (m *DevCreatePortIf) Reset()               { *m = DevCreatePortIf{} }
func (*DevCreatePortIf) GetMessageName() string { return "dev_create_port_if" }
func (*DevCreatePortIf) GetCrcString() string   { return "1eb00d01" }
func (*DevCreatePortIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DevCreatePortIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4               // m.DevIndex
	size += 32              // m.IntfName
	size += 2               // m.NumRxQueues
	size += 2               // m.NumTxQueues
	size += 2               // m.RxQueueSize
	size += 2               // m.TxQueueSize
	size += 2               // m.PortID
	size += 4               // m.Flags
	size += 4 + len(m.Args) // m.Args
	return size
}
func (m *DevCreatePortIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.DevIndex)
	buf.EncodeString(m.IntfName, 32)
	buf.EncodeUint16(m.NumRxQueues)
	buf.EncodeUint16(m.NumTxQueues)
	buf.EncodeUint16(m.RxQueueSize)
	buf.EncodeUint16(m.TxQueueSize)
	buf.EncodeUint16(m.PortID)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeString(m.Args, 0)
	return buf.Bytes(), nil
}
func (m *DevCreatePortIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DevIndex = buf.DecodeUint32()
	m.IntfName = buf.DecodeString(32)
	m.NumRxQueues = buf.DecodeUint16()
	m.NumTxQueues = buf.DecodeUint16()
	m.RxQueueSize = buf.DecodeUint16()
	m.TxQueueSize = buf.DecodeUint16()
	m.PortID = buf.DecodeUint16()
	m.Flags = DevPortFlags(buf.DecodeUint32())
	m.Args = buf.DecodeString(0)
	return nil
}

// DevCreatePortIfReply defines message 'dev_create_port_if_reply'.
type DevCreatePortIfReply struct {
	SwIfIndex   uint32 `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ErrorString string `binapi:"string[],name=error_string" json:"error_string,omitempty"`
}

func (m *DevCreatePortIfReply) Reset()               { *m = DevCreatePortIfReply{} }
func (*DevCreatePortIfReply) GetMessageName() string { return "dev_create_port_if_reply" }
func (*DevCreatePortIfReply) GetCrcString() string   { return "243c2374" }
func (*DevCreatePortIfReply) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DevCreatePortIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.SwIfIndex
	size += 4                      // m.Retval
	size += 4 + len(m.ErrorString) // m.ErrorString
	return size
}
func (m *DevCreatePortIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.SwIfIndex)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.ErrorString, 0)
	return buf.Bytes(), nil
}
func (m *DevCreatePortIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = buf.DecodeUint32()
	m.Retval = buf.DecodeInt32()
	m.ErrorString = buf.DecodeString(0)
	return nil
}

// DevDetach defines message 'dev_detach'.
type DevDetach struct {
	DevIndex uint32 `binapi:"u32,name=dev_index" json:"dev_index,omitempty"`
}

func (m *DevDetach) Reset()               { *m = DevDetach{} }
func (*DevDetach) GetMessageName() string { return "dev_detach" }
func (*DevDetach) GetCrcString() string   { return "afae52d6" }
func (*DevDetach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DevDetach) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.DevIndex
	return size
}
func (m *DevDetach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.DevIndex)
	return buf.Bytes(), nil
}
func (m *DevDetach) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DevIndex = buf.DecodeUint32()
	return nil
}

// DevDetachReply defines message 'dev_detach_reply'.
type DevDetachReply struct {
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ErrorString string `binapi:"string[],name=error_string" json:"error_string,omitempty"`
}

func (m *DevDetachReply) Reset()               { *m = DevDetachReply{} }
func (*DevDetachReply) GetMessageName() string { return "dev_detach_reply" }
func (*DevDetachReply) GetCrcString() string   { return "c8d74455" }
func (*DevDetachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DevDetachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.Retval
	size += 4 + len(m.ErrorString) // m.ErrorString
	return size
}
func (m *DevDetachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.ErrorString, 0)
	return buf.Bytes(), nil
}
func (m *DevDetachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ErrorString = buf.DecodeString(0)
	return nil
}

// DevRemovePortIf defines message 'dev_remove_port_if'.
type DevRemovePortIf struct {
	SwIfIndex uint32 `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *DevRemovePortIf) Reset()               { *m = DevRemovePortIf{} }
func (*DevRemovePortIf) GetMessageName() string { return "dev_remove_port_if" }
func (*DevRemovePortIf) GetCrcString() string   { return "529cb13f" }
func (*DevRemovePortIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DevRemovePortIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *DevRemovePortIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.SwIfIndex)
	return buf.Bytes(), nil
}
func (m *DevRemovePortIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = buf.DecodeUint32()
	return nil
}

// DevRemovePortIfReply defines message 'dev_remove_port_if_reply'.
type DevRemovePortIfReply struct {
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ErrorString string `binapi:"string[],name=error_string" json:"error_string,omitempty"`
}

func (m *DevRemovePortIfReply) Reset()               { *m = DevRemovePortIfReply{} }
func (*DevRemovePortIfReply) GetMessageName() string { return "dev_remove_port_if_reply" }
func (*DevRemovePortIfReply) GetCrcString() string   { return "c8d74455" }
func (*DevRemovePortIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DevRemovePortIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.Retval
	size += 4 + len(m.ErrorString) // m.ErrorString
	return size
}
func (m *DevRemovePortIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.ErrorString, 0)
	return buf.Bytes(), nil
}
func (m *DevRemovePortIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ErrorString = buf.DecodeString(0)
	return nil
}

func init() { file_dev_binapi_init() }
func file_dev_binapi_init() {
	api.RegisterMessage((*DevAttach)(nil), "dev_attach_44b725fc")
	api.RegisterMessage((*DevAttachReply)(nil), "dev_attach_reply_6082b181")
	api.RegisterMessage((*DevCreatePortIf)(nil), "dev_create_port_if_1eb00d01")
	api.RegisterMessage((*DevCreatePortIfReply)(nil), "dev_create_port_if_reply_243c2374")
	api.RegisterMessage((*DevDetach)(nil), "dev_detach_afae52d6")
	api.RegisterMessage((*DevDetachReply)(nil), "dev_detach_reply_c8d74455")
	api.RegisterMessage((*DevRemovePortIf)(nil), "dev_remove_port_if_529cb13f")
	api.RegisterMessage((*DevRemovePortIfReply)(nil), "dev_remove_port_if_reply_c8d74455")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DevAttach)(nil),
		(*DevAttachReply)(nil),
		(*DevCreatePortIf)(nil),
		(*DevCreatePortIfReply)(nil),
		(*DevDetach)(nil),
		(*DevDetachReply)(nil),
		(*DevRemovePortIf)(nil),
		(*DevRemovePortIfReply)(nil),
	}
}
