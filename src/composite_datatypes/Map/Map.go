package main

import "fmt"

func main() {

	/* **************************************
		Declaring a map
	****************************************/

	// Map is a implementation of a hash table
	// Use make() to create a map
	var idMap map[string]int // The square brackets contain the datatype of the key and the other
	// datatype is the value stored
	idMap = make(map[string]int) // This is the way to define a map using a make function

	/* ***************************************
	"idMap1 := make(map[string]int)" Another way of declaring map
	*************************************** */

	mapLiteral := map[string]int{"joe": 123} // This is a map literal

	// idMap is empty as no values have been added to it yet
	fmt.Println(idMap, mapLiteral)
	mapLiteral["joe"] = 124 // Access value of a map using this syntax
	idMap["jane"] = 456     // Adding new key value pair to the map
	idMap["joe"] = 123

	/* ***************************************
	idMap = mapLiteral This will completely overwrite the map so do not use
	*************************************** */

	fmt.Println(idMap)

	delete(idMap, "joe") // Delete a value completely from a map
	fmt.Println(idMap)

	// This way the value stored in idMap["joe"] is assigned to idVal and a boolean value p is returned
	// stating whether key "joe exists or not"
	idVal, p := idMap["joe"]
	fmt.Printf("Value of key \"joe\" after it has been deleted: %d and boolean value stating whether the key exists or not: %t\n", idVal, p)

	idMap["joe"] = 123
	idVal, p = idMap["joe"]
	fmt.Printf("Value of key \"joe\" after it has been readded: %d and boolean value stating whether the key exists or not: %t\n", idVal, p)

	fmt.Printf("Length of map: %d\n", len(idMap))

	// * **************************************
	// Iterating through a map

	// This is similar to the way array are iterated but in an array it returns an index and the value at
	// that index
	for key, value := range idMap {
		fmt.Println(key, value)
	}
	// ****************************************
}
