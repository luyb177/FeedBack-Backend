package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//	@title		木犀反馈系统 API
//	@version	1.0
//	@host		localhost:8080
//	@BasePath	/
func main() {
	initViper()
	app := InitApp()
	//app.t.StartDailyTask()
	app.r.Run(":8080")
}

type App struct {
	r *gin.Engine
}

func initViper() {
	cfile := pflag.String("config", "config/config.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cfile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
