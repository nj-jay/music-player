package main

import (

	"github.com/gin-gonic/gin"
	"github.com/nj-jay/music-player/server/global"
	"github.com/nj-jay/music-player/server/initalize"
	"github.com/nj-jay/music-player/server/middlewares"
	"github.com/nj-jay/music-player/server/routers"
	"log"

)

func main() {

	global.GMD_DB = initalize.GormMysql()

	initalize.MysqlTables(global.GMD_DB)

	global.GMD_RD = initalize.ConnectRedis()

	r := gin.Default()

	r.Use(middlewares.Cors())

	routers.LoadRouter(r)

	//err := http.ListenAndServeTLS(":8081", "conf/api.nj-jay.com_bundle.crt", "conf/api.nj-jay.com.key", r)

	err := r.Run(":8081")

	if err != nil {

		log.Fatal("run err")

	}

}