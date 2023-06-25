package dto

//资源参数及返回值
type ResourceRichText struct {
	Content string `json:"content"`
}

type UpdateRichtext struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
