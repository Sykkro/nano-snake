package main

import (
	"context"
	"sykkro/nano-snake/internal/engine"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	engine.Run(ctx)

}
