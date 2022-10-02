package models

import "time"

type Transaction struct {
  ID        			int       				`json:"id" gorm:"primary_key:auto_increment"`
  UserID      			int						`json:"user_id" form:"user_id" gorm:"type: int"`
  User      			UserResponse   					`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Attachment     		string    				`json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
  BookID     			[]int    				`json:"book_id" form:"book_id" gorm:"type: int"`
  Books     			[]Book   				`json:"booksPurchased" gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Total					  int						    `json:"totalPayment"`
  Status				  string					  `json:"status" gorm:"type:varchar(255)"`
  CreatedAt 			time.Time 				`json:"-"`
  UpdatedAt 			time.Time 				`json:"-"`
}

type TransactionResponse struct {
	ID        				int       				`json:"-"`
	UserID      			int						    `json:"user_id"`
	// User      				ProfileResponse    		`json:"-" gorm:"foreignKey:UserID"`
	BookID     				[]int    				    `json:"-"`
	Books     				[]Book   	        `json:"booksPurchased" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt 				time.Time 				`json:"-"`
	UpdatedAt 				time.Time 				`json:"-"`
  }

func (TransactionResponse) TableName() string {
	return "transactions"
}