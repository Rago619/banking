package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"fname"`
	City    string `json:"city_name" xml:"cname"`
	Zipcode string `json:"pin_code" xml:"pin"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

// func greet(w http.ResponseWriter, r *http.Request) { //first parameters sends the response back to the client and second is request to ther server
// 	fmt.Fprint(w, "Hello World") //sends the response back to the client
// }

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"ABC", "ND", "101001"},
	// 	{"DEF", "DN", "102001"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" { // user defines if info to be fetched in json format or xml format
		w.Header().Add("Content-Type", "application/xml") //change header parameters in postman to shift between json and xml
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	// w.Header().Add("Content-type", "applications/json") //writes to the header i.e. user will get to know the media type in response
	//if not specified, the media-type will be plain text
	// w.Header().Add("Content-type", "applications/xml")
	// xml.NewEncoder(w).Encode(customers) //encodes the data into json format (marhsaling)
}

// func to fetch customer details from the customer_id specified in the URL
// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r) // map is stored in the vars, something like below
// 	//map[string]string{
// 	//	"customer_id": "42",
// 	//	"order_id": "1001",
// 	//}
// 	fmt.Fprint(w, vars["customer_id"]) // customer_id value printed.
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Customer Post Created")
// }
