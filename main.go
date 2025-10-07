//simple go fiber taken from fiber's hello world
//added ability to send either name or date
package main

import (
	"github.com/gofiber/fiber/v2"
	"time"

	"os"
	"fmt"

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


	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // local dev
	}
	addr := "0.0.0.0:" + port
	fmt.Println("Server on http://" + "localhost:" + port)
	if err := app.Listen(addr); err != nil {
		panic(err)
	}

	app.Listen(":3000")

}