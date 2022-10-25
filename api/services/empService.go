package services

import (
	model "EmployeeAssisgnment/api/model"
)

func ValidateUser(login model.Login) (model.Login, bool) {

	if login.Username=="TusharN" && login.Password=="tusharnemade"{
		return login,true
	}else{
		return model.Login{},false
	}
}

func Validate(login model.Login) (error,[]model.Login) {
	err,user:=ValidateDetails(login)
	if err!=nil{
		return err,[]model.Login{}
	}
	return nil,user
}
func NewEmp() model.EmpDetails {
	return model.EmpDetails{}
}

func NewDeleteData() model.DeleteData {
	return model.DeleteData{}
}

func NewRestoreData() model.RestoreData {
	return model.RestoreData{}
}
func AddEmpService(emp model.EmpDetails) error {
	// Mongo DB
	err:=SaveEmployeeToDB(emp)
	if err !=nil{
		return err
	}
	return nil
}

func UpdateEmpService(empdetails interface{}) error {
	err:=UpdateEmpFromDB(empdetails)
	if err != nil{
		return err
	}
	return nil
}

func SearchEmpService(empdetails interface{}) (error,[]model.EmpDetails) {
	err,employeelist:=SearchEmpFromDB(empdetails)
	if err != nil{
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func ListEmpService(empdetails interface{}) (error,[]model.EmpDetails) {
	err,employeelist:=ListEmpFromDB(empdetails)
	if err != nil{
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func DeleteEmpService(deletedetails model.DeleteData) (error,string) {
	err,msg:=DeleteEmpFromDB(deletedetails)
	if err != nil{
		return err,""
	}
	return nil,msg
}

func RestoreEmpService(deletedetails model.RestoreData) (error,string) {
	err,msg:=RestoreEmpFromDB(deletedetails)
	if err != nil{
		return err,""
	}
	return nil,msg
}

func ViewDeletedEmpService() (error,[]model.EmpDetails) {
	err,employeelist:=ViewDeletedEmpFromDB()
	if err != nil{
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}