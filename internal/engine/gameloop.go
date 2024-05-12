package engine

import (
	"context"
	"sykkro/nano-snake/internal/model"
)

func (e *engine) moveSnake(d model.Direction, gg context.CancelFunc) <-chan model.Entity {
	// main cycle
	changes := make(chan model.Entity)

	go func() {
		defer close(changes)
		// 1. get head direction and use provided direction or head, if invalid (i.e. if reverse)
		h := e.b.GetSnakeHead()
		if !canTurn(h.Direction, d) {
			d = h.Direction
		}

		// 2. check if food on next square, if not -> move tail and enter collision check
		p := e.b.GetNextPos(h.Pos, d)
		m := e.b.GetEntityAt(p)

		if m.Kind == model.ENTITY_APPLE {
			// increment score & add new apple
			e.sc++
			changes <- e.addRandomApple(gg)

		} else {
			// move tail and check for collisions
			changes <- e.b.MoveTail()

			// 3. collision check: check if snake at next square, if so game over
			m = e.b.GetEntityAt(p) // (refetch as previous m might have been tail)
			if m.Kind == model.ENTITY_SNAKE {
				gg()
			}
		}

		// 4. move head to next square
		changes <- e.b.MoveHead(d)

	}()
	return changes
}

func (e *engine) addRandomApple(gg context.CancelFunc) model.Entity {
	a, o := e.b.GetRandomEmptyCell()
	if !o {
		gg()
	}
	return e.b.SetEntityAt(a, model.ENTITY_APPLE, model.DIRECTION_NONE /* ANY */)
}

func canTurn(from model.Direction, to model.Direction) bool {
	return from.GetOpposite() != to
}
