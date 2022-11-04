package services

import (

	model "EmployeeAssisgnment/api/model"
)

func ValidateUser(login model.Login) (model.Login, bool) {
	err,user:=ValidateDetails(login)
	if err!=nil{
		return model.Login{},false
	}
	if len(user)==0{
		return model.Login{},false
	}
	
	return user[0],true
}

func Validate(login model.Login) (error,[]model.Login) {
	err,user:=ValidateDetails(login)
	if err!=nil{
		return err,[]model.Login{}
	}
	return nil,user
}

func GetProfileService(login model.Login) (error,[]model.EmpDetails) {
	err,user:=GetProfileFromDB(login)
	if err!=nil{
		return err,[]model.EmpDetails{}
	}
	return nil,user
}


func NewEmp() model.EmpDetails {
	return model.EmpDetails{}
}

func AddEmpService(emp model.EmpDetails) error {
	// Mongo DB
	err:=SaveEmployeeToDB(emp)
	if err !=nil{
		return err
	}
	return nil
}

func UpdateEmpService(empdetails interface {}) error {
	err:=UpdateEmpFromDB(empdetails)
	if err != nil{
		return err
	}
	return nil
}
func UpdateLeavesService(leaves model.Leaves) error {
	err:=UpdateLeaveStatusToDB(leaves)
	if err != nil{
		return err
	}
	return nil
}

func GetManagers() (error,[]model.EmpDetails) {
	err,list:=GetManagersFromDB()
	if err != nil{
		return err,[]model.EmpDetails{}
	}
	return nil,list
}
func SearchEmpService(empdetails interface{}) (error,[]model.EmpDetails) {
	err,employeelist:=SearchEmpFromDB(empdetails)
	if err != nil{
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func AdminallEmpList(empdetails interface{}) (error,[]model.EmpDetails) {
	err,employeelist:=AdminallEmpListFromDB(empdetails)
	if err != nil{
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func GetLeaves(empdetails model.Email) (error,[]model.Leaves) {
	err,list:=GetLeavesFromDB(empdetails)
	if err != nil{
		return err,[]model.Leaves{}
	}
	return nil,list
}
func GetAppliedLeaves(email interface {}) (error,[]model.Leaves) {
	err,list:=GetAppliedLeavesFromDB(email)
	if err != nil{
		return err,[]model.Leaves{}
	}
	return nil,list
}
func ApplyLeavesService(leaves model.Leaves)(error, bool){
	err,msg:=StoreLeaves(leaves)
	if err != nil{
		return err,false
	}
	return nil,msg
}
func DeleteEmpPermanentlyService(deletedetails map[string]string) (error,string) {
	err,msg:=DeleteEmpFromDB(deletedetails)
	if err != nil{
		return err,""
	}
	return nil,msg
}
