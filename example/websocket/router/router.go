package router

import (
	"github.com/gin-gonic/gin"
	config2 "study_golang/example/websocket/config"
	"study_golang/example/websocket/server"
)

func InitRouter() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/add/:agentId", func(c *gin.Context) {
		agentId := c.Param("agentId")
		config2.Logger.Info("add ", agentId)
		err := server.Server.AddClient(agentId)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "success add",
		})
	})
	r.GET("/queryAll", func(c *gin.Context) {
		result := server.Server.QueryAllClient()
		c.JSON(200, gin.H{
			"total": len(result),
			"data":  result,
		})
	})
	r.GET("/stop/:agentId", func(c *gin.Context) {
		agentId := c.Param("agentId")
		config2.Logger.Info("stop ", agentId)

		err := server.Server.StopClient(agentId)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "success stop",
		})
	})

	config2.Logger.Infof("running port %s", config2.HttpPort)
	_ = r.Run(":" + config2.HttpPort)
}
