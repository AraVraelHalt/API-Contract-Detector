package storage

import (
  "database/sql"
  _ "github.com/lib/pq"
	"encoding/json"
  "log"
)

var DB *sql.DB

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

func GetLastSchema(endpoint string) (map[string]string, error) {
  row := DB.QueryRow(
      "SELECT schema FROM schemas WHERE endpoint=$1 ORDER BY created_at DESC LIMIT 1",
      endpoint,
  )

  var schemaJSON []byte
  err := row.Scan(&schemaJSON)
  
	if err != nil {
      return nil, err
  }

  var schema map[string]string
  json.Unmarshal(schemaJSON, &schema)

  return schema, nil
}
