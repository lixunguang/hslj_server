package dto

type GetLocationListParam struct {
	LocationCode int `json:"location_code"`
}

type GetLocationListRes struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Latitude  float64 `json:"latitude"` //点数据，mysql也支持更多的POINT LINESTRING
	Longitude float64 `json:"longitude"`
	Rating    int     `json:"rating"`
	//OpenTime time.Time
	//CloseTime time.Time
	//UserReviews []int
}

type AddLocationParam struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Frequency int `json:"frequency"`//频率1,2,3,4,5,6,7
	PeopleNum int `json:"people_num"`//参与人数 1代表0-10人 2 代表10~20人
	TimeStr string `json:"time_str"`
	ISAuth int `json:"is_auth"`
	Rating int `json:"rating"`
}

type AddLocationRes struct {

}

type LocationRes struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
