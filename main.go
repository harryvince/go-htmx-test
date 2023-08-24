package main

import (
	"app/lib"
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{})
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		results := lib.GetFacts()

		return c.Render("views/results", fiber.Map{
			"Results": results,
		})
	})

	app.Listen(":3000")
}
