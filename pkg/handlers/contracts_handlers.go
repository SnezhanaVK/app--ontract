// contracts.go
package handlers

import (
	"appContract/pkg/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "appContract/pkg/db/repository"

	"github.com/gorilla/mux"
)

// Contracts
func GetAllContracts(w http.ResponseWriter, r *http.Request) {
    if r.Method!=http.MethodGet{
        http.Error(w,"Invalid request method GetAllContracts",http.StatusBadRequest)
        return
    }
    contracts, err := db.DBgetContractAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var contractsResponse []map[string]interface{}
    for _, contract := range contracts {
        contractResponse := map[string]interface{}{
            "contract_id": contract.Id_contract,
            "name_contract": contract.Name_contract,
            "date_create_contract": contract.Data_contract_create,
            "user_id": contract.Id_user,
            "data_conclusion": contract.Data_conclusion,
            "data_start": contract.Data_contract_create,
            "data_end": contract.Data_end,
            "id_type": contract.Id_type,
            "name_type_contract": contract.Name_type,
            "id_counterparty": contract.Id_counterparty,
            "name_counterparty": contract.Name_counterparty,
            "id_status_contract": contract.Id_status_contract,
            "name_status_contract": contract.Name_status_contract,
            "id_teg": contract.Id_teg_contract,
            "name_teg": contract.Tegs_contract,
        }
        contractsResponse = append(contractsResponse, contractResponse)
    }
    data, err:=json.Marshal(contractsResponse)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}
func GetContractID(w http.ResponseWriter, r *http.Request) {
    if r.Method!=http.MethodGet{
        http.Error(w,"Invalid request method GetContract",http.StatusBadRequest)
        return
    }
    vars:=mux.Vars(r)
    contractId:=vars["contractID"]
    if contractId==""{
        http.Error(w,"Invalid contract_id",http.StatusBadRequest)
        return
    }
    id, err:= strconv.Atoi(contractId)
    if err != nil {
        http.Error(w, "Invalid contract_id", http.StatusBadRequest)
        return
    }
   
    contract, err := db.DBgetContractID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var contractsResponse []map[string]interface{}
    for _, contract := range contract {
        contractResponse := map[string]interface{}{
            "contract_id": contract.Id_contract,
            "name_contract": contract.Name_contract,
            "date_create_contract": contract.Data_contract_create,
            "user_id": contract.Id_user,
            "data_conclusion": contract.Data_conclusion,
            "data_start": contract.Data_contract_create,
            "data_end": contract.Data_end,
            "id_type": contract.Id_type,
            "name_type_contract": contract.Name_type,
            "id_counterparty": contract.Id_counterparty,
            "name_counterparty": contract.Name_counterparty,
            "id_status_contract": contract.Id_status_contract,
            "name_status_contract": contract.Name_status_contract,
            "id_teg": contract.Id_teg_contract,
            "name_teg": contract.Tegs_contract,
        }
        contractsResponse = append(contractsResponse, contractResponse)
    }
    data, err:=json.Marshal(contractsResponse)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}

func GetUserIDContracts(w http.ResponseWriter, r *http.Request) {

    if r.Method!=http.MethodGet{
    
        http.Error(w,"Invalid request method GetUserIDContracts",http.StatusBadRequest)
        return
    }
    vars:=mux.Vars(r)
    userId:=vars["userID"]
    if userId==""{
        http.Error(w,"Invalid user_id",http.StatusBadRequest)
        return
    }
    id, err:= strconv.Atoi(userId)
    if err != nil {
        log.Println(err)
        http.Error(w, "Invalid user_id", http.StatusBadRequest)
        return
    }
    contracts, err := db.DBgetContractUserId(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var contractsResponse []map[string]interface{}
    for _, contract := range contracts {
        contractResponse := map[string]interface{}{
            "contract_id": contract.Id_contract,
            "name_contract": contract.Name_contract,
            "date_create_contract": contract.Data_contract_create,
            "user_id": contract.Id_user,
            "data_conclusion": contract.Data_conclusion,
            "data_start": contract.Data_contract_create,
            "data_end": contract.Data_end,
            "id_type": contract.Id_type,
            "name_type_contract": contract.Name_type,
            "id_counterparty": contract.Id_counterparty,
            "name_counterparty": contract.Name_counterparty,
            "id_status_contract": contract.Id_status_contract,
            "name_status_contract": contract.Name_status_contract,
            "id_teg": contract.Id_teg_contract,
            "name_teg": contract.Tegs_contract,
        }
        contractsResponse = append(contractsResponse, contractResponse)
    }
    data, err:=json.Marshal(contractsResponse)  
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(data)
        
}



func PostCreateContract(w http.ResponseWriter, r *http.Request) {
  if r.Method!=http.MethodPost{
        http.Error(w,"Invalid request method PostCreateContract",http.StatusBadRequest)
        return
    }
    var contract models.Contracts
    err:=json.NewDecoder(r.Body).Decode(&contract)
    if err!=nil{
        log.Println(err)
        http.Error(w,"Invalid request body PostCreateContract",http.StatusBadRequest)
        return
    }
    err = db.DBaddContract(contract)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Contract created successfully"})
}

func PutChangeContract(w http.ResponseWriter, r *http.Request) {
    if r.Method!=http.MethodPut{
        http.Error(w,"Invalid request method PutChangeContract",http.StatusBadRequest)
        return
    }
    var contract models.Contracts
    err:=json.NewDecoder(r.Body).Decode(&contract)
    if err!=nil{
        http.Error(w,"Invalid request body PutChangeContract",http.StatusBadRequest)
        return
    }
    err = db.DBchangeContract(contract)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Contract updated successfully"})
}


func PutChangeContractUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method UpdateContractUser", http.StatusBadRequest)
        return
    }

    var data map[string]int
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    contract_id, ok := data["id_contract"]
    if !ok {
        http.Error(w, "Missing id_contract in request body", http.StatusBadRequest)
        return
    }

    userId, ok := data["id_user"]
    if !ok {
        http.Error(w, "Missing id_user in request body", http.StatusBadRequest)
        return
    }

    err = db.DBchangeContractUser(contract_id, userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Contract user updated successfully"})
}

func DeleteContract(w http.ResponseWriter, r *http.Request) {
    if r.Method!=http.MethodDelete{
        http.Error(w,"Invalid request method DeleteContract",http.StatusBadRequest)
        return
    }
    vars:=mux.Vars(r)
    contractId, err := strconv.Atoi(vars["contractID"])
    if err != nil {
        http.Error(w, "Invalid contract_id", http.StatusBadRequest)
        return
    }
    err =db.DBdeleteContract(contractId)
    if err !=nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Contract deleted successfully"})
}
