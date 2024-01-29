package main

import (
	"fmt"
	_ "math"
	"reflect"
)

type keyValue struct { // struct that contains required values to store individual units
	name   string
	course string
}

type result struct { // struct that contains the slice of keyValue to store full Data Structure
	kv []keyValue
}

func PopulateStruct(m map[string]interface{}, x interface{}) {
	rs := x.(*result) // Type assertion to get the result pointer

	_, ok1 := m["name"] // checking if the given keys exist
	_, ok2 := m["course"]

	if ok1 || ok2 { // if anyone of the following exists store it
		var a keyValue

		if ok1 {
			a.name = m["name"].(string)
		}

		if ok2 {
			a.course = m["course"].(string)
		}

		rs.kv = append(rs.kv, a) // update the slice
	}
	for _, y := range m { // iterate to find if there are any sub-maps in the given map

		if reflect.TypeOf(y).Kind() == reflect.Map {
			PopulateStruct(y.(map[string]interface{}), x) // if yes recursive call is made to the sub-graph

		}

	}
}

func setKeyValue(key string, value interface{}, source map[string]interface{}) bool {

	var ans bool // for storing ans i.e. whether the key exists or not
	if _, ok := source[key]; ok {
		source[key] = value // if the key exists update the value and return true
		return true
	} else {
		for _, y := range source { // else check if there are any submaps then recurssively make calls over them and store their result

			if reflect.TypeOf(y).Kind() == reflect.Map {
				ans = ans || setKeyValue(key, value, y.(map[string]interface{}))

			}

		}
	}
	return ans // since by default the value is false, f no sub call returns true it will return false
}

func RemoveKey(key string, source map[string]interface{}) bool {
	var ans bool // to check if the given key exists or not
	if _, ok := source[key]; ok {
		delete(source, key) // if found delete and return true
		return true
	} else { // else iterate over the map to check for submaps
		for _, y := range source {

			if reflect.TypeOf(y).Kind() == reflect.Map {
				ans = ans || RemoveKey(key, y.(map[string]interface{})) // recursive call to perform deletion on the sub map

			}	// OR is used because if a later recurssive call returns false, the ans is not overwritten. If we don't have duplicate keys
			// we can add a condition and return here only

		}
	}
	return ans
}

func main() {

	foods := map[string]interface{}{	// example input
		"Food": "delicious",
		"eggs": map[string]interface{}{
			"greeting": "Hello",
			"name":     "Kaif",
			"course":   "MCA",
		},
		"name":   "Soransh",
		"course": "BTech",
	}

	x := setKeyValue("Food", "Very Delicious", foods)// call to setKeyValue

	if x {
		fmt.Println("Changed")
	} else {
		fmt.Println("Not Found")

	}

	fmt.Println(foods)

	y := RemoveKey("greeting", foods) // call to RemoveKey

	if y {
		fmt.Println("Deleted")
	} else {
		fmt.Println("Not Found")
	}

	fmt.Println(foods)

	var r result

	PopulateStruct(foods, &r)	//call to populate struct

	fmt.Println(r)

}
