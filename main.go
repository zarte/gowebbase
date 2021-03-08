package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gowebbase/modules/utils"
	"gowebbase/routers"
	"io"
	"os"
	"reflect"
	"time"
)


//自定义验证器
type Booking struct {
	Colors []string `form:"colors[]"`
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}
func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func main() {
	//设置全局配置
	utils.GetConfigIni("./conf.ini")

	//设置路由 默认启动方式，包含 Logger、Recovery 中间件
	r := gin.Default()

	rr := routers.GinRouter(r)


	/*
	//自定义验证器注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	 */

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// 创建记录日志的文件
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	rr.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))


	//可以使用fvbock/endless来替换默认的ListenAndServe
	rr.Run(":" + utils.GetIniVal("port",""))
}