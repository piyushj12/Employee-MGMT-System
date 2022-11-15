package route

import (
	"github.com/gin-gonic/gin"
	handler "EmployeeAssisgnment/api/handler"
	
)

func Init(r,o,a,h,ah,he,e *gin.RouterGroup) {
	//route for employee registartion
	//o.POST("/register", handler.AddEmp())
	//route for employee/hr login
	o.POST("/login", handler.Login())
	ah.POST("/add", handler.AddEmp())
	ah.POST("/get/all/employees", handler.AdminEmpList())
	r.POST("/get/profile", handler.GetProfile())
	r.PUT("/edit/profile", handler.UpdateEmp())
	ah.GET("/get/managers", handler.GetManagers())
	r.POST("/get/leaves", handler.GetLeaves())
	r.POST("/apply/leaves", handler.ApplyLeaves())
	r.POST("/get/applied/leaves", handler.GetAppliedLeaves())
	r.PUT("/update/leave/status", handler.UpdateLeaveStatus())
	ah.DELETE("/delete/permanently/:id", handler.DeleteEmpPermanently())
	r.POST("/search", handler.SearchEmp())
	r.POST("/dashboard/data",handler.DashBoardDataHandler())
	he.POST("/capture/clockin",handler.CaptureClockinTime())
	he.POST("/capture/clockout",handler.CaptureClockoutTime())
	r.POST("/check/clockin/exists",handler.Isclockedin())
	r.POST("/check/clockout/exists",handler.Isclockedout())
	r.POST("/get/attendance",handler.GetAttendance())
	a.PUT("/update/array", handler.AddToArray())
	r.GET("/get/company/data", handler.GetCompantDataFromDB())
	a.PUT("/reset/employee/data", handler.AdminResetEmployee())
	
}
