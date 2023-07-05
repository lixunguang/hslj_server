package dao

//在参与活动后，可以对场地进行评价
type LocationRate struct {
	BaseModel

	PeopleNumber  int `gorm:"column:name"`     //参与人数
	PeopleTech    int `gorm:"column:desc"`     //技术水平
	SiteCondition int `gorm:"column:location"` //场地条件，比如灯光，风，场地大小等
	Experience    int `gorm:"column:location"` //个人主观体验

}

func (LocationRate) TableName() string {
	return "location_rate"
}
