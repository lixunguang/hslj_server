package dao

import (

	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/dto"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

type News struct {
	BaseModel
	Title     string `gorm:"column:title"`
	Content   string `gorm:"column:content"`
	PictureID int    `gorm:"column:picture_id"`
	Publisher string `gorm:"column:publisher"`
}

func (News) TableName() string {
	return "news"
}

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

	/*	result := mysqlDB.Save(&News{BaseModel:BaseModel{ID: news.ID,UpdatedAt: news.UpdatedAt},Title:news.Title,
			Content: news.Content,Publisher: news.Publisher,PictureID:news.PictureID})

		var newsModel News
		result := mysqlDB.Model(&newsModel).Select("ID", "Title","Content","PictureID","Publisher","UpdatedAt").Updates(News{BaseModel: BaseModel{ID: news.ID, UpdatedAt: news.UpdatedAt}, Title: news.Title,
			Content: news.Content, Publisher: news.Publisher, PictureID: news.PictureID})

	*/
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

/*
func GetTitleNews(ctx *gin.Context, number int) []dto.NewsItemRes {
	// 标题新闻
	mysqlDB := mysql.GetDB()

	var news []News
	result := mysqlDB.Order("updated_at desc").Limit(int(number)).Find(&news)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil
	}

	var newsItems []dto.NewsItemRes
	for _, val := range news {
		var item dto.NewsItemRes
		item.NewsID = val.ID
		item.Title = val.Title
		item.Author = val.Publisher
		item.DateStr = val.UpdatedAt.Format(util.FormatDate)
		newsItems = append(newsItems, item)
	}

	return newsItems
}
*/
func GetPictureNews(ctx *gin.Context, number int) []dto.PicNews {
	// 图片新闻
	mysqlDB := mysql.GetDB()

	var picNews []News
	result := mysqlDB.Order("updated_at desc").Limit(number).Where("picture_id <> ?", 0).Find(&picNews)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetPictureNews] fail,err=%+v", result.Error)
		return nil
	}

	var dtoPicNews []dto.PicNews
	for _, val := range picNews {
		var item dto.PicNews
		item.NewsID = val.ID
		item.Title = val.Title
		item.Date = val.UpdatedAt.Format(util.FormatDate)
		item.Author = val.Publisher

		dtoPicNews = append(dtoPicNews, item)
	}

	return dtoPicNews
}

func GetTitleNews(ctx *gin.Context, number int) []dto.TitleNews {
	// 图片新闻
	mysqlDB := mysql.GetDB()

	var titleNews []News
	result := mysqlDB.Order("updated_at desc").Limit(number).Where("picture_id = ?", 0).Find(&titleNews)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetPictureNews] fail,err=%+v", result.Error)
		return nil
	}

	var titleNewsRes []dto.TitleNews
	for _, val := range titleNews {
		var item dto.TitleNews
		item.NewsID = val.ID
		item.Title = val.Title
		item.Date = val.UpdatedAt.Format(util.FormatDate)
		item.Author = val.Publisher
		titleNewsRes = append(titleNewsRes, item)
	}

	return titleNewsRes
}

// 获取一条新闻详情
func GetNewsById(ctx *gin.Context, id dto.IDParam) (dto.NewsResObj, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var news dto.NewsResObj

	res := News{}
	result := mysqlDB.Where("id = ? ", id.ID).First(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return news, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	news.NewsID = res.ID
	news.Title = res.Title
	news.Author = res.Publisher
	news.DateStr = res.UpdatedAt.Format(util.FormatDate)
	news.Content = res.Content
	//news.PictureUrl = GetResourceContentFromID(ctx, res.PictureID)

	return news, nil
}

func GetNewsDetailById(ctx *gin.Context, id dto.IDParam) (News, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	res := News{}
	result := mysqlDB.Where("id = ? ", id.ID).First(&res)
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
