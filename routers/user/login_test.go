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

func TestLoginHandler(t *testing.T) {
	resp := OrdinaryResponse{}
	user := map[string]interface{}{
		"username": "Valiben",
	}

	err := utils.TestHandlerUnMarshalResp("POST", "/login", "json", user, &resp)
	if err != nil {
		t.Error(err)
		res,err := utils.TestOrdinaryHandler("POST", "/login", "json", user)
		fmt.Println(string(res))
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if resp.Code != 200 {
		t.Errorf(resp.Msg)
		return
	}

}

