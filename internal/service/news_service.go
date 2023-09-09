package service

import (
	"github.com/gin-gonic/gin"
	"hslj/pkg/util"

	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func NewsBanner(ctx *gin.Context) ([]dto.NewsLatestRes, cerror.Cerror) {

	var res []dto.NewsLatestRes

	news, err := dao.GetLatestNews(ctx, 3)

	for _, val := range news {
		var item dto.NewsLatestRes
		item.ID = val.ID
		item.DateStr = val.UpdatedAt.Format(util.FormatDate)
		item.Title = val.Title
		item.PictureUrl = val.PictureUrl

		res = append(res, item)
	}

	return res, err
}

func NewsLatest(ctx *gin.Context) ([]dto.NewsLatestRes, cerror.Cerror) {

	var res []dto.NewsLatestRes

	news, err := dao.GetLatestNews(ctx, 2)

	for _, val := range news {
		var item dto.NewsLatestRes
		item.ID = val.ID
		item.DateStr = val.UpdatedAt.Format(util.FormatDate)
		item.Title = val.Title
		item.PictureUrl = val.PictureUrl

		res = append(res, item)
	}

	return res, err
}

func NewsALL(ctx *gin.Context) ([]dto.NewsLatestRes, cerror.Cerror) {

	var res []dto.NewsLatestRes

	news, err := dao.GetLatestNews(ctx, 1000)

	for _, val := range news {
		var item dto.NewsLatestRes
		item.ID = val.ID
		item.DateStr = val.UpdatedAt.Format(util.FormatDate)
		item.Title = val.Title

		res = append(res, item)
	}

	return res, err
}

func NewsDetail(ctx *gin.Context, id dto.IDParam) (dto.NewsDetailRes, cerror.Cerror) {
	var res dto.NewsDetailRes

	news, err := dao.GetNewsDetail(ctx, id.ID)

	res.ID = news.ID
	res.Title = news.Title
	res.DateStr = news.UpdatedAt.Format(util.FormatDate)
	res.Author = news.Author
	res.Content = news.Content
	res.PictureUrl = news.PictureUrl

	return res, err
}
