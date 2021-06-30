package main

import (
	"blog/model"
	"blog/routes"
)

func main()  {
	model.InitDb()
	routes.InitRouter()
}

// https://github.com/dollarkillerx/goseaweed
// seaweedfs