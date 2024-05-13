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
func NewBoard(sd model.Direction, sp ...model.Position) *Board {
	if len(sp) < 2 {
		panic("2+ snake positions must be provided")
	}
	head := toByte(sp[0])
	tail := toByte(sp[len(sp)-1])

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

func (b *Board) SetEntityAt(p model.Position, e model.EntityKind, d model.Direction) model.Entity {
	if p.X > b.Width || p.Y > b.Height {
		panic("invalid position")
	}
	b.g.setAt(uint8(p.X), uint8(p.Y), byteOf(e, d))
	return b.GetEntityAt(p)
}

func (b *Board) GetNextPos(p model.Position, d model.Direction) model.Position {
	var xo, yo = dirToOffset(d)

	return model.Position{
		X: uint(int(p.X)+xo) % b.Width,
		Y: uint(int(p.Y)+yo) % b.Height,
	}
}

func (b *Board) GetRandomEmptyCell() (model.Position, bool) {

	rn := rand.Uint32()
	rx := uint(rn >> 16)
	ry := uint(rn & 0x0000FFFF)

	for y := range b.Height {
		for x := range b.Width {

			e := b.GetEntityAt(model.Position{X: (x + rx) % b.Width, Y: (y + ry) % b.Height})
			if e.Kind == model.ENTITY_NONE {
				return e.Pos, true
			}
		}
	}
	return model.Position{}, false
}

func (b *Board) GetEntityAt(p model.Position) model.Entity {
	a := b.g.getAt(uint8(p.X), uint8(p.Y))

	e, d := propsOf(a)

	return model.Entity{
		Kind:      e,
		Pos:       p,
		Direction: d,
	}
}

func (b *Board) GetSnakeHead() model.Entity {
	return b.GetEntityAt(toPos(b.g.head))
}

func (b *Board) MoveHead(d model.Direction) model.Entity {
	// 1. get head entity
	t := b.GetEntityAt(toPos(b.g.head))
	// 2. update current head direction
	b.SetEntityAt(t.Pos, model.ENTITY_SNAKE, d)
	// 3. move head to next pos
	p := b.GetNextPos(t.Pos, d)
	b.g.head = toByte(p)
	return b.SetEntityAt(p, model.ENTITY_SNAKE, d)
}

func (b *Board) MoveTail() model.Entity {
	t := b.GetEntityAt(toPos(b.g.tail))
	p := b.GetNextPos(t.Pos, t.Direction)
	b.g.tail = toByte(p)
	return b.SetEntityAt(t.Pos, model.ENTITY_NONE, model.DIRECTION_NONE /* ANY */)
}

func (b *Board) DumpEntities() []model.Entity {
	res := make([]model.Entity, b.Width*b.Height)
	var i int
	for y := range b.Height {
		for x := range b.Width {
			e := b.GetEntityAt(model.Position{X: x, Y: y})
			res[i] = e
			i++
		}
	}
	return res
}

func toByte(p model.Position) byte {
	return xyToByte(uint8(p.X), uint8(p.Y))
}

func toPos(b byte) model.Position {
	x, y := byteToXY(b)
	return model.Position{X: uint(x), Y: uint(y)}
}
