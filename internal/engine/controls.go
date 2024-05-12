package engine

import (
	"context"
	"sykkro/nano-snake/internal/model"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var arrow_controls = map[keys.KeyCode]model.Direction{
	keys.Up:    model.DIRECTION_UP,
	keys.Down:  model.DIRECTION_DOWN,
	keys.Left:  model.DIRECTION_LEFT,
	keys.Right: model.DIRECTION_RIGHT,
}

func (e *engine) setupControls(ctx context.Context, cancel context.CancelFunc) context.Context {
	go func() {
		defer cancel()
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {

			ks := key.String()
			if key.Code == keys.CtrlC {
				ks = "q"
			}

			switch ks {
			case "n":
				e.newGame(cancel)
			case "p":
				e.togglePause()
			case "q":
				stop = true
			default:
				dir, ok := arrow_controls[key.Code]
				if ok {
					e.switchDirection(dir)
				}
			}
			return
		})
	}()

	return ctx
}
