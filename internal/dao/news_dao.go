package dao

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
)

type News struct {
	BaseModel
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	PictureUrl string `gorm:"column:picture_url"`
	Author     string `gorm:"column:author"`
	ViewCount  string `gorm:"column:view_count"`
}

func (News) TableName() string {
	return "news"
}



func GetLatestNews(ctx *gin.Context, number int) ([]News,cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var news []News
	result := mysqlDB.Order("updated_at desc").Limit(number).Find(&news)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetLatestNews] fail,err=%+v", result.Error)
		return nil, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return news,nil
}

// 获取一条新闻详情
func GetNewsDetail(ctx *gin.Context, id int) (News, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	res := News{}
	result := mysqlDB.Where("id = ? ", id).First(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}


	return res, nil
}


// 获取新闻条数
func GetNewsCount(ctx *gin.Context) (int64, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var news []News
	var count int64

	result := mysqlDB.Model(&news).Count(&count)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail 2,err=%+v", result.Error)
		return 0, cerror.ErrorDataGet
	}

	return count, nil
}


/*
// 增加新闻
func AddNews(ctx *gin.Context, news News) (dto.UpdateNewsRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	// add

	result := mysqlDB.Create(&news)

	var res dto.UpdateNewsRes
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	res.ID = news.ID
	res.Title = news.Title

	return res, nil
}

func DelNews(ctx *gin.Context, id int) (int, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	result := mysqlDB.Delete(&News{}, id)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return id, cerror.NewCerror(common.InvalidID, result.Error.Error())
	}

	return id, nil
}

// 更新news
func UpdateNews(ctx *gin.Context, news News) (News, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var newsModel News
	param := News{
		BaseModel: BaseModel{
			ID:        news.ID,
			UpdatedAt: news.UpdatedAt,
		},
		Title:     news.Title,
		Content:   news.Content,
		Publisher: news.Publisher,
		PictureID: news.PictureID,
	}

	result := mysqlDB.Model(&newsModel).Where("id = ?", news.ID).Updates(param)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return news, cerror.NewCerror(common.InvalidID, result.Error.Error())
	}

	return news, nil
}
// 获取分页数据
func GetNewsPagedData(ctx *gin.Context, param dto.NewsAllParam) ([]News, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	// pageSize := fmt.Sprintf("%d", config.Vipper.Get("News.PageSize"))
	// ps, _ := strconv.ParseInt(pageSize, 10, 64)

	ps := param.PageSize

	var news []News

	offset := (param.CurrentPage - 1) * (ps)

	result := mysqlDB.Order("updated_at desc").Limit(int(ps)).Offset(int(offset)).Find(&news) // todo：不需要content信息，可以优化
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail,err=%+v, CurrentPage=%d", result.Error, param.CurrentPage)
		return nil, cerror.ErrorDataGet
	}

	return news, nil
}

*/
