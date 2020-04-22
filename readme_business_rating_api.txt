/////////////////////////////////////////////////////////////////////////////////////
////////////////////// SAMPLE API TO CREATE AND FIND BUSINESS ///////////////////////
////// IT ALSO DEMONSTRATES HOW TO USE NESTED STRUCTURES INSIDE GO STRUCTURE ////////
/////////////////////////////////////////////////////////////////////////////////////



Disclaimer:
Data shown in the REST requests and responses are not real data and should not be used.

Prerequisites:
Golang should be installed.
MySql should be installed and running.


How To Run The API:
c:\go_program\github>go run business_rating_api.go
2020/04/21 16:37:14 Connection Established
2020/04/21 16:37:16 Starting development server at http://127.0.0.1:10000/
2020/04/21 16:37:16 Quit the server with CONTROL-C.

After business_rating_api.go script is run as shown above, the MySql tables also get created.

Sample Request To Create A New Business:
POST   http://127.0.0.1:10000/businesss/new

{

	"name":"Nitin Super Stores",
	"description":"Retail of Indian and Asian  Groceries.",
	"category":"grocery",
	"highlights1":"Retail of Indian and Asian  Groceries.",
	"highlights2":"",
	"highlights3":"",
	"highlights4":"",


	"history":"Started Retail of Indian and Asian  Groceries In 1988.",
	"special_message": "",
	"offers_delivery": "Y",
	"accepts_credit_card": "Y",
	"drive_through_available": "N",
	"curb_side_pickup_available": "N",
	"business_verified": "Y",

	"average_price": 0,

	"address1":"Gala No 8, Laxmi Plaza, Laxmi Industrial Estate",
	"address2":"New Link Road, Andheri West",
	"city":"Mumbai",
	"state":"Maharashtra",
	"zip":"400054",
	"country":"India",
	"primary_phone":"9252224444",
	"secondary_phone":"9822044888",
	"primary_email":"nitin.patil1@test.com",
	"secondary_email":"nitin.patil2@test.com",
	"status":"Active",

	"reviews_list" : [
		{
			"username":"nitin",
			"email":"nitin.patil@test.com",
			"review_text":"This is a good grocery store.",
			"rating": 5
		},
		{
			"username":"amit",
			"email":"amit.patil@test.com",
			"review_text":"This is a good grocery store. Get fresh vegetables.",
			"rating": 5
		}
	],

	"business_hours" : [
		{
			"weekday":"monday",
			"hours":"8:00 am - 6:00 pm"
		},
		{
			"weekday":"tuesday",
			"hours":"8:00 am - 6:00 pm"
		},
		{
			"weekday":"wednesday",
			"hours":"8:00 am - 6:00 pm"
		},
		{
			"weekday":"thursday",
			"hours":"8:00 am - 6:00 pm"
		},
		{
			"weekday":"friday",
			"hours":"8:00 am - 6:00 pm"
		},
		{
			"weekday":"saturday",
			"hours":"closed"
		},
		{
			"weekday":"sunday",
			"hours":"closed"
		}
	]

}

------------------------------------------------------------------------------------------------------------

Sample JSON Response To Get One Business Using ID (Using PostMan):
GET  http://127.0.0.1:10000/businesss/3



{
	  "ID": 3,
	  "CreatedAt": "2020-04-21T21:27:31Z",
	  "UpdatedAt": "2020-04-21T21:27:31Z",
	  "DeletedAt": null,
	  "name": "Nitin Super Stores",
	  "description": "Retail of Indian and Asian  Groceries.",
	  "category": "grocery",
	  "highlights1": "Retail of Indian and Asian  Groceries.",
	  "highlights2": "",
	  "highlights3": "",
	  "highlights4": "",
	  "history": "Started Retail of Indian and Asian  Groceries In 1988.",
	  "special_message": "",
	  "offers_delivery": "Y",
	  "accepts_credit_card": "Y",
	  "drive_through_available": "N",
	  "curb_side_pickup_available": "N",
	  "business_verified": "Y",
	  "reviews_list": [
		{
		  "ID": 3,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "username": "nitin",
		  "email": "nitin.patil@test.com",
		  "review_text": "This is a good grocery store.",
		  "rating": 5
		},
		{
		  "ID": 4,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "username": "amit",
		  "email": "amit.patil@test.com",
		  "review_text": "This is a good grocery store. Get fresh vegetables.",
		  "rating": 5
		}
	  ],
	  "business_hours": [
		{
		  "ID": 8,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "monday",
		  "hours": "8:00 am - 6:00 pm"
		},
		{
		  "ID": 9,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "tuesday",
		  "hours": "8:00 am - 6:00 pm"
		},
		{
		  "ID": 10,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "wednesday",
		  "hours": "8:00 am - 6:00 pm"
		},
		{
		  "ID": 11,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "thursday",
		  "hours": "8:00 am - 6:00 pm"
		},
		{
		  "ID": 12,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "friday",
		  "hours": "8:00 am - 6:00 pm"
		},
		{
		  "ID": 13,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "saturday",
		  "hours": "closed"
		},
		{
		  "ID": 14,
		  "CreatedAt": "2020-04-21T21:27:31Z",
		  "UpdatedAt": "2020-04-21T21:27:31Z",
		  "DeletedAt": null,
		  "BusinessID": 3,
		  "weekday": "sunday",
		  "hours": "closed"
		}
	  ],
	  "rating": 5,
	  "average_price": 0,
	  "address1": "Gala No 8, Laxmi Plaza, Laxmi Industrial Estate",
	  "address2": "New Link Road, Andheri West",
	  "city": "Mumbai",
	  "state": "Maharashtra",
	  "zip": "400054",
	  "country": "India",
	  "primary_phone": "9252224444",
	  "secondary_phone": "9822044888",
	  "primary_email": "nitin.patil1@test.com",
	  "secondary_email": "nitin.patil2@test.com",
	  "status": "Active"
}

------------------------------------------------------------------------------------------------------------


Sample JSON Response (Using PostMan):
GET  http://127.0.0.1:10000/businesss/all
 
[
  {
    "ID": 1,
    "CreatedAt": "2020-04-21T21:14:28Z",
    "UpdatedAt": "2020-04-21T21:14:28Z",
    "DeletedAt": null,
    "name": "Nitin Super Stores",
    "description": "Retail of Indian and Asian  Groceries.",
    "category": "grocery",
    "highlights1": "Retail of Indian and Asian  Groceries.",
    "highlights2": "",
    "highlights3": "",
    "highlights4": "",
    "history": "Started Retail of Indian and Asian  Groceries In 1988.",
    "special_message": "",
    "offers_delivery": "Y",
    "accepts_credit_card": "Y",
    "drive_through_available": "N",
    "curb_side_pickup_available": "N",
    "business_verified": "Y",
    "reviews_list": [],
    "business_hours": [],
    "rating": 0,
    "average_price": 0,
    "address1": "Gala No 8, Laxmi Plaza, Laxmi Industrial Estate",
    "address2": "New Link Road, Andheri West",
    "city": "Mumbai",
    "state": "Maharashtra",
    "zip": "400054",
    "country": "India",
    "primary_phone": "9252224444",
    "secondary_phone": "9822044888",
    "primary_email": "nitin.patil1@test.com",
    "secondary_email": "nitin.patil2@test.com",
    "status": "Active"
  },
  {
    "ID": 2,
    "CreatedAt": "2020-04-21T21:18:58Z",
    "UpdatedAt": "2020-04-21T21:18:58Z",
    "DeletedAt": null,
    "name": "Nitin Super Stores",
    "description": "Retail of Indian and Asian  Groceries.",
    "category": "grocery",
    "highlights1": "Retail of Indian and Asian  Groceries.",
    "highlights2": "",
    "highlights3": "",
    "highlights4": "",
    "history": "Started Retail of Indian and Asian  Groceries In 1988.",
    "special_message": "",
    "offers_delivery": "Y",
    "accepts_credit_card": "Y",
    "drive_through_available": "N",
    "curb_side_pickup_available": "N",
    "business_verified": "Y",
    "reviews_list": [
      {
        "ID": 1,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "username": "nitin",
        "email": "nitin.patil@test.com",
        "review_text": "This is a good grocery store.",
        "rating": 5
      },
      {
        "ID": 2,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "username": "amit",
        "email": "amit.patil@test.com",
        "review_text": "This is a good grocery store. Get fresh vegetables.",
        "rating": 5
      }
    ],
    "business_hours": [
      {
        "ID": 1,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "monday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 2,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "tuesday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 3,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "wednesday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 4,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "thursday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 5,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "friday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 6,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "saturday",
        "hours": "closed"
      },
      {
        "ID": 7,
        "CreatedAt": "2020-04-21T21:18:58Z",
        "UpdatedAt": "2020-04-21T21:18:58Z",
        "DeletedAt": null,
        "BusinessID": 2,
        "weekday": "sunday",
        "hours": "closed"
      }
    ],
    "rating": 0,
    "average_price": 0,
    "address1": "Gala No 8, Laxmi Plaza, Laxmi Industrial Estate",
    "address2": "New Link Road, Andheri West",
    "city": "Mumbai",
    "state": "Maharashtra",
    "zip": "400054",
    "country": "India",
    "primary_phone": "9252224444",
    "secondary_phone": "9822044888",
    "primary_email": "nitin.patil1@test.com",
    "secondary_email": "nitin.patil2@test.com",
    "status": "Active"
  },
  {
    "ID": 3,
    "CreatedAt": "2020-04-21T21:27:31Z",
    "UpdatedAt": "2020-04-21T21:27:31Z",
    "DeletedAt": null,
    "name": "Nitin Super Stores",
    "description": "Retail of Indian and Asian  Groceries.",
    "category": "grocery",
    "highlights1": "Retail of Indian and Asian  Groceries.",
    "highlights2": "",
    "highlights3": "",
    "highlights4": "",
    "history": "Started Retail of Indian and Asian  Groceries In 1988.",
    "special_message": "",
    "offers_delivery": "Y",
    "accepts_credit_card": "Y",
    "drive_through_available": "N",
    "curb_side_pickup_available": "N",
    "business_verified": "Y",
    "reviews_list": [
      {
        "ID": 3,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "username": "nitin",
        "email": "nitin.patil@test.com",
        "review_text": "This is a good grocery store.",
        "rating": 5
      },
      {
        "ID": 4,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "username": "amit",
        "email": "amit.patil@test.com",
        "review_text": "This is a good grocery store. Get fresh vegetables.",
        "rating": 5
      }
    ],
    "business_hours": [
      {
        "ID": 8,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "monday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 9,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "tuesday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 10,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "wednesday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 11,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "thursday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 12,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "friday",
        "hours": "8:00 am - 6:00 pm"
      },
      {
        "ID": 13,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "saturday",
        "hours": "closed"
      },
      {
        "ID": 14,
        "CreatedAt": "2020-04-21T21:27:31Z",
        "UpdatedAt": "2020-04-21T21:27:31Z",
        "DeletedAt": null,
        "BusinessID": 3,
        "weekday": "sunday",
        "hours": "closed"
      }
    ],
    "rating": 5,
    "average_price": 0,
    "address1": "Gala No 8, Laxmi Plaza, Laxmi Industrial Estate",
    "address2": "New Link Road, Andheri West",
    "city": "Mumbai",
    "state": "Maharashtra",
    "zip": "400054",
    "country": "India",
    "primary_phone": "9252224444",
    "secondary_phone": "9822044888",
    "primary_email": "nitin.patil1@test.com",
    "secondary_email": "nitin.patil2@test.com",
    "status": "Active"
  }
]








-----------------------------------------------------------------------------------------------------

Updating A Single Business

PUT   http://127.0.0.1:10000/businesss/update/3

{
	"name":"Ganesh Super Stores",
	"highlights1":"Retail of Indian and Asian  Groceries.",
	"highlights2":"Fresh Vegetables",
	"highlights3":"Different Food Choices",
	"highlights4":"Fresh Indian Food"
}