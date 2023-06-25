package admin_dto

type AdminLoginParam struct {
	Password  string `json:"password" binding:"required"`
	Name      string `json:"login_id" `
	Dept      string `json:"dept" `
	IsTeacher bool   `json:"is_teacher"`
}

type LoginParam struct {
	LoginId   string `json:"login_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" `
	Dept      string `json:"dept" `
	IsTeacher bool   `json:"is_teacher"`
}

type LoginRes struct {
	Token      string `json:"token"`
	FirstLogin bool   `json:"first_login"`
}
