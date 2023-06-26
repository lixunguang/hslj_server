package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

//单个文件上传，not used
func UploadOneFile(ctx *gin.Context) cerror.Cerror {
	var error cerror.Cerror

	// 获取上传的文件
	file, header, err := ctx.Request.FormFile("files")
	if err != nil {
		fmt.Println(err)
		return error
	}
	defer file.Close()

	// 创建保存文件的目录
	err = os.MkdirAll("uploads", 0755)
	if err != nil {
		fmt.Println(err)
		return error
	}

	// 创建保存文件的路径
	filePath := fmt.Sprintf("uploads/%s", header.Filename)

	// 创建保存文件的目标文件
	targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return error
	}
	defer targetFile.Close()

	// 将上传的文件内容复制到目标文件中
	_, err = io.Copy(targetFile, file)
	if err != nil {
		fmt.Println(err)
		return error
	}
	return error
}

//获取原始名字（通过存储名字）
func GetOriginalNameFromStorageName(storageName string) string {
	var originalName string

	return originalName
}

//获取存储名字（通过原始名字）
func GetStorageNameFromOriginalName(originalName string) string {
	var storageName string
	var extName string //扩展名
	index := strings.LastIndex(originalName, ".")
	extName = originalName[index:]

	md5Str := util.EncodeMD5(originalName)
	md5Str = md5Str[0:6]
	now := time.Now()

	storageName = fmt.Sprintf("%d%02d%02d%02d%02d%02d%02d%s", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), md5Str)
	storageName = storageName + extName
	return storageName
}

type FileItem struct {
	Url     string `json:"url"`
	ID      int    `json:"id"`
	FullUrl string `json:"full_url"`
}

//上传多个文件接口
func UploadMultiFiles(ctx *gin.Context) (map[string]string, cerror.Cerror) {
	var error cerror.Cerror

	fileMap := make(map[string]string)

	err := ctx.Request.ParseMultipartForm(1024 * 1024 * 10) //10M
	if err != nil {
		log.Fatal(err)
		return fileMap, cerror.NewCerror(common.Failed, err.Error())
	}
	// 获取表单
	form := ctx.Request.MultipartForm
	logger.Infoc(ctx, "[service.UploadMultiFiles] form.File=%+v, form=%+v", form.File, form)
	fmt.Printf("[service.UploadMultiFiles] form.File=%+v, form=%+v", form.File, form)
	// 获取参数upload后面的多个文件名，存放到数组files里面，
	files := form.File["file"] //files-->file
	// 遍历数组，每取出一个file就拷贝一次
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Fatal(err)
			return fileMap, cerror.NewCerror(common.Failed, err.Error())
		}

		fileName := files[i].Filename
		fmt.Println("input param files:", fileName)

		newFileName := GetStorageNameFromOriginalName(fileName)

		err = os.MkdirAll("uploads", 0755)
		if err != nil {
			fmt.Println(err)
			return fileMap, cerror.NewCerror(common.Failed, err.Error())
		}
		filePath := fmt.Sprintf("uploads/%s", newFileName)

		out, err := os.Create(filePath)
		defer out.Close()
		if err != nil {
			log.Fatal(err)
			return fileMap, cerror.NewCerror(common.Failed, err.Error())
		}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
			return fileMap, cerror.NewCerror(common.Failed, err.Error())
		}

		fileMap[fileName] = filePath

	}

	return fileMap, error
}
