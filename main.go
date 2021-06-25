package main

import (
	"blog/model"
	"blog/routes"
)

func main()  {
	model.InitDb()
	routes.InitRouter()
}