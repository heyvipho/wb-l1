package main

import "fmt"

// Human - наследуемый тип
type Human struct {
	name string
	age  int
}

// CreateHuman создает эксземпляр Human
func CreateHuman(name string, age int) Human {
	return Human{
		name: name,
		age:  age,
	}
}

func (v *Human) GetName() string {
	return v.name
}

func (v *Human) GetAge() int {
	return v.age
}

// Action наследует Human
type Action struct {
	Human
}

func main() {
	human := CreateHuman("George", 27)
	action := Action{Human: human}

	fmt.Printf("Имя: %v\n", action.GetName())
	fmt.Printf("Возраст: %v\n", action.GetAge())
}
