package routers

import (
	"fmt"
	"gowebbase/routers/admin"
	"gowebbase/routers/api"
	"gowebbase/routers/user"
	"gowebbase/routers/demo"
	"gowebbase/routers/file"
	"gowebbase/routers/pay"
	"github.com/gin-gonic/gin"
)

func GinRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/")
	rr.GET("/first", func(c *gin.Context) {
		fmt.Println("first .........")
	})
	user.Routers(rr)
	// authorized.Use(AuthRequired())
	rr = r.Group("/admin")
	admin.Routers(rr)
	rr = r.Group("/api")
	api.Routers(rr)
	rr = r.Group("/demo")
	demo.Routers(rr)
	rr = r.Group("/file")
	file.Routers(rr)
	rr = r.Group("/pay")
	pay.Routers(rr)
	return r
}