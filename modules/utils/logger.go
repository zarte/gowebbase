package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Debug(str string)  {
	fmt.Fprintln(gin.DefaultWriter, "[DEBUG]",
		time.Now().Format("2006/01/02 - 15:04:05"),
		str)
}

func Info(str string)  {
	fmt.Fprintln(gin.DefaultWriter, "[INFO]",
		time.Now().Format("2006/01/02 - 15:04:05"),
		str)
}
func Warm(str string)  {
	fmt.Fprintln(gin.DefaultWriter, "[WARM]",
		time.Now().Format("2006/01/02 - 15:04:05"),
		str)
}