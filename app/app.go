package app

import (
	// "fmt"
	"banking/domain"
	"banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() { //APPLICATION LOGIC : Code which helps in starting our application

	//mux := http.NewServeMux(); // initialises a http server. Default server not used now
	router := mux.NewRouter() // intialises a http router using the gorilla mux library

	// WIRING THE COMPLETE APPLICATION
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//defining all the routes to the http server
	// router.HandleFunc("/greet", greet)                                       //two parameters : first is patter/route/endpoint and second is handler function
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet) //Method Matcher Methods(http.MethodGet) default is GET.
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)        //URL Segmenting. Take only numeric values since [0-9] specified

	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost) //method matcher for POST request

	//starting the http server
	err := http.ListenAndServe("localhost:9080", router) //starts the server on port 9080 and second parameter is handler function. here default mux is used. so nil.
	if err != nil {                                      //if there is an error present, print the error
		println(err.Error())
	}
}
