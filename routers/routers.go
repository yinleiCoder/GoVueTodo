package routers

import (
	"GoVueTodo/controller"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func SetupRouter() * gin.Engine {
	r := gin.Default()
	r.Use(TLSHandler())
	//告诉gin框架模板文件引用的静态文件从那儿找
	r.Static("/static","static")
	//告诉gin框架从哪里去找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)


	//v1
	v1Group := r.Group("v1")
	{
		//待办事项:
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有的代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		////查看某个代办事项
		//v1Group.GET("/todo/:id", func(c *gin.Context) {
		//
		//})

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}



func TLSHandler() gin.HandlerFunc  {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "yinleilei.club:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			return
		}
		c.Next()
	}
}