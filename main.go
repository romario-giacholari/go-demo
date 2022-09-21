package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/romario-giacholari/go-demo/config"
	"github.com/romario-giacholari/go-demo/entities"
	"github.com/romario-giacholari/go-demo/utilities"
)

const PI = math.Pi

func main() {
	toUpper := func(str string) string {
		return strings.ToUpper(str)
	}

	str := toUpper("awesome")
	fmt.Println(str)

	fmt.Println(PI)

	environmentVariables := config.Load()
	environment := environmentVariables["APP_ENV"]
	log.Println(environment)

	timeNow := time.Now().UTC()
	fmt.Println(timeNow)

	person := entities.Person{
		Name: "romario",
		Age:  26,
		Job: entities.Occupation{
			Title: "Software Engineer",
		},
	}

	fmt.Println(person.Job.Title)
	fmt.Println(utilities.Times(person.Age, 10))
	fmt.Println(person.Name, person.Age)
	fmt.Println(person.GetName())
	fmt.Println(person.IsOverAge())

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

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

	// go routines
	worker := func(jobs <-chan int, results chan<- int) {
		for n := range jobs {
			results <- n
		}
	}

	capacity := 1000000
	jobs := make(chan int, capacity)
	results := make(chan int, capacity)
	startTime := time.Now()

	go worker(jobs, results)

	for i := 0; i < capacity; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < capacity; j++ {
		fmt.Println(<-results)
	}

	close(results)

	duration := time.Since(startTime)
	fmt.Println("duration", duration)
}
