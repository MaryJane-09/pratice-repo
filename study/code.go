package main

import "fmt"

type Dog struct {
	Name string
	Age string
}

func (d Dog) Bark(){
	fmt.Println(d.Name, "say Woof!")
}

func main() {
	dog := Dog{
		Name: "Ben",
	}
	dog.Bark()

	dog2 := Dog{
		Name: "Jack",
	}
	dog2.Bark()
}