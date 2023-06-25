package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type WorkCommit struct {
	BaseModel
	LessonID int `gorm:"column:lesson_id"` //课程id
	UserID   int `gorm:"column:user_id"`   //用户id
}

func (WorkCommit) TableName() string {
	return "work_commit"
}

//获取作业提交id
func GetWorkCommitID(ctx *gin.Context, param dto.WorkCommitUserID) (int, cerror.Cerror) {
	var id int = 0

	mysqlDB := mysql.GetDB()

	var workCommit WorkCommit
	result := mysqlDB.Where("lesson_id = ? and user_id=? ", param.LessonID, param.UserID).First(&workCommit)

	if result.Error != nil {
		logger.Warnc(ctx, "[dao.GetWorkCommitID] fail 2,param=%+v, err=%+v", param, result.Error)
		if result.Error.Error() == "record not found" {
			return id, nil
		} else {
			return id, cerror.NewCerror(common.Failed, result.Error.Error())
		}
	}

	id = workCommit.ID

	return id, nil
}

//增加作业提交
func AddWorkCommit(ctx *gin.Context, param dto.WorkCommitParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	id, _ := GetIDByLoginID(ctx, param.LoginID)
	workCommit := WorkCommit{LessonID: param.LessonID, UserID: id}
	result := mysqlDB.Create(&workCommit)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return workCommit.ID, nil
}

//删除作业提交记录，返回已经删除的记录id
func DelWorkCommit(ctx *gin.Context, param dto.WorkCommitUserID) ([]int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res []int //workcommit ids
	//根据用户及章节查找到作业提交记录
	items, _ := GetWorkCommitRecords(ctx, param)
	if len(items) == 0 { //没有记录，直接返回
		return res, nil
	}

	var workCommits []WorkCommit
	for _, val := range items {
		var item WorkCommit
		item.ID = val.ID

		workCommits = append(workCommits, item)
		res = append(res, val.ID)
	}

	//删除记录
	result := mysqlDB.Delete(&workCommits)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return res, nil
}

//根据用户及章节查找到作业提交记录
func GetWorkCommitRecords(ctx *gin.Context, param dto.WorkCommitUserID) ([]WorkCommit, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var records []WorkCommit
	result := mysqlDB.Where("lesson_id = ? and user_id = ?", param.LessonID, param.UserID).Find(&records)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return records, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return records, nil
}
