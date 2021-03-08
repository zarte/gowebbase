package user

import (
	"github.com/gin-gonic/gin"
	"gowebbase/modules/middleware"
)


func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.GET("/userinfo",middleware.JWTAuthMiddleware(), Userinfo)
	rr.POST("/login", Login)
	return
}

