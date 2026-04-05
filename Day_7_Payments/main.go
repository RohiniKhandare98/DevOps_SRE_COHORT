package main

import (
	"fmt"

	"github.com/RohiniKhandare98/DevOps_SRE_COHORT/Day_7_Payments/pkg/payments"
	"github.com/RohiniKhandare98/DevOps_SRE_COHORT/Day_7_Payments/pkg/payments/creditcard"
	"github.com/RohiniKhandare98/DevOps_SRE_COHORT/Day_7_Payments/pkg/payments/upi"
)

func Checkout(method payments.PaymentMethod, amount float64) string {

	msg := method.Pay(amount)
	return msg
}

func main() {

	fmt.Println("payment interface example")

	rohiniUPI := upi.UPIPayment{UpiID: "rohini@okicici", App: "Gpay"}

	msg := Checkout(rohiniUPI, 50)

	fmt.Printf("Payment successful: %s", msg)

	rohiniCard := creditcard.CreditCardPayment{CardNumber: "1234567890123456", ExpiryDate: "01/2028", CVV: "456"}

	msg = Checkout(rohiniCard, 50)

	fmt.Printf("Payment successful: %s", msg)
}
