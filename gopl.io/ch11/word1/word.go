// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 303 (355).
//!+
// word1.go
// Пакет word предоставляет утилиты для игр со словами,
// Package word provides utilities for word games.
package word

import "testing"

// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//!-

*********************************************************************
Каждый тестовый файл должен импортировать пакет testing Тестовые функции
имеют следующую сигнатуру:
func TestName(t *testing.T) {
	// ...
}
Имена тестовых функций должны начинаться с T e s t; необязательный суффикс
должен начинаться с прописной буквы:
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }
