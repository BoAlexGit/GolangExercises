// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// trace.go
// See page 146.
// функция BigSlowOperation немедленно вызывает функцию trace, которая выполняет
// запись о входе в функцию и возвращает значение-функцию, которая при вызове
// выполняет запись о выходе из функции.
// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"log"
	"os"
	"time"
)

//!+main
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/
/*
Тот же шаблон может использоваться для других ресурсов, помимо сетевых под­
ключений, например, чтобы закрыть открытый файл:
io/ioutiL
package ioutil
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.CloseQ
	return ReadAll(f)
}
или для разблокирования мьютекса (раздел 9.2):
var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
*/

При каждом вызове b ig S lo w O p e r a tio n выполняется запись о времени входа в
нее и выхода, а также о времени работы данной функции (мы использовали t im e .
S le e p для имитации длительной работы функции):
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 вход в bigSlowOperation
2015/11/18 09:53:36 выход из bigSlowOperation (10.000589217s)
Отложенные функции выполняются после того, как инструкция возврата обновля­
ет переменные результатов функции. Поскольку анонимная функция может обращать­
ся к переменным охватывающей функции, включая именованные результаты, отло­
женная анонимная функция имеет доступ к результатам функции, в которой вызвана.
Рассмотрим функцию d o u b le :
func double(x int) int {
	return x + x
}
Присвоив имя ее результирующей переменной и добавив инструкцию d e f e r , мы
можем заставить функцию выводить свои аргументы и результат при каждом вызове:
func double(x int) (result int) {
	defer funcQ { fmt.Printf("double(%d) = %d\n,,J x, result) }()
		return x + x
	}
	_ = double(4)
	// Вывод:
	// "double(4) = 8”
	Этот трюк является излишеством для такой простой функции, как d o u b le , но мо­
	жет быть полезен в функции со многими операторами r e t u r n .
		Отложенная анонимная функция может даже изменять значения, которые возвра­
	щает охватывающая функция:
	func triple(x int) (result int) {
		defer funcQ { result += x }()
		return double(x)
	}
	fmt.Println(triple(4)) // "12"
	Поскольку отложенные функции не выполняются до самого конца выполнения
	функции, инструкция d e f e r в цикле заслуживает дополнительного изучения. У при­
	веденного ниже кода могут возникнуть проблемы из-за исчерпания доступных фай­
	ловых дескрипторов, поскольку ни один файл не будет закрыт, пока не будут обрабо­
	таны все файлы:
	for
	filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.CloseQ // Примечание: рискованно; может привести
		// к исчерпанию файловых дескрипторов
		// ...работа с f ...
		Одним из решений может быть перенос тела цикла, включая инструкцию d e f e r , в
		другую функцию, которая вызывается на каждой итерации.
		for _, filename := range filenames {
			if err := doFile(filename); err != nil {
				return err
			}
		}
		func doFile(filename string) error {
			f, err := os.Open(filename)
			if err != nil {
			return err
		}
		}
		defer f.CloseQ
		// ...работа с f...
		Приведенный ниже пример представляет собой усовершенствованную программу
		f e t c h (раздел 1.5), которая записывает HTTP-ответ в локальный файл, а не в стан­
		дартный вывод. Она выводит имя файла из последнего компонента пути URL, кото­
		рый получает с помощью функции p a t h . B ase.
