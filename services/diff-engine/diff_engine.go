package diffengine

import "fmt"

func DetectBreakingChanges(oldSchema, newSchema map[string]string) {
	for key, oldType := range oldSchema {
		newType, exists := newSchema[key]

		if !exists {
			fmt.Printf("❌ BREAKING: field '%s' removed\n", key)
			continue
		}

		if oldType != newType {
			fmt.Printf("❌ BREAKING: field '%s' type changed (%s → %s)\n", key, oldType, newType)
		}
	}
}
