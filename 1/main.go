package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Action struct {
	Human Human
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) GetPhone() string {
	return h.Phone
}

func main() {
	a := Action{
		Human: Human{
			Name:  "leha",
			Age:   22,
			Phone: "8988-888-88-88",
		},
	}

	fmt.Println(a)
	fmt.Println(a.Human.Age)
	fmt.Println(a.Human.GetName())
	fmt.Println(a.Human.GetAge())
	fmt.Println(a.Human.GetPhone())
}
