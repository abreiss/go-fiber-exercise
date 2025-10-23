//this app adds a JSON response with your name and a dynamic timestamp
// and ensures the JSON output is minified
package main

import (
"github.com/gofiber/fiber/v2" 	// go framework
"time" 							//used for timestamp
"os"							//read port and env variables
"fmt"							//used for printing
)
//structure of json response
// {“message”: “My name is ___”, “timestamp”: 123456789}


func main() {
	//create fiber web app
	app := fiber.New()
	//get route for root path
	//when users visit (http get req), return the json
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":     "My name is Nico Reiss",
			//time is in unix miliseconds
			"timestamp":   time.Now().UnixMilli(),
			//deployed_at date.
			"deployed_at": time.Now().Format("2006-01-02"),
		})
	})
	//if port is not specified set it to 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // local dev
	}
	//list to all  ip address at port
	addr := "0.0.0.0:" + port
	//print where the server is hosted 
	fmt.Println("Server on http://" + "localhost:" + port)
	//inf loop unless error
	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}