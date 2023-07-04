package dto

//
type NewsLatestRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DateStr string `json:"date"`
}

type NewsAllRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DateStr string `json:"date"`
}

//新闻详情
type NewsDetailRes struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	DateStr    string `json:"date"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"` //暂时不用
}
