package controller

import (
	"GoVueTodo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
url___>controller------>logic里面的业务逻辑层---->model层
请求来了--->控制器---->业务逻辑---》模型层的增删改查
 */

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",nil)
}

func CreateTodo(c *gin.Context) {
	//前端页面填写事项并点击提交
	//1.从请求中吧数据拿出来
	var todo model.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err := model.CreateATodo(&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
	//3.返回响应
}


func GetTodoList(c *gin.Context) {
	//查询todovia表的所有数据
	todoList, err:=model.GetAllTodoList()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,todoList)
	}
}


func UpdateATodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := model.GetATodo(id)

	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err =model.UpdateATodo(todo);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := model.DeleteATodo(id);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			id: "deleted",
		})
	}
}
