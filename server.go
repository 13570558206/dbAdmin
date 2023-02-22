package main

import (
	"dbAdmin/frame"
	"dbAdmin/router"
)

func init() {
	frame.Initialize()
}

func main() {

	if err := router.Register().Start(":8080"); err != nil {
		panic("路由启动失败:" + err.Error())
	}
}
