package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Model "gowebbase/models"
	"gowebbase/modules/utils"
	"net/http"
)

func userlist(c *gin.Context)  {
	var params Model.CommonMap = Model.CommonMap{}
	params["username"] = c.DefaultQuery("username", "")
	params["roleid"] = c.DefaultQuery("roleid", "")
	utils.ParsePageAndPageSize(c, params)
	userModel := new(Model.User)
	total, err :=userModel.Total(params)
	if err != nil {
		fmt.Println(err)
	}
	userlist, err :=userModel.GetUserPage(params)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "code": http.StatusOK,"data":map[string]interface{}{
		"total": total,
		"data":  userlist,
	}})
	return
}
