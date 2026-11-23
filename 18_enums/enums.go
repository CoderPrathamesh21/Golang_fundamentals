package main

import "fmt"

// enumerated types
// type OrderStatus int
type OrderStatus string

const (
	CONFIRMED  OrderStatus = "confirmed"
	RECEIVED   OrderStatus = "received"
	SHIPPED    OrderStatus = "shipped"
	DISPATCHED OrderStatus = "dispatched"
	DELIVERED  OrderStatus = "delivered"
	CANCELED   OrderStatus = "canceled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(DISPATCHED)
}
