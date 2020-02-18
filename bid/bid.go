package bid

import(
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
	"encoding/json"
	// "math/rand"
	"strconv"
	"github.com/fakturk/auction-bid-tracker/item"
	"github.com/fakturk/auction-bid-tracker/user"
)

type Bid struct {
	UserID string `json:"userid"`
	ItemID string `json:"itemid"`
	Amount string `json:"amount"`
  }

var bids []Bid


func AddMockBids(){
	//adding mock data for test usage 
	bids = append(bids, Bid{UserID: "498081",ItemID: "727888", Amount: "34"})
	bids = append(bids, Bid{UserID: "498081",ItemID: "727887", Amount: "35"})
	bids = append(bids, Bid{UserID: "498082",ItemID: "727887", Amount: "54"})
	bids = append(bids, Bid{UserID: "498082",ItemID: "727888", Amount: "72"})

}

func GetBids(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"GetBids\n")
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(bids)
  }

  func GetBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bid:=FindBid(params["userid"] ,params["itemid"] )
	json.NewEncoder(w).Encode(bid)
  }

  func FindBid(userid,itemid string) Bid{
	var b Bid
	for _, bid := range bids {
		if bid.UserID == userid && bid.ItemID ==itemid {
			return bid
		}
	  }
	return b
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

	if item.FindItem(bid.ItemID)!=(item.Item{}) && user.FindUser(bid.UserID)!=(user.User{}) {
		if FindBid(params["userid"] ,params["itemid"] )!=(Bid{}) {
			BidUpdate(params["userid"],params["itemid"],params["amount"])
		} else{
			bids = append(bids, bid)
		}
		
		json.NewEncoder(w).Encode(&bid)
	} else {
		fmt.Println("inside else")
		if item.FindItem(bid.ItemID)==(item.Item{}) {
			fmt.Println("inside item not found")
			bid.ItemID  = "Item Not Found"
		} 

		if user.FindUser(bid.UserID)==(user.User{}) {
			fmt.Println("inside item not found")
			bid.UserID  = "User Not Found"
		}
		fmt.Println(bid)
		json.NewEncoder(w).Encode(&bid)

	}	
	
  }

  func UpdateBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	BidUpdate(params["userid"],params["itemid"],params["amount"])
	
	json.NewEncoder(w).Encode(bids)
  }

  func BidUpdate(userid,itemid,amount string){
	for index, bid := range bids {
		if bid.UserID == userid && bid.ItemID == itemid {
		  bids = append(bids[:index], bids[index+1:]...)
		  var bid Bid
		  bid.UserID = userid
		  bid.ItemID = itemid
		  bid.Amount = amount
		  bids = append(bids, bid)
		  return
		}
	}
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
	  if bid.ItemID == params["itemid"] {
		bidAmount,_:=strconv.Atoi(bid.Amount)
		winnerAmount,_:=strconv.Atoi(winner.Amount)
		if  bidAmount> winnerAmount {
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

	var allBids []Bid
	for _, bid := range bids {
	  if bid.ItemID == params["itemid"] {

		allBids = append(allBids,bid)
		
	  }
	}

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