package main

import (
	"go-chat/router"
	"go-chat/utils"

	"github.com/spf13/viper"
)

func main() {
	// 项目配置
	utils.InitConfig()
	utils.InitMySQL()

	r := router.InitRouter()
	println("server run port: ", viper.GetString("port.server"))
	r.Run(viper.GetString("port.server")) // listen and serve on 0.0.0.0:8080
}
