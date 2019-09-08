// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Вариативная функция (функция с переменным количеством аргументов)
// /ch5/sum
// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"os"
)

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//!-

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice
}


Неявно вызывающая функция выделяет память для массива, копирует в него аргу­
менты и передает в функцию срез, представляющий весь массив. Последний вызов,
показанный выше, таким образом, ведет себя так же, как приведенный ниже, в кото­
ром показано, как вызывать вариативные функции, когда аргументы уже находятся
в срезе: следует добавить троеточие после последнего аргумента:
values := []int{l, 2, 3, 4}
fmt.Println(sum(values...)) И
"10"
Хотя параметр . . . i n t ведет себя в теле функции, как срез, тип вариативной функ­
ции отличается от типа функции с параметром, являющимся обычным срезом:
func f(...int) {}
func g([]int) {}
fmt.Printf("%T\n", f) 11 "func(...int)"
fmt.Printf("%T\n", g) // "func([]int)"
Вариативные функции часто используются для форматирования строк. Показан­
ная далее функция e r r o r f создает форматированное сообщение об ошибке с номе­
ром строки в его начале. Суффикс f является широко распространенным соглаше­
нием об именовании вариативных функций, которые принимают строку формата в
стиле P r i n t f .
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Стр. %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
linenum, name := 12, "count"
errorf(linenum, "не определен %s", name) // "Стр. 12: не определен count"
Тип i n t e r f a c e { } означает, что данная функция может принимать любые значе­
ния в качестве последних аргументов, как будет пояснено в главе 7, “Интерфейсы”.
