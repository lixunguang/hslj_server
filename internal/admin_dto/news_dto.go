package admin_dto

//新闻输入参数
type NewsAddParam struct {
	Title     string `json:"title"`
	DateStr   string `json:"date"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	PictureID int    `json:"picture_id"`
}

//新闻详情
type NewsResObj struct {
	NewsID     int    `json:"news_id"`
	Title      string `json:"title"`
	DateStr    string `json:"date"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"`
}

//新闻列表返回项
type NewsItemRes struct {
	NewsID  int    `json:"news_id"`
	Title   string `json:"title"`
	DateStr string `json:"date"`
	//Date      time.Time `json:"date"`
	Author    string `json:"author"`
	PictureID int    `json:"-"`
}

type PicNews struct {
	NewsID     int    `json:"news_id"`
	PictureUrl string `json:"pic_url"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Author     string `json:"author"`
}

type TitleNews struct {
	NewsID int    `json:"news_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Author string `json:"author"`
}

type NewsRes struct {
	News NewsResObj `json:"news"`
	//	PicNews PicNews `json:"pic"`
}

//标题新闻
type NewsLatestRes struct {
	News []NewsResObj `json:"news_list"`
}

//banner新闻
type BannerNewsLatestRes struct {
	PicNews []PicNews `json:"banner_list"`
}

//newall 接口入参数
type NewsAllParam struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
}

//newall 接口返回值
type NewsAllRes struct {
	TotalNumber int64 `json:"total_number"`
	//PageSize    int64 `json:"page_size"`
	//CurrentPage int64 `json:"current_page"`

	//Pages    []int64    `json:"pages"`
	NewsList []NewsItemRes `json:"news_list"`
}

//addnews 接口返回值
type UpdateNewsRes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//admin news dto
type DelNewsParam struct {
	ID int `json:"id"`
}

type DelNewsRes struct {
	ID int `json:"id"`
}

type NewsUpdateParam struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	DateStr   string `json:"date"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	PictureID int    `json:"picture_id"`
}
