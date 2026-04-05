package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (d Dog) Sound() {
	fmt.Println(d.Name, "says: Woof")
}

type Cat struct {
	Name string
}

func (c Cat) Sound() {
	fmt.Println(c.Name, "says: Meow")
}

type Cow struct {
	Name string
}

func (c Cow) Sound() {
	fmt.Println(c.Name, "says: Moo")
}

func MakeAnimalSpeak(a Animal) {
	a.Sound()
}

func main() {

	animals := []Animal{
	 Dog{Name: "Bruno"},
	 Cat{Name: "Misty"},
	 Cow{Name: "Gauri"},
	}
	for _, animal := range animals{
		MakeAnimalSpeak(animal)
	}
}
