package dto

type GetLocationListParam struct {
	LocationCode int `json:"location_code"`
}

type GetLocationListRes struct {
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Latitude  float64 `json:"latitude"`
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
