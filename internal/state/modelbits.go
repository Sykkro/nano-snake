package state

import "sykkro/nano-snake/internal/model"

type SnakeDirection = byte

const (
	SNAKE_UP    SnakeDirection = 0b100
	SNAKE_DOWN  SnakeDirection = 0b110
	SNAKE_LEFT  SnakeDirection = 0b111
	SNAKE_RIGHT SnakeDirection = 0b101
)

var dir_to_snake = map[model.Direction]SnakeDirection{
	model.DIRECTION_UP:    SNAKE_UP,
	model.DIRECTION_DOWN:  SNAKE_DOWN,
	model.DIRECTION_LEFT:  SNAKE_LEFT,
	model.DIRECTION_RIGHT: SNAKE_RIGHT,
}

var snake_to_dir = map[SnakeDirection]model.Direction{
	SNAKE_UP:    model.DIRECTION_UP,
	SNAKE_DOWN:  model.DIRECTION_DOWN,
	SNAKE_LEFT:  model.DIRECTION_LEFT,
	SNAKE_RIGHT: model.DIRECTION_RIGHT,
}

func dirToOffset(d model.Direction) (x int, y int) {
	switch d {
	case model.DIRECTION_UP:
		y = -1
	case model.DIRECTION_DOWN:
		y = 1
	case model.DIRECTION_LEFT:
		x = -1
	case model.DIRECTION_RIGHT:
		x = 1
	}
	return
}

func byteOf(e model.EntityKind, d model.Direction) byte {
	switch e {
	case model.ENTITY_NONE:
		return 0b0
	case model.ENTITY_APPLE:
		return 0b010
	case model.ENTITY_SNAKE:
		return dir_to_snake[d]
	default:
		panic("usupported model type")
	}
}

func propsOf(b byte) (e model.EntityKind, d model.Direction) {
	e = model.ENTITY_NONE
	if b>>2 == 0b1 {
		e = model.ENTITY_SNAKE
		d = snake_to_dir[b]

	} else if b>>1 == 0b1 {
		e = model.ENTITY_APPLE
	}
	return
}
