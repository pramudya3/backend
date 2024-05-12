package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/backend/payment/domain"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

type transactionHandler struct {
	txUsecase domain.TransactionUsecase
}

func NewTransactionHandler(txUc domain.TransactionUsecase) *transactionHandler {
	return &transactionHandler{
		txUsecase: txUc,
	}
}

func (t *transactionHandler) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
		userId := sessionContainer.GetUserID()
		id, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		payload := &domain.TransactionSend{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		newTx := &domain.Transaction{
			UserID:        uint64(id),
			ToUserID:      &payload.ToUserID,
			Amount:        payload.Amount,
			PaymentMethod: payload.PaymentMethod,
		}

		resp, err := t.txUsecase.Send(c, newTx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		if resp.Status != "Success" {
			c.JSON(http.StatusBadRequest, domain.ResponseSuccess(resp))
			return
		}

		c.JSON(http.StatusOK, domain.ResponseSuccess(resp))
	}
}

func (t *transactionHandler) Withdraw() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
		userId := sessionContainer.GetUserID()
		id, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		payload := &domain.TransactionWithdraw{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		newTx := &domain.Transaction{
			UserID:        uint64(id),
			Amount:        payload.Amount,
			PaymentMethod: payload.PaymentMethod,
		}

		resp, err := t.txUsecase.Withdraw(c, newTx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		if resp.Status != "Success" {
			c.JSON(http.StatusBadRequest, domain.ResponseSuccess(resp))
			return
		}

		c.JSON(http.StatusOK, domain.ResponseSuccess(resp))
	}
}
