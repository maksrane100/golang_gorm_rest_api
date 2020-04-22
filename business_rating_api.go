/////////////////////////////////////////////////////////////////////////////////////
////////////////////// SAMPLE API TO CREATE AND FIND BUSINESS ///////////////////////
////// IT ALSO DEMONSTRATES HOW TO USE NESTED STRUCTURES INSIDE GO STRUCTURE ////////
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

type Business struct {
		gorm.Model
		//Id      int    `json:"id"`
		Name    					string `json:"name"`
		Description    				string `json:"description"`
		Category    				string `json:"category"`
		Highlights1    				string `json:"highlights1"`
		Highlights2    				string `json:"highlights2"`
		Highlights3    				string `json:"highlights3"`
		Highlights4    				string `json:"highlights4"`		
		History    					string `json:"history"`
		SpecialMessage    			string `json:"special_message"`
		OffersDelivery    			string `json:"offers_delivery"`
		AcceptsCreditCard    		string `json:"accepts_credit_card"`
		
		DriveThroughAvailable    	string `json:"drive_through_available"`
		CurbSidePickupAvailable    	string `json:"curb_side_pickup_available"`
		BusinessVerified    		string `json:"business_verified"`


		Reviews 					[]Reviews `json:"reviews_list"`
		BusinessHours 				[]BusinessHours `json:"business_hours"`
		Rating      				float32   `json:"rating"`
		AveragePrice      			float32   `json:"average_price"`
		
		Address1    				string `json:"address1"`
		Address2    				string `json:"address2"`
		City    					string `json:"city"`
		State    					string `json:"state"`
		Zip    						string `json:"zip"`
		Country    					string `json:"country"`
		

		PrimaryPhone    			string `json:"primary_phone"` 
		SecondaryPhone    			string `json:"secondary_phone"` 
		PrimaryEmail    			string `json:"primary_email"` 
		SecondaryEmail    			string `json:"secondary_email"` 
  
		Status    					string `json:"status"` 
}

type Reviews struct {
	gorm.Model
	
	BusinessID  uint
	Username    string `json:"username"`
	Email    	string `json:"email"`
	ReviewText  string `json:"review_text"`
	Rating      int   `json:"rating"`
}

type BusinessHours struct {
	gorm.Model
	
	BusinessID  uint
	Weekday    	string `json:"weekday"`
	Hours    	string `json:"hours"`
}




func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to business lookup api!")
    fmt.Println("Endpoint Hit: Business Lookup Api")
}


func createNewBusiness(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    
	fmt.Println("in createNewBusiness()")
	
	reqBody, _ := ioutil.ReadAll(r.Body)
    var business Business
    json.Unmarshal(reqBody, &business)
	
	var total int = 0
	var total_businesses int = 0
    for _, elem := range business.Reviews {
        total += elem.Rating
		total_businesses += 1
    }
	fmt.Println(total);
	business.Rating=float32(total/total_businesses)
    db.Create(&business)
    fmt.Println("Endpoint Hit: Creating New Business")
    json.NewEncoder(w).Encode(business)
}



func returnAllBusinesss(w http.ResponseWriter, r *http.Request){
	businesss := []Business{}
	//businesssArray := []Business{}
	
	db.Find(&businesss)
	businesssArray := make([]Business, len(businesss))
	i:=0
	for _, elem := range businesss {
        reviews := []Reviews{}
		db.Where("business_id = ?", elem.ID).Find(&reviews)
		businessHours := []BusinessHours{}
		db.Where("business_id = ?", elem.ID).Find(&businessHours)
		fmt.Println(elem)
		fmt.Println(elem.ID)
		fmt.Println(elem.Name)
		fmt.Println(reviews)
		fmt.Println(businessHours)
		elem.Reviews=reviews
		elem.BusinessHours=businessHours
		businesssArray[i] = elem
		i++
    }
	fmt.Println("Endpoint Hit: returnAllBusinesss")
	json.NewEncoder(w).Encode(businesssArray)
}


func returnSingleBusiness(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	b := Business{}
	db.First(&b, key)
	reviews := []Reviews{}
	db.Where("business_id = ?", key).Find(&reviews)
	businessHours := []BusinessHours{}
	db.Where("business_id = ?", key).Find(&businessHours)
	b.Reviews=reviews
	b.BusinessHours=businessHours
	json.NewEncoder(w).Encode(b)
 
}


func updateSingleBusinessNew(w http.ResponseWriter, r *http.Request){
		fmt.Println("In updateSingleBusinessNew()")
		vars := mux.Vars(r)
		key := vars["id"]
		//businesss := []Business{}
		var inputbusiness Business
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &inputbusiness)

		b := Business{}
		db.First(&b, key)
		fmt.Println("Found business")
		
		if inputbusiness.Name != "" {
			b.Name=inputbusiness.Name
		}
		if inputbusiness.Description != "" {
			b.Description=inputbusiness.Description
		}
		if inputbusiness.Category != "" {
			b.Category=inputbusiness.Category
		}
		if inputbusiness.Highlights1 != "" {
			b.Highlights1=inputbusiness.Highlights1
		}
		if inputbusiness.Highlights2 != "" {
			b.Highlights2=inputbusiness.Highlights2
		}
		if inputbusiness.Highlights3 != "" {
			b.Highlights3=inputbusiness.Highlights3
		}
		if inputbusiness.Highlights4 != "" {
			b.Highlights4=inputbusiness.Highlights4
		}
		if inputbusiness.History != "" {
			b.History=inputbusiness.History
		}
		if inputbusiness.SpecialMessage != "" {
			b.SpecialMessage=inputbusiness.SpecialMessage
		}
		if inputbusiness.OffersDelivery != "" {
			b.OffersDelivery=inputbusiness.OffersDelivery
		}
		if inputbusiness.AcceptsCreditCard != "" {
			b.AcceptsCreditCard=inputbusiness.AcceptsCreditCard
		}		
		if inputbusiness.DriveThroughAvailable != "" {
			b.DriveThroughAvailable=inputbusiness.DriveThroughAvailable
		}	
		if inputbusiness.CurbSidePickupAvailable != "" {
			b.CurbSidePickupAvailable=inputbusiness.CurbSidePickupAvailable
		}	
		if inputbusiness.BusinessVerified != "" {
			b.BusinessVerified=inputbusiness.BusinessVerified
		}	
		if inputbusiness.AveragePrice > 0 {
			b.AveragePrice=inputbusiness.AveragePrice
		}			
		if inputbusiness.Address1 != "" {
			b.Address1=inputbusiness.Address1
		}
		if inputbusiness.Address2 != "" {
			b.Address2=inputbusiness.Address2
		}
		if inputbusiness.City != "" {
			b.City=inputbusiness.City
		}
		if inputbusiness.State != "" {
			b.State=inputbusiness.State
		}
		if inputbusiness.Zip != "" {
			b.Zip=inputbusiness.Zip
		}
		if inputbusiness.Country != "" {
			b.Country=inputbusiness.Country
		}
		if inputbusiness.PrimaryPhone != "" {
			b.PrimaryPhone=inputbusiness.PrimaryPhone
		}
		if inputbusiness.SecondaryPhone != "" {
			b.SecondaryPhone=inputbusiness.SecondaryPhone
		}
		if inputbusiness.PrimaryEmail != "" {
			b.PrimaryEmail=inputbusiness.PrimaryEmail
		}
		if inputbusiness.SecondaryEmail != "" {
			b.SecondaryEmail=inputbusiness.SecondaryEmail
		}
		if inputbusiness.Status != "" {
			b.Status=inputbusiness.Status
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
	myRouter.HandleFunc("/businesss/new", createNewBusiness).Methods("POST")
	myRouter.HandleFunc("/businesss/all", returnAllBusinesss)
	myRouter.HandleFunc("/businesss/{id}", returnSingleBusiness)
	myRouter.HandleFunc("/businesss/update/{id}", updateSingleBusinessNew).Methods("PUT")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    // Define your user name and password for my sql.
    db, err = gorm.Open("mysql", "root:@(127.0.0.1:3306)/business_management?parseTime=true")

    if err!=nil{
		log.Println("Connection Failed to Open")
    }else{
		log.Println("Connection Established")
    }

    db.AutoMigrate(&Business{}, &Reviews{}, &BusinessHours{})
    handleRequests()
}
