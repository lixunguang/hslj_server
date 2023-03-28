package util

import (
	"fmt"
	"lxg_jz/config"
)

func GetFtpUser() string {
	return fmt.Sprintf("%s", config.Vipper.Get("FTP.User"))
}

func GetFtpPassword() string {
	return fmt.Sprintf("%s", config.Vipper.Get("FTP.Password"))
}
