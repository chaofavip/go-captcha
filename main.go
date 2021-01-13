package main

import (
	"captchaCode/routers"
)

func main() {
	//路由初始化
	router := routers.InitRouter()
	//输出配置得端口(默认是8080，注意端口执行要加冒号)
	_ = router.Run(":8006")
}
