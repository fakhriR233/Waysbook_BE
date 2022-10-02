package handlers

import (
	dto "_waysbook/dto/result"
	transactiondto "_waysbook/dto/transaction"
	"_waysbook/models"
	"_waysbook/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
  }

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	transaction, err := h.TransactionRepository.FindTransactions()
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK, 
		Data: transaction,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(transaction)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
  
	// request := new(transactiondto.TransactionRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//   w.WriteHeader(http.StatusBadRequest)
	//   response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	//   json.NewEncoder(w).Encode(response)
	//   return
	// }

	total, _ := strconv.Atoi(r.FormValue("totalPayment"))

	var booksID []int
	for _, r := range r.FormValue("book_id") {
		if int(r-'0') >= 0 {
			booksID = append(booksID, int(r-'0'))
		}
	}
  
	request := transactiondto.TransactionRequest{
		UserID:    				userId,
		Attachment:    			r.FormValue("attachment"),
		BookID:     			booksID,
		Total:    				total,
		Status:      			"Pending",
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	book, _ := h.TransactionRepository.FindBooksById(booksID)

	transaction := models.Transaction{
		Attachment:    			request.Attachment,
		Total:    				request.Total,
		Books: 					book,				
		UserID:    				userId,
		Status:      			"Pending",
	}

  
	transaction, err = h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	transaction, _ = h.TransactionRepository.GetTransaction(transaction.ID)
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(transactiondto.TransactionUpdateRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transactionDataOld, _ := h.TransactionRepository.GetTransaction(id)
  
	transaction := models.Transaction{}
  
	if request.UserID != 0 {
		transaction.UserID = request.UserID
		transactionDataNew, _ := h.TransactionRepository.GetTransaction(transaction.UserID)
		transaction.User = transactionDataNew.User
	}else {
		transaction.UserID = transactionDataOld.UserID
		transaction.User = transactionDataOld.User
	}
	
	if request.Attachment != "" {
	transaction.Attachment = request.Attachment
	}else {
		transaction.Attachment = transactionDataOld.Attachment
	}

	var booksID []int
	for _, r := range r.FormValue("book_id") {
		if int(r-'0') >= 0 {
			booksID = append(booksID, int(r-'0'))
		}
	}
	  
	if request.BookID != nil {
		transaction.BookID = booksID
		}else {
		transaction.BookID = transactionDataOld.BookID
	}

	if request.Total != 0 {
		transaction.Total = request.Total
	}else {
		transaction.Total = transactionDataOld.Total
	}
	if request.Status != "" {
		transaction.Status = request.Status
	}else {
		transaction.Status = transactionDataOld.Status
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransactionUpdate(data)}
	json.NewEncoder(w).Encode(response)
  }

func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	data, err := h.TransactionRepository.DeleteTransaction(transaction,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteTransaction(data)}
	json.NewEncoder(w).Encode(response)
  }

func convertResponseTransaction(u models.Transaction) models.Transaction {
	return models.Transaction{
		ID:				u.ID,
	  User:    			u.User,
	  Attachment:    	u.Attachment,
	  BookID:			u.BookID,
	  Status:      		u.Status,
	}
}

func convertResponseTransactionUpdate(u models.Transaction) transactiondto.TransactionUpdateResponse {
	return transactiondto.TransactionUpdateResponse{
		ID:				u.ID,
	  User:    			u.User,
	  Attachment:    	u.Attachment,
	  BookID:			u.BookID,
	  Status:      		u.Status,
	}
}

func convertResponseDeleteTransaction(u models.Transaction) transactiondto.TransactionDeleteResponse {
	return transactiondto.TransactionDeleteResponse{
	  ID:    u.ID,
	}
}