package models

import "time"

type Transaction struct {
  ID        			int       				`json:"id" gorm:"primary_key:auto_increment"`
  UserID      			int						`json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  User      			UserResponse    		`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Attachment     		string    				`json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
  Book     				[]int    				`json:"book_id" form:"book_id"`
  BooksPurchased     	[]BookResponse    		`json:"booksPurchased" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Total					int						`json:"totalPayment" gorm:"type:varchar(25)"`
  Status				string					`json:"status" gorm:"type:varchar(255)"`
  CreatedAt 			time.Time 				`json:"-"`
  UpdatedAt 			time.Time 				`json:"-"`
}

type TransactionResponse struct {
	ID        				int       				`json:"id"`
	UserID      			int						`json:"-"`
	User      				UserResponse    		`json:"user" gorm:"foreignKey:UserID"`
	Attachment     			string    				`json:"attachment"`
	Book     				[]int    				`json:"-"`
	BooksPurchased     		[]BookResponse    		`json:"booksPurchased" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total					int						`json:"totalPayment" gorm:"type:varchar(25)"`
	Status					string					`json:"status" gorm:"type:varchar(255)"`
	CreatedAt 				time.Time 				`json:"-"`
	UpdatedAt 				time.Time 				`json:"-"`
  }

func (TransactionResponse) TableName() string {
	return "transactions"
}