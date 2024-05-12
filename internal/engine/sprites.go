package engine

import (
	"sykkro/nano-snake/internal/model"
)

const (
	color_reset = "\u001b[0m"

	color_black         = "\u001b[30m"
	color_red           = "\u001b[31m"
	color_green         = "\u001b[32m"
	color_yellow        = "\u001b[33m"
	color_blue          = "\u001b[34m"
	color_magenta       = "\u001b[35m"
	color_cyan          = "\u001b[36m"
	color_white         = "\u001b[37m"
	color_light_black   = "\u001b[30;1m"
	color_light_red     = "\u001b[31;1m"
	color_light_green   = "\u001b[32;1m"
	color_light_yellow  = "\u001b[33;1m"
	color_light_blue    = "\u001b[34;1m"
	color_light_magenta = "\u001b[35;1m"
	color_light_cyan    = "\u001b[36;1m"
	color_light_white   = "\u001b[37;1m"
	bg_color_dark_grey  = "\u001b[48:5:233m"
)

const (
	sprite_header = "▒" // for reference: "░▒▓█"
	sprite_none   = ":" //" " //"░"
	sprite_apple  = "⬢" //"●" //"⧇"
	sprite_snake  = "■" //"█"
)

func getSprite(m model.Entity) (sprite string) {

	sprite = color(sprite_none, color_black)
	if m.Kind == model.ENTITY_APPLE {
		sprite = color(sprite_apple, color_light_red)
	} else if m.Kind == model.ENTITY_SNAKE {
		sprite = color(sprite_snake, color_light_green)
		// switch m.Direction {
		// case model.DIRECTION_DOWN:
		// 	sprite = color("v", color_light_green)
		// case model.DIRECTION_UP:
		// 	sprite = color("^", color_light_green)
		// case model.DIRECTION_LEFT:
		// 	sprite = color("<", color_light_green)
		// case model.DIRECTION_RIGHT:
		// 	sprite = color(">", color_light_green)
		// }
	}
	return bg_color_dark_grey + sprite
}

func color(s string, c string) string {
	return c + s + color_reset
}
