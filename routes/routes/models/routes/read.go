package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
)

func Read(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	result, err := utils.ScanCurrentFolder("./models", "file", ".go", search)
	if err != nil {
		utils.RespondingInternalError(ctx, &models.BaseResponse{Data: err.Error()})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Data: result})
}
