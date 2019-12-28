package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetRengokuRankingRank represents the MSG_MHF_GET_RENGOKU_RANKING_RANK
type MsgMhfGetRengokuRankingRank struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetRengokuRankingRank) Opcode() network.PacketID {
	return network.MSG_MHF_GET_RENGOKU_RANKING_RANK
}

// Parse parses the packet from binary
func (m *MsgMhfGetRengokuRankingRank) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetRengokuRankingRank) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}