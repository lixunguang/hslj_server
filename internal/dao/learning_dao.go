package dao

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Learning struct {
	BaseModel
	Title          string `gorm:"column:title"`
	Desc           string `gorm:"column:desc"`
	Author         string `gorm:"column:author"`
	CategoryID     int    `gorm:"column:category_id"`
	CoverPictureID int    `gorm:"column:cover_picture_id"`
}

func (Learning) TableName() string {
	return "learning"
}

func AddLearning(ctx *gin.Context, param dto.Learning) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	learning := Learning{Title: param.Title, Desc: param.Desc, Author: param.Author, CoverPictureID: param.CoverPictureID, CategoryID: param.CategoryID}
	result := mysqlDB.Create(&learning)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return learning.ID, nil
}

func DelLearning(ctx *gin.Context, id int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Delete(&Learning{}, id)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return id, nil
}

func UpdateLearning(ctx *gin.Context, param admin_dto.UpdateLearning) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var learning Learning
	learning.ID = param.ID
	learning.Author = param.Author
	learning.Title = param.Title
	learning.Desc = param.Desc
	learning.CoverPictureID = param.PictureID
	learning.CategoryID = param.CategoryID

	result := mysqlDB.Model(&Learning{}).Where("id = ?", param.ID).Updates(learning)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return param.ID, nil
}

//获取n个学习资料，根据一定的推荐标准，比如：时间
func GetNlearning(ctx *gin.Context, number int) []Learning {
	mysqlDB := mysql.GetDB()

	var res []Learning

	result := mysqlDB.Order("updated_at desc").Limit(int(number)).Find(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil
	}

	return res
}

//获取学习资料详情
func GetLearning(ctx *gin.Context, id int) (Learning, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var learning Learning

	result := mysqlDB.Where("id = ?", id).First(&learning)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return learning, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return learning, nil
}

//获取n个学习资料，根据一定的推荐标准，比如：时间
func GetNCategoryLearning(ctx *gin.Context, categoryID int, number int) []Learning {
	mysqlDB := mysql.GetDB()

	var res []Learning

	result := mysqlDB.Where("category_id=?", categoryID).Order("updated_at desc").Limit(int(number)).Find(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil
	}

	return res
}

//分页获取学习资源
func GetLearningCategoryResourcePagedData(ctx *gin.Context, pageParam dto.LearningCategoryParam) ([]dto.CategoryResourceItem, cerror.Cerror) {
	var res []dto.CategoryResourceItem

	mysqlDB := mysql.GetDB()

	ps := pageParam.PageSize
	offset := (pageParam.CurrentPage - 1) * (ps)

	var learningList []Learning
	result := mysqlDB.Where("category_id=?", pageParam.CategoryID).Limit(int(ps)).Offset(int(offset)).Find(&learningList)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail,err=%+v, CurrentPage=%d", result.Error, pageParam.CurrentPage)
		return res, cerror.ErrorDataGet
	}

	for _, val := range learningList {
		var item dto.CategoryResourceItem

		item.Name = val.Title
		item.Desc = val.Desc
		item.LearningID = val.ID
		item.Author = val.Author
		item.PictureUrl = GetResourceContentFromID(ctx, val.CoverPictureID)

		res = append(res, item)
	}

	return res, nil
}

//获取分类下的学习资源条数
func GetLearningCategoryResourceCount(ctx *gin.Context, id int) (int64, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var learning []Learning
	var count int64

	result := mysqlDB.Where("category_id=?", id).Model(&learning).Count(&count)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail 2,err=%+v", result.Error)
		return 0, cerror.ErrorDataGet
	}

	return count, nil
}
