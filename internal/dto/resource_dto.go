package dto

//资源参数及返回值
type ResourceRes struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
	Desc    string `json:"desc"`
	Title   string `json:"title"`
}

type Resource struct {
	Title   string `json:"title"`
	Type    int8   `json:"type" binding:"required"`
	Desc    string `json:"desc"`
	Content string `json:"content" binding:"required"`
}

type UpdateResourceParam struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content" `
	Type    int8   `json:"type"`
}
