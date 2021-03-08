package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gowebbase/modules/utils"
	"io/ioutil"
	"net/http"
)
type login struct {
	Username     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Passwd string `form:"passwd" json:"passwd" xml:"passwd" binding:"required"`
}
func Login(r *gin.Context)  {
	var json login
	if err := r.ShouldBindJSON(&json); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.Username != "manu" || json.Passwd != "123" {
		r.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	token,err:=utils.GenToken(1,json.Username,utils.JwtUserInfo{Other:""})
	if err !=nil{
		fmt.Println(err)
		r.JSON(http.StatusOK, gin.H{"msg": "jwt fail", "code": 4})
		return
	}
	r.JSON(http.StatusOK, utils.ReturnComJson{
		Code: http.StatusOK,
		Msg:  "success",
		Data: gin.H{"username": json.Username, "token": token},
	})
	fmt.Println("login .........")
	return
}

func Demo(r *gin.Context)  {
	result, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		r.JSON(http.StatusOK, gin.H{"msg": "fail", "code": http.StatusOK})
		return
	} else {
		var user map[string]interface{}
		json.Unmarshal(result, &user)
		if user["username"]==nil {
			r.JSON(http.StatusOK, "")
			return
		}

		token,err:=utils.GenToken(1,user["username"].(string),utils.JwtUserInfo{Other:""})
		if err !=nil{
			fmt.Println(err)
			r.JSON(http.StatusOK, gin.H{"msg": "jwt fail", "code": 4})
			return
		}
		r.JSON(http.StatusOK, utils.ReturnComJson{
			Code: http.StatusOK,
			Msg:  "success",
			Data: gin.H{"username": user["username"], "token": token},
		})
	}
	fmt.Println("login .........")
	return
}