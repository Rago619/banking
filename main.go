package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {

	//defining all the routes to the http server

	http.HandleFunc("/greet", greet) //two parameters : first is patter/route/endpoint and second is handler function
	http.HandleFunc("/customers", getAllCustomers)

	//starting the http server
	err := http.ListenAndServe("localhost:9080", nil) //starts the server on port 9080 and second parameter is handler function. here default mux is used. so nil.
	if err != nil {
		println(err.Error())
	}
}

func greet(w http.ResponseWriter, r *http.Request) { //first parameters sends the response back to the client and second is request to ther server
	fmt.Fprint(w, "Hello World") //sends the response back to the client
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"ABC", "ND", "101001"},
		{"DEF", "DN", "102001"},
	}
	json.NewEncoder(w).Encode(customers)
}
