package repositories

import (
	"_waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindBooksById(BookID []int) ([]models.Book, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Books").Find(&transactions).Error

	return transactions, err
}

func (r *repository) FindBooksById(BookID []int) ([]models.Book, error) {
	var Books []models.Book
	err := r.db.Find(&Books, BookID).Error

	return Books, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
var transaction models.Transaction
// not yet using category relation, cause this step doesnt Belong to Many
err := r.db.Preload("User").Preload("Profile").Preload("Books").First(&transaction, ID).Error

return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Profile").Preload("Books").Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Books").Raw("UPDATE transactions SET user_id=?, attachment=?, book_id=?, total_payment=?, status=? WHERE id=?", transaction.UserID, transaction.Attachment, transaction.BookID,transaction.Total, transaction.Status,ID).Scan(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}