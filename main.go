package main

import (
	"GoVueTodo/dao"
	"GoVueTodo/model"
	"GoVueTodo/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//绑定模型
	dao.DB.AutoMigrate(&model.Todo{})

	r := routers.SetupRouter()
	//r.Run()
	r.RunTLS(":8080", "1_www.yinleilei.club_bundle.crt","2_www.yinleilei.club.key")

}
