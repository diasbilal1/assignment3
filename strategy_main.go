package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (cc *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card\n", amount)
}

type PayPalPayment struct{}

func (pp *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal\n", amount)
}

type KaspiBankPayment struct{}

func (kb *KaspiBankPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Kaspi Bank\n", amount)
}

type BitcoinPayment struct{}

func (bp *BitcoinPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Bitcoin\n", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ProcessPayment(amount float64) {
	pc.strategy.Pay(amount)
}

func main() {
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}
	kaspiBank := &KaspiBankPayment{}
	bitcoin := &BitcoinPayment{}

	fmt.Println()
	context := NewPaymentContext(creditCard)
	context.ProcessPayment(100.0)

	fmt.Println()
	context.SetStrategy(payPal)
	context.ProcessPayment(50.0)

	fmt.Println()
	context.SetStrategy(kaspiBank)
	context.ProcessPayment(75.0)

	fmt.Println()
	context.SetStrategy(bitcoin)
	context.ProcessPayment(25.0)

	fmt.Println()
	// Add more SetStrategy calls
	context.SetStrategy(creditCard)
	context.ProcessPayment(60.0)

	fmt.Println()
	context.SetStrategy(payPal)
	context.ProcessPayment(30.0)

	fmt.Println()
	context.SetStrategy(kaspiBank)
	context.ProcessPayment(45.0)
}
