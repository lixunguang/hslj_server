package admin_dto

//课程资源表 add 入参
type CourseResourceParam struct {
	CourseID   int `json:"course_id" binding:"required"`
	ResourceID int `json:"resource_id" binding:"required"`
	Type       int `json:"resource_type" binding:"required"`
	Index      int `json:"index"`
}

type AddLessonRes struct {
	CourseId int `json:"course_id"`
	LessonId int `json:"lesson_id"`
}
