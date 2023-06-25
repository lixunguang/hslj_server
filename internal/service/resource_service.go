package service

import (
	"edu-imp/config"
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddResource(ctx *gin.Context, param dto.Resource) (int, cerror.Cerror) {

	var res int = 0

	if param.Type == common.RichType { //富文本类型时，先插入富文本表
		id, err := dao.AddResourceRichText(ctx, dto.ResourceRichText{param.Content})
		if err != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err=")
			return res, err
		}

		idstr := fmt.Sprintf("%d", id)
		param.Content = idstr

		res, err = dao.AddResource(ctx, param)
		if err != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err=")
			return res, err
		}

	} else {
		var err cerror.Cerror
		res, err = dao.AddResource(ctx, param)
		if err != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err=")
			return res, err
		}
	}

	return res, nil
}

func DelResource(ctx *gin.Context, param admin_dto.ResouceDelParam) (int, cerror.Cerror) {

	var res int = 0

	if param.Type == common.RichType { //富文本类型时，先插入富文本表

		contentID, _ := dao.GetRichtextIdByID(ctx, param.ID)

		var arr = []int{param.ID}
		res, err := dao.DelResource(ctx, arr)
		if err != nil {
			fmt.Println(res)
		}

		//删除富文本
		delRes, delErr := dao.DelResourceRichtext(ctx, contentID)
		if delErr != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err=")
			return delRes, delErr
		}
	} else { //除了richtype之外其他的类型
		filePath, _ := dao.GetFilePathByID(ctx, param.ID)

		var arr = []int{param.ID}
		res, err := dao.DelResource(ctx, arr)
		if err != nil {
			fmt.Println(res)
		}

		filePath = config.GetExecDirectory() + "/" + filePath
		if util.IsExist(filePath) {
			//删除文件
			isDel, err2 := util.RemoveFile([]string{filePath})

			fmt.Println(isDel, err2)
		}

	}

	return res, nil
}

func UpdateResource(ctx *gin.Context, param dto.UpdateResourceParam) (int, cerror.Cerror) {

	var res int = 0

	if param.Type == common.RichType { //富文本类型时，先插入富文本表
		contentID, _ := dao.GetRichtextIdByID(ctx, param.ID)
		dao.UpdateResource(ctx, param)

		var param2 dto.UpdateRichtext
		param2.ID = contentID
		param2.Content = param.Content
		dao.UpdateResourceRichtext(ctx, param2)
	}

	return res, nil
}
