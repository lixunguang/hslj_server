package admin_dto

type LearningRes struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"`
	Category   string `json:"category"`
}

type UpdateLearning struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Author     string `json:"author"`
	PictureID  int    `json:"picture_id"`
	CategoryID int    `json:"category_id"`
}

type LearningDetailRes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Author      string `json:"author"`
	PictureName string `json:"picture_name"`
	PictureID   int    `json:"picture_id"`
	PictureUrl  string `json:"picture_url"`
	Category    string `json:"category"`

	ResourceList []ResourceRes `json:"resource_list"`
}

type AddLearningResourceParam struct {
	LearningID int `json:"learning_id"`
	//Title      string `json:"title"`
	//Desc       string `json:"desc"`
	ResourceID int `json:"resource_id"`
	Index      int `json:"index"`
}

type DelLearningResourceParam struct {
	LearningID int `json:"learning_id"`
	ResourceID int `json:"resource_id"`
}
