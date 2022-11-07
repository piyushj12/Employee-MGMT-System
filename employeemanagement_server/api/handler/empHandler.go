package handler

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	service "EmployeeAssisgnment/api/services"
	model "EmployeeAssisgnment/api/model"
	helper "EmployeeAssisgnment/api/helpers"
)


// func TokenGeneration() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		login := model.Login{}
// 		c.Bind(&login)
		
// 		_, isValidUser := service.ValidateUser(login)
// 		if isValidUser {
// 			token, err := helper.GenerateToken(login, 24*time.Hour)
// 			if err != nil {
// 				fmt.Print("error while generating token:", err)
// 			}
// 			c.Header("Authorization", token)

// 			c.JSON(http.StatusOK, gin.H{"Authorization": token})
// 		} else {
// 			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"Error": "Please Enter valid username and password"})
// 		}
// 	}
// }

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := model.Login{}
		c.Bind(&login)
		
		err, user := service.Validate(login)
		if err!=nil{
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"Error": err})
		}else{
			if len(user)>0{
				fmt.Println(user[0])
				token, err := helper.GenerateToken(user[0], 24*time.Hour)
				if err != nil {
					c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"Error": err})
				}
				c.Header("Authorization", token)
				c.JSON(http.StatusOK, gin.H{"Authorization": token})
			}else{
				c.JSON(http.StatusOK, gin.H{"Error": "Please Enter valid username and password"})
			}
		}
		
	}
}

func AddEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := service.NewEmp()
		c.Bind(&requestBody)
		err:=service.AddEmpService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK,gin.H{"message":"employee added sucessfully"})
		}
		
	}
}


// Update Emp Update Handler
func UpdateEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var UpdateObj interface {}
		c.Bind(&UpdateObj)
		err:=service.UpdateEmpService(UpdateObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":"employee update sucessfully"})
		}
		
	}
}


//Get manager list
func GetManagers() gin.HandlerFunc {
	return func(c *gin.Context) {
		err,list:=service.GetManagers()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, list)
		}
		
	}
}
//Search Emp Handler
func SearchEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var SearchObj interface{}
		c.Bind(&SearchObj)
		err,employeelist:=service.SearchEmpService(SearchObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}

//List Employee
func AdminEmpList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ListObj interface{}
		c.Bind(&ListObj)
		err,employeelist:=service.AdminallEmpList(ListObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}
//get dashboarddata
func DashBoardDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var email interface{}
		c.Bind(&email)
		err,details:=service.DashBoardDataService(email)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, details)
		}
		
	}
}

//Get leaves
func GetLeaves() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ListObj model.Email
		c.Bind(&ListObj)
		err,employeelist:=service.GetLeaves(ListObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}

func GetAppliedLeaves() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ListObj interface {}
		c.Bind(&ListObj)
		err,list:=service.GetAppliedLeaves(ListObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, list)
		}
		
	}
}
//Delete Employee
func DeleteEmpPermanently() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody:=make(map[string]string) 
		requestBody["empid"]=c.Param("id")
		err,msg:=service.DeleteEmpPermanentlyService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":msg})
		}
		
	}
}

//get profile handler
func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := model.Login{}
		c.Bind(&login)
		err,profile:=service.GetProfileService(login)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, profile)
		}
		
	}
}

//apply leaves
func ApplyLeaves() gin.HandlerFunc {
	return func(c *gin.Context) {
		leaves := model.Leaves{}
		c.Bind(&leaves)
		err,ack:=service.ApplyLeavesService(leaves)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, ack)
		}
		
	}
}

func UpdateLeaveStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var UpdateObj model.Leaves
		c.Bind(&UpdateObj)
		err:=service.UpdateLeavesService(UpdateObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":"Leave approved sucessfully"})
		}
		
	}
}

//capture clock in time
func CaptureClockinTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj interface{}
		c.Bind(&obj)
		err,status:=service.CaptureClockinTimeService(obj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, status)
		}
		
	}
}

//capture clock Out time
func CaptureClockoutTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj interface{}
		c.Bind(&obj)
		err,status:=service.CaptureClockoutTimeService(obj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, status)
		}
		
	}
}

//check clock in
func Isclockedin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj model.Attendance
		c.Bind(&obj)
		err,status:=service.IsclockedinService(obj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, status)
		}
		
	}
}

//check clock out
func Isclockedout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj model.Attendance
		c.Bind(&obj)
		err,status:=service.IsclockedoutService(obj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, status)
		}
		
	}
}

//Get attendance of emplyoee
func GetAttendance() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj interface{}
		c.Bind(&obj)
		err,status:=service.GetAttendanceService(obj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, status)
		}
		
	}
}