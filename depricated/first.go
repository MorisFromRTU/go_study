package main

import (
	"fmt"
	"os"
	"strconv"
)

// создается структура (похоже на создание объекта)
type Student struct{
	Name string
	IQ int
}
// вместо конструктора используется функция, возвращающая экземпляр нужного типа
func NewStudent(name string, iq int) Student{
	return Student{
		Name: name,
		IQ: iq,
	}
}
// функция структуры Student
func (stu *Student) AddIQ() {
	stu.IQ += 10
}

func structs(){
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <name> <iq>")
		return
	}

	name:= os.Args[1]
	IQ, _ := strconv.Atoi(os.Args[2])

	// Создаем экземпляр структуры 
	moris := &Student{name, IQ}
	// Создаем экземпляр структуры (способ 2)
	egor := &Student{}
	// Создаем экземпляр структуры (способ 3)
	sveta := new(Student)

	sveta.Name = "Sveta"
	sveta.IQ = 10
	moris.AddIQ()
	sveta.AddIQ()
	fmt.Println(moris.Name, moris.IQ)
	fmt.Println(sveta.Name, sveta.IQ)
	fmt.Println(egor.IQ)
}

type Dog struct{
	Name string
	Age int
}

func (d *Dog) Bark(){
	fmt.Println("Baaark!!")
}

type ColoredDog struct{
	*Dog
	Color string
}
// аналог наследования, мы передаем как поле структуры другую структуру 
func composition(){
	dog := new(Dog)
	dog.Name = "doggy"
	dog.Age = 12
	mike := new(ColoredDog)
	mike.Dog = dog
	mike.Color = "brown"
	
	fmt.Println("My name is", mike.Name, "I am", mike.Color, "I can do this..")
	mike.Bark()

}

// Массивы

var scores[10]int

func arr(){
	items := [3]int{3,4,5}
	fmt.Println(items[0])
}

func main() {
	arr()
}
