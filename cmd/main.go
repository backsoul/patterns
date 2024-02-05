package main

import (
	"fmt"
	"log"

	database "github.com/backsoul/pattern/internal"
)

func main() {
	// Initialize the database
	dbFactory, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	db, err := dbFactory.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Use GetNow method
	nowQuery, err := dbFactory.GetNow()
	if err != nil {
		log.Fatal(err)
	}

	// Execute the query
	rows, err := db.Query(nowQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var now string
		if err := rows.Scan(&now); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Current time:", now)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database operations completed!")
}
