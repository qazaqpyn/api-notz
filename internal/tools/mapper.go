package tools

import (
	"errors"
	"reflect"
)

func AddUserIdToModel(model interface{}, userId string) ([]map[string]interface{}, error) {
	val := reflect.ValueOf(model)
	if val.Kind() != reflect.Slice {
		return nil, errors.New("input type is not correct")
	}

	result := make([]map[string]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		newMap := make(map[string]interface{})

		// Copy existing key-value pairs
		keys := elem.MapKeys()
		for _, key := range keys {
			newMap[key.String()] = elem.MapIndex(key).Interface()
		}

		// Add userId to the new map
		newMap["userId"] = userId

		// Store the new map in the result slice
		result[i] = newMap
	}

	return result, nil
}
