package models

import "time"

type Book struct {
  ID        					int       					`json:"id" gorm:"primary_key:auto_increment"`
  Title      					string						`json:"title" gorm:"type: varchar(255)"`
  PublicationDate      			string    					`json:"publicationDate" gorm:"type: varchar(255)"`
  Pages     					int    						`json:"pages" gorm:"type: int"`
  ISBN     						int    						`json:"ISBN" gorm:"type: int"`
  Author						string						`json:"author" gorm:"type: varchar(255)"`
  Price     					int    						`json:"price" gorm:"type: int"`
  Description					string						`json:"description" gorm:"type:text" form:"desc"`
  BookAttachment				string						`json:"bookAttachment" gorm:"type: varchar(255)"`
  Thumbnail						string						`json:"thumbnail" gorm:"type: varchar(255)"`
  CreatedAt 					time.Time 					`json:"-"`
  UpdatedAt 					time.Time 					`json:"-"`
}

type BookResponse struct {
  ID        					int       					`json:"id"`
  Title      					string						`json:"title"`
  PublicationDate      			string    					`json:"publicationDate"`
  Pages     					int    						`json:"pages"`
  ISBN     						int    						`json:"ISBN"`
  Author						string						`json:"author"`
  Price     					int    						`json:"price"`
  Description					string						`json:"description"`
  BookAttachment				string						`json:"bookAttachment"`
  Thumbnail						string						`json:"thumbnail"`
  CreatedAt 					time.Time 					`json:"-"`
  UpdatedAt 					time.Time 					`json:"-"`
}

type BookUserResponse struct {
  ID        					int       					`json:"id"`
  Title      					string						`json:"title"`
  PublicationDate      			string    					`json:"publicationDate"`
  Pages     					int    						`json:"pages"`
  ISBN     						int    						`json:"ISBN"`
  Author						string						`json:"author"`
  Price     					int    						`json:"price"`
  Description					string						`json:"description"`
  BookAttachment				string						`json:"bookAttachment"`
  Thumbnail						string						`json:"thumbnail"`
  CreatedAt 					time.Time 					`json:"-"`
  UpdatedAt 					time.Time 					`json:"-"`
}

func (BookResponse) TableName() string {
	return "books"
}

func (BookUserResponse) TableName() string {
	return "books"
}