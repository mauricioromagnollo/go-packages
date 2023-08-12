package gotype

import "fmt"

// GetTypeOf returns the type of the value passed as parameter as string.
/*
	Example:
		var value bool = true
		fmt.Println(GetTypeOf(value))
		Output: "bool"
*/
func GetTypeOf(value interface{}) string {
	return fmt.Sprintf("%T", value)
}
