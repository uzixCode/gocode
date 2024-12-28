package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
)

func ReadName(ctx *gin.Context) {
	result, err := utils.GetAllStructNamesInFolder("./models")
	if err != nil {
		utils.RespondingInternalError(ctx, &models.BaseResponse{Data: err.Error()})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Data: result})
}
