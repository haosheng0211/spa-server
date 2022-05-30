package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

var (
	host      string
	port      string
	directory string
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Hostname to listen on")
	flag.StringVar(&port, "port", "3000", "Port to listen on")
	flag.StringVar(&directory, "directory", "/data", "Directory to serve files from")
	flag.Parse()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	app.Use(logger.New())

	app.Static("/", directory)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile(fmt.Sprintf("%s/index.html", directory))
	})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", host, port)))
}
