package controller

import (
	"edu-imp/internal/service"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {

	service.UploadOneFile(ctx)

	util.SuccessJson(ctx, nil)
}

//上传文件，资源id
func UploadFiles(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "UploadFiles Controller")

	fileMap, err := service.UploadFiles(ctx)

	if err == nil {
		util.SuccessJson(ctx, fileMap)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AddUser Controller", fileMap, err)
}
