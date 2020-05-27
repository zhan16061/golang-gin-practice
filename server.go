package main

import (
	v1 "golang-nick/api/v1"
	"golang-nick/controller"
	"golang-nick/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://127.0.0.1:3000"}
	server.Use(cors.New(config))

	// server.GET("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.FindAll())
	// })

	// server.POST("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.Save(ctx))
	// })

	v1.ApplyRoutes(server)
	server.Run(":8080")
}
