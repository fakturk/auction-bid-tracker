package main

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":8000", router)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"GetUsers")
  }