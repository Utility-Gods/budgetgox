package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// User represents an individual using the budgeting app.
type User struct {
	ID       string // Unique identifier for the user
	Name     string // Name of the user
	Email    string // Email of the user
	Password string // Hashed password (store passwords securely!)
}

// Account represents a financial account belonging to a user (e.g., checking, savings, credit card).
type Account struct {
	ID      string  // Unique identifier for the account
	UserID  string  // ID of the user the account belongs to
	Name    string  // Name of the account (e.g., "Checking Account")
	Type    string  // Type of account (e.g., "Checking", "Savings", "Credit Card")
	Balance float64 // Current balance of the account
}

// Transaction represents a financial transaction in an account.
type Transaction struct {
	ID         string  // Unique identifier for the transaction
	AccountID  string  // ID of the account this transaction belongs to
	CategoryID string  // ID of the category this transaction belongs to
	Amount     float64 // Amount of the transaction (positive for income, negative for expenses)
	Date       string  // Date of the transaction (choose appropriate date type)
	Note       string  // Optional note or description for the transaction
}

// Category represents a category for transactions (e.g., groceries, rent, utilities).
type Category struct {
	ID   string // Unique identifier for the category
	Name string // Name of the category
}

// Budget represents a budget set by the user for a particular category.
type Budget struct {
	ID         string  // Unique identifier for the budget
	UserID     string  // ID of the user this budget belongs to
	CategoryID string  // ID of the category this budget applies to
	Amount     float64 // Budgeted amount for the category
	Period     string  // Budget period (e.g., "Monthly", "Yearly")
}

func main() {
	fmt.Println("Server is running on port 8080")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("index.html"))
		tmp1.Execute(w, nil)

	}

	h2 := func (w http.ResponseWriter, r *http.Request) {
		log.Print("Request received")
		log.Print(r.Method)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-entry",h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
