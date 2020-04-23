package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)


func InitMySQL() (err error) {
	//dsn := "root:yl13795950539@tcp(127.0.0.1:3306)/govuetodo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:yl13795950539@tcp(cd-cdb-60i9tisg.sql.tencentcdb.com:62524)/govuetodo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close(){
	DB.Close()
}
