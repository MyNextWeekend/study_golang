package main

import (
	config2 "study_golang/example/websocket/config"
	"study_golang/example/websocket/router"
)

func main() {
	config2.InitLog()
	config2.InitConfig()
	router.InitRouter()
}
