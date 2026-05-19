package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	e := NewExchange()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Exchange!")
	})

	app.Get("/get-bids", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": true,
			"data":   e.Bids,
		})
	})

	app.Get("/get-asks", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data":   e.Asks,
			"status": true,
		})
	})

	type MarketOrder struct {
		OrderType string  `json:"order_type"`
		Quantity  float64 `json:"quantity"`
	}

	app.Post("/market-order", func(c *fiber.Ctx) error {
		body := MarketOrder{}
		if err := c.BodyParser(&body); err != nil {
			return c.JSON(fiber.Map{
				"status": false,
				"data":   err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status": true,
			"data":   e.marketOrder(body.OrderType, body.Quantity),
		})
	})

	type LimitOrder struct {
		OrderType string  `json:"order_type"`
		Rate      float64 `json:"rate"`
		Quantity  float64 `json:"quantity"`
	}

	app.Post("/limit-order", func(c *fiber.Ctx) error {
		body := LimitOrder{}
		if err := c.BodyParser(&body); err != nil {
			return c.JSON(fiber.Map{
				"status": false,
				"data":   err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status": true,
			"data":   e.limitOrder(body.OrderType, Order{Qty: body.Quantity, Rate: body.Rate}),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
