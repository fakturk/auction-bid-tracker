package main

import(
	"github.com/gorilla/mux"
	"net/http"	
	"github.com/fakturk/auction-bid-tracker/user"
	"github.com/fakturk/auction-bid-tracker/item"
	"github.com/fakturk/auction-bid-tracker/bid"
	
)


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

	router.HandleFunc("/bids", bid.GetBids).Methods("GET")
	router.HandleFunc("/bids/{userid}/{itemid}", bid.GetBid).Methods("GET")
	router.HandleFunc("/bids/{userid}/{itemid}/{amount}", bid.AddBid).Methods("POST")
	router.HandleFunc("/bids/{userid}/{itemid}/{amount}", bid.UpdateBid).Methods("PUT")
	router.HandleFunc("/bids/{userid}/{itemid}", bid.DeleteBid).Methods("DELETE")

	router.HandleFunc("/bids/{itemid}", bid.WinnerBidByItemID).Methods("GET")
	


	http.ListenAndServe(":8000", router)

}
