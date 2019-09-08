package main

import "fmt"

func main() {
	// инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"name1":     "Alex",
		"name2":     "Andy",
		"name3":     "Olga",
		"name4":     "Irine",
		"name5":     "Kate",
		"lastName": "Kitchens",
	}

	// сразу с нужной ёмкостью
	profile := make(map[string]string, 10)

	// количество элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	// если ключа нет - вернёт значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	fmt.Printf("%#v\n", user)
	// удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
