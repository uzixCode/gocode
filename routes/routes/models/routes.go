package models

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/routes/routes/models/routes"
)

func Routes(r *gin.Engine) {
	protected := r.Group("/models")
	protected.GET("", routes.Read)
	protected.GET("/names", routes.ReadName)
	protected.POST("", routes.Create)
	protected.GET("/:id", routes.Detail)
	protected.DELETE("/:id", routes.Delete)
}
