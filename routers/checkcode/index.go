package checkcode

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gowebbase/modules/utils"
	"net/http"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")

	rr.GET("/createcode", createcode)
	rr.GET("/checkcode", checkcode)
	return
}
func createcode(ctx *gin.Context) {
	//config struct for digits

	var store = base64Captcha.DefaultMemStore
	// 如果 Boy 对象没有完全实现 Person 的方法 此处向上转型会报错
	// 而且一旦通过接口限定 golang 不会帮你把变量隐式转为指针
	// 所以这里必须要加 & 或者用 new
	dirverstring := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      40,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source: 		 "1234567890qwertyuioplkjhgfdsazxcvbnm",
		Fonts: []string{"actionj.ttf"},
	}

	c := base64Captcha.NewCaptcha(dirverstring.ConvertFonts(),store)
	id, b64s, err := c.Generate()
	//config struct for audio
	/*
	//声音验证码配置
	var configA = base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height:             60,
		Width:              240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	 */
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 4,  "msg":err.Error()})
		return
	}
	utils.Debug("测试")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": b64s, "captchaId": id, "msg": "success"})
	return
}

func checkcode(ctx *gin.Context) {
	id := ctx.DefaultQuery("codeid", "")
	code := ctx.DefaultQuery("code", "")
	if id=="" || code == ""{
		ctx.JSON(http.StatusOK, gin.H{"code": 4,  "msg":"缺少字段"})
	}
	if !verifycode(id,code,false){
		ctx.JSON(http.StatusOK, gin.H{"code": 4,  "msg":"验证失败"})
	}else{
		ctx.JSON(http.StatusOK, gin.H{"code": 200,  "msg":"验证成功"})
	}
}


func verifycode(id string, code string,clear bool) bool {
	var store = base64Captcha.DefaultMemStore
	return store.Verify(id,code,clear)
}
