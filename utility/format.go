package utility

import (
	"encoding/json"
	"time"
)

func FormatDate(date, currentISOFormat, newISOFormat string) (string, error) {
	t, err := time.Parse(currentISOFormat, date)
	if err != nil {
		return date, err
	}
	return t.Format(newISOFormat), nil
}

func NumberFormat(t interface{}) float64 {
	num, ok := t.(float64)
	if !ok {
		numInt, ok := t.(int)
		if ok {
			num = float64(numInt)
		}
		return num
	}
	return num
}

func Add(num1, num2 interface{}) float64 {
	first, ok := num1.(float64)
	if !ok {
		firstInt, ok := num1.(int)
		if ok {
			first = float64(firstInt)
		}
	}
	second, ok := num2.(float64)
	if !ok {
		secondInt, ok := num1.(int)
		if ok {
			second = float64(secondInt)
		}
	}
	return first + second
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	// Convert the struct to JSON bytes
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Create an empty map
	result := make(map[string]interface{})

	// Unmarshal the JSON bytes into the map
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Convert int values to int instead of float64
	ConvertIntValues(result)

	return result, nil
}

func ConvertIntValues(m map[string]interface{}) {
	for key, value := range m {
		switch v := value.(type) {
		case float64:
			if intValue := int(v); float64(intValue) == v {
				m[key] = intValue
			}
		case map[string]interface{}:
			ConvertIntValues(v)
		}
	}
}
