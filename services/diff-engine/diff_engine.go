package diffengine

import (
	"fmt"

	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
)

func DetectBreakingChanges(endpoint string, oldSchema, newSchema map[string]string) {
	for key, oldType := range oldSchema {
		newType, exists := newSchema[key]

		if !exists {
			msg := fmt.Sprintf("field '%s' removed", key)
			fmt.Println("❌ BREAKING: ", msg)
			storage.SaveChange(endpoint, msg)
			continue
		}

		if oldType != newType {
			msg := fmt.Sprintf("field '%s' type changed (%s → %s)", key, oldType, newType)
			fmt.Println("❌ BREAKING: ", msg)
			storage.SaveChange(endpoint, msg)
		}
	}
}
