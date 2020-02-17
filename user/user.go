package user

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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"GetUsers\n")
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(users)
  }

  func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

  func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
  }


  func AddUserWithName(w http.ResponseWriter, r *http.Request) {
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

  func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
	  if user.ID == params["id"] {
		users = append(users[:index], users[index+1:]...)
		break
	  }
	}
	json.NewEncoder(w).Encode(users)
  }