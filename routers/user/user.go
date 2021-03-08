package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Userinfo(r *gin.Context)  {
	username,err  :=r.Get("username")
	if !err {
		r.JSON(http.StatusOK, gin.H{"msg": "get username fail", "code": 4})
	}
	r.JSON(http.StatusOK, gin.H{"msg": "userinfo:" + username.(string), "code": http.StatusOK})
	return
}
