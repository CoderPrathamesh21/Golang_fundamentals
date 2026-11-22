package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
	email string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

// constructor
func newOrder(id string, amount float32, status string) *order {
	myOrder1 := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder1
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	myOrder1 := newOrder("Ord110", 500.00, "Shipped")
	fmt.Println(myOrder1)
	fmt.Println(myOrder1.id)

	fmt.Println("==============================")

	myOrder := order{
		id:     "ord101",
		amount: 100.00,
		status: "Received",
	}

	myOrder.createdAt = time.Now()

	myOrder.changeStatus("Confirmed!!!")

	//to get any field from struct
	fmt.Println(myOrder.status)

	fmt.Println("Order struct: ", myOrder)

	fmt.Println("==============================")

	newCustomer := customer{
		name:  "John",
		phone: "+21 654789",
		email: "john@hotmail.com",
	}

	myOrder2 := order{
		id:        "Ord102",
		amount:    400.00,
		status:    "Dispatched",
		createdAt: time.Now(),
		customer:  newCustomer, //struct embedding
	}

	myOrder2.customer.name = "Bob"
	fmt.Println(myOrder2.customer)

	fmt.Println(myOrder2.getAmount())
	fmt.Println("Order struct: ", myOrder2)
	fmt.Println("==============================")

	//another way to create struct
	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

}
