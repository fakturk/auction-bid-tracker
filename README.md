# auction-bid-tracker
Auction Bid Tracker golang implementation

You have been asked with building part of a simple online auction system which will allow users to concurrently bid on items for sale. The system needs to be built in Go and/or Python.

Please, provide a bid-tracker interface and concrete implementation with the following functionality:

- [x] record a user’s bid on an item;
- [x] get the current winning bid for an item;
- [x] get all the bids for an item;
- [x] get all the items on which a user has bid;
- [x] build simple REST API to manage bids.

You are not required to implement a GUI (or CLI) or persistent store (events are for reporting only). You may use any appropriate libraries to help.

# Structure
We have three different structs for holding each value which are user, item and bid. When we are adding a bid to an item from a user we use user id and item id together.

| User  | Item  | Bid     |
| ----- | ----- | ------- |
| id    | id    | user id |
| name  | name  | user id |
|       |       | amount  |

# Postman usage

There is a postman collection added for 18 requests and "auction-bid-tracker.postman_collection.json" file can be implemented in the Postman and directly called from there.

Also  some mock data added inside of the program for test usage.

![Postman ](/images/auction-bid-tracker-postman.png)
 
# How To Run

go run main.go

# Usage
(Part I - Solutions to given problems)
- record a user’s bid on an item :

    - request type: POST

    - host: localhost:8000/bids/{userid}/{itemid}/{amount}

- get the current winning bid for an item:

    - request type: GET

    - host: localhost:8000/winner/{itemid}

- get all the bids for an item:

    - request type: GET

    - host: localhost:8000/bids/{itemid}

- get all the items on which a user has bid

    - request type: GET

    - host: localhost:8000/items/user/{userid}

(Part II - REST API Management)

- add user

    - request type: POST

    - host: localhost:8000/users

- add user with name

    - request type: POST

    - host: localhost:8000/users/{name}

- get users

    - request type: GET

    - host: localhost:8000/users

- get user by id

    - request type: GET

    - host: localhost:8000/users/id/{idnumber}

- delete user by id

    - request type: DELETE

    - host: localhost:8000/users/id/{idnumber}

- add item

    - request type: POST

    - host: localhost:8000/items

- add item with name

    - request type: POST

    - host: localhost:8000/items/{name}

- get items

    - request type: GET

    - host: localhost:8000/items

- get item by id

    - request type: GET

    - host: localhost:8000/items/id/{idnumber}

- delete item by id

    - request type: DELETE

    - host: localhost:8000/items/id/{idnumber}

- add bid

    - request type: POST

    - host: localhost:8000/bids/{userid}/{itemid}/{amount}

- update bid

    - request type: PUT

    - host: localhost:8000/bids/{userid}/{itemid}/{amount}


- get bids

    - request type: GET

    - host: localhost:8000/bids

- get bid (by user id and item id)

    - request type: GET

    - host: localhost:8000/bids/{userid}/{itemid}

- delete bid 

    - request type: DELETE

    - host: localhost:8000/bids/{userid}/{itemid}