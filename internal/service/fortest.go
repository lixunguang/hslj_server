package service

import (
	"fmt"
	"os"
)

type LoginParamXX struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Len      int    `json:length`
}

func XXX() {

	envstr := os.Environ()
	fmt.Printf("%v\n", envstr)

}
