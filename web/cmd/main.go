package main

import (
	"mc-agones-web/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":3001")
}