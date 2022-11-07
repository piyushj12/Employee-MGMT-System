package route

import (
	"github.com/gin-gonic/gin"
	handler "EmployeeAssisgnment/api/handler"
)

func Init(r,o *gin.RouterGroup) {
	//route for employee registartion
	//o.POST("/register", handler.AddEmp())
	//route for employee/hr login
	o.POST("/login", handler.Login())
	r.POST("/add", handler.AddEmp())
	r.POST("/get/all/employees", handler.AdminEmpList())
	r.POST("/get/profile", handler.GetProfile())
	r.PUT("/edit/profile", handler.UpdateEmp())
	r.GET("/get/managers", handler.GetManagers())
	r.POST("/get/leaves", handler.GetLeaves())
	r.POST("/apply/leaves", handler.ApplyLeaves())
	r.POST("/get/applied/leaves", handler.GetAppliedLeaves())
	r.PUT("/update/leave/status", handler.UpdateLeaveStatus())
	r.DELETE("/delete/permanently/:id", handler.DeleteEmpPermanently())
	r.POST("/search", handler.SearchEmp())
	r.POST("/dashboard/data",handler.DashBoardDataHandler())
	r.POST("/capture/clockin",handler.CaptureClockinTime())
	r.POST("/capture/clockout",handler.CaptureClockoutTime())
	r.POST("/check/clockin/exists",handler.Isclockedin())
	r.POST("/check/clockout/exists",handler.Isclockedout())
	r.POST("/get/attendance",handler.GetAttendance())
	
}
