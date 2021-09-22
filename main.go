package main

import (
	"singo/conf"
	"singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	err := r.Run(":3000")
	if err != nil {
		panic(err.Error())
	}
}
