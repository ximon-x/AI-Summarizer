package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func getSummary(c *fiber.Ctx) error {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	var jsonData map[string]interface{}

	if err := c.BodyParser(&jsonData); err != nil {
		return err
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),

		openai.ChatCompletionRequest{

			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: jsonData["content"].(string),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v \n\n", err)

		return c.JSON(
			fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"summary":    "Something went wrong...",
			},
		)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return c.JSON(
		fiber.Map{
			"statusCode": http.StatusOK,
			"summary":    resp.Choices[0].Message.Content,
		})

}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	app.Post("/summary/", getSummary)

	log.Fatal(app.Listen(":8000"))
}
