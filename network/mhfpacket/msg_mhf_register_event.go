package mhfpacket

import (
	"erupe-ce/common/byteframe"
	"erupe-ce/network"
	"erupe-ce/network/clientctx"
)

// MsgMhfRegisterEvent represents the MSG_MHF_REGISTER_EVENT
type MsgMhfRegisterEvent struct {
	AckHandle uint32
	Unk0      uint16
	WorldID   uint16
	LandID    uint16
	Unk3      uint8
	Unk4      uint8
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfRegisterEvent) Opcode() network.PacketID {
	return network.MSG_MHF_REGISTER_EVENT
}

// Parse parses the packet from binary
func (m *MsgMhfRegisterEvent) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	m.WorldID = bf.ReadUint16()
	m.LandID = bf.ReadUint16()
	m.Unk3 = bf.ReadUint8()
	m.Unk4 = bf.ReadUint8()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfRegisterEvent) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return nil
}
