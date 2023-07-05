package dto

type User struct {
	LoginID        string `json:"login_id"`
	Name           string `json:"name"`
	Password       string `json:"password"`

}

type UserRes struct {
	LoginID string `json:"login_id"`
	Name    string `json:"name"`

}

type UserRecordRes struct {
	LoginID string `json:"login_id"`
	Name    string `json:"name"`
	DataStr string `json:"data_str"`
}

type AddUserRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


