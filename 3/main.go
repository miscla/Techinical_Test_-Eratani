package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type UserSpend struct {
	ID             int     `json:"id"`
	Country        string  `json:"country"`
	CreditCardType string  `json:"credit_card_type"`
	CreditCard     string  `json:"credit_card_number"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	TotalSpend     float64 `json:"total_spend"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []UserSpend `json:"data"`
}

type UserInput struct {
	Country        string `json:"country"`
	CreditCardType string `json:"credit_card_type"`
	CreditCard     string `json:"credit_card_number"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type PostResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ID      int    `json:"id,omitempty"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := "host=localhost port=5432 user=mac dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully")
}

func getTopSpendingUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := `
        SELECT 
            u.id,
            u.country,
            u.credit_card_type,
            u.credit_card_number,
            u.first_name,
            u.last_name,
            SUM(b.total_buy) AS total_spend
        FROM 
            public."user" u
        INNER JOIN 
            public.belanja b ON u.id = b.id_user
        GROUP BY 
            u.id, u.country, u.credit_card_type, u.credit_card_number, u.first_name, u.last_name
        ORDER BY 
            total_spend DESC
        LIMIT 10
    `

	rows, err := db.Query(query)
	if err != nil {
		response := Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer rows.Close()

	var users []UserSpend
	for rows.Next() {
		var user UserSpend
		err := rows.Scan(
			&user.ID,
			&user.Country,
			&user.CreditCardType,
			&user.CreditCard,
			&user.FirstName,
			&user.LastName,
			&user.TotalSpend,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		response := Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{
		Status:  "success",
		Message: "Data retrieved successfully",
		Data:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response := PostResponse{
			Status:  "error",
			Message: "Invalid request body: " + err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validasi input
	if input.Country == "" || input.CreditCardType == "" || input.CreditCard == "" ||
		input.FirstName == "" || input.LastName == "" {
		response := PostResponse{
			Status:  "error",
			Message: "All fields are required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	query := `
        INSERT INTO public."user" (country, credit_card_type, credit_card_number, first_name, last_name)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `

	var userID int
	err = db.QueryRow(query,
		input.Country,
		input.CreditCardType,
		input.CreditCard,
		input.FirstName,
		input.LastName,
	).Scan(&userID)

	if err != nil {
		response := PostResponse{
			Status:  "error",
			Message: "Failed to insert data: " + err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := PostResponse{
		Status:  "success",
		Message: "User created successfully",
		ID:      userID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/api/top-spending-users", getTopSpendingUsers)
	http.HandleFunc("/api/users", createUser)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("GET  /api/top-spending-users - Get top spending users")
	fmt.Println("POST /api/users - Create new user")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
