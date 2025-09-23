package main

import (
	"fmt"
	"log"
	"net/http"
	"payment-service/config"
	"payment-service/db"
	handlers "payment-service/handler"
	"payment-service/repository"
	"payment-service/service"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.InitDB(cfg.DBUrl)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	repo := repository.NewPaymentRepository(database)
	svc := service.NewPaymentService(repo)
	handler := &handlers.PaymentHandler{Service: svc}

	r := mux.NewRouter()

	r.HandleFunc("/payments", handler.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id:[0-9]+}", handler.GetPayment).Methods("GET")
	r.HandleFunc("/payments", handler.ListPayments).Methods("GET")
	r.HandleFunc("/payments/{id:[0-9]+}/status", handler.UpdatePaymentStatus).Methods("PUT")
	r.HandleFunc("/payments/{id:[0-9]+}/refund", handler.RefundPayment).Methods("POST")

	addr := ":" + cfg.Port
	fmt.Println("Server running on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
