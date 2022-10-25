package services

import (
	database "EmployeeAssisgnment/api/database"
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"errors"
	"gopkg.in/mgo.v2/bson"
	model "EmployeeAssisgnment/api/model"
)




func ValidateDetails(login model.Login) (error,[]model.Login){
	var user []model.Login
	if err := database.Collection().Find(bson.M{"username":login.Username,"password":login.Password,"empstatus":"active"}).All(&user); err != nil {
		return err,[]model.Login{}
	}
	fmt.Println(user)
	return nil,user
}
//save employee details in db
func SaveEmployeeToDB(empDetails model.EmpDetails) error {
	//check if name present in emp object
	if empDetails.Firstname == "" || empDetails.Lastname == "" || empDetails.Department==""{
		return errors.New("Name not Present,Enter all required filed")
	}
	//generate random number
	rand.Seed(time.Now().Unix())
	ranNum:=rand.Intn(1000)
	var pad string
	if len(empDetails.Department) >= 4{
		pad=empDetails.Lastname[:2]+empDetails.Department[:4]
	}else{
		pad=empDetails.Lastname[:2]+empDetails.Department[:len(empDetails.Department)]
	}
	fmt.Println(pad)
	empDetails.EmpID = empDetails.Firstname  + pad + strconv.Itoa(ranNum)
	empDetails.Empstatus="Pending"
	//query to insert
	err:=database.Collection().Insert(empDetails)
	if err != nil{
		return err
	}
	return nil
}

//update
func UpdateEmpFromDB(empdetails interface{}) error {
	origin:= empdetails.(map[string]interface {})
		query := bson.M{
			"empid": origin["empid"],
		}

		_,ok:=origin["skills"]
		if ok{
			doc1:=bson.M{"$addToSet":bson.M{"skills":bson.M{"$each":origin["skills"]}}}
			err:=database.Collection().Update(query, doc1)
			if err != nil{
				return err
			}
		}else{
			doc := bson.M{
				"$set": empdetails,
			}
			err:=database.Collection().Update(query, doc)
			if err != nil{
				return err
			}
		}
		
		return nil
}

func SearchEmpFromDB(empdetails interface{}) (error,[]model.EmpDetails){
	var employeelist []model.EmpDetails
	origin:= empdetails.(map[string]interface {})
	query:= make([]map[string]interface{},0)
	for key,value:= range origin{
		if key!="skills"{
			
			query=append(query,map[string]interface{}{key:bson.M{"$regex":value,"$options":"i"}})
		}
		if key=="skills"{
			doc := bson.M{"skills":bson.M{"$in":value}}
			query=append(query,doc)
		}
		
	}
	fmt.Println(query)
	if err := database.Collection().Find(bson.M{"$or":query,"empstatus":"Activated"}).All(&employeelist); err != nil {
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func ListEmpFromDB(empdetails interface{}) (error,[]model.EmpDetails){
	var employeelist []model.EmpDetails
	origin:= empdetails.(map[string]interface {})
	
	for key,value:= range origin{
		origin[key]=bson.M{"$regex":value,"$options":"i"}
	}
	origin["empstatus"]="Activated"
	if err := database.Collection().Find(origin).All(&employeelist); err != nil {
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func DeleteEmpFromDB(deletedetails model.DeleteData) (error,string){
     if deletedetails.PermanentlyDelete ==true{
		err:=database.Collection().Remove(bson.M{"empid": deletedetails.EmpID})
		if err!=nil{
			return err,""
		}
		return nil,"Permanently deleted employee"
	 }else{
		query:=bson.M{"empid":deletedetails.EmpID}
		UpdateQuery:=bson.M{"$set":bson.M{"empstatus":"Deactivated"}}
		err:=database.Collection().Update(query, UpdateQuery)
			if err != nil{
				return err,""
			}
			return nil,"Employee Status changed to deactivated"
	 }
	
	
}

func RestoreEmpFromDB(restoredetails model.RestoreData) (error,string){
	
	   query:=bson.M{"empid":restoredetails.EmpID}
	   UpdateQuery:=bson.M{"$set":bson.M{"empstatus":"Activated"}}
	   err:=database.Collection().Update(query, UpdateQuery)
		   if err != nil{
			   return err,""
		   }
		   return nil,"Employee Status changed to Activated"
}
   
func ViewDeletedEmpFromDB() (error,[]model.EmpDetails){
	var employeelist []model.EmpDetails
	if err := database.Collection().Find(bson.M{"empstatus":"Deactivated"}).All(&employeelist); err != nil {
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}