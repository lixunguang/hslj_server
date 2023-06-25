package dto

//作业提交所对应的资源 Type取值：
//WorkCommitAttachment  = 43 //作业提交的文件
//WorkCommitText        = 44 //作业提交附言

type WorkCommitResource struct {
	WorkCommitID int
	ResourceID   int
	Type         int //
	//	CanModify    int
}

type WorkCommitResourceID struct {
	ID           int
	WorkCommitID int
	ResourceID   int
	Type         int //
	//	CanModify    int
}
