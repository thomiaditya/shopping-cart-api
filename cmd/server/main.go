package main

import (
	"github.com/thomiaditya/shop-api/internal/app"
)

func main() {
	if err := app.NewApp().Start(); err != nil {
		panic(err)
	}
}
