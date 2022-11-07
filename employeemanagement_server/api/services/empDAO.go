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
			query=append(query,map[string]interface{}{key:bson.M{"$regex":"^"+fmt.Sprintf("%v", value),"$options":"i"}})
	}
	fmt.Println(query)
	if err := database.Collection().Find(bson.M{"$or":query}).All(&employeelist); err != nil {
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
func DashBoardData(email interface{}) (error, map[string]interface {}){
		origin:= email.(map[string]interface {})
		details:=make(map[string]interface {}) 
		k:=make(map[string]interface {}) 
		var empdetails []model.EmpDetails
		if err := database.Collection().Find(origin).All(&empdetails); err != nil {
			return err,k
		}
		query := []bson.M{ 
			bson.M{"$match":bson.M{"email":origin["email"],"status":"pending"}},
			bson.M{"$count":"pending"},
		  }
		  var list []interface{}
		  if err := database.Leaves().Pipe(query).All(&list); err != nil {
			return err,k
		}
		var managerdetails []model.EmpDetails
		if err := database.Collection().Find(bson.M{"email":empdetails[0].Manageremail}).All(&managerdetails); err != nil {
			return err,k
		}
		type pending struct{
			Pending int64 `json:"pending"`
		}
		var pendingap pending
		pendingap.Pending=0
		fmt.Println(empdetails)
		fmt.Println(managerdetails)
		fmt.Println(list)
		details["holidays"]=empdetails[0].Holidays
		details["remholidays"]=empdetails[0].Remholidays
		details["department"]=empdetails[0].Department
		if len(list)>0{
			details["approval"]=list[0]
		}else{
			details["approval"]=pendingap
		}
		
		details["manager"]=managerdetails[0].Firstname+" "+managerdetails[0].Lastname
		return nil,details
		
}
func GetAttendanceFromDB(attendance interface{}) (error,[]model.Attendance){
	origin:= attendance.(map[string]interface {})
	months:=make(map[int64]string)
	months[0]="January"
	months[1]="February"
	months[2]="March"
	months[3]="April"
	months[4]="May"
	months[5]="June"
	months[6]="July"
	months[7]="August"
	months[8]="September"
	months[9]="October"
	months[10]="November"
	months[11]="December"
	t:=time.Now()
	m:=(origin["month"].(float64))+1
	minDate:= time.Date(int(origin["year"].(float64)), time.Month(m), 1, 0, 0, 0, 0, t.Location())
	fmt.Println(minDate)
	maxDate:=minDate.AddDate(0,1,0)
	fmt.Println(maxDate)
	var result []model.Attendance
	if err := database.Attendance().Find(bson.M{"clockin":bson.M{"$gte":minDate,"$lt":maxDate},"empid":origin["empid"]}).All(&result); err != nil {
		return err,[]model.Attendance{}
	}
	return nil,result
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


func CaptureClockinToDB(attendance interface{})(error,bool){
	origin:= attendance.(map[string]interface {})
	origin["clockin"]=time.Now()
	origin["status"]=false
	err:=database.Attendance().Insert(origin)
	if err != nil{
		return err,false
	}
	return nil,true
}

func CaptureClockoutToDB(attendance interface{})(error,bool){
	origin:= attendance.(map[string]interface {})
	origin["clockout"]=time.Now()
	err2:=database.Attendance().Update(bson.M{"empid":origin["empid"],"clockin":bson.M{"$gte":Bod(time.Now())},"status":false}, bson.M{"$set":bson.M{"clockout":origin["clockout"],"status":true}})
			if err2 != nil{
				return err2,false
			}
	return nil,true
}
func Bod(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
func Eod(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 23, 59, 59, 59, t.Location())
}
func IsclockedinDB(attendance model.Attendance)(error, map[string]interface{}){
	var ans []model.Attendance
	result:=make(map[string]interface{})
	if err := database.Attendance().Find(bson.M{"clockin":bson.M{"$gte":Bod(time.Now())},"empid":attendance.EmpID}).All(&ans); err != nil {
		return err,result
	}
	if len(ans)>0{
		result["status"]=true
		result["clockin"]=ans[0].Clockin
		return nil,result
	}else{
		result["status"]=false
		return nil,result
	}
	
}

func IsclockedoutDB(attendance model.Attendance)(error, map[string]interface{}){
	var ans []model.Attendance
	result:=make(map[string]interface{})
	if err := database.Attendance().Find(bson.M{"clockout":bson.M{"$gte":Bod(time.Now())},"empid":attendance.EmpID,"status":true}).All(&ans); err != nil {
		return err,result
	}
	if len(ans)>0{
		result["status"]=true
		result["clockin"]=ans[0].Clockin
		result["clockout"]=ans[0].Clockout
		return nil,result
	}else{
		result["status"]=false
		return nil,result
	}
	
}