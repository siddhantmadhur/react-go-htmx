package main

import (
	"github.com/gofiber/template/html/v2"

	"embed"

	"github.com/gofiber/fiber/v2"
)

//go:embed assets/*
//go:embed components/*
//go:embed views/*

var f embed.FS

func main() {

	engine := html.New("./components", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "views/index.html")
	app.Get("/api/hello-world", func(c *fiber.Ctx) error {
		return c.Render("hello", fiber.Map{
			"Name": "Siddhant Madhur",
		})
	})
	app.Get("/people/:name", func(c *fiber.Ctx) error {
		return c.Render("people", fiber.Map{
			"Name": c.Params("name"),
		})
	})
	app.Static("/assets/bundle.css", "./assets/bundle.css")
	app.Static("/assets/bundle.js", "./assets/bundle.js")
	app.Listen(":8080")
}
