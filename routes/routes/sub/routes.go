package sub

import (
	"github.com/gin-gonic/gin"
	"github.com/uzixCode/gocode/routes/routes/sub/routes/entah"
)

func Routes(r *gin.Engine) {
	protected := r.Group("/sib")
	entah.Routes(protected)

}
