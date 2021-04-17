package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/apis"
	repoMysql "github.com/masterraf21/reksti-ordering-backend/repositories/mysql"
	"github.com/masterraf21/reksti-ordering-backend/usecases"

	"github.com/masterraf21/reksti-ordering-backend/configs"
	"github.com/masterraf21/reksti-ordering-backend/utils/mysql"
)

// Server represents server
type Server struct {
	Reader      *sql.DB
	Writer      *sql.DB
	Port        string
	ServerReady chan bool
}

func main() {
	reader, writer := configureMySQL()
	serverReady := make(chan bool)
	server := Server{
		Reader:      reader,
		Writer:      writer,
		Port:        configs.Server.Port,
		ServerReady: serverReady,
	}
	server.Start()
}

func configureMySQL() (*sql.DB, *sql.DB) {
	readerConfig := mysql.Option{
		Host:     configs.MySQL.ReaderHost,
		Port:     configs.MySQL.ReaderPort,
		Database: configs.MySQL.Database,
		User:     configs.MySQL.ReaderUser,
		Password: configs.MySQL.ReaderPassword,
	}

	writerConfig := mysql.Option{
		Host:     configs.MySQL.WriterHost,
		Port:     configs.MySQL.WriterPort,
		Database: configs.MySQL.Database,
		User:     configs.MySQL.WriterUser,
		Password: configs.MySQL.WriterPassword,
	}

	reader, writer, err := mysql.SetupDatabase(readerConfig, writerConfig)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mysql", err)
	}

	log.Println("MySQL connection is successfully established!")

	return reader, writer
}

// Start will start server
func (s *Server) Start() {
	port := configs.Server.Port
	if port == "" {
		port = "8000"
	}

	r := new(mux.Router)

	orderRepo := repoMysql.NewOrderRepo(s.Reader, s.Writer)
	orderDetailsRepo := repoMysql.NewOrderDetailsRepo(s.Reader, s.Writer)
	orderUsecase := usecases.NewOrderUsecase(
		orderRepo,
		orderDetailsRepo,
	)

	menuRepo := repoMysql.NewMenuRepo(s.Reader, s.Writer)
	menuTypeRepo := repoMysql.NewMenuTypeRepo(s.Reader, s.Writer)
	menuUsecase := usecases.NewMenuUsecase(menuRepo, menuTypeRepo)

	ratingRepo := repoMysql.NewRatingRepo(s.Reader, s.Writer)

	customerRepo := repoMysql.NewCustomerRepo(s.Reader, s.Writer)
	customerUsecase := usecases.NewCustomerUsecase(customerRepo)

	paymentRepo := repoMysql.NewPaymentRepo(s.Reader, s.Writer)
	paymentTypeRepo := repoMysql.NewPaymentTypeRepo(s.Reader, s.Writer)
	paymentUsecase := usecases.NewPaymentUsecase(paymentRepo)
	paymentTypeUsecase := usecases.NewPaymentTypeUsecase(paymentTypeRepo)

	apis.NewOrderAPI(r, orderUsecase)
	apis.NewMenuAPI(r, menuUsecase)
	apis.NewRatingAPI(r, ratingRepo)
	apis.NewCustomerAPI(r, customerUsecase)
	apis.NewPaymentAPI(r, paymentUsecase, paymentTypeUsecase)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("Starting server on port %s!", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Shutting Down Server...")
			log.Fatal(err.Error())
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
