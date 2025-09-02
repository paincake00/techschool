package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) SayMyName() {
	fmt.Println("My name is", h.Name)
}

func (h *Human) SetAge(age int) {
	h.Age = age
	fmt.Println("Set age is", h.Age)
}

type Action struct {
	Human
	Value string
}

func (a Action) DoAction() {
	fmt.Println("Do action:", a.Value)
}

func main() {
	action := Action{
		Human: Human{
			Name: "Bob",
			Age:  20,
		},
		Value: "sleep",
	}

	action.Human.SayMyName()
	action.Human.SetAge(30)

	action.DoAction()
}
