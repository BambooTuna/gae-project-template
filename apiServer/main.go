package main

import (
	"fmt"
	"github.com/BambooTuna/gae-project-template/apiServer/interfaces"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	serverPort := config.GetEnvString("PORT", "18080")
	r := gin.Default()

	sampleHandler := interfaces.SampleHandler{}
	r.POST("/sample", sampleHandler.PostSampleRoute())

	//if gin.Mode() == gin.DebugMode {
	//	docs.SwaggerInfo.Schemes = []string{"http"}
	//	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", serverPort)
	//	docs.SwaggerInfo.BasePath = ""
	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}

	log.Fatal(r.Run(fmt.Sprintf(":%s", serverPort)))
}
