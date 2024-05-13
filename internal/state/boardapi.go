package state

import (
	"math/rand"
	"sykkro/nano-snake/internal/model"
)

const (
	_width  = 16
	_height = 16
)

type Board struct {
	g      gamestate
	Width  uint
	Height uint
}

// NewBoard initializes a new board, given the provided snake direction and positions on board
func NewBoard(sd model.Direction, sp ...model.Point) *Board {
	if len(sp) < 2 {
		panic("2+ snake positions must be provided")
	}
	head := sp[0].ToByte()
	tail := sp[len(sp)-1].ToByte()

	res := &Board{
		Width:  _width,
		Height: _height,
		g:      newGameState(head, tail, _width, _height),
	}
	// Initialize snake
	for _, p := range sp {
		res.SetEntityAt(p, model.ENTITY_SNAKE, sd)
	}

	return res
}

func (b *Board) SetEntityAt(p model.Point, e model.EntityKind, d model.Direction) model.Entity {
	if p.X > b.Width || p.Y > b.Height {
		panic("invalid position")
	}
	b.g.setAt(uint8(p.X), uint8(p.Y), byteOf(e, d))
	return b.GetEntityAt(p)
}

func (b *Board) GetNextPos(p model.Point, d model.Direction) model.Point {
	var xo, yo = dirToOffset(d)

	return model.Point{
		X: uint(int(p.X)+xo) % b.Width,
		Y: uint(int(p.Y)+yo) % b.Height,
	}
}

func (b *Board) GetRandomEmptyCell() (model.Point, bool) {

	rn := rand.Uint32()
	rx := uint(rn >> 16)
	ry := uint(rn & 0x0000FFFF)

	for y := range b.Height {
		for x := range b.Width {

			e := b.GetEntityAt(model.Point{X: (x + rx) % b.Width, Y: (y + ry) % b.Height})
			if e.Kind == model.ENTITY_NONE {
				return e.Pos, true
			}
		}
	}
	return model.Point{}, false
}

func (b *Board) GetEntityAt(p model.Point) model.Entity {
	a := b.g.getAt(uint8(p.X), uint8(p.Y))

	e, d := propsOf(a)

	return model.Entity{
		Kind:      e,
		Pos:       p,
		Direction: d,
	}
}

func (b *Board) GetSnakeHead() model.Entity {
	return b.GetEntityAt(model.PointFromByte(b.g.head))
}

func (b *Board) MoveHead(d model.Direction) model.Entity {
	// 1. get head entity
	t := b.GetEntityAt(model.PointFromByte(b.g.head))
	// 2. update current head direction
	b.SetEntityAt(t.Pos, model.ENTITY_SNAKE, d)
	// 3. move head to next pos
	p := b.GetNextPos(t.Pos, d)
	b.g.head = p.ToByte()
	return b.SetEntityAt(p, model.ENTITY_SNAKE, d)
}

func (b *Board) MoveTail() model.Entity {
	t := b.GetEntityAt(model.PointFromByte(b.g.tail))
	p := b.GetNextPos(t.Pos, t.Direction)
	b.g.tail = p.ToByte()
	return b.SetEntityAt(t.Pos, model.ENTITY_NONE, model.DIRECTION_NONE /* ANY */)
}

func (b *Board) DumpEntities() []model.Entity {
	res := make([]model.Entity, b.Width*b.Height)
	var i int
	for y := range b.Height {
		for x := range b.Width {
			e := b.GetEntityAt(model.Point{X: x, Y: y})
			res[i] = e
			i++
		}
	}
	return res
}
