package main

import (
	"fmt"
	_ "math"
	"reflect"
)

func setKeyValue(key string, value interface{}, source map[string]interface{}) bool {

	var ans bool
	if _, ok := source[key]; ok {
		source[key] = value
		return true
	} else {
		for _, y := range source {

			if reflect.TypeOf(y).Kind() == reflect.Map {
				ans = setKeyValue(key, value, y.(map[string]interface{}))

			}

		}
	}
	return ans
}

func RemoveKey(key string, source map[string]interface{}) bool {
	var ans bool
	if _, ok := source[key]; ok {
		delete(source, key)
		return true
	} else {
		for _, y := range source {

			if reflect.TypeOf(y).Kind() == reflect.Map {
				ans = RemoveKey(key, y.(map[string]interface{}))

			}

		}
	}
	return ans
}

func main() {

	foods := map[string]interface{}{
		"Food": "delicious",
		"eggs": map[string]interface{}{
			"greeting": "Hello",
			"name":     "Kaif",
			"course":   "MCA",
		},
		"name":   "Soransh",
		"course": "BTech",
	}

	x := setKeyValue("Food", "Very Delicious", foods)

	if x {
		fmt.Println("Changed")
	} else {
		fmt.Println("Not Found")

	}

	fmt.Println(foods)

	y := RemoveKey("greeting", foods)

	if y {
		fmt.Println("Deleted")
	} else {
		fmt.Println("Not Found")
	}

	fmt.Println(foods)

}
