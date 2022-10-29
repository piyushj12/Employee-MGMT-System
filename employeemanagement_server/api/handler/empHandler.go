package handler

import (
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
		var UpdateObj interface{}
		c.Bind(&UpdateObj)
		err:=service.UpdateEmpService(UpdateObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":"employee update sucessfully"})
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

//Delete Employee
func DeleteEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody:=service.NewDeleteData()
		c.Bind(&requestBody)
		err,msg:=service.DeleteEmpService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":msg})
		}
		
	}
}

//Restore Employee
func RestoreEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody:=service.NewRestoreData()
		c.Bind(&requestBody)
		err,msg:=service.RestoreEmpService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":msg})
		}
		
	}
}

//VIEW DELETED EMPLOYEE
func ViewDeletedEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		err,employeelist:=service.ViewDeletedEmpService()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}