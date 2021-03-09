package pay

import (
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	Model "gowebbase/models"
	"net/http"
	"crypto/rsa"
)

type tradeFrom struct {
	OrderNo     string `form:"orderNo" json:"orderNo"  binding:"required"`
	ProductName     string `form:"productName" json:"productName"  binding:"required"`
	OrderPeriod     string `form:"orderPeriod" json:"orderPeriod"  binding:"required"`
	OrderPrice     string `form:"orderPrice" json:"orderPrice"  binding:"required"`
	AppId     string `form:"appId" json:"appId"  binding:"required"`
	Sign     string `form:"sign" json:"sign"  binding:"required"`
}

func Trade(r *gin.Context)  {
	var tradefrom tradeFrom
	if err := r.ShouldBindJSON(&tradefrom); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//获取商户配置
    agentInfo := Model.GetOneAgentByAppid(tradefrom.AppId)
    if agentInfo.Id==0 {
		r.JSON(http.StatusBadRequest, gin.H{"error": "appid no exist!"})
		return
	}

	//验签
	signarg := "appId="+tradefrom.AppId+"&orderNo="+tradefrom.OrderNo+"&orderPeriod="+tradefrom.OrderPeriod+"&orderPrice="+tradefrom.OrderPrice+"&productName="+tradefrom.ProductName
	if !checkSign(tradefrom.Sign,signarg,agentInfo.Pubsec) {
		r.JSON(http.StatusBadRequest, gin.H{"error": "check sign fail!"})
		return
	}

	//判断订单是否存在




	return
}
func checkSign(sign string,content string,pubkey string) bool {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(content))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	msgHashSum := msgHash.Sum(nil)
	block, _ := pem.Decode([]byte(pubkey))
	if block == nil {
		fmt.Println("system err 10001")
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return false
	}
	pub := pubInterface.(*rsa.PublicKey)
	//rsa 签名验证
	err = rsa.VerifyPSS(pub, crypto.SHA256, msgHashSum, []byte(sign),nil)
	if err != nil {
		fmt.Println(err)
		return false
	}else{
		return true
	}
}


