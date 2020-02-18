package item

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
	"encoding/json"
	"math/rand"
  	"strconv"
)

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
  }

var items []Item

func AddMockItems(){
	//adding mock data for test usage 
	items = append(items, Item{ID: "727887", Name: "Golang-Book"})
	items = append(items, Item{ID: "727888", Name: "Iphone"})
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"GetItems\n")
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(items)
  }

  func GetItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item :=FindItem(params["id"])
	json.NewEncoder(w).Encode(item)
  }

  func FindItem(id string) Item{
	var i Item
	for _, item := range items {
		if item.ID == id {
		  return item
		}
	  }
	return i
  }

  func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(item)
	item.ID = strconv.Itoa(rand.Intn(1000000))
	items = append(items, item)
	json.NewEncoder(w).Encode(&item)
  }


  func AddItemWithName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(item)
	item.ID = strconv.Itoa(rand.Intn(1000000))
	item.Name = params["name"]
	items = append(items, item)
	json.NewEncoder(w).Encode(&item)
  }

  func DeleteItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
	  if item.ID == params["id"] {
		items = append(items[:index], items[index+1:]...)
		break
	  }
	}
	json.NewEncoder(w).Encode(items)
  }