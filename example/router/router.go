package router

import (
	"github.com/gin-gonic/gin"
	"study_golang/example/config"
	"study_golang/example/server"
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
		config.Logger.Info("add ", agentId)
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
		config.Logger.Info("stop ", agentId)

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

	config.Logger.Infof("running port %s", config.HttpPort)
	_ = r.Run(":" + config.HttpPort)
}
