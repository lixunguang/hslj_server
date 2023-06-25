package dao

import (
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	BaseModel

	LoginID     int    `gorm:"column:login_id"`
	Name        string `gorm:"column:name"`
	Password    string `gorm:"column:password"`
	Role        int    `gorm:"column:role"`
	PhoneNumber string `gorm:"column:phone_number"`
	CreatorID   int    `gorm:"column:creator_id"`
	IsDeleted   int8   `gorm:"column:is_deleted"`
}

func (Admin) TableName() string {
	return "admin"
}

func AddAdmin(ctx *gin.Context, param dto.Admin) (dto.AdminRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	admin := Admin{Name: param.Name, Password: param.Password}
	result := mysqlDB.Create(&admin)

	var res dto.AdminRes

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.ErrorDataAdd
	}

	res.Name = admin.Name
	res.ID = admin.ID

	return res, nil
}

func GetAdmin(ctx *gin.Context, param dto.AdminParam) ([]Admin, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var admin Admin
	admin.Name = param.Name

	var adminArray []Admin
	result := mysqlDB.Where(&admin).Find(&adminArray)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return adminArray, cerror.ErrorDataGet
	}

	return adminArray, nil
}

func DelAdmin(ctx *gin.Context, param dto.AdminParam) ([]Admin, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var admin Admin
	admin.Name = param.Name

	var adminArray []Admin
	result := mysqlDB.Where(&admin).Find(&adminArray)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return adminArray, cerror.ErrorDataGet
	}

	var ids []int
	if len(adminArray) > 0 {

		for _, v := range adminArray {
			ids = append(ids, v.ID)
		}

		result := mysqlDB.Delete(&Admin{}, ids)

		if result.Error != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
			return adminArray, cerror.ErrorDataDel
		}
	}

	return adminArray, nil
}

// 校验用户名及密码
func CheckAdmin(ctx *gin.Context, name string, password string) cerror.Cerror {
	db := mysql.GetDB()

	// query
	u := Admin{}

	result := db.Where("name = ? ", name).First(&u)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail,err=%+v, loginId=%d", result.Error, name)
		return cerror.ErrorUserNotExist
	}

	if u.Name == "" { // 用户不存在
		return cerror.ErrorUserNotExist
	}

	result = db.Where("name = ? and password = ?", name, password).First(&u)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, loginId=%d", result.Error, name)
		return cerror.ErrorPassword
	}

	if u.Name == "" { // 密码错误
		return cerror.ErrorPassword
	}

	return cerror.ErrorUserAuthSucc
}
