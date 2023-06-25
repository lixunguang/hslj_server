package controller

import (
	"edu-imp/internal/dto"
)

//用户相关接口

type UserFake struct {
	Name             string `json:"name"`
	Avatar           string `json:"avatar"`
	Email            string `json:"email"`
	Job              string `json:"job"`
	JobName          string `json:"jobName"`
	Organization     string `json:"organization"`
	OrganizationName string `json:"organizationName"`
	Location         string `json:"location"`
	LocationName     string `json:"locationName"`
	Introduction     string `json:"introduction"`
	PersonalWebsite  string `json:"personalWebsite"`
	Phone            string `json:"phone"`
	RegistrationDate string `json:"registrationDate"`
	AccountId        string `json:"accountId"`
	Certification    int    `json:"certification"`

	Role int `json:"role"`
	dto.UserRes
}
