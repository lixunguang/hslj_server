package util

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"unicode/utf8"
)

func GetIntValue(old interface{}) int {
	if new, ok := old.(int); ok {
		return new
	}
	return 0
}

func GetStringValue(old interface{}) string {
	if new, ok := old.(string); ok {
		return new
	}
	return ""
}

func IsUtf8(s string) bool {
	return utf8.ValidString(s)
}

func IsGbk(s string) bool {
	if IsUtf8(s) {
		return false
	}
	data := []byte(s)
	length := len(data)
	i := 0
	for i < length {
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func TransformGbkToUtf8(s string) (string, error) {
	decoder := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GB18030.NewDecoder())
	content, err := ioutil.ReadAll(decoder)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func GbkToUtf8(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())

	content, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	return string(content), nil
}
