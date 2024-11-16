package main

import (
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"fname"`
	City    string `json:"city_name" xml:"cname"`
	Zipcode string `json:"pin_code" xml:"pin"`
}

func main() {
	//APPLICATION LOGIC : Code which helps in starting our application
	//defining all the routes to the http server

	http.HandleFunc("/greet", greet) //two parameters : first is patter/route/endpoint and second is handler function
	http.HandleFunc("/customers", getAllCustomers)

	//starting the http server
	err := http.ListenAndServe("localhost:9080", nil) //starts the server on port 9080 and second parameter is handler function. here default mux is used. so nil.
	if err != nil {
		println(err.Error())
	}
}
