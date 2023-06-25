package dao

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Course struct {
	BaseModel
	Title     string         `gorm:"column:title"`
	Desc      string         `gorm:"column:desc"`
	DeletedAt gorm.DeletedAt //支持软删
}

func (Course) TableName() string {
	return "course"
}

//获取n个课程，根据一定的推荐标准，比如：时间
func GetNCourse(ctx *gin.Context, number int) []Course {
	mysqlDB := mysql.GetDB()

	var res []Course

	result := mysqlDB.Order("updated_at desc").Limit(int(number)).Find(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil
	}

	return res
}

//获取分页数据
func GetCoursePagedData(ctx *gin.Context, param dto.CourseAllParam) ([]Course, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	//pageSize := fmt.Sprintf("%d", config.Vipper.Get("AddCourseParam.PageSize"))
	//ps, _ := strconv.ParseInt(pageSize, 10, 64)

	ps := param.PageSize
	var course []Course

	offset := (param.CurrentPage - 1) * (ps)

	result := mysqlDB.Order("updated_at desc").Limit(int(ps)).Offset(int(offset)).Find(&course)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail,err=%+v, CurrentPage=%d", result.Error, param.CurrentPage)
		return nil, cerror.ErrorDataGet
	}

	return course, nil
}

//获取课程总条数
func GetCourseCount(ctx *gin.Context) (int64, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var course []Course
	var count int64

	result := mysqlDB.Model(&course).Count(&count)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail 2,err=%+v", result.Error)
		return 0, cerror.ErrorDataGet
	}

	return count, nil
}

func GetCourseByID(ctx *gin.Context, id int) Course {

	mysqlDB := mysql.GetDB()

	var course Course

	result := mysqlDB.Where("id = ? ", id).First(&course)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return course
	}

	return course
}

//--------------------------------------------------------------------------------------------------admin
// 后台使用的接口。获取课程列表
func GetAllCourse(ctx *gin.Context) ([]Course, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res []Course

	result := mysqlDB.Find(&res)
	//result := mysqlDB.Unscoped().Find(&res) //查找所有记录
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return res, nil
}

// 后台使用的接口
func DelCourse(ctx *gin.Context, id int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Delete(&Course{}, id)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return id, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return id, nil
}

// 后台使用的接口：获取课程信息
func GetCourse(ctx *gin.Context, id int) (Course, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var course Course

	result := mysqlDB.Where("id = ?", id).Find(&course)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return course, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return course, nil
}

func AddCourse(ctx *gin.Context, param admin_dto.AddCourseParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	course := Course{Title: param.Title, Desc: param.Desc}
	result := mysqlDB.Create(&course)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return course.ID, nil
}

func UpdateCourse(ctx *gin.Context, param admin_dto.UpdateCourseParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var course Course
	course.ID = param.ID

	result := mysqlDB.Model(&course).Where("id = ?", param.ID).Updates(Course{Title: param.Title, Desc: param.Desc})

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return param.ID, nil
}
