package mhfpacket

import ( 
 "errors" 

 	"erupe-ce/network/clientctx"
	"erupe-ce/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEntryFesta represents the MSG_MHF_ENTRY_FESTA
type MsgMhfEntryFesta struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEntryFesta) Opcode() network.PacketID {
	return network.MSG_MHF_ENTRY_FESTA
}

// Parse parses the packet from binary
func (m *MsgMhfEntryFesta) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEntryFesta) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}