package admin

import (
	"github.com/gin-gonic/gin"
	"gowebbase/modules/middleware"
)


func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/userlist",middleware.JWTAuthAdminMiddleware(), userlist)
	return
}

