package main

import (
	"fmt"
	"github.com/fsayed/magazine"
)

func main() {
	//var subscriber magazine.Subscriber
	//tedious way of setting values to fields
	//s.Name = "Jare Dunn"
	//s.Rate = 4.99
	//s.Active = true
	//fmt.Println(s.Rate)

	//just like maps and lists, Go provides struct literals

	//subscriber = magazine.Subscriber{Name: "Jared Dunn", Rate: 4.99, Active: true}

		
	//fmt.Println("Name:", s.Name)
	//fmt.Println("Rate:", s.Rate)
	//fmt.Println("Active:", s.Active)
	//--------------------------------//


	address := magazine.Address{Street: "4141 Yonge Street", City: "Toronto", State: "ON", PostalCode: "M2P2A8"}
	subscriber := magazine.Subscriber{Name: "Jared Dunn", Rate: 4.99, Active: true}
	subscriber.HomeAddress = address

	fmt.Println(subscriber.HomeAddress)

	//---------------------------------

	//you can also daisy chain structs using dot operators. u can access inner struct using outer struct

	subscriber1 := magazine.Subscriber{Name: "Bertram Guilfoyle"}
	subscriber1.HomeAddress.Street = "101 Yonge Street"
	subscriber1.HomeAddress.City = "Toronto"
	subscriber1.HomeAddress.State = "ON"
	subscriber1.HomeAddress.PostalCode  = "M2P2A8"

	fmt.Println("Name: ", subscriber1.Name)
	fmt.Println("Street: ", subscriber1.HomeAddress.Street)
	fmt.Println("City: " , subscriber1.HomeAddress.City)
	fmt.Println("State: ", subscriber1.HomeAddress.State)
	fmt.Println("Postal: ", subscriber1.HomeAddress.PostalCode)

	//not this is so much typing, how many times did we type HomeAddress above. You can use anonymous structs for this reason
	//anonymous struct has a benefit that its fields are promoted to outer struct and u can access then as if they are part of
	//the outer struct
	//subscriber := magazine.Subscriber{Name: "Richard Hendricks"}
	//subscriber.Street = "123 Oak Street"
	//
	//In the magazine.go file, the Subscriber struct would look like this:
	//struct Subscriber {
	//	Nane string
	//	Rate float64
	//	Active bool
	//	Address
	//}
	// In the above code, we have removed the field name HomeAddress and just left the name of the inner sturct Address
	
}
	
