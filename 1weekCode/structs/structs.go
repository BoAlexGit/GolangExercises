package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id int
	// Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	// полное объявление структуры
	var acc Account = Account{
		Id: 1,
		// Name: "rvasily",
		Person: Person{
			Name:    "Василий",
			Address: "Москва",
		},
	}
	fmt.Printf("%#v\n", acc)

	// короткое объявление структуры
	acc.Owner = Person{2, "Romanov Vasily", "Moscow"}

	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}

//structs.go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}


//struct-fields.go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
