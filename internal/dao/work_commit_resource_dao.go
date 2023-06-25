package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

//用户提交资源表

type WorkCommitResource struct {
	BaseModel
	WorkCommitID int `gorm:"column:work_commit_id"` //用户提交记录id
	ResourceID   int `gorm:"column:resource_id"`    //资源id
	Type         int `gorm:"column:type"`           //资源所属类型，如提交作业图片，留言 取值为43和44
	CanModify    int `gorm:"column:can_modify"`     //是否可以再次修改，1 可以，2 不可以
}

func (WorkCommitResource) TableName() string {
	return "work_commit_resource"
}

//增加作业提交资源
func AddWorkCommitResource(ctx *gin.Context, param dto.WorkCommitResource) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	workCommitResource := WorkCommitResource{WorkCommitID: param.WorkCommitID, ResourceID: param.ResourceID, Type: param.Type} //, CanModify: param.CanModify
	result := mysqlDB.Create(&workCommitResource)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return workCommitResource.ID, nil
}

//删除作业提交资源 返回删除的文件id和富文本id
func DelWorkCommitResource(ctx *gin.Context, workCommitResource []WorkCommitResource) ([]int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var resIDs []int
	for _, val := range workCommitResource {
		resIDs = append(resIDs, val.ID)
	}
	result := mysqlDB.Delete(&workCommitResource)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return resIDs, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return resIDs, nil
}

func GetWorkCommitResourceRecords(ctx *gin.Context, workCommitIDs []int) ([]dto.WorkCommitResourceID, []dto.WorkCommitResourceID, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var fileCommitResource []dto.WorkCommitResourceID
	var richTextCommitResource []dto.WorkCommitResourceID

	var records []WorkCommitResource
	result := mysqlDB.Where("work_commit_id IN ?", workCommitIDs).Find(&records)

	for _, val := range records {
		if val.Type == common.WorkCommitAttachment {
			fileCommitResource = append(fileCommitResource, dto.WorkCommitResourceID{ID: val.ID, WorkCommitID: val.WorkCommitID, ResourceID: val.ResourceID})
		} else if val.Type == common.WorkCommitText {
			richTextCommitResource = append(richTextCommitResource, dto.WorkCommitResourceID{ID: val.ID, WorkCommitID: val.WorkCommitID, ResourceID: val.ResourceID})
		} else {
			logger.Errorc(ctx, "wrong workcommit type: %d %d", val.Type, val.ResourceID)
		}

	}

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return fileCommitResource, richTextCommitResource, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return fileCommitResource, richTextCommitResource, nil
}

//获取作业提交资源
func GetWorkCommitResourceByCommitID(ctx *gin.Context, id int) ([]dto.WorkCommitResource, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res []dto.WorkCommitResource

	var resources []WorkCommitResource
	result := mysqlDB.Where("work_commit_id = ? ", id).Find(&resources)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}
	for _, val := range resources {
		var item dto.WorkCommitResource

		item.WorkCommitID = val.WorkCommitID
		item.ResourceID = val.ResourceID
		item.Type = val.Type
		//item.CanModify = val.CanModify

		res = append(res, item)
	}

	return res, nil
}
