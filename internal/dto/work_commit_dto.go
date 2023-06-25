package dto

//作业提交信息
//对外接口使用loginID
type WorkCommitParam struct {
	LessonID int    `json:"lesson_id" binding:"required"`
	LoginID  string `json:"login_id" binding:"required"`
}

//内部使用userid
type WorkCommitUserID struct {
	LessonID int `json:"lesson_id"`
	UserID   int `json:"user_id"`
}

//作业删除入参
type WorkCommitDelParam struct {
	LessonID   int `json:"lesson_id" binding:"required"`
	UserID     int `json:"user_id" binding:"required"`
	ResourceId int `json:"resource_id binding:"required""`
}

//作业提交项：包括文件及留言类型
//WorkCommitAttachment  = 43 //作业提交的文件
//WorkCommitText        = 44 //作业提交附言
type WorkCommitItem struct {
	FileName  string `json:"file_name"`
	Content   string `json:"file_url"` //resource content
	ContentID int    `json:"content_id"`
	Type      int    `json:"type"` //43：作业提交的文件，44 ：作业提交附言
}

//作业提交返回值
type WorkCommitRes struct {
	LessonID          int              `json:"lesson_id"`
	UserLoginID       string           `json:"login_id"`
	WorkCommitFiles   []WorkCommitItem `json:"work_commit_files"`
	WorkCommitMessage string           `json:"work_commit_message"`
}

type WorkCommitAddItem struct { //对应于workcommitresource表的内容
	ContentID int `json:"content_id"` //提交作业文件时，用int类型
	//Type      int    `json:"type"`       //43：作业提交的文件，44 ：作业提交附言
}

//作业提交增加参数
type WorkCommitAddParam struct {
	LessonID          int                 `json:"lesson_id" binding:"required"`
	UserLoginID       string              `json:"login_id" binding:"required"`
	WorkCommitFiles   []WorkCommitAddItem `json:"work_commit_files"`
	WorkCommitMessage string              `json:"work_commit_message"`
}
