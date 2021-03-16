package Model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowebbase/modules/config"
	"strconv"
)

type CommonMap map[string]interface{}


type BaseModel struct {
	Page     int `xorm:"-"`
	PageSize int `xorm:"-"`
}

func (model *BaseModel) pageLimitOffset() int {
	return (model.Page - 1) * config.GetIniValInt("PageSize","")
}



// ParsePageAndPageSize 解析查询参数中的页数和每页数量
func ParsePageAndPageSize(ctx *gin.Context, params CommonMap) {
	rpage := ctx.DefaultQuery("page","0")
	rpageSize := ctx.DefaultQuery("page_size","0")

	page,err:= strconv.Atoi(rpage)
	if err != nil{
		fmt.Println(err)
		page = 1
	}
	pageSize,err:= strconv.Atoi(rpageSize)
	if err != nil{
		fmt.Println(err)
		pageSize = config.GetIniValInt("PageSize","")
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = config.GetIniValInt("PageSize","")
	}

	params["Page"] = page
	params["PageSize"] = pageSize
}