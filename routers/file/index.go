package file

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/upload", Function)

	return
}
func Function(ctx *gin.Context) {
	if file, err := ctx.FormFile("file"); err == nil {
		//获取文件的后缀名
		extString := path.Ext(file.Filename)
		fmt.Println("111", extString)
		//允许上传文件的格式
		allowExtMap := map[string]bool {
			".rar":  true,
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extString]; !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"message": "上传文件格式不支持",
			})
		}
		// 根据时间鹾生成文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		fileName := fileNameStr + extString
		filePath := filepath.Join(utils.Mkdir("static/upload"), "/", fileName)
		fmt.Println("22",filePath)
		//client, err := oss.New("Endpoint", "yourAccessKeyId", "yourAccessKeySecret")
		client, err := oss.New("http://oss-cn-shenzhen.aliyuncs.com", "LTAI4Ff9jV7DfiPrJT36a", "zZOpRqGtKNQl30Su6Ytj12b3IEF")
		if err != nil {
			fmt.Println("阿里云上传错误", err)
			return
		}
		//指定存储空间
		//bucket, err := client.Bucket("yourBucketName")
		bucket, err := client.Bucket("shuiping-code")
		if err != nil {
			fmt.Println("存储空间错误")
			os.Exit(-1)
		}
		//打开文件
		fileHandle, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"message": "打开文件错误",
			})
			return
		}
		defer fileHandle.Close()
		fileByte,_:= ioutil.ReadAll(fileHandle)
		//上传到oss上
		err = bucket.PutObject(filePath, bytes.NewReader(fileByte))
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code":0,
				"message": "解析错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":0,
			"message": "上传成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":0,
			"message": "上传失败",
		})
	}
	return
}
