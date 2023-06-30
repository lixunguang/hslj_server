package util

import (

	"github.com/gin-gonic/gin"
	"hslj/pkg/cerror"
	"net/http"
)

const (
	SuccessCode    = 0
	SuccessMessage = "success"
	FailCode       = -1
	HttpStatus     = http.StatusOK
	HeaderData     = "_header_bind_data"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewSuccessHttpResponse(data interface{}) *HttpResponse {
	return &HttpResponse{
		Code:    SuccessCode,
		Data:    data,
		Message: SuccessMessage,
	}
}

func NewFailHttpResponse(err cerror.Cerror) *HttpResponse {
	if err == nil {
		return &HttpResponse{
			Code:    FailCode,
			Data:    nil,
			Message: "error message is nil",
		}
	}
	return &HttpResponse{
		Data:    nil,
		Code:    err.Code(),
		Message: err.Error(),
	}
}

func NewFailMessageHttpResponse(err string) *HttpResponse {
	return &HttpResponse{
		Code:    FailCode,
		Data:    nil,
		Message: err,
	}
}

func SuccessJson(c *gin.Context, data interface{}) {
	c.JSON(
		HttpStatus,
		NewSuccessHttpResponse(data),
	)

}

func FailJson(c *gin.Context, err cerror.Cerror) {
	c.JSON(
		HttpStatus,
		NewFailHttpResponse(err),
	)
}

func NewFailMessage(err cerror.Cerror, data interface{}) *HttpResponse {

	return &HttpResponse{
		Code:    err.Code(),
		Data:    data,
		Message: err.Error(),
	}
}

func FailJsonData(c *gin.Context, err cerror.Cerror, data interface{}) {
	c.JSON(
		HttpStatus,
		NewFailMessage(err, data),
	)
}
