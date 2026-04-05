package main

import (
	"fmt"
	"time"
)
// check() -> checks services are healty or not 
// checker -> implements all check functions 
type Checker interface {
	Check()
}

type HttpService struct {
	Name    string
	URL     string
	isHealthy bool
}

func (h HttpService) Check() {

	// checks services are healty or not
	fmt.Println("checking if ", h.URL, "is healthy or not")
	time.Sleep(3 * time.Second)
	h.isHealthy = true

	fmt.Printf("HTTP => %s, %s, %v\n", h.Name, h.URL, h.isHealthy)
}

type DBService struct {
	Name    string
	Host    string
	Port    int
	isHealthy bool
}

func (db DBService) Check() {

	fmt.Println("checking if ", db.Host, "is healthy or not")
	time.Sleep(3 * time.Second)

	db.isHealthy = true

	fmt.Printf("DB => %s, %s:%d, %v\n", db.Name, db.Host, db.Port, db.isHealthy)
}

type TCPService struct {
	Name    string
	Addr    string
	isHealthy bool
}

func (t TCPService) Check() {
	fmt.Println("checking if ", t.Addr, "is healthy or not")
	time.Sleep(3 * time.Second)

	t.isHealthy = true

	fmt.Printf("TCP: %s %s %v\n", t.Name, t.Addr, t.isHealthy)
}

type UDPService struct {
	Name    string
	Addr    string
	isHealthy bool
}

func (u UDPService) Check(){
	fmt.Println("checking if ", u.Addr, "is healthy or not")
	time.Sleep(3 * time.Second)

	u.isHealthy = true

	fmt.Printf("TCP: %s %s %v\n", u.Name, u.Addr, u.isHealthy)
}

func main() {

	fmt.Println("welcome to interfaces")


	// if want to check multiple services ... we have to make array of struct services --> and then run loop 
	// thats why interface are here , checks automatically who has check() function
	
	
	// now we implement checker interface array 
	
//	services array => all sevices structures and their methods 
	
	services := []Checker{
		HttpService{Name: "example.com", URL: "https://example.com", isHealthy: false},
		DBService{Name: "postgre", Host: "localhost", Port: 5432, isHealthy: false},
		TCPService{Name: "kafka", Addr: "192.168.10.65:8765", isHealthy: false},
		UDPService{Name: "udp", Addr: "192.168.10.65:8765", isHealthy: false},
	}

 // run check functions of all services -- now it gets very simple 
	for _, svc := range services {
		svc.Check()
	}

}

