package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   00233257753432,
		},
	}
	fmt.Printf("%+v\n", john)

	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "filipe@domain.com", phone: 233213442, address: "Waterford, Ireland"}
	fmt.Println(myContact)

}
