package models

import "time"

// User model struct
type User struct {
  ID          	int							`json:"id"`
  FullName 		string						`json:"fullname" gorm:"type: varchar(255)"`
  Email		    string 						`json:"email" gorm:"type: varchar(255)"`
  Password 	  	string						`json:"password" gorm:"type: varchar(255)"`
  Status 	  	string						`json:"status" gorm:"type: varchar(50)"`
  Profile     ProfileResponse   `json:"profile"`
  Transaction []TransactionResponse `json:"transaction"`
  CreatedAt 	time.Time					`json:"created_at"`
  UpdatedAt 	time.Time					`json:"updated_at"`
}

type UserResponse struct {
	ID          	int			`json:"id"`
	FullName 		  string		`json:"fullname" gorm:"type: varchar(255)"`
	Email		      string 		`json:"email" gorm:"type: varchar(255)"`
  }

func (UserResponse) TableName() string {
	return "users"
}