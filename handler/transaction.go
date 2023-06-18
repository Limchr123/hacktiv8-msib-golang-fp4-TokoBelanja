package handler

import (
	"net/http"
	"strconv"
	"tokoBelanja/helper"
	"tokoBelanja/transaction"
	"tokoBelanja/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.ServiceTransaction
}

func NewtransactionHandler(service transaction.ServiceTransaction) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.TransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	input.UserID = currentUser.ID

	newUser, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newUser)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransaction(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("product_id"))
	userID, _ := strconv.Atoi(c.Query("user_id"))

	sosmed, err := h.transactionService.GetTransaction(productID, userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, transaction.FormatterGetCampaign(sosmed))
	c.JSON(http.StatusOK, response)
}
