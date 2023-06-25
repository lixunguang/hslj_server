package service

import (
	"edu-imp/config"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLearningCategory(ctx *gin.Context) ([]dto.LearningCategory, cerror.Cerror) {

	return dao.GetLearningCategory(ctx)

}

//首页获取推荐学习资源
func GetLearningRecommend(ctx *gin.Context) []dto.LearningRecommendRes {

	var res []dto.LearningRecommendRes

	showNumberStr := fmt.Sprintf("%d", config.Vipper.Get("learning.ShowNumber"))
	showNumber, _ := strconv.ParseInt(showNumberStr, 10, 32)
	learningList := dao.GetNlearning(ctx, int(showNumber))

	for _, val := range learningList {
		var learningRecommend dto.LearningRecommendRes
		//learningRecommend.CategoryID = val.CategoryID
		learningRecommend.LearningID = val.ID
		learningRecommend.Name = val.Title
		learningRecommend.Desc = val.Desc
		learningRecommend.Author = val.Author
		learningRecommend.PictureUrl = dao.GetResourceContentFromID(ctx, val.CoverPictureID)

		res = append(res, learningRecommend)
	}

	return res
}

//获取资源详细信息
func GetLearningByID(ctx *gin.Context, id int) dto.LearningInfoRes {

	var res dto.LearningInfoRes

	//var learningRes dao.Learning
	learningRes, _ := dao.GetLearning(ctx, id)

	res.LearningRes.Author = learningRes.Author
	res.LearningRes.Title = learningRes.Title
	res.LearningRes.Desc = learningRes.Desc

	resourceList, _ := dao.GetLearningResourceList(ctx, id)
	for _, v := range resourceList {
		var item dto.LearningItem
		item.Title = v.Title
		item.Desc = v.Desc
		item.ResourceUrl = dao.GetResourceContentFromID(ctx, v.ResourceID)
		item.Index = v.Index

		res.ResourceList = append(res.ResourceList, item)
	}

	return res
}

//根据分类id获取推荐的学习资源
func GetLearningCategoryResourceRecommend(ctx *gin.Context, param dto.IDParam) ([]dto.LearningRecommendRes, cerror.Cerror) {

	var res []dto.LearningRecommendRes

	showNumberStr := fmt.Sprintf("%d", config.Vipper.Get("learning.ShowNumber"))
	showNumber, _ := strconv.ParseInt(showNumberStr, 10, 32)
	learningList := dao.GetNCategoryLearning(ctx, int(param.ID), int(showNumber))

	for _, val := range learningList {
		var learningRecommend dto.LearningRecommendRes
		//learningRecommend.CategoryID = val.CategoryID
		learningRecommend.LearningID = val.ID
		learningRecommend.Name = val.Title
		learningRecommend.Desc = val.Desc
		learningRecommend.Author = val.Author
		learningRecommend.PictureUrl = dao.GetResourceContentFromID(ctx, val.CoverPictureID)

		res = append(res, learningRecommend)
	}

	return res, nil
}

//分页获取分类下的所有学习资源
func GetLearningCategoryResourcePagedData(ctx *gin.Context, param dto.LearningCategoryParam) (dto.LearningCategoryResourceRes, cerror.Cerror) {

	categoryResource, err := dao.GetLearningCategoryResourcePagedData(ctx, param)
	count, _ := dao.GetLearningCategoryResourceCount(ctx, param.CategoryID)

	var res dto.LearningCategoryResourceRes
	res.TotalNumber = count
	res.CategoryResourceList = categoryResource

	return res, err

}
