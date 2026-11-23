package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	//gateway stripe
	//gateway razorpay
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorPayPaymentGw := razorpay{}
	// razorPayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

type refund struct {
	gateway paymenter
}

func (r refund) makeRefund(amount float32, account string) {
	r.gateway.refund(amount, account)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Making fake payment for testing purpose")
}

type payPal struct{}

func (p payPal) pay(amount float32) {
	fmt.Println("Making payment using payPal", amount)
}

func (p payPal) refund(amount float32, account string) {
	fmt.Println("Making refund of", amount, "in acccount", account, " using payPal")
}

func main() {
	//stripePaymentGw := stripe{}
	//razorPayPaymentGw := razorpay{}
	//fakeGateway := fakePayment{}
	payPalGateway := payPal{}

	newPayment := payment{
		//gateway: razorPayPaymentGw,
		//gateway: fakeGateway,
		gateway: payPalGateway,
	}
	newPayment.makePayment(100)

	newRefund := refund{
		gateway: payPalGateway,
	}
	newRefund.makeRefund(200, "AccABC321")

}
