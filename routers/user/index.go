package user

import (
	"github.com/gin-gonic/gin"
	"gowebbase/modules/middleware"
)


func Routers(r *gin.RouterGroup) *gin.RouterGroup{
	rr := r.Group("")
	rr.GET("/selfinfo",middleware.JWTAuthMiddleware(), Selfinfo)
	rr.GET("/userinfo",middleware.JWTAuthMiddleware(), Userinfo)
	rr.POST("/login", Login)
	rr.POST("/loginout",middleware.JWTAuthMiddleware(), Loginout)
	rr.POST("/register", Register)
	return rr
}

