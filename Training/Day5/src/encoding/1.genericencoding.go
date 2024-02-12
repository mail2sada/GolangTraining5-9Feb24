package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{
        "roll-no": 1,
        "name": "Anil Kumar",
        "class": "10A",
        "Marks": {
        "Science": 100,
        "Maths": 99,
        "MentalAbility": 100,
        "Total": 299,
        "Average": 99.66666666666667
        }}`
	b := []byte(jsonString)

	m := make(map[string]interface{})

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Received error while Unmarshaling...")
	} else {
		fmt.Println(m)

		for key, val := range m {
			fmt.Println(key, val)
			switch val.(type) {
			case int:
				fmt.Println("Integer")
			case string:
				fmt.Println("string")
			case map[string]interface{}:
				mp := val.(map[string]interface{})
				for k, v := range mp {
					fmt.Println(k, v)
				}
			default:
				fmt.Println("Not handled")
			}
		}
	}

}
