package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/demo", Function)

	rr.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	return
}
func Function(c *gin.Context) {
	fmt.Println("demo .........")
    c.DefaultQuery("firstname", "Guest")
	c.Query("lastname")
	 c.PostForm("message")
	c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

	return
}
