package api


import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/api", Function)
	return
}
func Function(c *gin.Context) {
	fmt.Println("api .........")
	return
}