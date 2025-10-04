//simple go fiber taken from fiber's hello world
//added ability to send either name or date
package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)


func main() {
	app := fiber.New()
	timestamp := time.Now().Unix()
	//looks like its unix formatting
	//route path, does a get, returns the send string
	// the / significes path to the endpoint like /auth

	/*its gotta look like
		{
		“message”: “My name is ___”,
		“timestamp”: 12312344
		}

	*/
	app.Get("/", func(c *fiber.Ctx) error {
		payload := fiber.Map{
			"message":   "My name is Nico Reiss",
			"timestamp": timestamp,
		}
		return c.JSON(payload)
	})

	app.Listen(":3000")
}