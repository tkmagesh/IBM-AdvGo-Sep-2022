package main

import "fmt"

type Employee struct {
	Id  int
	Org *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	org := Organization{
		Name: "IBM",
		City: "Bengaluru",
	}
	e1 := Employee{
		Id:  100,
		Org: &org,
	}

	e2 := e1
	//e2.Id = 200
	e2.Org.City = "Pune"
	fmt.Println(e1.Org.City)
	fmt.Println(e2.Org.City)

}
