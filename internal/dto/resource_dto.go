package dto

type ResouceIDParam struct {
	ID int `json:"id"`
}

type ResouceDelParam struct {
	ID   int `json:"id"`
	Type int `json:"type"`
}

type ResourceParam struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type UpdateRichtextParam struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type ResourceRes struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	ContentUrl string `json:"content_url"`
	Type       int    `json:"type"`
	ID         int    `json:"id"`
	Index      int    `json:"index"`
}
