package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/service"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

//上传文件，资源id
func UploadFiles(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "UploadFiles Controller")

	fileMap, err := service.UploadMultiFiles(ctx)

	if err == nil {
		util.SuccessJson(ctx, fileMap)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AddUser Controller", fileMap, err)
}
