package models

import "time"

// User model struct
type User struct {
  ID          	int							`json:"id"`
  FullName 		string						`json:"fullname" gorm:"type: varchar(255)"`
  Email		    string 						`json:"email" gorm:"type: varchar(255)"`
  Password 	  	string						`json:"password" gorm:"type: varchar(255)"`
  Gender 	  	string						`json:"gender" gorm:"type: varchar(255)"`
  Phone 	 	string						`json:"phone" gorm:"type: varchar(255)"`
  Avatar 	  	string						`json:"avatar" gorm:"type: varchar(255)"`
  PurchasedBook  []BookUserResponse 	`json:"purchasedBook" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  CreatedAt 	time.Time					`json:"created_at"`
  UpdatedAt 	time.Time					`json:"updated_at"`
}

type UserResponse struct {
	ID          	int			`json:"id"`
	FullName 		string		`json:"fullname" gorm:"type: varchar(255)"`
	Email		    string 		`json:"email" gorm:"type: varchar(255)"`
	Password 	  	string		`json:"password" gorm:"type: varchar(255)"`
	Gender 	  		string		`json:"gender" gorm:"type: varchar(255)"`
	Phone 	 		string		`json:"phone" gorm:"type: varchar(255)"`
	Avatar 	  		string		`json:"avatar" gorm:"type: varchar(255)"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
  }

func (UserResponse) TableName() string {
	return "users"
}