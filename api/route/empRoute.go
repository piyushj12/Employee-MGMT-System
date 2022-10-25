package route

import (
	"github.com/gin-gonic/gin"
	handler "EmployeeAssisgnment/api/handler"
)

func Init(r,o *gin.RouterGroup) {
	//route for employee registartion
	o.POST("/register", handler.AddEmp())
	//route for employee/hr login
	o.POST("/login", handler.Login())
	r.POST("/add", handler.AddEmp())
	r.PUT("/update", handler.UpdateEmp())
	r.POST("/search", handler.SearchEmp())
	r.POST("/list", handler.ListEmp())
	r.POST("/delete", handler.DeleteEmp())
	r.POST("/restore",handler.RestoreEmp())
	r.GET("/viewdeleted",handler.ViewDeletedEmp())
	//o.POST("/token",handler.TokenGeneration())
}
