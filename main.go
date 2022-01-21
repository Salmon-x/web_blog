package main

import (
	"blog/db"
	"blog/model"
	"blog/routes"
)

func main()  {
	model.InitDb()
	model.NosOne()
	db.RedisInit()
	if model.Db != nil {
		sql_db, _ := model.Db.DB()
		defer sql_db.Close()
	}
	routes.InitRouter()
}

// https://github.com/dollarkillerx/goseaweed
// seaweedfs
