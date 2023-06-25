package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func OneNews(ctx *gin.Context, id dto.IDParam) (admin_dto.NewsDetailRes, cerror.Cerror) {

	var res admin_dto.NewsDetailRes

	newRes, err := dao.GetNewsDetailById(ctx, id)

	res.Author = newRes.Publisher
	res.Date = newRes.UpdatedAt.Format(util.FormatDate)
	res.Title = newRes.Title
	res.Content = newRes.Content
	res.PictureID = newRes.PictureID
	res.PictureUrl = dao.GetResourceContentFromID(ctx, newRes.PictureID)
	res.PictureTitle, _ = dao.GetResourceTitleByID(ctx, newRes.PictureID)
	return res, err
}

func AddNews(ctx *gin.Context, newsParam dto.NewsAddParam) (dto.UpdateNewsRes, cerror.Cerror) {

	time, err := time.Parse(util.FormatDate, newsParam.DateStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	news := dao.News{Title: newsParam.Title, Content: newsParam.Content, PictureID: newsParam.PictureID, Publisher: newsParam.Author, BaseModel: dao.BaseModel{UpdatedAt: time}}

	return dao.AddNews(ctx, news)
}

func DelNews(ctx *gin.Context, newsID int) (dto.DelNewsRes, cerror.Cerror) {
	var res dto.DelNewsRes
	var err cerror.Cerror

	res.ID, err = dao.DelNews(ctx, newsID)

	return res, err
}

func UpdateNews(ctx *gin.Context, newsParam dto.NewsUpdateParam) (dto.UpdateNewsRes, cerror.Cerror) {

	time, err1 := time.Parse(util.FormatDate, newsParam.DateStr)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	news := dao.News{Title: newsParam.Title, Content: newsParam.Content, PictureID: newsParam.PictureID,
		Publisher: newsParam.Author, BaseModel: dao.BaseModel{ID: newsParam.ID, UpdatedAt: time}}

	res, err := dao.UpdateNews(ctx, news)

	var newsRes dto.UpdateNewsRes
	newsRes.ID = res.ID
	newsRes.Title = res.Title

	return newsRes, err
}

//获取新闻列表
func NewsTitleALL(ctx *gin.Context) ([]dto.TitleNews, cerror.Cerror) {

	newsNumber := -1
	newsRes := dao.GetTitleNews(ctx, newsNumber)

	return newsRes, nil

}

type FileItem struct {
	Url string `json:"url"`
	ID  int    `json:"id"`
}

//功能同service.UploadFiles，此处做适配
func UploadFiles(ctx *gin.Context) ([]FileItem, cerror.Cerror) {

	var res []FileItem

	//step1:上传文件,得到新文件名字
	upLoadRes, _ := service.UploadMultiFiles(ctx)

	for key, value := range upLoadRes {

		//step2:记录到资源数据库，title（原始名字），desc（"作业文件"），type（设置为0），content（文件路径，包括文件名）
		var resourceParam dto.Resource
		resourceParam.Title = key
		resourceParam.Desc = "作业文件"
		resourceParam.Content = value
		resourceParam.Type = 0

		id, err := dao.AddResource(ctx, resourceParam)
		if err != nil {
			return res, err
		}

		var fileItem FileItem
		fileItem.ID = id
		fileItem.Url = common.GetCDNAddr() + value
		res = append(res, fileItem)

		//res[value] = id
	}

	fmt.Println("service UploadFiles return: ", res)
	return res, nil
}
