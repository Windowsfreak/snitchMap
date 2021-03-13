package ptypes

import (
	pk "github.com/Windowsfreak/go-mc/net/packet"
)

// JoinGame encodes global/world information from the server.
type JoinGame struct {
	PlayerEntity pk.Int
	Gamemode     pk.UnsignedByte
	Dimension    pk.Int
	Difficulty   pk.UnsignedByte
	MaxPlayers   pk.UnsignedByte // Now ignored
	LevelType    pk.String
	RDI          pk.Boolean // Reduced Debug Info

	//Dimension    pk.Int
	//Hardcore     pk.Boolean
	//PrevGamemode pk.UnsignedByte
	//WorldCount   pk.VarInt
	//WorldNames   pk.Identifier
	//DimensionCodec pk.NBT
	//WorldName    pk.Identifier
	//HashedSeed   pk.Long
	//ViewDistance pk.VarInt
	//ERS          pk.Boolean // Enable respawn screen
	//IsDebug      pk.Boolean
	//IsFlat       pk.Boolean
}

func (p *JoinGame) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.PlayerEntity, &p.Gamemode, &p.Dimension, &p.Difficulty, &p.MaxPlayers, &p.LevelType, &p.RDI)
}
