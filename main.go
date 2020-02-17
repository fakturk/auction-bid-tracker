package main

import(
	"github.com/gorilla/mux"
	"net/http"	
	"github.com/fakturk/auction-bid-tracker/user"
	"github.com/fakturk/auction-bid-tracker/item"
	
)





type Bid struct {
	UserID string `json:"userid"`
	ItemID string `json:"itemid"`
	Amount string `json:"amount"`
  }

var bids []Bid

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/users", user.GetUsers).Methods("GET")
	router.HandleFunc("/users/id/{id}", user.GetUserByID).Methods("GET")
	router.HandleFunc("/users", user.AddUser).Methods("POST")
	router.HandleFunc("/users/{name}", user.AddUserWithName).Methods("POST")
	router.HandleFunc("/users/id/{id}", user.DeleteUserByID).Methods("DELETE")

	router.HandleFunc("/items", item.GetItems).Methods("GET")
	router.HandleFunc("/items/id/{id}", item.GetItemByID).Methods("GET")
	router.HandleFunc("/items", item.AddItem).Methods("POST")
	router.HandleFunc("/items/{name}", item.AddItemWithName).Methods("POST")
	router.HandleFunc("/items/id/{id}", item.DeleteItemByID).Methods("DELETE")

	http.ListenAndServe(":8000", router)

}
