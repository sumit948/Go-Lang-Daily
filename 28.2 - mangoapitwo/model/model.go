package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	EmpId      primitive.ObjectID `json:"_empid,omitempty" bson:"_id,omitempty"`
	EmpName    string             `json:"empName,omitempty"`
	EmpDept    string             `json:"empdept,omitempty"`
	Empcontact int64              `json:"contact,omitempty"`
	IsAvilable bool               `json:"isavailable,omitempty"`
}
