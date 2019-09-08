// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 130.
// WaitForServer пытается соединиться с сервером заданного URL.
// Попытки предпринимаются в течение минуты с растущими интервалами.
// Сообщает об ошибке, если все попытки неудачны,

// The wait program waits for an HTTP server to start responding.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//!+
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

//!-

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	//!+main
	// (In function main.)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
	//!-main
}

/*
/****************************************************************

package io
import "errors"
// EOF - это ошибка, возвращаемая функцией Read,
// когда больше нет данных для чтения,
var EOF = errors.New("EOF")


in := bufio.NewReader(os.Stdin)
for {
err := in.ReadRuneQ
if err == io.EOF {
break 11 Чтение завершено
>
if err != nil {
return fmt.Errorf("сбой чтения: %v", err)
>
// ...использование г...
}



	// (В функции main.)
if err := WaitForServer(url); err != nil {
fmt.Fprintf(os.Stderr, "Сервер не работает: %v\n", err)
os.Exit(l)
}

if err := WaitForServer(url); err != nil {
log.Fatalf("Сервер не работает: %v\n", err)
}
//Формат по умолчанию полезен при работе “долгоиграющего” сервера и менее
//удобен для интерактивного инструмента:
2006/01/02 15:04:05 Сервер не работает: неверный домен: bad.gopl.io
log.SetPrefix("wait: ")
log.SetFlags(G)
//В некоторых случаях достаточно просто записать ошибку в журнал и продолжить
//работу, возможно, с уменьшенной функциональностью. При этом вновь можно выби­
//рать между использованием пакета lo g , который добавляет обычный префикс:
if err := Ping(); err != nil {
log.Printf("ошибка ping: %v; сеть отключена", err)
}
//и выводом непосредственно в стандартный поток ошибок:
if err := PingQ; err != nil {
fmt.Fprintf(os.Stderr, "ошибка ping: %v; сеть отключена\п", err)
}
//(Все функции пакета lo g добавляют символ новой строки, если таковой отсутствует.)
//Последняя стратегия — в редких случаях можно безопасно полностью игнориро­
//вать ошибку:
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
return fmt.Errorf("ошибка создания временного каталога: %v", err)
}
// ... используем временный каталог ...
os.RemoveAll(dir) // Игнорируем ошибки; $TMPDIR периодически очищается
**************************************************************************/