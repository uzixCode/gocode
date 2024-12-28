package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
)

func Read(ctx *gin.Context) {
	// search := ctx.DefaultQuery("search", "")
	route, _, err := utils.ScanningRoute("./routes/routes.go")
	if err != nil {
		utils.RespondingInternalError(ctx, &models.BaseResponse{Message: err.Error()})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Data: route})
}
