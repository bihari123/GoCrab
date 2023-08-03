package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import the database driver of your choice here
	_ "github.com/denisenkom/go-mssqldb" // Microsoft SQL Server
	_ "github.com/go-sql-driver/mysql"   // MySQL
	_ "github.com/lib/pq"                // PostgreSQL
	// Add other database drivers as needed
)

// Database configuration
const (
	dbDriver = "mysql"                                  // Change this to your desired database driver
	dbSource = "root:password@tcp(localhost:3306)/mydb" // Replace with your actual connection string
)

func main() {
	// Open a connection to the database
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Ping the database to ensure the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Query the database and crawl the records
	rows, err := db.Query(
		"SELECT * FROM users",
	) // Replace "your_table" with the name of the table you want to crawl
	if err != nil {
		log.Fatalf("Failed to query the database: %v", err)
	}
	defer rows.Close()

	// Get the column names from the result set
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to get column names: %v", err)
	}

	// Prepare a slice to hold the scan destinations (interface{})
	scanArgs := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))

	// Assign the address of each RawBytes value to the corresponding scan destination
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Loop through the rows and process the data
	for rows.Next() {
		// Scan the values into the scan destinations
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalf("Failed to scan row data: %v", err)
		}

		// Process the data (you can replace this with your custom logic)
		for i, col := range values {
			if col == nil {
				fmt.Printf("%s: NULL\n", columns[i])
			} else {
				fmt.Printf("%s: %s\n", columns[i], string(col))
			}
		}
		fmt.Println("-----------")
	}

	// Check for any errors during row iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error occurred while iterating rows: %v", err)
	}
}

/*
package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import the database driver of your choice here
	_ "github.com/denisenkom/go-mssqldb" // Microsoft SQL Server
	_ "github.com/go-sql-driver/mysql"   // MySQL
	_ "github.com/lib/pq"                // PostgreSQL
	// Add other database drivers as needed
)

// Database configuration
const (
	dbDriver = "mysql"                                  // Change this to your desired database driver
	dbSource = "root:password@tcp(localhost:3306)/mydb" // Replace with your actual connection string
)

func main() {
	// Open a connection to the database
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Ping the database to ensure the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Query the database and crawl the records
	rows, err := db.Query(
		"SELECT * FROM users",
	) // Replace "your_table" with the name of the table you want to crawl
	if err != nil {
		log.Fatalf("Failed to query the database: %v", err)
	}
	defer rows.Close()

	// Loop through the rows and process the data
	for rows.Next() {
		var id int
		var name string
		var email string
		// Scan the values from the row into the variables
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatalf("Failed to scan row data: %v", err)
		}

		// Process the data (you can replace this with your custom logic)
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// Check for any errors during row iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error occurred while iterating rows: %v", err)
	}
}
*/
/*
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Email    string
}

func main() {
	// Replace the following connection parameters with your database credentials
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Generate and insert 100 mock rows
	for i := 1; i <= 100; i++ {
		user := User{
			Username: faker.Username(),
			Email:    faker.Email(),
		}

		// Prepare the INSERT statement
		stmt, err := db.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
		if err != nil {
			log.Fatalf("Failed to prepare the statement: %v", err)
		}

		// Execute the INSERT statement
		_, err = stmt.Exec(user.Username, user.Email)
		if err != nil {
			log.Fatalf("Failed to insert row: %v", err)
		}

		fmt.Printf("Inserted row %d: Username: %s, Email: %s\n", i, user.Username, user.Email)
	}
}
*/
