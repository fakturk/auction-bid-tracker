package main

import(
	"github.com/gorilla/mux"
	"net/http"	
	"github.com/fakturk/auction-bid-tracker/user"
	
)



type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
  }

var items []Item

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
	http.ListenAndServe(":8000", router)

}
