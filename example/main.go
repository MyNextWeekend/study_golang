package main

import (
	"study_golang/example/config"
	"study_golang/example/router"
)

func main() {
	config.InitLog()
	config.InitConfig()
	router.InitRouter()
}
