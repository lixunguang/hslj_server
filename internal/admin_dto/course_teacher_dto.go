package admin_dto

//教师-课程表 add 入参
type CourseTeacherParam struct {
	CourseID  int `json:"course_id" binding:"required"`
	TeacherId int `json:"teacher_id"`
}
type TeacherParam struct {
	Name string `json:"name"`
}

type UpdateTeacherParam struct {
	LoginID string `json:"login_id"`
	//	ID             int    `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	OrganizationID int    `json:"organization_id"`
	Introduce      string `json:"introduce"`
}
type AddTeacherParam struct {
	LoginID        string `json:"login_id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	OrganizationID int    `json:"organization_id"`
	Introduce      string `json:"introduce"`
}

type TeacherRes struct {
	LoginID string `json:"login_id"`
	//ID             int    `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	OrganizationID int    `json:"organization_id"`
	Introduce      string `json:"introduce"`
}
