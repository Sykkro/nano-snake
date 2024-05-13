package engine

import (
	"context"
	"fmt"
	"strings"
	"sykkro/nano-snake/internal/model"
	"sykkro/nano-snake/internal/scene"
	"sykkro/nano-snake/internal/state"
	"time"
)

const run_frequency = 150 * time.Millisecond

type engine struct {
	b  *state.Board      // game board
	nd model.Direction   // next direction
	rb chan model.Entity // render buffer
	sc uint16            // score
	on bool              // playing/pausing
}

func Run(ctx context.Context) {
	e := makeEngine()
	defer close(e.rb)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// bind arrow controls:
	ctx = e.setupControls(ctx, cancel)

	// start a new game:
	e.newGame(cancel)

	// render scene:
	go e.render()

	// main game loop:
	go e.loop(cancel)

	<-ctx.Done()
}

func (e *engine) redraw() {
	scene.ClearCanvas()
	// write header where score will be:
	// note: we're doubling X positions for stretching viewport and make a grid-like layout:
	scene.WriteAt(0, 0, strings.Repeat(sprite_header, int(e.b.Width)*2))
	scene.WriteAt(0, 1, strings.Repeat(sprite_header, int(e.b.Width)*2))
	scene.WriteAt(0, 2, strings.Repeat(sprite_header, int(e.b.Width)*2))
	// dump all
	for _, es := range e.b.DumpEntities() {
		e.rb <- es
	}
}

func (e *engine) render() {
	var hh uint8 = 3 // header height
	// run while render buffer is open:
	for x := range e.rb {
		scene.WriteAt(uint8(x.Pos.X*2), uint8(x.Pos.Y)+hh, getSprite(x))
		scene.WriteAt(uint8(x.Pos.X*2+1), uint8(x.Pos.Y)+hh, getSprite(model.Entity{}))
		scene.WriteAt(uint8(e.b.Width)-6, 1, fmt.Sprintf(" Score: %.3d ", e.sc))
		scene.MoveCursor(uint8(e.b.Width)*2, uint8(e.b.Height)+hh)
	}
}

func (e *engine) loop(cancel context.CancelFunc) {

	// main game loop (run "forever")
	for range time.Tick(run_frequency) {
		if e.on { // not paused
			for x := range e.moveSnake(e.nd, cancel) {
				e.rb <- x
			}
		}
	}
}

func makeEngine() *engine {
	e := &engine{}
	e.rb = make(chan model.Entity)
	return e
}

func (e *engine) newGame(cancel context.CancelFunc) {

	e.nd = model.DIRECTION_UP
	sp := []model.Point{
		{X: 8, Y: 6},
		{X: 8, Y: 7},
		{X: 8, Y: 8},
	}
	e.b = state.NewBoard(e.nd, sp...)
	e.sc = 0
	e.on = true
	e.addRandomApple(cancel)
	go e.redraw()
}

func (e *engine) togglePause() {
	e.on = !e.on
}

func (e *engine) switchDirection(d model.Direction) {
	if e.on { // disable controls if paused
		e.nd = d
	}
}
