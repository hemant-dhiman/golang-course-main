package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	emain   string
	zipCode int
}

func main() {
	//alex := person{"Alex", "Anderson"} // 1 way of declaring struct
	//or

	//alex := person{firstName: "Alex", lastName: "Anderson"} // 2nd way of declaring struct
	//fmt.Println(alex)

	// var alex person // 3rd way of declaring struct

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			emain:   "email.com",
			zipCode: 123456,
		},
	}
	fmt.Printf("%+v", jim)
}
