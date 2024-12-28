package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
)

func Detail(ctx *gin.Context) {
	file, err := utils.LoadFile(fmt.Sprintf("./models/%v", ctx.Param("id")))
	if err != nil {
		utils.RespondingNotFound(ctx, &models.BaseResponse{Message: "File tidak ditemukan"})
		return
	}
	result, err := utils.ModelsTOJson(file)
	if err != nil {
		utils.RespondingNotFound(ctx, &models.BaseResponse{Message: "Gagal membaca models"})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Data: result})
}
