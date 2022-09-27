package main

import (
	"fmt"
)

//go:generate echo "Hi there!"
func main() {

	fmt.Println("done")
	/* products := models.Products{
		models.Product{101, "Pen", 10, 100, "Stationary"},
		models.Product{104, "Pencil", 5, 200, "Stationary"},
		models.Product{103, "Marker", 50, 10, "Stationary"},
		models.Product{105, "Mouse", 900, 5, "IT"},
		models.Product{102, "Scribble-Pad", 25, 50, "Stationary"},
	}

	fmt.Println("Initial List")
	for _, product := range products {
		fmt.Println(product)
	}

	stationaryProductPredicate := func(p models.Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println("Stationary Products")
	for _, product := range stationaryProducts {
		fmt.Println(product)
	}

	customers := models.Customers{
		models.Customer{100, "Magesh", "Bengaluru"},
		models.Customer{102, "Suresh", "Pune"},
		models.Customer{103, "Ramesh", "Bengaluru"},
		models.Customer{104, "Rajesh", "Bengaluru"},
		models.Customer{105, "Ganesh", "Mysuru"},
	}

	bengaluruCustomers := customers.Filter(func(c models.Customer) bool {
		return c.City == "Bengaluru"
	})
	fmt.Println(bengaluruCustomers) */
}
