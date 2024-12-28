package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/routes/routes/routing/routes"
)

func Routes(r *gin.Engine) {
	protected := r.Group("/routes")
	protected.GET("", routes.Read)
	protected.PUT("/method", routes.UpdateMethod)
	protected.PUT("/path", routes.UpdatePath)
	// protected.GET("/names", routes.ReadName)
	// protected.GET("/:id", routes.Detail)
	// protected.DELETE("/:id", routes.Delete)
}
