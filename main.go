package main

import (
	"go/router"
	"go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()

	r.Run(":8080")
}
