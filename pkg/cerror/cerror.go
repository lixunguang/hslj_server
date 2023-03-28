package cerror

import "fmt"

type Cerror interface {
	Error() string
	Code() int
}

type cerror struct {
	C   int    //code编号
	Msg string //错误消息
}

//构造函数
func NewCerror(code int, msg string) Cerror {
	return &cerror{
		C:   code,
		Msg: msg,
	}
}

//返回错误信息
func (this *cerror) Error() string {
	return fmt.Sprintf("%s", this.Msg)
}

func (this *cerror) Code() int {
	return this.C
}

// 通用方法
func NewECerror(code int) func(err error) Cerror {
	return func(err error) Cerror {
		return NewCerror(code, err.Error())
	}
}

func NewSCerror(code int) func(err string) Cerror {
	return func(err string) Cerror {
		return NewCerror(code, err)
	}
}
