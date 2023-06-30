package middleware

import (
	"hslj/pkg/logger"
	"hslj/pkg/request"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 添加trace-id, 如果父亲有则 [父亲-儿子]
func AddTraceId(ctx *gin.Context) {
	traceId := ctx.Request.Header.Get(request.TraceIdHeader)
	if traceId == "" {
		traceId = logger.GenerateTraceId()
	} else {
		traceId = fmt.Sprintf("%s-%s", traceId, logger.GenerateTraceId())
	}
	ctx.Set(logger.TraceIdKey, traceId)
	ctx.Next()
}
