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
}

type AddLocationRes struct {
}

type LocationRes struct {
}
