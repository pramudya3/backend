package handler

import (
	"net/http"
	"strconv"

	"github.com/pramudya3/backend/payment/domain"
	"github.com/supertokens/supertokens-golang/recipe/session"

	"github.com/gin-gonic/gin"
)

type paymentAccountHandler struct {
	uc domain.PaymentAccountUsecase
}

func NewpaymentAccountHandler(uc domain.PaymentAccountUsecase) *paymentAccountHandler {
	return &paymentAccountHandler{
		uc: uc,
	}
}

func (p *paymentAccountHandler) CreatePaymentAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
		userId := sessionContainer.GetUserID()
		id, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		payload := &domain.CreatePaymentAccount{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		newAccount := &domain.PaymentAccount{
			User_ID: uint64(id),
			Type:    payload.Type,
			Balance: payload.Balance,
		}

		if err := p.uc.AddNewPaymentAccount(c, newAccount); err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		c.Status(http.StatusCreated)
	}
}
