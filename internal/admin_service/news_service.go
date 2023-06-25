package admin_service

import (
	"edu-imp/config"
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/dao"
	"edu-imp/pkg/cerror"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取前NewsNumber条新闻,NewsNumber为配置项
func BannerNewsLatest(ctx *gin.Context) ([]admin_dto.PicNews, cerror.Cerror) {

	pictureNewsNumberStr := fmt.Sprintf("%d", config.Vipper.Get("News.PictureNewsNumber"))
	pictureNewsNumber, _ := strconv.ParseInt(pictureNewsNumberStr, 10, 32)
	dtoPicNews := dao.GetPictureNews(ctx, int(pictureNewsNumber))

	return dtoPicNews, nil
}

func BannerNewsALL(ctx *gin.Context) ([]admin_dto.PicNews, cerror.Cerror) {

	pictureNewsNumber := -1
	dtoPicNews := dao.GetPictureNews(ctx, pictureNewsNumber)

	return dtoPicNews, nil
}

func OneNews(ctx *gin.Context, id admin_dto.IDParam) (admin_dto.NewsResObj, cerror.Cerror) {
	return dao.GetNewsById(ctx, id)
}
