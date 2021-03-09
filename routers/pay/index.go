package pay

import (
	"github.com/gin-gonic/gin"
)


func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/gateway", Gateway)
	return
}

