/////////////////////////////////////////////////////////////////////////////////////
////////////////////// SAMPLE API TO CREATE AND FIND ORDERS /////////////////////////
////// IT ALSO DEMONSTRATES HOW TO USE NESTED STRUCTURES INSIDE GO STRUCTURE ////////
/////////////////////////////////////////////////////////////////////////////////////



Disclaimer:
Data shown in the REST requests and responses are not real data and should not be used.


Prerequisites:
Golang should be installed.
MySql should be installed and running.


How To Run The API:
c:\go_program\github>go run  gorm_order_api.go
2020/04/21 19:01:15 Connection Established
2020/04/21 19:01:17 Starting development server at http://127.0.0.1:10000/
2020/04/21 19:01:17 Quit the server with CONTROL-C.

After gorm_order_api.go script is run as shown above, the MySql tables also get created.

Sample Request To Create A New Order:
POST    http://127.0.0.1:10000/orders/new

{
	"firstname":"Nitin",
	"lastname":"Patil",
	"username":"Nitin.Patil",
	"password":"Nitin",
	"gender":"Male",
	"age": 24,

	"order_items" : [
		{
			"name":"computer",
			"description":"good computer",
			"price":399
		},
			{
			"name":"wireless headphones",
			"description":"good wireless headphones",
			"price":49.99
		},
			{
			"name":"soap",
			"description":"good soap",
			"price":1.99
		},
			{
			"name":"wireless mouse",
			"description":"good wireless mous",
			"price": 12.99
		}
	],

	"credit_card_details": {
		
		"name":"Nitin Patil",
		"type":"VISA",
		"number":"45564564556XXXXXX",
		"expiration_month":"1",
		"expiration_year":"2024",
		"cvv":"1044"
	},
	"total_price":100,
	"payment_status":"paid",




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
	"status":"Active"
}


------------------------------------------------------------------------------------------------------------

Sample JSON Response To Get One Order Using ID (Using PostMan):
GET  http://127.0.0.1:10000/orders/4



{
	  "ID": 4,
	  "CreatedAt": "2020-04-21T06:10:23Z",
	  "UpdatedAt": "2020-04-21T06:47:39Z",
	  "DeletedAt": null,
	  "firstname": "Rani",
	  "lastname": "Patil",
	  "username": "Rani.Patil",
	  "password": "Rani",
	  "gender": "Male",
	  "age": 24,
	  "order_items": [
		{
		  "ID": 6,
		  "CreatedAt": "2020-04-21T06:10:23Z",
		  "UpdatedAt": "2020-04-21T06:10:23Z",
		  "DeletedAt": null,
		  "OrderID": 4,
		  "name": "computer",
		  "description": "good computer",
		  "price": 399
		},
		{
		  "ID": 7,
		  "CreatedAt": "2020-04-21T06:10:23Z",
		  "UpdatedAt": "2020-04-21T06:10:23Z",
		  "DeletedAt": null,
		  "OrderID": 4,
		  "name": "wireless headphones",
		  "description": "good wireless headphones",
		  "price": 49.99
		},
		{
		  "ID": 8,
		  "CreatedAt": "2020-04-21T06:10:23Z",
		  "UpdatedAt": "2020-04-21T06:10:23Z",
		  "DeletedAt": null,
		  "OrderID": 4,
		  "name": "soap",
		  "description": "good soap",
		  "price": 1.99
		},
		{
		  "ID": 9,
		  "CreatedAt": "2020-04-21T06:10:23Z",
		  "UpdatedAt": "2020-04-21T06:10:23Z",
		  "DeletedAt": null,
		  "OrderID": 4,
		  "name": "wireless mouse",
		  "description": "good wireless mous",
		  "price": 12.99
		}
	  ],
	  "credit_card_details": {
		"ID": 4,
		"CreatedAt": "2020-04-21T06:10:23Z",
		"UpdatedAt": "2020-04-21T06:10:23Z",
		"DeletedAt": null,
		"OrderID": 4,
		"name": "Nitin Patil",
		"type": "VISA",
		"number": "45564564556XXXXXX",
		"expiration_month": "1",
		"expiration_year": "2024",
		"cvv": "1044"
	  },
	  "total_price": 463.96997,
	  "payment_status": "paid",
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
GET  http://127.0.0.1:10000/orders/all
 
[
  {
    "ID": 1,
    "CreatedAt": "2020-04-21T05:41:33Z",
    "UpdatedAt": "2020-04-21T05:41:33Z",
    "DeletedAt": null,
    "firstname": "Nitin",
    "lastname": "Patil",
    "username": "Nitin.Patil",
    "password": "Nitin",
    "gender": "Male",
    "age": 24,
    "order_items": [
      {
        "ID": 1,
        "CreatedAt": "2020-04-21T05:41:33Z",
        "UpdatedAt": "2020-04-21T05:41:33Z",
        "DeletedAt": null,
        "OrderID": 1,
        "name": "computer",
        "description": "good computer",
        "price": 399
      }
    ],
    "credit_card_details": {
      "ID": 1,
      "CreatedAt": "2020-04-21T05:41:33Z",
      "UpdatedAt": "2020-04-21T05:41:33Z",
      "DeletedAt": null,
      "OrderID": 1,
      "name": "Nitin Patil",
      "type": "VISA",
      "number": "45564564556XXXXXX",
      "expiration_month": "1",
      "expiration_year": "2024",
      "cvv": "1044"
    },
    "total_price": 100,
    "payment_status": "paid",
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
    "CreatedAt": "2020-04-21T05:43:09Z",
    "UpdatedAt": "2020-04-21T05:43:09Z",
    "DeletedAt": null,
    "firstname": "Nitin",
    "lastname": "Patil",
    "username": "Nitin.Patil",
    "password": "Nitin",
    "gender": "Male",
    "age": 24,
    "order_items": [
      {
        "ID": 2,
        "CreatedAt": "2020-04-21T05:43:09Z",
        "UpdatedAt": "2020-04-21T05:43:09Z",
        "DeletedAt": null,
        "OrderID": 2,
        "name": "computer",
        "description": "good computer",
        "price": 399
      },
      {
        "ID": 3,
        "CreatedAt": "2020-04-21T05:43:09Z",
        "UpdatedAt": "2020-04-21T05:43:09Z",
        "DeletedAt": null,
        "OrderID": 2,
        "name": "wireless headphones",
        "description": "good wireless headphones",
        "price": 49.99
      }
    ],
    "credit_card_details": {
      "ID": 2,
      "CreatedAt": "2020-04-21T05:43:09Z",
      "UpdatedAt": "2020-04-21T05:43:09Z",
      "DeletedAt": null,
      "OrderID": 2,
      "name": "Nitin Patil",
      "type": "VISA",
      "number": "45564564556XXXXXX",
      "expiration_month": "1",
      "expiration_year": "2024",
      "cvv": "1044"
    },
    "total_price": 100,
    "payment_status": "paid",
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
    "CreatedAt": "2020-04-21T06:03:32Z",
    "UpdatedAt": "2020-04-21T06:03:32Z",
    "DeletedAt": null,
    "firstname": "Nitin",
    "lastname": "Patil",
    "username": "Nitin.Patil",
    "password": "Nitin",
    "gender": "Male",
    "age": 24,
    "order_items": [
      {
        "ID": 4,
        "CreatedAt": "2020-04-21T06:03:32Z",
        "UpdatedAt": "2020-04-21T06:03:32Z",
        "DeletedAt": null,
        "OrderID": 3,
        "name": "computer",
        "description": "good computer",
        "price": 399
      },
      {
        "ID": 5,
        "CreatedAt": "2020-04-21T06:03:32Z",
        "UpdatedAt": "2020-04-21T06:03:32Z",
        "DeletedAt": null,
        "OrderID": 3,
        "name": "wireless headphones",
        "description": "good wireless headphones",
        "price": 49.99
      }
    ],
    "credit_card_details": {
      "ID": 3,
      "CreatedAt": "2020-04-21T06:03:32Z",
      "UpdatedAt": "2020-04-21T06:03:32Z",
      "DeletedAt": null,
      "OrderID": 3,
      "name": "Nitin Patil",
      "type": "VISA",
      "number": "45564564556XXXXXX",
      "expiration_month": "1",
      "expiration_year": "2024",
      "cvv": "1044"
    },
    "total_price": 100,
    "payment_status": "paid",
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
    "ID": 4,
    "CreatedAt": "2020-04-21T06:10:23Z",
    "UpdatedAt": "2020-04-21T06:47:39Z",
    "DeletedAt": null,
    "firstname": "Rani",
    "lastname": "Patil",
    "username": "Rani.Patil",
    "password": "Rani",
    "gender": "Male",
    "age": 24,
    "order_items": [
      {
        "ID": 6,
        "CreatedAt": "2020-04-21T06:10:23Z",
        "UpdatedAt": "2020-04-21T06:10:23Z",
        "DeletedAt": null,
        "OrderID": 4,
        "name": "computer",
        "description": "good computer",
        "price": 399
      },
      {
        "ID": 7,
        "CreatedAt": "2020-04-21T06:10:23Z",
        "UpdatedAt": "2020-04-21T06:10:23Z",
        "DeletedAt": null,
        "OrderID": 4,
        "name": "wireless headphones",
        "description": "good wireless headphones",
        "price": 49.99
      },
      {
        "ID": 8,
        "CreatedAt": "2020-04-21T06:10:23Z",
        "UpdatedAt": "2020-04-21T06:10:23Z",
        "DeletedAt": null,
        "OrderID": 4,
        "name": "soap",
        "description": "good soap",
        "price": 1.99
      },
      {
        "ID": 9,
        "CreatedAt": "2020-04-21T06:10:23Z",
        "UpdatedAt": "2020-04-21T06:10:23Z",
        "DeletedAt": null,
        "OrderID": 4,
        "name": "wireless mouse",
        "description": "good wireless mous",
        "price": 12.99
      }
    ],
    "credit_card_details": {
      "ID": 4,
      "CreatedAt": "2020-04-21T06:10:23Z",
      "UpdatedAt": "2020-04-21T06:10:23Z",
      "DeletedAt": null,
      "OrderID": 4,
      "name": "Nitin Patil",
      "type": "VISA",
      "number": "45564564556XXXXXX",
      "expiration_month": "1",
      "expiration_year": "2024",
      "cvv": "1044"
    },
    "total_price": 463.96997,
    "payment_status": "paid",
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

Updating A Single Order

PUT    http://127.0.0.1:10000/orders/update/4

{
  
  "firstname": "Sameer",
  "lastname": "Patil",
  "username": "Sameer.Patil",
  "password": "Sameer",
  "OrderItems": [
    {
      "ID": 6,
      "CreatedAt": "2020-04-21T06:10:23Z",
      "UpdatedAt": "2020-04-21T06:10:23Z",
      "DeletedAt": null,
      "OrderID": 4,
      "name": "computer",
      "description": "good computer",
      "price": 599
    }
    ]
}