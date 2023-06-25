package dto

//lesson参数
type LearningCategory struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

//分页获取分类下的资源
type LearningCategoryParam struct {
	CategoryID  int `json:"id"`
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
}

type UpdateLearningCategoryParam struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
