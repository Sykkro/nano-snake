package model

type Sprite uint

const (
	SPRITE_NONE Sprite = iota
	SPRITE_EMPTY
	SPRITE_SNAKE_BODY
	SPRITE_APPLE
)

type EntityKind uint

const (
	ENTITY_NONE EntityKind = iota
	ENTITY_SNAKE
	ENTITY_APPLE
)

type Direction uint

const (
	DIRECTION_UP Direction = iota
	DIRECTION_RIGHT
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_NONE
)

type Position struct {
	X uint
	Y uint
}
type Entity struct {
	Kind      EntityKind
	Pos       Position
	Direction Direction
}

func (d Direction) GetOpposite() Direction {
	switch d {
	case DIRECTION_UP:
		return DIRECTION_DOWN
	case DIRECTION_RIGHT:
		return DIRECTION_LEFT
	case DIRECTION_LEFT:
		return DIRECTION_RIGHT
	case DIRECTION_DOWN:
		return DIRECTION_UP
	default:
		panic("unsupported direction")
	}
}
