package dao

import (
	"lxg_jz/internal/model/db"
	"time"
)

type User struct {
	ID        string    `gorm:"column:id"`
	Role      int8      `gorm:"column:role"`
	Name      string    `gorm:"column:name"`
	Password  string    `gorm:"column:password"`
	CreatorID int       `gorm:"column:creator_id"`
	IsDeleted int8      `gorm:"column:is_deleted"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}

func AddUser() {
	//获取DB
	db := db.GetDB()

	//add
	user := User{Name: "nnnnn", Role: 18}
	db.Create(&user)

}

func CheckUser(id string, password string) {
	//获取DB
	db := db.GetDB()

	//query
	u := User{}
	db.Where("age = ?", 22).First(&u)

}
