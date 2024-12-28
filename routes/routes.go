package routes

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/routes/routes/models"
	"github.com/uzixCode/gocode/routes/routes/routing"
	"github.com/uzixCode/gocode/routes/routes/sub"
)

func Routes(r *gin.Engine) {
	gocodeFolder, err := getGocodeFolder()
	if err == nil {
		r.Static("/web", gocodeFolder)

	}
	models.Routes(r)
	routing.Routes(r)
	sub.Routes(r)
}

func getGocodeFolder() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Documents", "gocode"), nil
}
