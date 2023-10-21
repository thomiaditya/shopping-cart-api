package main

import (
	"context"

	"github.com/thomiaditya/shop-api/internal/app"
)

func main() {
	ctx := context.Background()

	if err := app.NewApp().Start(ctx); err != nil {
		panic(err)
	}
}
