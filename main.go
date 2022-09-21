package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/romario-giacholari/go-demo/config"
	"github.com/romario-giacholari/go-demo/entities"
	"github.com/romario-giacholari/go-demo/utilities"
)

const PI = math.Pi

func main() {
	fmt.Println(PI)

	environmentVariables := config.Load()
	environment := environmentVariables["APP_ENV"]
	log.Println(environment)

	timeNow := time.Now().UTC()
	fmt.Println(timeNow)

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
	flexibleSlice := []string{"one", "two", "three"}
	for index, value := range flexibleSlice {
		fmt.Println(index, value)
	}

	// returns new slice - it does not modify original
	newFlexibleSlixe := append(flexibleSlice, "four")
	fmt.Println(newFlexibleSlixe)

	associative := map[string]string{
		"foo": "bar",
		"baz": "boo",
	}

	associative["new"] = "foobar"
	fmt.Println(associative)

	for key, value := range associative {
		fmt.Println(key, value)
	}

	value, ok := associative["is"]
	fmt.Printf("value:%v type:%T", value, value)
	fmt.Println()
	fmt.Printf("ok:%v type:%T", ok, ok)
	fmt.Println()
}
