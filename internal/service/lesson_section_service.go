package service

import (
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func AddLessonSection(ctx *gin.Context, section dto.LessonSection) (int, cerror.Cerror) {
	return dao.AddLessonSection(ctx, section)
}
