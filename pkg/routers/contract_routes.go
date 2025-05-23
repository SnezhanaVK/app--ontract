package routers

// contract_routes.go

import (
	"appContract/pkg/handlers"

	"github.com/gorilla/mux"
)

func ContractRoutes(router *mux.Router) {

	router.HandleFunc("/api/contracts", handlers.GetAllContracts).Methods("GET")
	router.HandleFunc("/api/contracts/user/{userID}", handlers.GetUserIDContracts).Methods("GET")
	router.HandleFunc("/api/contracts/{contractID}", handlers.GetContractID).Methods("GET")
	router.HandleFunc("/api/contracts/byType/{idType}", handlers.GetAllContractsByType).Methods("GET")
	router.HandleFunc("/api/contracts/byDateCreate", handlers.PostAllContractsByDateCreate).Methods("POST")
	router.HandleFunc("/api/contracts/byTeg/{id_teg_contract}", handlers.GetAllContractsByTegs).Methods("GET")
	router.HandleFunc("/api/contracts/byStatus/{id_status_contract}", handlers.GetAllContractsByStatus).Methods("GET")
	router.HandleFunc("/api/contracts/create", handlers.PostCreateContract).Methods("POST")
	router.HandleFunc("/api/contracts/{contractID}", handlers.PutChangeContract).Methods("PUT")
	router.HandleFunc("/api/contracts/userchange", handlers.PutChangeContractUser).Methods("PUT")
	router.HandleFunc("/api/contracts/{contractID}", handlers.DeleteContract).Methods("DELETE")

}
