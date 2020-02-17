package main

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
	"encoding/json"
	"math/rand"
  	"strconv"
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
	router.HandleFunc("/users/id/{id}", getUserByID).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users/{name}", addUserWithName).Methods("POST")
	http.ListenAndServe(":8000", router)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"GetUsers\n")
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(users)
  }

  func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range users {
	  if user.ID == params["id"] {
		json.NewEncoder(w).Encode(user)
		return
	  }
	}
	json.NewEncoder(w).Encode(&User{})
  }

  func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
  }


  func addUserWithName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	var user User
	_ = json.NewDecoder(r.Body).Decode(user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	user.Name = params["name"]
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
  }