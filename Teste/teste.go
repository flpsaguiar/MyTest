package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

func main() {
	p := Person{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City:   "Gotham",
			State:  "NY",
			Zip:    "01313",
		},
	}
	p.Talk()
	p.Location()
}

/*


type T struct{
	 name string
	 value int
}

//aqui a função printT está sendo assinada para a struct T
func (c* T) printT(){
	 fmt.Println(c.name)
	 fmt.Println(c.value)
	 fmt.Println(&c.name)
}

//função que espera como parâmetro uma struct
func prinT2(c* T){
	 fmt.Println(c.name)
	 fmt.Println(c.value)
}

func prinT3() (int, float64){
	 return 666, 1.1
}

func main(){
	 var t T
	 t.name = "Teste"
	 t.value = 10
	 t.printT()
	 fmt.Println(prinT3())
}
*/
