package model

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
	EmpID string `json:"empid"`
	Role string `json:"role"`
	Empstatus string `json:"empstatus"`
	
}

type Email struct {
	Email string `json:"email"`
	Status string `json:"status"`
	Id string `json:"id"`
}
type Name struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
type Leaves struct {
	Lid string `json:"lid"`
	Email string `json:"email"`
	Holidays int64 `json:"holidays"`
	Remholidays int64 `json:"remholidays"`
	Manageremail string `json:"manageremail"`
	Status string `json:"status"`
	Todate int64 `json:"todate"`
	Fromdate int64 `json:"fromdate"`
	Applieddate int64 `json:"applieddate"`
	Numdays int64 `json:"numdays"`
	Comment string `json:"comment"`
	Details EmpDetails  `json:"details"`
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
	Manageremail string `json:"manageremail"`
	Holidays int64 `json:"holidays"`
	Remholidays int64 `json:"remholidays"`
	Gender string `json:"gender"`
	Role string `json:"role"`
}

// Created global list which will be available throughout the application
var empList []EmpDetails
