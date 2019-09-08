//functions.go
//Функция в Go объявляется через служебное слово func, после которого следует имя
//функции, перечисление параметров в скобках и тип возвращаемого значения после скобок. Результат из
//функции возвращается, с помощью оператора return, как и в других языках.
//
package main

import "fmt"

// обычное объявление
func singleIn(in int) int {
	return in
}

// много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// именованный результат
func namedReturn() (out int) {
	out = 2
	return
}

// несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// несколько именованных результатов
func multipleNamedReturn(ok bool) (rez int, err error) {
	rez = 1
	if ok {
		err = fmt.Errorf("some error happend")
		// аналогично return rez, err
		return 3, fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

// не фиксированное количество параметров
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

//В этом случае на вход, в данном случае в переменную in, поступит
// slice интов, по которым вы сможете
//проитерироваться, используя цикл for range. Обратим внимание на
// одну тонкость при вызове функции c вариативным числом параметров.

func main() {
	// fmt.Println(multipleNamedReturn(false))
	// return

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return
}
//У нас есть слайс nums, если передать в сумму просто слайс, то компилятор
// в этом случае будет ругаться,
//потому что это отдельный тип. А ожидается какое-то количество
// повторяющихся одиночных параметров.
//Поэтому необходимо использовать троеточие после имени слайса.