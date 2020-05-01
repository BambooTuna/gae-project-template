package main

import (
	"fmt"
	"github.com/BambooTuna/gae-project-template/apiServer/interfaces"
	"github.com/BambooTuna/gae-project-template/apiServer/swagger/docs"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

func main() {
	serverPort := config.GetEnvString("PORT", "18080")
	r := gin.Default()

	sampleHandler := interfaces.SampleHandler{}
	r.POST("/sample", sampleHandler.PostSampleRoute())

	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		docs.SwaggerInfo.Host = ""
		docs.SwaggerInfo.BasePath = ""
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", serverPort)))
}
