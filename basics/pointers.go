package main

func main() {
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // новый указатель на переменную a

	// получение указателя на переменнут типа int
	// инициализировано значением по-умолчанию
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c и a не изменились

	c = d   // теперь с указывает туда же, куда d
	*c = 14 // с = 14 -> d = 14, a = 12
	tutoial()
}
//pointers.go
package main

import "fmt"

//
func tutoial() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
