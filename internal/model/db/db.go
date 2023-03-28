package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lxg_jz/config"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

func OpenDB() {

	userName := fmt.Sprintf("%s", config.Vipper.Get("DataBase.UserName"))
	password := fmt.Sprintf("%s", config.Vipper.Get("DataBase.Password"))
	host := fmt.Sprintf("%s", config.Vipper.Get("DataBase.Host"))
	port := fmt.Sprintf("%s", config.Vipper.Get("DataBase.Port"))
	Dbname := fmt.Sprintf("%s", config.Vipper.Get("DataBase.DBName"))
	timeout := fmt.Sprintf("%s", config.Vipper.Get("DataBase.Timeout"))

	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", userName, password, host, port, Dbname, timeout)

	var err error
	_db, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	/*
		defer func() {
			sqlDB, err := _db.DB()
			if err != nil {

			}

			//设置数据库连接池参数
			sqlDB.SetMaxOpenConns(100)   //设置数据库连接池最大连接数
			sqlDB.SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。


			sqlDB.Close()
		}()

	*/

	/*
		// Migrate the schema
		_db.AutoMigrate(&Product{})

		// Create
		_db.Create(&Product{Code: "D42", Price: 100})

		// Read
		var product Product
		db.First(&product, 1) // find product with integer primary key
		db.First(&product, "code = ?", "D42") // find product with code D42

		// Update - update product's price to 200
		db.Model(&product).Update("Price", 200)
		// Update - update multiple fields
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - delete product
		db.Delete(&product, 1)
	*/
}

func GetDB() *gorm.DB {
	return _db
}
