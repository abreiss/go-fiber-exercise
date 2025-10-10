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
	//timestamp := time.Now().Unix()
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
		nowMs := time.Now().UnixMilli()
		json := fmt.Sprintf(`{"message":"My name is Nico Reiss","timestamp":%d}`, nowMs)
		c.Response().Header.SetContentType("application/json")
		c.Response().SetBodyRaw([]byte(json))
		return nil
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // local dev
	}
	addr := "0.0.0.0:" + port
	fmt.Println("Server on http://" + "localhost:" + port)
	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}