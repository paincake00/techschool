package main

import "fmt"

// Human представляет человека с именем и возрастом
type Human struct {
	Name string
	Age  int
}

// SayMyName выводит имя человека
func (h Human) SayMyName() {
	fmt.Println("My name is", h.Name)
}

// SetAge устанавливает возраст
func (h *Human) SetAge(age int) {
	h.Age = age
	fmt.Println("Set age is", h.Age)
}

// Action описывает действие у человека
type Action struct {
	Human
	Value string
}

// DoAction выполняет действие
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

	// вызов методов родительского Human у Action
	action.Human.SayMyName()
	action.Human.SetAge(30)

	action.DoAction()
}
