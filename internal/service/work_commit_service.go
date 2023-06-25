package service

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

//获取作业提交信息，
func WorkCommit(ctx *gin.Context, param dto.WorkCommitParam) (dto.WorkCommitRes, cerror.Cerror) {
	var res dto.WorkCommitRes

	id, _ := dao.GetIDByLoginID(ctx, param.LoginID)
	var param2 dto.WorkCommitUserID
	param2.LessonID = param.LessonID
	param2.UserID = id
	//step1: 获取提交id
	id, err := dao.GetWorkCommitID(ctx, param2)
	if err != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", err.Error())
		return res, cerror.NewCerror(common.Failed, err.Error())
	}

	if id == 0 { //没有提交记录
		//todo:
		//todo:这种处理方法对吗
		//todo:
		return res, cerror.NewCerror(0, "record not found")
	}

	//step2: 根据id获取资源信息的id
	resources, err2 := dao.GetWorkCommitResourceByCommitID(ctx, id)
	if err2 != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", err2.Error())
		return res, cerror.NewCerror(common.Failed, err2.Error())
	}

	//返回值
	res.UserLoginID = param.LoginID
	res.LessonID = param.LessonID

	for _, val := range resources {
		if val.Type == common.WorkCommitAttachment { //文件
			var item dto.WorkCommitItem
			item.Type = val.Type
			item.ContentID = val.ResourceID
			item.Content = dao.GetResourceContentFromID(ctx, val.ResourceID)
			item.FileName, _ = dao.GetResourceTitleByID(ctx, val.ResourceID)
			res.WorkCommitFiles = append(res.WorkCommitFiles, item)

		} else if val.Type == common.WorkCommitText { //附言
			res.WorkCommitMessage = dao.GetResourceContentFromID(ctx, val.ResourceID)
		}

	}

	return res, nil
}

//作业提交，此接口会覆盖已有的数据 ，客户端应该给出用户提示
func WorkCommitAdd(ctx *gin.Context, param dto.WorkCommitAddParam) ([]int, cerror.Cerror) {
	var res []int
	var workCommit dto.WorkCommitParam
	workCommit.LoginID = param.UserLoginID
	workCommit.LessonID = param.LessonID

	var workCommituserID dto.WorkCommitUserID
	workCommituserID.UserID, _ = dao.GetIDByLoginID(ctx, param.UserLoginID)
	workCommituserID.LessonID = param.LessonID

	//step0:删除已有记录
	DelWorkCommitAllInfoByLessonIDUserID(ctx, workCommituserID)

	//step1:记录 更新workcommit表
	commitID, err := dao.AddWorkCommit(ctx, workCommit)
	if err != nil {
		logger.Errorc(ctx, "WorkCommitAdd error")
		return res, err
	}

	var workCommitRes dto.WorkCommitResource

	//step2: 文件 更新workcommitResource表
	for _, item := range param.WorkCommitFiles { //文件
		workCommitRes.WorkCommitID = commitID
		workCommitRes.ResourceID = item.ContentID
		workCommitRes.Type = common.WorkCommitAttachment

		//更新work commit resource表
		resourceID, _ := dao.AddWorkCommitResource(ctx, workCommitRes)
		res = append(res, resourceID)
	}

	//step3:富文本  更新resource_richtext 富文本资源表
	var resourceParam dto.Resource
	resourceParam.Content = param.WorkCommitMessage
	resourceParam.Desc = "富文本"
	resourceParam.Type = common.RichType
	resourceParam.Title = "富文本"
	id, err := AddResource(ctx, resourceParam)

	if err != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err")
		return res, err
	}
	workCommitRes.WorkCommitID = commitID
	workCommitRes.ResourceID = id
	workCommitRes.Type = common.WorkCommitText

	resourceID, _ := dao.AddWorkCommitResource(ctx, workCommitRes)
	res = append(res, resourceID)

	return res, nil
}

func GetResourceIDsBy() {

}

//作业提交；删除以前的记录
//todo:增加事务
func DelWorkCommitAllInfoByLessonIDUserID(ctx *gin.Context, workCommitUserId dto.WorkCommitUserID) (int, cerror.Cerror) {

	///////////////////////////////////////////////////////////0 删除提交记录
	commitIDs, _ := dao.DelWorkCommit(ctx, workCommitUserId) //

	if len(commitIDs) > 0 {
		fileCommitResource, richTextCommitResource, _ := dao.GetWorkCommitResourceRecords(ctx, commitIDs)

		var workCommitResource []dao.WorkCommitResource
		for _, val := range fileCommitResource {
			var item dao.WorkCommitResource
			item.ID = val.ID

			workCommitResource = append(workCommitResource, item)
		}

		for _, val := range richTextCommitResource {
			var item dao.WorkCommitResource
			item.ID = val.ID

			workCommitResource = append(workCommitResource, item)
		}

		///////////////////////////////////////////////////////////1 删除提交记录：所关联的资源
		resIDs, _ := dao.DelWorkCommitResource(ctx, workCommitResource) //
		fmt.Println(resIDs)

		var resourceIDs []int
		for _, val := range fileCommitResource {
			resourceIDs = append(resourceIDs, val.ResourceID)
		}

		for _, val := range richTextCommitResource {
			resourceIDs = append(resourceIDs, val.ResourceID)
		}
		///////////////////////////////////////////////////////////2 删除提交记录：删除资源
		var resourceIDs2 []int
		for _, val := range richTextCommitResource {
			resourceIDs2 = append(resourceIDs2, val.ResourceID)
		}
		if len(resourceIDs2) > 0 {
			//richtextID, _ := dao.GetRichtextIdByID(ctx, resourceIDs2[0])
			///////////////////////////////////////////////////////////3 删除提交记录：所关联的资源：富文本
			//dao.DelResourceRichtext(ctx, richtextID)
		}

		///////////////////////////////////////////////////////////4 删除提交记录：所关联的资源：文件
		var fileIDs []int
		for _, val := range fileCommitResource {
			fileIDs = append(fileIDs, val.ResourceID)
		}
		//filePaths, _ := dao.GetFilePathByID(ctx, fileIDs)
		//util.RemoveFile(filePaths)

		//	dao.DelResource(ctx, resourceIDs) //

	}

	//step2:删除文件

	return 1, nil
}
