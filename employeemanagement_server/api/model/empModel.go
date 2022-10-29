package model

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
	EmpID string `json:"empid"`
	Role string `json:"role"`
	Empstatus string `json:"empstatus"`
}

type EmpDetails struct {
	// Id   string `json:"id" bson:"_id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Department string `json:"department"`
	Address string `json:"address"`
	EmpID string `json:"empid"`
	Empstatus string `json:"empstatus"`
	Email string `json:"email"`
	Password string `json:"password"`
	Contact string `json:"contact"`
	Level string `json:"level"`
	Role string `json:"role"`
}

type DeleteData struct{
	EmpID string `json:"empid"`
	PermanentlyDelete bool `json:"permanentlyDelete"`
}

type RestoreData struct{
	EmpID string `json:"empid"`
}
// Created global list which will be available throughout the application
var empList []EmpDetails
