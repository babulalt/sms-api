package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	FirstName      string
	LastName       string
	Gender         string
	DOB            string
	MobileNumber   string
	Email          string
	ParanetName    string
	ParanetMobile  string
	ParentRelation string
	ParentEmail    string
	CID            uint
	Course         Course
}
