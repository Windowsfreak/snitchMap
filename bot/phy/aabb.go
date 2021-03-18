package phy

import (
	"math"

	"github.com/Windowsfreak/go-mc/bot/world"
)

type MinMax struct {
	Min, Max float64
}

// Extends adjusts the bounds of the MinMax. A negative number will reduce the
// minimum bound, whereas a positive number will increase the maximum bound.
func (mm MinMax) Extend(delta float64) MinMax {
	if delta < 0 {
		return MinMax{
			Min: mm.Min + delta,
			Max: mm.Max,
		}
	}

	return MinMax{
		Min: mm.Min,
		Max: mm.Max + delta,
	}
}

// Contract reduces both the minimum and maximum bound by the provided amount,
// such that the difference between the bounds decreases for positive values.
func (mm MinMax) Contract(amt float64) MinMax {
	return MinMax{
		Min: mm.Min + amt,
		Max: mm.Max - amt,
	}
}

// Expand changes the minimum and maximum bounds by the provided amount, such
// that the difference between the bounds increases for positive values.
func (mm MinMax) Expand(amt float64) MinMax {
	return MinMax{
		Min: mm.Min - amt,
		Max: mm.Max + amt,
	}
}

// Offset adds the provided value to both the minimum and maximum value.
func (mm MinMax) Offset(amt float64) MinMax {
	return MinMax{
		Min: mm.Min + amt,
		Max: mm.Max + amt,
	}
}

// AABB implements Axis Aligned Bounding Box operations.
type AABB struct {
	X, Y, Z MinMax
	Block   world.BlockStatus
}

// Extend adjusts the minimum (for negative values) or maximum bounds (for
// positive values) by the provided scalar for each dimension.
func (bb AABB) Extend(dx, dy, dz float64) AABB {
	return AABB{
		X:     bb.X.Extend(dx),
		Y:     bb.Y.Extend(dx),
		Z:     bb.Z.Extend(dx),
		Block: bb.Block,
	}
}

// Contract reduces the difference between the min/max bounds (for positive
// values) for each dimension.
func (bb AABB) Contract(x, y, z float64) AABB {
	return AABB{
		X:     bb.X.Contract(x),
		Y:     bb.Y.Contract(y),
		Z:     bb.Z.Contract(z),
		Block: bb.Block,
	}
}

// Expand increases both the minimum and maximum bounds by the provided amount
// (for positive values) for each dimension.
func (bb AABB) Expand(x, y, z float64) AABB {
	return AABB{
		X:     bb.X.Expand(x),
		Y:     bb.Y.Expand(y),
		Z:     bb.Z.Expand(z),
		Block: bb.Block,
	}
}

// Offset moves both the minimum and maximum bound by the provided value for
// each dimension.
func (bb AABB) Offset(x, y, z float64) AABB {
	return AABB{
		X:     bb.X.Offset(x),
		Y:     bb.Y.Offset(y),
		Z:     bb.Z.Offset(z),
		Block: bb.Block,
	}
}

func (bb AABB) XOffset(o AABB, xOffset float64) float64 {
	if o.Y.Max > bb.Y.Min && o.Y.Min < bb.Y.Max && o.Z.Max > bb.Z.Min && o.Z.Min < bb.Z.Max {
		if xOffset > 0.0 && o.X.Max <= bb.X.Min {
			xOffset = math.Min(bb.X.Min-o.X.Max, xOffset)
		} else if xOffset < 0.0 && o.X.Min >= bb.X.Max {
			xOffset = math.Max(bb.X.Max-o.X.Min, xOffset)
		}
	}
	return xOffset
}

func (bb AABB) YOffset(o AABB, yOffset float64) float64 {
	if o.X.Max > bb.X.Min && o.X.Min < bb.X.Max && o.Z.Max > bb.Z.Min && o.Z.Min < bb.Z.Max {
		if yOffset > 0.0 && o.Y.Max <= bb.Y.Min {
			yOffset = math.Min(bb.Y.Min-o.Y.Max, yOffset)
		} else if yOffset < 0.0 && o.Y.Min >= bb.Y.Max {
			yOffset = math.Max(bb.Y.Max-o.Y.Min, yOffset)
		}
	}
	return yOffset
}

func (bb AABB) ZOffset(o AABB, zOffset float64) float64 {
	if o.X.Max > bb.X.Min && o.X.Min < bb.X.Max && o.Y.Max > bb.Y.Min && o.Y.Min < bb.Y.Max {
		if zOffset > 0.0 && o.Z.Max <= bb.Z.Min {
			zOffset = math.Min(bb.Z.Min-o.Z.Max, zOffset)
		} else if zOffset < 0.0 && o.Z.Min >= bb.Z.Max {
			zOffset = math.Max(bb.Z.Max-o.Z.Min, zOffset)
		}
	}
	return zOffset
}

func (bb AABB) Intersects(o AABB) bool {
	return true &&
		bb.X.Min < o.X.Max && bb.X.Max > o.X.Min &&
		bb.Y.Min < o.Y.Max && bb.Y.Max > o.Y.Min &&
		bb.Z.Min < o.Z.Max && bb.Z.Max > o.Z.Min
}
