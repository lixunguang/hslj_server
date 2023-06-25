package service

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {

}

func TestLogout(t *testing.T) {

}

func TestCheckExpired(t *testing.T) {
	res := CheckExpired("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Nzg2OTcwOTgsInVzZXJuYW1lIjoiMTIzIn0.LzP7eOumvo7DDEXP4lJAtvmzYU6SjjrdrkMgFq2fch4")
	fmt.Println(res)
}
