package admin_dto

type NewsDetailRes struct {
	NewsID       int    `json:"news_id"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	PictureID    int    `json:"picture_id"`
	PictureUrl   string `json:"picture_url"`
	PictureTitle string `json:"pictrue_title"`
}
