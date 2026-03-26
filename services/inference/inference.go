package inference

import (
    "encoding/json"
    "fmt"
)

/**
* InferSchema is a stub for converting request/response to JSON schema
*/
func InferSchema(payload []byte) map[string]string {
  var data map[string]interface{}
  json.Unmarshal(payload, &data)

  schema := make(map[string]string)

  for key, value := range data {
  	switch value.(type) {
      case string:
      	schema[key] = "string"
      case float64:
      	schema[key] = "number"
      case bool:
      	schema[key] = "boolean"
      case []interface{}:
      	schema[key] = "array"
      case map[string]interface{}:
      	schema[key] = "object"
      default:
      	schema[key] = "unknown"
    }
	}

	fmt.Println("Inferred schema:", schema)
  return schema
}
