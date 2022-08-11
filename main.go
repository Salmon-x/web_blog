package main

import (
	_ "blog/initialize"
	"blog/model"
	"blog/routes"
	_ "blog/utils"
)

func main() {
	model.NosOne()
	routes.InitRouter()
}
