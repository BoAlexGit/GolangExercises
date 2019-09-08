// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.
// forEachNode вызывает функции pre(x) и post(x) для каждого узла х
// в дереве с корнем п. Обе функции необязательны.
// рге вызывается до посещения дочерних узлов, a post - после,

// Outline2 prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
/*
/********************************************************************
func square(n int) int
{ return n * n }
func negative(n int) int
{ return -n }
func product(m, n int) int { return m * n }
166

f := square
fmt.Println(f(3))
// "9м
f = negative
fmt.Println(f(3))
// "-3"
fmt.Printf("%T\n", f) // "func(int) int"
f = product
// Ошибка компиляции: нельзя присваивать
// func(int, int) int переменной func(int) int
//Нулевым значением типа функции является n i l . Вызов нулевой функции приво­
//дит к аварийной ситуации:
var f func(int) int
f(3)
// Аварийная ситуация: вызов nil-функции
//Значение-функцию можно сравнить с nil:
var f func(int) int
if f != nil {
f(3 )
}


func addl(r rune) rune { return r + 1 }
fmt.Println(strings.Map(addl, "HAL-9000")) // "IBM.:111"
fmt.Println(strings.Map(addl, "VMS"))
// "WNT"
fmt.Println(strings.Map(addl, "Admix"))
// "Benjy"

var depth int
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf(M%*s<%s>\nM, depth*2,n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2,n.Data)
	}
}
*//////////////////////////////////////////////////////////////////////
