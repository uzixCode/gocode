package entah

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/routes/routes/sub/routes/entah/routes"
)

func Routes(r *gin.RouterGroup) {
	protected := r.Group("/entih")
	protected.DELETE("/ujik", routes.Read)

}
