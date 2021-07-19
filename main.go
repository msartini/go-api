package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx, *fiber, Ctx) {
		ctx.Send("Hello Fiber")
	})

}
