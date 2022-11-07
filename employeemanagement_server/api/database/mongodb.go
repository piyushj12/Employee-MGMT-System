package helpers

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitDB() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("EmployeeDetails")
}

func Collection() *mgo.Collection {
	return db.C("EmployeeData")
}
func Leaves() *mgo.Collection {
	return db.C("LeavesData")
}
func Attendance() *mgo.Collection {
	return db.C("Attendance")
}