package request

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary //兼容标准库，效率比标准库高
)

// 解码
func JsonStringDecode(str string, result interface{}) error {
	if len(str) == 0 {
		return nil
	}
	err := json.Unmarshal([]byte(str), result)
	if err != nil {
		return err
	}
	return nil
}

// 解码
func JsonDecode(str []byte, result interface{}) error {
	if len(str) == 0 {
		return nil
	}
	err := json.Unmarshal(str, result)
	if err != nil {
		return err
	}
	return nil
}

// 编码
func JsonEncode(result interface{}) ([]byte, error) {
	if result == nil {
		return []byte{}, nil
	}
	slice, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return slice, err
}

// 比如 "error.message" 收集的是{"error":{"message":""}}的信息,是个很好的工具这个 gjson
func CheckJsonHasPattern(body []byte, pattern []string) (bool, string) {
	states := gjson.GetManyBytes(body, pattern...)
	for _, state := range states {
		if state.Exists() {
			return true, state.String()
		}
	}
	return false, ""
}
