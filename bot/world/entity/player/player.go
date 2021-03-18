package player

import "github.com/Windowsfreak/go-mc/bot/world/entity"

type Pos struct {
	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool
}

func (p Pos) PosEqual(other Pos) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}
func (p Pos) LookEqual(other Pos) bool {
	return p.Yaw == other.Yaw && p.Pitch == other.Pitch
}
func (p Pos) Equal(other Pos) bool {
	return p.PosEqual(other) && p.LookEqual(other) && p.OnGround == other.OnGround
}

// Player includes the player's status.
type Player struct {
	entity.Entity
	UUID [2]int64 //128bit UUID

	Pos Pos

	HeldItem int //拿着的物品栏位

	Health         float32 //血量
	Food           int32   //饱食度
	FoodSaturation float32 //食物饱和度

	Level int32
}
