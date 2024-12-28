package routes

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils"
	"github.com/uzixCode/gocode/utils/changecase"
)

func Create(ctx *gin.Context) {
	var data models.Model
	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{})
		return
	}
	result, err := utils.GenerateStructsFromJSON(data)
	if err != nil {
		utils.RespondingNotFound(ctx, &models.BaseResponse{Message: "Gagal generate models"})
		return
	}
	path := fmt.Sprintf("./models/%v.go", changecase.ToSnake(data.Name))
	utils.CreateOrReplaceFile(path, result)
	cm := exec.Command("go", "fmt", path)
	output, err := cm.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return
	}

	// Print the output
	fmt.Println(string(output))
	utils.Responding(ctx, &models.BaseResponse{Data: result})
}
