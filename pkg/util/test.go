package util

import (
	"encoding/json"
	"fmt"
)

func CheckResult(result interface{}, err interface{}) {
	if err != nil {
		fmt.Printf("\n================error================\n\n"+
			"%s"+
			"\n\n=================end=================\n\n", err)
		return
	}
	bytes, err := json.Marshal(&result)
	if err != nil {
		fmt.Printf("\n================error================\n\n"+
			"%s"+
			"\n\n=================end=================\n\n", err)
		return
	}
	fmt.Printf("\n================success================\n\n"+
		"%s"+
		"\n\n==================end==================\n\n", string(bytes))
}
