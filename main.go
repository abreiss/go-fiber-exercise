//simple go fiber taken from fiber's hello world
//added ability to send either name or date
package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)


func main() {
	app := fiber.New()
	//name := "Nico Reiss"
	timestamp := time.Now()
	//fmt.Println(t.Format("20060102150405"))
	//route path, does a get, returns the send string
	// the / significes path to the endpoint like /auth

	/*its gotta look like
		{
		“message”: “My name is ___”,
		“timestamp”: 12312344
		}

	*/
	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString(timestamp)
		return c.JSON(timestamp)
	})

	app.Listen(":3000")
}