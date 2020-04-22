/////////////////////////////////////////////////////////////////////////////////////
////////////////////// SAMPLE API TO CREATE AND FIND ORDERS / ///////////////////////
/////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"net/http"
	//"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Order struct {
		gorm.Model
		//Id      int    `json:"id"`
		Firstname    string `json:"firstname"`
		Lastname    string `json:"lastname"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Gender    string `json:"gender"`
		Age      int    `json:"age"`

		OrderItems []OrderItem `json:"order_items"`

        CreditCardDetails CreditCard `json:"credit_card_details"`
		
		TotalPrice      float32   `json:"total_price"`
		PaymentStatus    string `json:"payment_status"`
		
		Address1    string `json:"address1"`
		Address2    string `json:"address2"`
		City    		string `json:"city"`
		State    string `json:"state"`
		Zip    string `json:"zip"`
		Country    string `json:"country"`
		

		PrimaryPhone    string `json:"primary_phone"` 
		SecondaryPhone    string `json:"secondary_phone"` 
		PrimaryEmail    string `json:"primary_email"` 
		SecondaryEmail    string `json:"secondary_email"` 
  
		Status    string `json:"status"` 
}

type OrderItem struct {
	gorm.Model
	//Id      int    `json:"id"`
	OrderID  uint
	Name    string `json:"name"`
	Description    string `json:"description"`
	Price      float32   `json:"price"`
}

type CreditCard struct {
	gorm.Model
	//Id      int    `json:"id"`
	OrderID  uint
	Name    string `json:"name"`
	Type    string `json:"type"`
	Number    string `json:"number"`
	ExpirationMonth    string `json:"expiration_month"`
	ExpirationYear    string `json:"expiration_year"`
	Cvv    string `json:"cvv"`
}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to order management api!")
    fmt.Println("Endpoint Hit: Order management api")
}


func createNewOrder(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    var order Order
    json.Unmarshal(reqBody, &order)
	
	var total float32 = 0
    for _, elem := range order.OrderItems {
        total += elem.Price
    }
	fmt.Println(total);
	order.TotalPrice=total
    db.Create(&order)
    fmt.Println("Endpoint Hit: Creating New Order")
    json.NewEncoder(w).Encode(order)
}

func returnAllOrders(w http.ResponseWriter, r *http.Request){
	orders := []Order{}
	db.Find(&orders)
	ordersArray := make([]Order, len(orders))
	i:=0
	for _, elem := range orders {
        orderItems := []OrderItem{}
		db.Where("order_id = ?", elem.ID).Find(&orderItems)
		creditCard := CreditCard{}
		db.Where("order_id = ?", elem.ID).Find(&creditCard)
		fmt.Println(elem)
		fmt.Println(orderItems)
		fmt.Println(creditCard)
		elem.OrderItems=orderItems
		elem.CreditCardDetails=creditCard
		ordersArray[i] = elem
		i++
    }
	fmt.Println("Endpoint Hit: returnAllOrders")
	json.NewEncoder(w).Encode(ordersArray)
}


func returnSingleOrder(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	b := Order{}
	db.First(&b, key)
	orderItems := []OrderItem{}
	db.Where("order_id = ?", key).Find(&orderItems)
	creditCard := CreditCard{}
	db.Where("id = ?", key).Find(&creditCard)
	b.OrderItems=orderItems
	b.CreditCardDetails=creditCard
	json.NewEncoder(w).Encode(b)
 
}


func updateSingleOrderNew(w http.ResponseWriter, r *http.Request){
		fmt.Println("In updateSingleOrderNew()")
		vars := mux.Vars(r)
		key := vars["id"]
		//orders := []Order{}
		var inputorder Order
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &inputorder)

		b := Order{}
		db.First(&b, key)
		fmt.Println("Found order")
		
		if inputorder.Firstname != "" {
			b.Firstname=inputorder.Firstname
		}
		if inputorder.Lastname != "" {
			b.Lastname=inputorder.Lastname
		}
		if inputorder.Username != "" {
			b.Username=inputorder.Username
		}
		if inputorder.Password != "" {
			b.Password=inputorder.Password
		}
		if inputorder.Gender != "" {
			b.Gender=inputorder.Gender
		}

		if inputorder.Age > 0 {
			b.Age=inputorder.Age
		}
		




		if inputorder.TotalPrice > 0 {
			b.TotalPrice=inputorder.TotalPrice
		}
		if inputorder.PaymentStatus != "" {
			b.PaymentStatus=inputorder.PaymentStatus
		}
		if inputorder.Address1 != "" {
			b.Address1=inputorder.Address1
		}
		if inputorder.Address2 != "" {
			b.Address2=inputorder.Address2
		}
		if inputorder.City != "" {
			b.City=inputorder.City
		}

		if inputorder.State != "" {
			b.State=inputorder.State
		}
		if inputorder.Zip != "" {
			b.Zip=inputorder.Zip
		}
		if inputorder.Country != "" {
			b.Country=inputorder.Country
		}
		if inputorder.PrimaryPhone != "" {
			b.PrimaryPhone=inputorder.PrimaryPhone
		}
		if inputorder.SecondaryPhone != "" {
			b.SecondaryPhone=inputorder.SecondaryPhone
		}

		if inputorder.PrimaryEmail != "" {
			b.PrimaryEmail=inputorder.PrimaryEmail
		}
		if inputorder.SecondaryEmail != "" {
			b.SecondaryEmail=inputorder.SecondaryEmail
		}
		if inputorder.Status != "" {
			b.Status=inputorder.Status
		}
		if inputorder.CreditCardDetails.Type != "" {
			b.CreditCardDetails.Type=inputorder.CreditCardDetails.Type
		}
		if inputorder.CreditCardDetails.Number != "" {
			b.CreditCardDetails.Number=inputorder.CreditCardDetails.Number
		}		
		
		db.Save(&b)
		json.NewEncoder(w).Encode(b)
}



func handleRequests(){
    log.Println("Starting development server at http://127.0.0.1:10000/")
    log.Println("Quit the server with CONTROL-C.")
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/orders/new", createNewOrder).Methods("POST")
	myRouter.HandleFunc("/orders/all", returnAllOrders)
	myRouter.HandleFunc("/orders/{id}", returnSingleOrder)
	myRouter.HandleFunc("/orders/update/{id}", updateSingleOrderNew).Methods("PUT")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    // Define your user name and password for mysql.
    db, err = gorm.Open("mysql", "root:@(127.0.0.1:3306)/order_management?parseTime=true")
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function

    if err!=nil{
		log.Println("Connection Failed to Open")
    }else{
		log.Println("Connection Established")
    }

    db.AutoMigrate(&Order{}, &OrderItem{}, &CreditCard{})
    handleRequests()
}
