// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Echo.go
// See page 308.
//!+
// Echo выводит аргументы командной строки,
// Echo prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)
//  "опустить символы новой строки"
var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)
// Изменяется на время тестирования
var out io.Writer = os.Stdout // modified during testing

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}

//!-

/****************************************************************
import "math/rand"
// randomPalindrome возвращает палиндром, длина и содержимое
// которого задаются генератором псевдослучайных чисел rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // Случайная длина до 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // Случайная руна до '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
func TestRandomPalindromes(t *testing.T) {
	// Инициализация генератора псевдослучайных чисел,
	seed := time.Now().UTC().UnixNano()
	t .Logf("ГПСЧ инициализирован: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf(IsPalindrome(%q) = false", p)
		}
	}
}

 */
