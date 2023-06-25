package dao

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Teacher struct {
	BaseModel
	LoginID        string `gorm:"column:login_name"`
	Name           string `gorm:"column:name"`      //姓名
	Introduce      string `gorm:"column:introduce"` //介绍
	Password       string `gorm:"column:password"`
	PhoneNumber    string `gorm:"column:phone_number"`
	OrganizationID int    `gorm:"column:organization_id"` //创建本用户的管理员id
	IsDeleted      int8   `gorm:"column:is_deleted"`
}

func (Teacher) TableName() string {
	return "teacher"
}

func GetTeacherNameByID(ctx *gin.Context, teacherID []int) []string {
	mysqlDB := mysql.GetDB()

	var teacher []Teacher
	var userNames []string
	result := mysqlDB.Where("id IN ? ", teacherID).Find(&teacher) //有多个的话

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return userNames
	}

	for _, val := range teacher {
		userNames = append(userNames, val.Name)
	}
	return userNames
}

//这个函数需要重构，因为名字可能重复，目前假设名字唯一
func GetTeacherByName(ctx *gin.Context, name string) ([]Admin, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var teacher Teacher
	teacher.Name = name

	var adminArray []Admin
	result := mysqlDB.Where(&teacher).Find(&adminArray)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return adminArray, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return adminArray, nil

}

//这个函数需要重构，因为名字可能重复，目前假设名字唯一
func AddTeacher(ctx *gin.Context, teacher Teacher) (dto.AdminRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Create(&teacher)

	var res dto.AdminRes

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	res.Name = teacher.Name
	res.ID = teacher.ID

	return res, nil
}

//这个函数需要重构，因为名字可能重复，目前假设名字唯一
func DelTeacher(ctx *gin.Context, loginID string) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Where("login_name = ?", loginID).Delete(&Teacher{})

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return loginID, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return loginID, cerror.ErrorUserNotExist
	}

	return loginID, nil
}

//这个函数需要重构，因为名字可能重复，目前假设名字唯一
func UpdateTeacher(ctx *gin.Context, param admin_dto.UpdateTeacherParam) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	teacher := Teacher{Name: param.Name, LoginID: param.LoginID, Password: param.Password, OrganizationID: param.OrganizationID, Introduce: param.Introduce}
	result := mysqlDB.Where("login_name = ?", param.LoginID).Updates(teacher)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return param.Name, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return param.Name, nil
}

//这个函数需要重构，因为名字可能重复，目前假设名字唯一
func AllTeacher(ctx *gin.Context) ([]Teacher, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var teachers []Teacher
	result := mysqlDB.Order("updated_at desc").Find(&teachers)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return teachers, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return teachers, nil
}
