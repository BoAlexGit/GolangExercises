// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// word_test.go
//!+test
package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

//!-test

// The tests below are expected to fail.
// See package gopl.io/ch11/word2 for the fix.

//!+more
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

//!-more

/************************************************************************
$ go test
--- FAIL: TestFrenchPalindrome (0.00s)
word_test.go:28: IsPalindrome("ete") = false
--- FAIL: TestCanalPalindrome (0.00s)
word_test.go:35: IsPalindrome("A man, a plan, a canal: Panama")
= false
--- FAIL
FAIL gopl.io/chll/wordl 0.014s



$ go test -v
=== RUN TestPalindrome
--- PASS: TestPalindrome (0.00s)
=== RUN TestNonPalindrome
--- PASS: TestNonPalindrome (0.00s)
=== RUN TestFrenchPalindrome
--- FAIL: TestFrenchPalindrome (0.00s)
word__test.go:28: IsPalindrome("ete") = false
=== RUN TestCanalPalindrome
--- FAIL: TestCanalPalindrome (0.00s)
word_test.go:35: IsPalindrome("A man, a plan, a canal: Panama")
= false
FAIL
exit status 1
FAIL gopl.io/chll/wordl 0.017s
А флаг -run , аргументом которого является регулярное выражение, приводит к тому,
что команда go t e s t выполняет только те тесты, имена функций которых соответ­
ствуют шаблону:
$ go test -v -run="French|Canal"
=== RUN TestFrenchPalindrome
--- FAIL: TestFrenchPalindrome (0.00s)
word_test.go:28: IsPalindrome("ete") = false
===== RUN TestCanalPalindrome
--- FAIL: TestCanalPalindrome (0.00s)
word_test.go:35: IsPalindrome("A man, a plan, a canal: Panama")
= false
FAIL
exit status 1
FAIL gopl.io/chll/wordl 0.014s

*************************************************************************
Находящийся в том же каталоге файл word_test.go содержит две тестовые
функции с именами TestPalindrome и TestNonPalindrome. Каждая из них прове­
ряет, дает ли IsPalindrome правильный ответ для одного ввода, и сообщает о сбоях
с помощью t.Error:
package word
import "testing"
func TestMyPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
		if !IsPalindrome("kayak") {
			t.Error(`IsPalindrome("kayak") = false`)
		}
	}
	func TestMyNonPalindrome(t *testing.T) {
		if IsPalindrome("palindrome") {
			t.Error(`IsPalindrome("palindrome") = true`)

		}
	}
	Команда go t e s t (или go b u i l d ) без аргумента, указывающего пакет, работает с
	пакетом в текущем каталоге. Мы можем строить и выполнять тесты с помощью сле­
	дующей команды:

$ cd $GOPATH/src/gopl.io/chll/word1
$ go test
ok
gopl.io/chll/wordl 0.008s

 */
