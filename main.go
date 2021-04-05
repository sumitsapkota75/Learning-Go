package main

import "fmt"

func changeValue(str *string) {
	*str = "Changed"
}

func changeValue2(str string) string {
	str = "Changed"
	print("from pointer", str)
	return str
}

func main() {
	toChange := "Hello"
	fmt.Println(toChange)
	a := 7

	changeValue2(toChange)
	fmt.Println(&a)

}
