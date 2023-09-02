package dto

//最近新闻
type NewsLatestRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DateStr string `json:"date"`
	PictureUrl string `json:"picture_url"`
}

//所有新闻
type NewsAllRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DateStr string `json:"date"`
	PictureUrl string `json:"picture_url"`
}

//新闻详情
type NewsDetailRes struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	DateStr    string `json:"date"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"`
}
