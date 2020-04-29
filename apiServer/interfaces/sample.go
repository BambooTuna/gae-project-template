package interfaces

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SampleHandler struct {}

type sampleRequestJson struct {
	Text string `json:"text"`
}

type sampleResponseJson struct {
	Message string `json:"message"`
}

// Sample handler godoc
// @Summary Sample handler
// @Description This is sample handler
// @Accept  json
// @Produce  json
// @Param SignRequestJson body sampleRequestJson true "please set request text"
// @Success 200 {object} sampleResponseJson
// @Failure 400 {object} sampleResponseJson
// @Router /sample [post]
func (s SampleHandler) PostSampleRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestJson sampleRequestJson
		if err := ctx.BindJSON(&requestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, sampleResponseJson{Message:"bind json failed: "+err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, sampleResponseJson{Message:"request text is:"+requestJson.Text})
	}
}
