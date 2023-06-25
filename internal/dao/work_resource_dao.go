package dao

type WorkResource struct {
	BaseModel
	WorkID     int `gorm:"column:work_id"`
	ResourceID int `gorm:"column:resource_id"`
}
