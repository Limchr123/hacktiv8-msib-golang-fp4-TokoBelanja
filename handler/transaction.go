package handler

import (
	"net/http"
	"tokoBelanja/helper"
	"tokoBelanja/transactionhistory"
	"tokoBelanja/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transactionhistory.Service
}

func NewtransactionHandler(service transactionhistory.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transactionhistory.TransactionInput

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
