package handler

import (
	"net/http"
	"strconv"

	"github.com/pramudya3/backend/payment/domain"
	"github.com/supertokens/supertokens-golang/recipe/session"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	phUsecase domain.TransactionUsecase
}

func NewTransactionHandler(phUc domain.TransactionUsecase) *transactionHandler {
	return &transactionHandler{
		phUsecase: phUc,
	}
}

func (h *transactionHandler) Fetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
		userId := sessionContainer.GetUserID()
		id, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		res, err := h.phUsecase.FetchTransaction(c, uint64(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
		}

		c.JSON(http.StatusOK, domain.ResponseSuccess(res))
	}
}
