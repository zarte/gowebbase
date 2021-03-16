package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	Model "gowebbase/models"
	"gowebbase/modules/utils"
	"io/ioutil"
	"net/http"
)
type login struct {
	Username     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Passwd string `form:"passwd" json:"passwd" xml:"passwd" binding:"required"`
}
type register struct {
	Username     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Passwd string `form:"passwd" json:"passwd" xml:"passwd" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
	Checkcode string `form:"checkcode" json:"checkcode" binding:"required"`
}
func Login(r *gin.Context)  {
	var json login
	if err := r.ShouldBindJSON(&json); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"msg": "login parmas fail", "code": 4})
		return
	}
	info,err := Model.GetUserByName(json.Username)
	if err !=nil {
		r.JSON(http.StatusOK, gin.H{"msg": "login fail 10001", "code": 4})
		return
	}
	if info.Username =="" {
		r.JSON(http.StatusOK, gin.H{"msg": "login fail 10002", "code": 4})
		return
	}
	if utils.MD5(utils.MD5(json.Passwd) + info.Salt) != info.Passwd {
		r.JSON(http.StatusOK, gin.H{"msg": "login fail 10003", "code": 4})
		return
	}
	//获取权限
	roles,err := Model.GetUroles(info.Id)
	token,err:=utils.GenToken(info.Id,json.Username,utils.JwtUserInfo{Other:"",Role:roles})
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
	return
}

func Register(r *gin.Context)  {
	var json register
	if err := r.ShouldBindJSON(&json); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !utils.RegUsername([]byte(json.Username)){
		r.JSON(http.StatusOK, gin.H{"msg": "用户名格式不符要求", "code": 4})
		return
	}
	if !utils.RegPasswd([]byte(json.Passwd)){
		r.JSON(http.StatusOK, gin.H{"msg": "密码格式不符要求", "code": 4})
		return
	}
	if !utils.RegEmail([]byte(json.Email)){
		r.JSON(http.StatusOK, gin.H{"msg": "邮箱格式不符要求", "code": 4})
		return
	}
	//判断用户名是否重复
	if Model.CheckUserNameExist(json.Username) {
		r.JSON(http.StatusOK, gin.H{"msg": "该用户名已存在", "code": 4})
		return
	}
	var info Model.User
	info.Username = json.Username
	info.Salt = utils.CreateCaptcha6()
	info.Passwd = utils.MD5(utils.MD5(json.Passwd)+info.Salt)
	info.Roles = "0"
	info.Email = json.Email
	uid,err := Model.AddUserInfo(info)
	if err != nil{
		r.JSON(http.StatusOK, gin.H{"msg": "注册失败,10001", "code": 4})
		return
	}
	//添加默认角色
	Model.UpdateUrole(uid,[]int{0})
	r.JSON(http.StatusOK, gin.H{"msg": "success", "code": 200})
	return
}

func Loginout(r *gin.Context)  {

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