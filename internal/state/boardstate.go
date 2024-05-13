package state

import (
	"sykkro/nano-snake/internal/binary"
)

const (
	// state per [x,y] in the board:
	// - first bit == snake (0b1**)
	// - if snake, last bits are direction (0b*11)
	// - otherwise, fruit if set (0b01*)
	cellbits uint = 3
)

type gamestate struct {
	head    byte // bits represent coordinates : x x x x y y y y
	tail    byte // bits represent coordinates : x x x x y y y y
	rowsize uint
	board   []byte // flat board state
}

func newGameState(head byte, tail byte, width uint, height uint) gamestate {
	if width > 16 || height > 16 {
		panic("unsupported board size: max width and height is 16")
	}
	return gamestate{
		head:    head,
		tail:    tail,
		rowsize: width,
		board:   make([]byte, cellbits*width*height/8), // e.g. a 16 * 16 board with 3 bits per cell takes 96 bytes
	}
}

func (g *gamestate) setAt(x uint8, y uint8, v byte) {
	i := cellbits*g.rowsize*uint(y) + cellbits*uint(x)

	for n := range cellbits {
		o := i + n          // bit to set, absolute offset
		b := o / 8          // byte offset to work with
		p := 7 - uint8(o%8) // bit, within byte offset, to manipulate (7..0)

		if binary.HasBit(v, uint8(n)) {
			g.board[b] = binary.SetBit(g.board[b], p)
		} else {
			g.board[b] = binary.ClearBit(g.board[b], p)
		}
	}
}

func (g *gamestate) getAt(x uint8, y uint8) byte {
	i := cellbits*g.rowsize*uint(y) + cellbits*uint(x)
	var r byte
	for n := range cellbits {
		o := i + n          // bit to get, absolute offset
		b := o / 8          // byte offset to work with
		p := 7 - uint8(o%8) // bit, within byte offset, to get (7..0)

		if binary.HasBit(g.board[b], p) {
			r = binary.SetBit(r, uint8(n))
		}
	}
	return r
}
