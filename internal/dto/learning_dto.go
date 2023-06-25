package dto

//add 接口入参数
type Learning struct {
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	Author         string `json:"author"`
	CoverPictureID int    `json:"cover_picture_id"`
	CategoryID     int    `json:"category_id"`
}

type LearningRes struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Author string `json:"author"`
}

//相对于LearningRes，字段更全
type LearningInfoRes struct {
	LearningRes
	ResourceList []LearningItem `json:"resource_list"`
}

type LearningRecommendRes struct {
	Name       string `json:"title"`
	Desc       string `json:"desc"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"`
	LearningID int    `json:"learning_id"`
}

type CategoryResourceItem struct {
	Name       string `json:"name"`
	Author     string `json:"author"`
	Desc       string `json:"desc"`
	PictureUrl string `json:"picture_url"`
	LearningID int    `json:"learning_id"`
}

//接口返回值
type LearningCategoryResourceRes struct {
	TotalNumber int64 `json:"total_number"`

	CategoryResourceList []CategoryResourceItem `json:"category_resource_list"`
}
