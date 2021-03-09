package pay

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
支付网关
 */
func Gateway(r *gin.Context)  {
	method,err  :=r.Get("method")
	if !err {
		r.JSON(http.StatusOK, gin.H{"msg": "get method fail", "code": 4})
	}
	//路由转发
	switch method {
	case "trade.page.pay":
		Trade(r)
		break
	default:
		r.JSON(http.StatusOK, gin.H{"msg": "get method fail", "code": 4})
		break
	}
	return
}

