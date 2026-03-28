package storage

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Change struct {
	Endpoint  string    `json:"endpoint"`
	Change    string    `json:"change"`
	CreatedAt time.Time `json:"created_at"`
}

func InitDB() {
	connStr := "postgres://postgres:postgres@localhost:5432/contracts?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func SaveSchema(endpoint string, schema map[string]string) {
	jsonSchema, err := json.Marshal(schema)

	if err != nil {
		log.Println("Error marshalling schema: ", err)
		return
	}

	_, err = DB.Exec(
		"INSERT INTO schemas (endpoint, schema) VALUES ($1, $2)",
		endpoint, jsonSchema,
	)

	if err != nil {
		log.Println("Error saving schema:", err)
	}
}

func SaveChange(endpoint, change string) {
	_, err := DB.Exec("INSERT INTO changes (endpoint, change) VALUES ($1, $2)", endpoint, change)

	if err != nil {
		log.Println("Error saving change: ", err)
	}
}

func GetLastSchema(endpoint string) (map[string]string, error) {
	row := DB.QueryRow("SELECT schema FROM schemas WHERE endpoint=$1 ORDER BY created_at DESC LIMIT 1", endpoint)

	var schemaJSON []byte
	err := row.Scan(&schemaJSON)

	if err != nil {
		return nil, err
	}

	var schema map[string]string
	json.Unmarshal(schemaJSON, &schema)

	return schema, nil
}

func GetChanges() ([]Change, error) {
	rows, err := DB.Query("SELECT endpoint, change, created_at FROM changes ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var changes []Change
	for rows.Next() {
		var c Change
		err := rows.Scan(&c.Endpoint, &c.Change, &c.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		changes = append(changes, c)
	}
	return changes, nil
}
