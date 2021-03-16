package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Model "gowebbase/models"
	"net/http"
	"strconv"
)

func Userinfo(r *gin.Context)  {
	uidstr :=r.DefaultQuery("uid", "")
	if uidstr =="" {
		r.JSON(http.StatusOK, gin.H{"msg": "userid no set", "code": 4})
	}
	uid,err :=strconv.Atoi(uidstr)
	if err != nil {
		r.JSON(http.StatusOK, gin.H{"msg": err.Error(), "code": 4})
	}
	//查询数据库
	var userModel Model.User
	userInfo,err := userModel.GetUserById(uid)
	if err!=nil {
		fmt.Println(err)
		r.JSON(http.StatusOK, gin.H{"msg": "get userInfo fail", "code": 4})
	}
	r.JSON(http.StatusOK, gin.H{"msg": "success", "code": http.StatusOK,"data":userInfo})
	return
}

func Selfinfo(r *gin.Context)  {
	userid,bres  :=r.Get("userid")
	if !bres {
		r.JSON(http.StatusOK, gin.H{"msg": "get userid fail", "code": 4})
	}
	//查询数据库
	var userModel Model.User
	userInfo,err := userModel.GetUserById(userid.(int))
	if err!=nil {
		fmt.Println(err)
		r.JSON(http.StatusOK, gin.H{"msg": "get userInfo fail", "code": 4})
	}
	r.JSON(http.StatusOK, gin.H{"msg": "success", "code": http.StatusOK,"data":userInfo})
	return
}