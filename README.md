# banking

go 
postman
dbeaver


Objectives
    1. Mechanism of HTTP Web Server
    2. Handler Functions and Request Multiplexer (Router)
    3. Request and Response Headers
    4. Marshaling data structures to JSON and XML representations

Note : All ports are interfaces

Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

Ports are constant. But adapters can be changed. Adapter is just a basis of conversion to what is specified in the port.

Go uses duck typing. A struct doesn’t have to implement an interface in Go. Duck typing tends to happen at run time. If it looks like a duck, and it quacks like a duck, then it is a duck. If it has a set of methods that match an interface, then you can use it wherever that interface is needed without explicitly defining that your types implement that interface.


*** Implemeting Hexgonal Architecture ***
--> Retrieving a list of customers 

--> Business logic has a  dependecy of the repository. Business logic needs to implement the service layer interface.