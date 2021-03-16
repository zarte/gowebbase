package user

import (
	"fmt"
	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"testing"
)

func init() {
	router := gin.Default()  // 这需要写到init中，启动gin框架
	rr := router.Group("/")
	Routers(rr)
	utils.SetRouter(router)  //把启动的engine 对象传入到test框架中
}

type OrdinaryResponse struct {
	Code  int `json:"code"`
	Msg string `json:"msg"`
}

func TestUserinfoHandler(t *testing.T) {
	resp := OrdinaryResponse{}
	user := map[string]interface{}{
		"username": "Valiben",
	}
	utils.AddHeader("Authorization","sdfd")
	err := utils.TestHandlerUnMarshalResp("Get", "/userinfo", "FORM", user, &resp)
	if err != nil {
		t.Error(err)
		res,err := utils.TestOrdinaryHandler("Get", "/userinfo", "FORM", user)
		if err != nil {
			return
		}
		fmt.Println(string(res))
		return
	}

	if resp.Code != 200 {
		t.Errorf(resp.Msg)
		return
	}

}

/*
func TestLoginHandler(t *testing.T) {
	resp := OrdinaryResponse{}
	user := map[string]interface{}{
		"username": "Valiben",
		"password": "123456",
		"age":      13,
	}
	err := utils.TestHandlerUnMarshalResp("POST", "/login", "json", user, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}

	if resp.Errno != "0" {
		t.Errorf("TestLoginHandler: response is not expected\n")
		return
	}
}
*/