package main

import (
	"go/router"
	"go/utils"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	// InitTimer()
	r := router.Router()
	// 不写127.0.0.1 win每次执行会被防火墙拦截
	r.Run(viper.GetString("port.server"))
}

// 初始化定时器
// func InitTimer() {
// 	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
// }
