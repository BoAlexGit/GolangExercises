// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// echo_test.go
// Test of echo command.  Run with: go test gopl.io/ch11/echo

//!+
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)",
			test.newline, test.sep, test.args)

		out = new(bytes.Buffer) // captured output
		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

//!-
/****************************************************************************
Обратите внимание, что тестовый код находится в том же пакете, что и основ­
ной рабочий код. Хотя имя пакета — m ain, и в нем определена функция m ain, при
тестировании этот пакет действует, как библиотека, которая предоставляет функцию
T e s tE c h o для тест-драй вера; его функция m ain игнорируется.
Организуя тест в виде таблицы, мы можем легко добавлять новые тестовые при­
меры. Давайте посмотрим, что произойдет, если тест не будет пройден, просто доба­
вив в таблицу следующую строку :
// Примечание: неверно указанное ожидание!
{ tr u e , V 1, [ ] s t r i n g { " a u, " b ", " с "} , "а b с \ п " Ь
При этом go t e s t выводит следующее:
$ go test gopl.io/chll/echo
--- FAIL: TestEcho (0.00s)
e c h o _ te s t.g o :3 1 : e c h o ( tr u e , V ' , ["a""b" V ] )
= " а ,Ь ,с " ,

требуется "a b c\n "
FAIL
FAIL
gopl.io/chll/echo 0.006s

 */
