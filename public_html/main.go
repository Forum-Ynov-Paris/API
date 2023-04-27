package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type Message struct {
	Title      string       `json:"title"`
	Tag        string       `json:"tag"`
	Content    string       `json:"text"`
	UUID       string       `json:"uuid"`
	Commentary []Commentary `json:"commentaire"`
}

type Commentary struct {
	Content string `json:"text"`
	UUID    string `json:"uuid"`
}

var Data []Message

func Init() {
	Data = []Message{
		{
			Title:   "Hello",
			Tag:     "Golang",
			Content: "Hello World",
			UUID:    "1",
			Commentary: []Commentary{
				{
					Content: "Hello World",
					UUID:    "1",
				},
			},
		},
	}
}

func main() {

	Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/data", func(c *fiber.Ctx) error {
		return c.SendString("Hello data 👋!" + os.Getenv("DATA_ENV"))
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0" + port))

}
