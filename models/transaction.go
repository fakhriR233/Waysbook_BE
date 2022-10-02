package models

import "time"

type Transaction struct {
  ID        			int       				`json:"id" gorm:"primary_key:auto_increment"`
  UserID      			int						`json:"user_id" form:"user_id" gorm:"type: int"`
  User      			UserResponse   					`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Attachment     		string    				`json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
  BookID     			[]int    				`json:"-" form:"book_id" gorm:"-"`
  Books     			[]Book   				`json:"booksPurchased" gorm:"many2many:book_transaction;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Total					int						`json:"totalPayment"`
  Status				string					`json:"status" gorm:"type:varchar(255)"`
  CreatedAt 			time.Time 				`json:"-"`
  UpdatedAt 			time.Time 				`json:"-"`
}

type TransactionResponse struct {
	ID        				int       				`json:"-"`
	UserID      			int						    `json:"-"`
	User      				ProfileResponse    		`json:"-" gorm:"foreignKey:UserID"`
	BookID     				[]int    				    `json:"-"`
	Books     				[]BookResponse    	`json:"booksPurchased" gorm:"many2many:book_transaction;foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt 				time.Time 				`json:"-"`
	UpdatedAt 				time.Time 				`json:"-"`
  }

func (TransactionResponse) TableName() string {
	return "transactions"
}