package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Producter struct {
	PID    string  `json:"PID"`
	PrName string  `json:"PrName"`
	Pprice float64 `json:"Pprice"`
}

var storagecenter []Producter

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage called: Welcome to Kloudone")
}

func getstoragecenter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Storage center called")
	json.NewEncoder(w).Encode(storagecenter)

}
func updateproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Producter
	_ = json.NewDecoder(r.Body).Decode(&product)

	params := mux.Vars(r)
	_deleteItemAtId(params["pid"])

	storagecenter = append(storagecenter, product)
	json.NewEncoder(w).Encode(product)

}

func createproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Producter
	_ = json.NewDecoder(r.Body).Decode(&product)
	storagecenter = append(storagecenter, product)
	json.NewEncoder(w).Encode(product)

}
func deleteproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	_deleteItemAtId(params["pid"])
	json.NewEncoder(w).Encode(storagecenter)

}

func _deleteItemAtId(pid string) {
	for index, item := range storagecenter {
		if item.PID == pid {
			storagecenter = append(storagecenter[:index], storagecenter[index+1:]...)
			break
		}
	}
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/storagecenter", getstoragecenter).Methods("GET")
	router.HandleFunc("/storagecenter", createproduct).Methods("POST")
	router.HandleFunc("/storagecenter/{pid}", deleteproduct).Methods("DELETE")
	router.HandleFunc("/storagecenter/{pid}", updateproduct).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func main() {
	storagecenter = append(storagecenter, Producter{
		PID:    "1",
		PrName: "IPHONE12",
		Pprice: 123.32,
	})
	handleRequests()
}
