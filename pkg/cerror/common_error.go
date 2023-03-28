package cerror

var (
	InvalidParams             = NewCerror(100001, "参数错误")
	NotFound                  = NewCerror(100002, "找不到")
	UnauthorizedAuthNotExist  = NewCerror(100003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewCerror(100004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewCerror(100005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewCerror(100006, "鉴权失败，Token 生成失败")
)
