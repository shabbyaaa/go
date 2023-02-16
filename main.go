package main

import (
	"go/router"
)

func main() {
	r := router.Router()

	r.Run(":8080")
}
