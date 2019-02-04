package main

import (
	"./app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", app.GetInstructions)
		v1.GET("/breaking/:channel/:start/:stop", app.GetInstruction)
		v1.POST("/instructions", app.PostInstruction)
		v1.PUT("/instructions/:id", app.UpdateInstruction)
		v1.DELETE("/instructions/:id", app.DeleteInstruction)
	}

	return router
}

func main() {
	router := SetupRouter()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	router.Use(cors.Default())
	router.Run(":8080")
}
