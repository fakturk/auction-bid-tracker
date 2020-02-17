package bid

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
	"encoding/json"
	// "math/rand"
	"strconv"
	"github.com/fakturk/auction-bid-tracker/item"
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

  func WinnerBidByItemID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,"WinnerBids\n")
	fmt.Println("inside winner func")

	params := mux.Vars(r)
	fmt.Println(params)
	var winner Bid
	winner.Amount="0"
	fmt.Println(bids)
	for _, bid := range bids {
		fmt.Println(bid)
	  if bid.ItemID == params["itemid"] {
		  fmt.Println("inside check itemid: ",bid.ItemID)
		// json.NewEncoder(w).Encode(bid)
		bidAmount,_:=strconv.Atoi(bid.Amount)
		winnerAmount,_:=strconv.Atoi(winner.Amount)
		if  bidAmount> winnerAmount {
			fmt.Println("bid is bigger")
			winner=bid
		}
		
	  }
	}
	fmt.Println(winner)
	json.NewEncoder(w).Encode(winner)
  }

  func BidsByItemID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// fmt.Println(params)
	var allBids []Bid
	for _, bid := range bids {
	  if bid.ItemID == params["itemid"] {
		//   fmt.Println("inside check itemid: ",bid.ItemID)
		// json.NewEncoder(w).Encode(bid)
		allBids = append(allBids,bid)
		
	  }
	}
	// fmt.Println(winner)
	json.NewEncoder(w).Encode(allBids)
  }

  func ItemByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var allItems []item.Item
	for _, bid := range bids {
	  if bid.UserID == params["userid"] {
		item:=item.FindItem(bid.ItemID)
		allItems = append(allItems,item)
		
	  }
	}
	json.NewEncoder(w).Encode(allItems)
  }