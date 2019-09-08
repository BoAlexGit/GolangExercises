// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline1 prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-
/*
/***************************************************************************

// обычная функция
func doNothing() {
	fmt.Println("i’m regular function")
}


// обычное объявление
func singleIn(in int) int {
	return in
}
При этом, если у вас несколько параметров одного типа, то вы можете указать тип только один раз, а
сами параметры перечислить через запятую. Либо же указывать тип для каждого из параметров.
// много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}
В Go можно указывать сразу указывать переменную, в которую будет возвращен результат.
// именованный результат
func namedReturn() (out int) {
	out = 2
	return
}
При этом вы можете записать туда значение и написать пустой return без имени этой переменной. И она
вернет вам то, что вы записали. В данном случае это переменная out. Впрочем, можно написать,например,
return 3, даже в таком случае функция будет корректно работать. В Go функции могут возвращать
несколько результатов. Чаще всего всего это используется для возврата ошибки в качестве второго пара-
метра.
// несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}
************************************************************/