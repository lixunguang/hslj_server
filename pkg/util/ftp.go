package util

import (
	"edu-imp/config"
	"fmt"
)

func GetFtpUser() string {
	return fmt.Sprintf("%s", config.Vipper.Get("FTP.User"))
}

func GetFtpPassword() string {
	return fmt.Sprintf("%s", config.Vipper.Get("FTP.Password"))
}
