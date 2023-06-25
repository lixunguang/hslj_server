package dto

type LearningItemParam struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	LearningID int    `json:"learning_id"`
	ResourceID int    `json:"resource_id"`
	Index      int    `json:"index"`
}

type LearningItem struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	ResourceUrl string `json:"resource_url"`
	Index       int    `json:"index"`
}
