package main

import "fmt"

//////////////////////////////
// 1️⃣ SRP - Single Responsibility
//////////////////////////////

// Order represents a purchase
type Order struct {
	ID     int
	Amount float64
}

// OrderService is only responsible for order creation logic
type OrderService struct{}

func (os OrderService) CreateOrder(amount float64) Order {
	return Order{ID: 1, Amount: amount}
}

//////////////////////////////
// 2️⃣ OCP - Open/Closed Principle
//////////////////////////////

// PaymentMethod is open for extension
type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) {
	fmt.Println("💳 Paid", amount, "using Credit Card")
}

type UpiPayment struct{}

func (u UpiPayment) Pay(amount float64) {
	fmt.Println("📱 Paid", amount, "using UPI")
}

// ✅ If tomorrow we add CryptoPayment, no need to change existing code.

//////////////////////////////
// 3️⃣ LSP - Liskov Substitution Principle
//////////////////////////////

// DiscountCalculator interface
type DiscountCalculator interface {
	Calculate(amount float64) float64
}

// PercentageDiscount can replace DiscountCalculator safely
type PercentageDiscount struct{}

func (p PercentageDiscount) Calculate(amount float64) float64 {
	return amount * 0.90 // 10% discount
}

// FlatDiscount can also replace DiscountCalculator safely
type FlatDiscount struct{}

func (f FlatDiscount) Calculate(amount float64) float64 {
	return amount - 50 // Flat ₹50 discount
}

// ✅ Both types behave correctly as DiscountCalculator without breaking logic.

//////////////////////////////
// 4️⃣ ISP - Interface Segregation Principle
//////////////////////////////

// Instead of one huge interface, we split responsibilities

type EmailNotifier interface {
	SendEmail(order Order)
}

type SMSNotifier interface {
	SendSMS(order Order)
}

type EmailService struct{}

func (e EmailService) SendEmail(order Order) {
	fmt.Println("📧 Email sent for Order:", order.ID)
}

type SMSService struct{}

func (s SMSService) SendSMS(order Order) {
	fmt.Println("📩 SMS sent for Order:", order.ID)
}

// ✅ A service that only needs email isn't forced to implement SMS.

//////////////////////////////
// 5️⃣ DIP - Dependency Inversion Principle
//////////////////////////////

// Database is an abstraction — high-level code depends on interface, not implementation
type Database interface {
	Save(order Order)
}

type MySQLDB struct{}

func (db MySQLDB) Save(order Order) {
	fmt.Println("💾 Order saved to MySQL DB:", order.ID)
}

type PostgresDB struct{}

func (db PostgresDB) Save(order Order) {
	fmt.Println("💾 Order saved to Postgres DB:", order.ID)
}

//////////////////////////////
// 🧪 Bringing Everything Together
//////////////////////////////

func main() {
	// SRP: Create order
	orderService := OrderService{}
	order := orderService.CreateOrder(500)

	// OCP: Choose payment method without changing code
	var payment PaymentMethod = CreditCardPayment{}
	payment.Pay(order.Amount)

	// LSP: Choose any discount strategy
	var discount DiscountCalculator = PercentageDiscount{}
	finalAmount := discount.Calculate(order.Amount)
	fmt.Println("💰 Final Amount after discount:", finalAmount)

	// ISP: Use only the services you need
	emailNotifier := EmailService{}
	emailNotifier.SendEmail(order)

	// DIP: Save order without knowing actual DB details
	var db Database = PostgresDB{}
	db.Save(order)
}
