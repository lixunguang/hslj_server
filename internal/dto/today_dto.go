package dto

type Today struct {
	LunarDesc  string `json:"lunar_desc"`
	Txt        string `json:"txt"`
	PictureUrl string `json:"picture_url"`
}

type TodayShort struct {
	Desc string `json:"desc"`
}
