package main

import "draw-service/router"

func main() {
	e := router.Router()
	e.Run(":8080")
}
