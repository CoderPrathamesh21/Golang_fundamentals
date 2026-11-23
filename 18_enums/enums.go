package main

import "fmt"

// enumerated types
// type OrderStatus int
type OrderStatus string

const (
	CONFIRMED  OrderStatus = "confirmed"
	RECEIVED               = "received"
	SHIPPED                = "shipped"
	DISPATCHED             = "dispatched"
	DELIVERED              = "delivered"
	CANCELED               = "canceled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(DISPATCHED)
}
