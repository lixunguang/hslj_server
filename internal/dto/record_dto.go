package dto

type AddRecordParam struct {
	UserID     int `json:"user_id"`
	LocationID int `json:"location_id"`
}

type LocationRecordRes struct {
	LocationName string `json:"location_name"`
	RecordDate   string `json:"record_date"`
}

type UserRecordRes struct {
	UserName   string `json:"user_name"`
	RecordDate string `json:"record_date"`
}
