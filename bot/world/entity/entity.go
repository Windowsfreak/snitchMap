package entity

import (
	"bytes"

	"github.com/Windowsfreak/go-mc/data/entity"
	item "github.com/Windowsfreak/go-mc/data/item"
	"github.com/Windowsfreak/go-mc/nbt"
	pk "github.com/Windowsfreak/go-mc/net/packet"
	"github.com/google/uuid"
)

// BlockEntity describes the representation of a tile entity at a position.
type BlockEntity struct {
	ID string `nbt:"id"`

	// global co-ordinates
	X int `nbt:"x"`
	Y int `nbt:"y"`
	Z int `nbt:"z"`

	// sign-specific.
	Color string `nbt:"color"`
	Text1 string `nbt:"Text1"`
	Text2 string `nbt:"Text2"`
	Text3 string `nbt:"Text3"`
	Text4 string `nbt:"Text4"`
}

//Entity represents an instance of an entity.
type Entity struct {
	ID   int32
	Data int32
	Base *entity.Entity

	UUID uuid.UUID

	X, Y, Z          float64
	Pitch, Yaw       int8
	VelX, VelY, VelZ int16
	OnGround         bool

	IsLiving  bool
	HeadPitch int8
}

// The Slot data structure is how Minecraft represents an item and its associated data in the Minecraft Protocol
type Slot struct {
	Present bool
	ItemID  item.ID
	Count   int8
	Damage  int16
	NBT     interface{}
}

//Decode implement packet.FieldDecoder interface
func (s *Slot) Decode(r pk.DecodeReader) error {
	var itemID pk.Short
	if err := itemID.Decode(r); err != nil {
		return err
	}
	s.Present = itemID != -1
	if s.Present {
		s.ItemID = item.ID(itemID)
		if err := (*pk.Byte)(&s.Count).Decode(r); err != nil {
			return err
		}
		if err := (*pk.Short)(&s.Damage).Decode(r); err != nil {
			return err
		}
		if err := nbt.NewDecoder(r).Decode(&s.NBT); err != nil {
			return err
		}
	}
	return nil
}

func (s Slot) Encode() []byte {
	if !s.Present {
		return pk.Boolean(false).Encode()
	}

	var b bytes.Buffer
	b.Write(pk.Short(s.ItemID).Encode())
	b.Write(pk.Byte(s.Count).Encode())
	b.Write(pk.Short(s.Damage).Encode())

	if s.NBT != nil {
		nbt.NewEncoder(&b).Encode(s.NBT)
	} else {
		b.Write([]byte{nbt.TagEnd})
	}

	return b.Bytes()
}

func (s Slot) String() string {
	return item.ByID[item.ID(s.ItemID)].DisplayName
}
