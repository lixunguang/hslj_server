package dao

import (

	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/dto"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
)

type User struct {
	BaseModel
	LoginID string `gorm:"column:login_id"`
	// 学号 ,注意区分ID（程序内部使用）和LoginID（对外接口使用），通过GetIDByLoginID和GetLoginIDByID转换 一个为int类型，一个为string类型！！！！
	Name           string `gorm:"column:name"`            // 姓名
	OrganizationID int    `gorm:"column:organization_id"` // 所属院系
	Password       string `gorm:"column:password"`
	PhoneNumber    string `gorm:"column:phone_number"`
	CreatorID      int    `gorm:"column:creator_id"` // 创建本用户的管理员id
	IsDeleted      int8   `gorm:"column:is_deleted"`
}

func (User) TableName() string {
	return "user"
}

// 删除用户
func DelUser(ctx *gin.Context, login_id string) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Where("login_id = ?", login_id).Delete(&User{})

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return login_id, cerror.NewCerror(common.InvalidID, result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return login_id, cerror.ErrorUserNotExist
	}

	return login_id, nil
}

// 删除用户
func UpdateUser(ctx *gin.Context, param dto.User) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()
	var user User
	user.OrganizationID = param.OrganizationID
	user.LoginID = param.LoginID
	user.Name = param.Name
	user.Password = param.Password

	result := mysqlDB.Model(&User{}).Where("login_id = ?", param.LoginID).Updates(user)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return "", cerror.ErrorDataUpdate
	}

	if result.RowsAffected == 0 {
		return param.LoginID, cerror.ErrorUserNotExist
	}

	return param.LoginID, nil
}

func AllUser(ctx *gin.Context) ([]User, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var users []User
	result := mysqlDB.Order("updated_at desc").Find(&users)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return users, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return users, nil
}

func GetUser(ctx *gin.Context, loginID string) (dto.UserRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var user User
	result := mysqlDB.Where("login_id = ? ", loginID).First(&user)

	var userRes dto.UserRes
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return userRes, cerror.ErrorDataGet
	}

	userRes.Name = user.Name
	userRes.OrganizationID = user.OrganizationID
	userRes.LoginID = user.LoginID

	return userRes, nil
}

func CreateUser(ctx *gin.Context, param dto.User) (dto.AddUserRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	user := User{LoginID: param.LoginID, Name: param.Name, Password: param.Password, OrganizationID: param.OrganizationID}
	result := mysqlDB.Create(&user)

	var res dto.AddUserRes

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.ErrorDataAdd
	}

	res.ID = user.ID
	res.Name = user.Name
	return res, nil
}

// 根据userid获取username
func GetUserNameByUserID(ctx *gin.Context, userID []int) []string {
	mysqlDB := mysql.GetDB()

	var users []User
	var userNames []string
	result := mysqlDB.Where("id IN ? ", userID).Find(&users) // 有多个的话
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return userNames
	}

	for _, val := range users {
		userNames = append(userNames, val.Name)
	}
	return userNames
}

// 根据id获取loginid
func GetLoginIDByID(ctx *gin.Context, userID int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()
	var id int = 0
	var user User
	user.ID = userID

	result := mysqlDB.First(&user)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return id, cerror.NewCerror(common.InvalidID, result.Error.Error())
	}

	return id, nil
}

// 根据loginid获取id
func GetIDByLoginID(ctx *gin.Context, loginID string) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()
	var id int = 0
	var user User

	result := mysqlDB.Where("login_id=?", loginID).First(&user)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return id, cerror.NewCerror(common.InvalidID, result.Error.Error())
	}

	id = user.ID
	return id, nil
}
