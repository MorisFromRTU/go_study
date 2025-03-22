package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	name  string
	age   uint16
	color string
	breed string
}

func NewDog(name string, age uint16, color string, breed string) Dog {
	return Dog{
		name:  name,
		age:   age,
		color: color,
		breed: breed,
	}
}

type Cat struct {
	name  string
	age   uint16
	color string
}

func NewCat(name string, age uint16, color string) Cat {
	return Cat{
		name:  name,
		age:   age,
		color: color,
	}
}

func (cat Cat) Speak() string {
	return fmt.Sprintf("%s says meaw meaw", cat.name)
}

func (dog Dog) Speak() string {
	return fmt.Sprintf("%s says bark bark", dog.name)
}

func MakeAnimalsSpeak(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func main() {
	barsik := NewCat("Barsik", 10, "yellow")
	rex := NewDog("Rex", 12, "black", "ovcharka")
	var animals []Animal = []Animal{barsik, rex}
	MakeAnimalsSpeak(animals)
}
