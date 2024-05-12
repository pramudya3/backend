package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pramudya3/backend/payment/domain"

	"gorm.io/gorm"
)

type App struct {
	router http.Handler
	db     *gorm.DB
	Config Config
}

func New() *App {
	cfg := LoadConfig()

	db, err := initGorm(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{}, &domain.Transaction{}, &domain.PaymentAccount{})

	routes := loadRoutes(db)

	return &App{
		Config: cfg,
		db:     db,
		router: routes,
	}
}

func (a *App) Start() {
	server := &http.Server{
		Addr:    fmt.Sprintf("%v", a.Config.ServerAddr),
		Handler: a.router,
	}

	go func() {
		// service connections
		log.Println("Server starting at ", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	// 	select {
	// 	case <-ctx.Done():
	// 		log.Println("timeout of 5 seconds.")
	// 		log.Println("Server exiting")
	// 	}
}
