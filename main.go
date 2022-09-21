package main

import (
	"fmt"

	"github.com/romario-giacholari/go-demo/entities"
	"github.com/romario-giacholari/go-demo/utilities"
)

func main() {
	person := entities.Person{
		Name: "romario",
		Age:  26,
	}

	fmt.Println(utilities.Times(person.Age, 10))
	fmt.Println(person.Name, person.Age)
	fmt.Println(person.GetName())
	fmt.Println(person.IsOverAge())

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// first element from arr
	flexibleSlice := arr[0:1]
	fmt.Println(len(flexibleSlice))

	associative := map[string]string{
		"foo": "bar",
		"baz": "boo",
	}

	for key, value := range associative {
		fmt.Println(key, value)
	}

	value, ok := associative["is"]
	fmt.Printf("value:%v type:%T", value, value)
	fmt.Println()
	fmt.Printf("ok:%v type:%T", ok, ok)
	fmt.Println()
}
