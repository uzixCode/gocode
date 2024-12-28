package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
)

func UpdatePath(ctx *gin.Context) {
	var dataMap map[string]interface{}
	if err := ctx.ShouldBindJSON(&dataMap); err != nil {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{})
		return
	}
	location, ok := dataMap["route_location"]
	if !ok {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{Message: "Location Cannot be empty"})
		return
	}
	method, ok := dataMap["method"]
	if !ok {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{Message: "Method can't be empty"})
		return
	}
	path, ok := dataMap["path"]
	if !ok {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{Message: "Path can't be empty"})
		return
	}
	to, ok := dataMap["change"]
	if !ok {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{Message: "Path can't be empty"})
		return
	}

	err := utils.ModifyRoutesFile(location.(string), method.(string), path.(string), to.(string))
	if err != nil {
		utils.RespondingInternalError(ctx, &models.BaseResponse{Message: err.Error()})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{})
}
