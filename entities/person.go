package entities

type Person struct {
	Name string
	Age  int
}

func (person *Person) GetName() string {
	return person.Name
}

func (person *Person) IsOverAge() bool {
	return person.Age >= 18
}
