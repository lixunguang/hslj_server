package dto

type User struct {
	LoginID        string `json:"login_id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	OrganizationID int    `json:"organization_id"`
}

type UserRes struct {
	LoginID string `json:"login_id"`
	Name    string `json:"name"`

	OrganizationID int `json:"organization_id"`
}

type AddUserRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
