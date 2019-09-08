// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// /ch5/squares
// squares возвращает
// функцию, которая при каждом вызове
// возвращает квадрат очередного числа,

// See page 135.

// The squares program demonstrates a function value with state.
package main

import "fmt"

//!+
// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

//!-

/*
//Литерал функции позволяет определить функцию в точке использования. В каче­
//стве примера приведенный выше вызов strings .Мар можно переписать следующим
//образом:
strings.Map(func(r rune) rune { return r + 1 }, "HAL9000")
*/