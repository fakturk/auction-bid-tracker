package main

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
  }

var users []User

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
	router.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":8000", router)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"GetUsers")
  }