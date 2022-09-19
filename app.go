package main

import "fmt"

type Person struct
{
	name string
	age int
}

func (person *Person) getName() string {
	return person.name
}

func (person *Person) isOverAge() bool {
	return person.age > 18
}

func times (number int, times int) float32 {
	return float32(number * times)
}

func main() {
	var person = Person { 
		name: "romario", 
		age: 26 }

	fmt.Println(times(person.age, 10))
    fmt.Println(person.name, person.age)
	fmt.Println(person.getName())
	fmt.Println(person.isOverAge())

	var arr = [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// first element from arr
	var flexibleSlice = arr[0:1]
	fmt.Println(len(flexibleSlice))

	var associative = map[string]string {
		"foo": "bar",
		"baz": "boo"}
	for key, value := range associative {
        fmt.Println(key, value)
    }

	var value, ok = associative["is"]
	fmt.Printf("value:%v type:%T", value, value)
	fmt.Println()
	fmt.Printf("ok:%v type:%T", ok, ok)
	fmt.Println()
}