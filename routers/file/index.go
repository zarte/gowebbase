package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowebbase/modules/config"
	"gowebbase/modules/utils"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/upload", upload)
	rr.POST("/uploadimg", uploadImg)
	return
}
func upload(ctx *gin.Context) {
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
			return
		}
		// 根据时间鹾生成文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		fileName := fileNameStr + extString
		filePath := config.GetIniVal("Uploaddir","")
		if !utils.Mkdir(filePath) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"message": "创建文件夹失败",
			})
			return
		}
		fmt.Println("22",filePath)
		//client, err := oss.New("Endpoint", "yourAccessKeyId", "yourAccessKeySecret")
		dstFile,err := os.OpenFile(filePath+fileName,os.O_WRONLY | os.O_CREATE,0777)
		if err != nil{
			fmt.Printf("打开目标文件错误，错误信息=%v\n",err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"message": "打开目标文件错误",
			})
			return
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
		_,err = io.Copy(dstFile,fileHandle)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code":0,
				"message": "保存文件失败",
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

func uploadImg(ctx *gin.Context)  {
	if file, err := ctx.FormFile("file"); err == nil {
		//获取文件的后缀名
		extString := path.Ext(file.Filename)
		//允许上传文件的格式
		allowExtMap := map[string]bool {
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".bmp":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extString]; !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"message": "上传文件格式不支持",
			})
			return
		}
		// 根据时间鹾生成文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		fileName := fileNameStr + extString
		filePath := config.GetIniVal("Uploaddir","")
		if !utils.Mkdir(filePath) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"message": "创建文件夹失败",
			})
			return
		}
		dstFile,err := os.OpenFile(filePath+fileName,os.O_WRONLY | os.O_CREATE,0777)
		if err != nil{
			fmt.Printf("打开目标文件错误，错误信息=%v\n",err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"message": "打开目标文件错误",
			})
			return
		}
		defer dstFile.Close()
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
		_,err = io.Copy(dstFile,fileHandle)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code":0,
				"message": "保存文件失败",
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
