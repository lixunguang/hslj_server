package util

import (
<<<<<<< HEAD
	"hslj/config"
=======
>>>>>>> 688a4df92fb5de6d3a3c38567476cf81c98e2521
	"fmt"
	"hslj/config"
)

func GetFtpUser() string {
	return fmt.Sprintf("%s", config.Vipper.Get("FTP.User"))
}

func GetFtpPassword() string {
	return fmt.Sprintf("%s", config.Vipper.Get("FTP.Password"))
}
