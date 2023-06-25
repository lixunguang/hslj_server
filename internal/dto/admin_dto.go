package dto

type Admin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AdminParam struct {
	Name string `json:"name"`
}
