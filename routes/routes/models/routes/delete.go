package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
)

func Delete(ctx *gin.Context) {
	err := utils.DeleteFile(fmt.Sprintf("./models/%v", ctx.Param("id")))
	if err != nil {
		utils.RespondingNotFound(ctx, &models.BaseResponse{Message: "File tidak ditemukan"})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{})
}
