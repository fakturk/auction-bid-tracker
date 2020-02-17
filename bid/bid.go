package bid

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
	"encoding/json"
	// "math/rand"
  	// "strconv"
)

type Bid struct {
	UserID string `json:"userid"`
	ItemID string `json:"itemid"`
	Amount string `json:"amount"`
  }

var bids []Bid

func GetBids(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"GetBids\n")
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(bids)
  }

  func GetBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, bid := range bids {
	  if bid.UserID == params["userid"] && bid.ItemID == params["itemid"] {
		json.NewEncoder(w).Encode(bid)
		return
	  }
	}
	json.NewEncoder(w).Encode(&Bid{})
  }


  func AddBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	var bid Bid
	_ = json.NewDecoder(r.Body).Decode(bid)
	// bid.ID = strconv.Itoa(rand.Intn(1000000))
	bid.UserID = params["userid"]
	bid.ItemID = params["itemid"]
	bid.Amount = params["amount"]
	bids = append(bids, bid)
	json.NewEncoder(w).Encode(&bid)
  }

  func UpdateBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, bid := range bids {
	  if bid.UserID == params["userid"] && bid.ItemID == params["itemid"] {
		bids = append(bids[:index], bids[index+1:]...)
		var bid Bid
		_ = json.NewDecoder(r.Body).Decode(&bid)
		// bid.ID = params["id"]
		bid.UserID = params["userid"]
		bid.ItemID = params["itemid"]
		bid.Amount = params["amount"]
		bids = append(bids, bid)
		json.NewEncoder(w).Encode(&bid)
		return
	  }
	}
	json.NewEncoder(w).Encode(bids)
  }

  func DeleteBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, bid := range bids {
	  if bid.UserID == params["userid"] && bid.ItemID == params["itemid"] {
		bids = append(bids[:index], bids[index+1:]...)
		break
	  }
	}
	json.NewEncoder(w).Encode(bids)
  }