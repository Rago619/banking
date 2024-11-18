package domain // BUSINESS LOGIC / DOMAIN

type Customer struct { // business object
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface { //secondary port / interface
	FindAll() ([]Customer, error) // returns a list of customers from the server side and an error if present
}
