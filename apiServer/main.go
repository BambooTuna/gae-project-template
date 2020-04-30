package main

import (
	"fmt"
	"github.com/BambooTuna/gae-project-templete/apiServer/interfaces"
	"github.com/BambooTuna/gae-project-templete/apiServer/swagger/docs"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

func main() {
	apiVersion := "v1"
	serverPort := config.GetEnvString("PORT", "8080")
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	sampleHandler := interfaces.SampleHandler{}
	api := r.Group(apiVersion)
	api.POST("/sample", sampleHandler.PostSampleRoute())

	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Schemes = []string{"http"}
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", serverPort)
		docs.SwaggerInfo.BasePath = apiVersion
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", serverPort)))
}
