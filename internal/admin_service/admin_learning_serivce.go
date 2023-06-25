package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func AddLearning(ctx *gin.Context, param dto.Learning) (int, cerror.Cerror) {
	return dao.AddLearning(ctx, param)
}

func DelLearning(ctx *gin.Context, param dto.IDParam) (int, cerror.Cerror) {
	//todo:删除关联？
	return dao.DelLearning(ctx, param.ID)

}

func UpdateLearning(ctx *gin.Context, param admin_dto.UpdateLearning) (int, cerror.Cerror) {

	return dao.UpdateLearning(ctx, param)
}

//根据分类id获取推荐的资源
func GetLearningAll(ctx *gin.Context) ([]admin_dto.LearningRes, cerror.Cerror) {

	var res []admin_dto.LearningRes

	showNumber := 500 //最大500条记录
	learningList := dao.GetNlearning(ctx, showNumber)

	for _, val := range learningList {
		var learningRes admin_dto.LearningRes

		categoryRes, _ := dao.GetLearningCategoryByID(ctx, val.CategoryID)

		learningRes.Category = categoryRes.Title
		learningRes.ID = val.ID
		learningRes.Title = val.Title
		learningRes.Desc = val.Desc
		learningRes.Author = val.Author
		learningRes.PictureUrl = dao.GetResourceContentFromID(ctx, val.CoverPictureID)

		res = append(res, learningRes)
	}

	return res, nil
}

//获取资源详细信息
func GetLearningByID(ctx *gin.Context, id int) admin_dto.LearningDetailRes {

	var res admin_dto.LearningDetailRes

	getRes, _ := dao.GetLearning(ctx, id)
	res.ID = getRes.ID
	res.Desc = getRes.Desc
	res.Title = getRes.Title
	res.Author = getRes.Author
	res.PictureName, _ = dao.GetResourceTitleByID(ctx, getRes.CoverPictureID)
	res.PictureUrl = dao.GetResourceContentFromID(ctx, getRes.CoverPictureID)
	res.PictureID = getRes.CoverPictureID
	categoryRes, _ := dao.GetLearningCategoryByID(ctx, getRes.CategoryID)
	res.Category = categoryRes.Title

	res.ResourceList, _ = GetLearningResource(ctx, id)

	return res
}

func GetLearningCategory(ctx *gin.Context) ([]dto.LearningCategory, cerror.Cerror) {

	return dao.GetLearningCategory(ctx)

}

func AddLearningCategory(ctx *gin.Context, param dto.LearningCategory) (int, cerror.Cerror) {

	return dao.AddLearningCategory(ctx, param)
}

func UpdateLearningCategory(ctx *gin.Context, param dto.UpdateLearningCategoryParam) (int, cerror.Cerror) {

	return dao.UpdateLearningCategory(ctx, param)
}

func DelLearningCategory(ctx *gin.Context, id int) (int, cerror.Cerror) {

	return dao.DelLearningCategory(ctx, id)
}

func GetLearningResource(ctx *gin.Context, id int) ([]admin_dto.ResourceRes, cerror.Cerror) {

	var res []admin_dto.ResourceRes

	resList, err := dao.GetLearningResourceList(ctx, id)
	for _, val := range resList {
		var item admin_dto.ResourceRes
		item.ID = val.ResourceID
		item.Desc = val.Desc
		item.Title = val.Title
		item.ContentUrl = dao.GetResourceContentFromID(ctx, val.ResourceID)
		item.Index = val.Index
		resourceType, _ := dao.GetResourceTypeByID(ctx, val.ResourceID)
		item.Type = int(resourceType)

		res = append(res, item)
	}

	return res, err
}

func AddLearningResource(ctx *gin.Context, param admin_dto.AddLearningResourceParam) (int, cerror.Cerror) {

	title, _ := dao.GetResourceTitleByID(ctx, param.ResourceID)
	desc, _ := dao.GetResourceDescByID(ctx, param.ResourceID)

	return dao.AddLearningResource(ctx, title, desc, param)
}

func DelLearningResource(ctx *gin.Context, param admin_dto.DelLearningResourceParam) (int, cerror.Cerror) {

	return dao.DelLearningResource(ctx, param)
}
