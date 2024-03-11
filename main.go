package main

import (
	router "github.com/prithvi009/easedeploy/routes"
)

func main() {

	// Start the server
	r := router.SetupRouter()
	r.Run(":8080")

}
