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

	// unprotected routes
	signupRoutes(db, router.Group("/"))

	// protected routes
	checkDb := true
	router.Use(midd.VerifySession(&sessmodels.VerifySessionOptions{CheckDatabase: &checkDb}))

	userRoutes(db, router.Group("/user"))
	paymentAccountRoutes(db, router.Group("/payment-account"))
	paymentHistoryRoutes(db, router.Group("/payment-history"))

	return router
}

func signupRoutes(db *gorm.DB, g *gin.RouterGroup) {
	paymentRepo := repository.NewPaymentAccountRepository(db)
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(userRepo, paymentRepo)
	uc := handler.NewHandlerUser(usecase)

	g.POST("/signup", uc.Signup())
	g.POST("/login", uc.Login())
}

func userRoutes(db *gorm.DB, g *gin.RouterGroup) {
	paymentRepo := repository.NewPaymentAccountRepository(db)
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(userRepo, paymentRepo)
	uc := handler.NewHandlerUser(usecase)

	g.GET("/my-profile", uc.FindByUsername())
	g.GET("/logout", uc.Logout())
}

func paymentAccountRoutes(db *gorm.DB, g *gin.RouterGroup) {
	repo := repository.NewPaymentAccountRepository(db)
	usecase := usecase.NewPaymentTypeUsecase(repo)
	uc := handler.NewpaymentAccountHandler(usecase)

	g.POST("/", uc.CreatePaymentAccount())
}

func paymentHistoryRoutes(db *gorm.DB, g *gin.RouterGroup) {
	repo := repository.NewTransactionRepository(db)
	usecase := usecase.NewPaymentHistoryUsecase(repo)
	uc := handler.NewTransactionHandler(usecase)

	g.GET("/", uc.Fetch())
}
