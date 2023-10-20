package dto

//所有动作
type ActionAllRes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//动作详情
type ActionDetailRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
