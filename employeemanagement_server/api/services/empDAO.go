package services

import (
	database "EmployeeAssisgnment/api/database"
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"errors"
	"strings"
	"gopkg.in/mgo.v2/bson"
	model "EmployeeAssisgnment/api/model"
	"github.com/google/uuid"
)




func ValidateDetails(login model.Login) (error,[]model.Login){
	var user []model.Login
	if err := database.Collection().Find(bson.M{"email":login.Email,"password":login.Password,"empstatus":"active"}).All(&user); err != nil {
		return err,[]model.Login{}
	}
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
	empDetails.EmpID = strings.ReplaceAll(empDetails.Firstname  + pad + strconv.Itoa(ranNum)," ","") 
	empDetails.Empstatus="Pending"
	empDetails.Password=generatePassword()
	empDetails.Empstatus="active"
	empDetails.Remholidays=empDetails.Holidays
	err:=database.Collection().Insert(empDetails)
	if err != nil{
		return err
	}
	return nil
}

//profile
func GetProfileFromDB(login model.Login) (error,[]model.EmpDetails){
	var user []model.EmpDetails
	if err := database.Collection().Find(bson.M{"email":login.Email}).All(&user); err != nil {
		return err,[]model.EmpDetails{}
	}
	return nil,user
}
//update
func UpdateEmpFromDB(empdetails interface {}) error {
	origin:= empdetails.(map[string]interface {})
	for key, value := range origin {
		if value==""|| value==0 || value==nil{
			delete(origin,key)
		}
    }
	err:=database.Collection().Update(bson.M{"email":origin["email"]}, bson.M{"$set":origin})
			if err != nil{
				return err
			}
		return nil
}

func UpdateLeaveStatusToDB(leaves model.Leaves) error {
	err2:=database.Leaves().Update(bson.M{"lid":leaves.Lid}, bson.M{"$set":bson.M{"status":leaves.Status}})
			if err2 != nil{
				return err2
	}
	if leaves.Status=="approved"{
		var user []model.Leaves
		if err := database.Collection().Find(bson.M{"email":leaves.Email}).All(&user); err != nil {
			return err
		}
		newdays:=user[0].Remholidays-leaves.Numdays
		err2:=database.Collection().Update(bson.M{"email":leaves.Email}, bson.M{"$set":bson.M{"remholidays":newdays}})
			if err2 != nil{
				return err2
			}
	}
	return nil
}
func GetManagersFromDB()(error, [] model.EmpDetails){
	var user []model.EmpDetails
	if err := database.Collection().Find(bson.M{}).All(&user); err != nil {
		return err,[]model.EmpDetails{}
	}
	return nil,user
}

func GetLeavesFromDB(empdetails model.Email) (error,[]model.Leaves){
	var list []model.Leaves
	if err := database.Collection().Find(bson.M{"email":empdetails.Email}).All(&list); err != nil {
		return err,[]model.Leaves{}
	}
	return nil,list
}
func GetAppliedLeavesFromDB(empdetails interface {}) (error,[]model.Leaves){
	var list []model.Leaves
	origin:= empdetails.(map[string]interface {})
	field := fmt.Sprintf("%v",origin["field"] )
	if origin["status"]=="applied"{
		if err := database.Leaves().Find(bson.M{field:origin["email"]}).All(&list); err != nil {
			return err,[]model.Leaves{}
		}
	}else{
		query := []bson.M{ 
			bson.M{"$match":bson.M{field:origin["email"],"status":origin["status"]}},
			bson.M{"$lookup": bson.M{"from": "EmployeeData","localField":"email","foreignField":"email","pipeline":[]bson.M{bson.M{"$project":bson.M{"firstname":1,"lastname":1,"remholidays":1}}},"as":"details"}},
			bson.M{"$unwind":bson.M{"path": "$details"}},
		  }
		if err := database.Leaves().Pipe(query).All(&list); err != nil {
			return err,[]model.Leaves{}
		}
	}
	
	return nil,list
}
func StoreLeaves(leaves model.Leaves) (error,bool){
	leaves.Applieddate=time.Now().UnixMilli()
	leaves.Status="pending"
	leaves.Lid=uuid.New().String()
	err:=database.Leaves().Insert(leaves)
	if err != nil{
		return err,false
	}
	return nil,true
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

func AdminallEmpListFromDB(empdetails interface{}) (error,[]model.EmpDetails){
	var employeelist []model.EmpDetails
	origin:= empdetails.(map[string]interface {})
	
	if err := database.Collection().Find(origin).All(&employeelist); err != nil {
		return err,[]model.EmpDetails{}
	}
	return nil,employeelist
}

func DeleteEmpFromDB(deletedetails map[string]string) (error,string){
		
		err:=database.Collection().Remove(deletedetails)
		if err!=nil{
			return err,""
		}
		return nil,"Permanently deleted employee"
}

func generatePassword() string {

	rand.Seed(time.Now().Unix())
    minSpecialChar := 1
    minNum := 1
    minUpperCase := 1
    passwordLength := 8
	lowerCharSet   := "abcdedfghijklmnopqrst"
    upperCharSet   := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet := "!@#$%&*"
    numberSet      := "0123456789"
    allCharSet     := lowerCharSet + upperCharSet + specialCharSet + numberSet

    var password strings.Builder

    //Set special character
    for i := 0; i < minSpecialChar; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < minNum; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < minUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}


// query := []bson.M{ // NOTE: slice of bson.M here
// 	bson.M{
// 		"$match":bson.M{
// 			"manageremail":empdetails.Email,
// 			"status":empdetails.Status
// 		}
// 	},
// 	bson.M{
// 	  "$lookup": bson.M{
// 		"from": "EmployeeData",
// 		"localField":"email",
// 		"foreignField":"email",
// 		"pipeline":[]bson.M{
// 			bson.M{
// 				"$project":bson.M{
// 					"firstname":1,
// 					"lastname":1
// 				}
// 			}
// 		},
// 		"as":"details"
// 	},
// 	},
// 	bson.M{
// 	  "$unwind":bson.M{
// 		"path": "$details"
// 	  }
// 	},
//   }