package main

import (
    "context"
    "log"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	 "github.com/Utility-Gods/budgetox/ent"
    _ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASS")
	ssl_mode := os.Getenv("SSL_MODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, pass, ssl_mode)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	 CreateUser(context.Background(), client); 
	 
	log.Println("user was created: ")
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Create().
		SetID("1").
		SetName("sid").
		SetPassword("password").
        SetAge(30).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %w", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}