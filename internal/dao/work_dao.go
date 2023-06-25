package dao

type Work struct {
	BaseModel
	UserID   int `gorm:"column:user_id"`
	LessonID int `gorm:"column:lesson_id"`
}
