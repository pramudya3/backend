package cmd

import (
	"github.com/pramudya3/backend/payment/handler"
	midd "github.com/pramudya3/backend/payment/handler/middleware"
	"github.com/pramudya3/backend/payment/repository"
	"github.com/pramudya3/backend/payment/usecase"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func loadRoutes(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), midd.Cors())

	checkDb := true
	router.Use(midd.VerifySession(&sessmodels.VerifySessionOptions{CheckDatabase: &checkDb}))
	tx := router.Group("/transaction")
	loadTransactionRoutes(db, tx)

	return router
}

func loadTransactionRoutes(db *gorm.DB, g *gin.RouterGroup) {
	paRepo := repository.NewPaymentAccountRepository(db)
	txRepo := repository.NewTransactionRepository(db)
	paUc := usecase.NewPaymentTypeUsecase(paRepo)
	uc := usecase.NewTransactionUsecase(paUc, txRepo)

	tx := handler.NewTransactionHandler(uc)

	g.POST("/send", tx.Send())
	g.POST("/withdraw", tx.Withdraw())
}
