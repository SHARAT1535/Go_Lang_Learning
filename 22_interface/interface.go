// package main

// import "fmt"


// type paymenter interface {
// 	pay(amount float32)
// 	refund(amount float32)  
// } 

// type payment struct{
// 	gateway paymenter
// }

// func (p payment) makepayment(amount float32) {
// 	razorPaymentGW:= raxorpay{}
// 	strippaymentgw := strip{}
// 	razorPaymentGW.pay(amount)
// 	strippaymentgw.pay(amount)
// 	p.gateway.pay(amount)
// }


// type raxorpay struct{}
// type strip struct{
// 	//payment
// }
// type juspay struct{}

// func (j juspay)pay(amount float32){
// 	fmt.Println("maling payment using juspay",amount)
// }

// func (s strip)pay(amount float32){
// 	fmt.Println("making payment using stripṇ")
// }

// func (r raxorpay) pay(amount float32) {
// 	fmt.Println("making payment using raxzorpay", amount)
// }

// type fakePayment struct{}

// func (f fakePayment) pay(amount float32) {
// 	fmt.Println("fake payment of amount", amount)
// }
// type paypal struct{}

// func (p paypal) pay(amount float32) {
// 	fmt.Println("making payment using paypal", amount)
// } 

// func (p paypal) refund(amount float32) {
// 	fmt.Println("refunding payment using paypal", amount)
// }
// func (p raxorpay) refund(amount float32) {
// 	fmt.Println("refunding payment using raxorpay", amount)
// }

// func main() {
// 	// newpayment := payment{}
// 	// newpayment.makepayment(100)
// 	paypalgw:= raxorpay{}
// 	newPayment := payment{ gateway: paypalgw}
// 	newPayment.makepayment(100)

// }

package main

import "fmt"

// Interface
type Speaker interface {
    Speak() string
}

// Struct
type Person struct {
    Name string
}

// Method implementation
func (p Person) Speak() string {
    return "Hello, I am " + p.Name
}

func main() {
    var s Speaker
    p := Person{Name: "Sharat"}

    s = p
    fmt.Println(s.Speak())
}

