package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"shortened_link_creation_service/cmd/service_provider"
)

func main() {
	repoType := dbSelection()

	sp := service_provider.NewServiceProvider(repoType)

	// запуск сервера
	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}

/*Сервис принимает длинный URL, пришедший в POST-запросе.
Сервис проверят, нет ли уже такого URL в базе данных. Если входящий длинный URL уже существует в системе, то в ней остался и сгенерированный короткий вариант.
Если длинного URL нет в базе данных или срок действия сокращённой ссылки истёк, необходимо создать новый токен и отправить короткий URL в качестве ответа, сохранив результат в базе данных.
Сервис отправляет короткий URL в качестве ответа. Статус HTTP 201, если создана новая запись или 200, если запись уже была в базе данных.*/

/*
хеширование - это алгоритм преобразования исходных данных в последовательность заранее известной длины.
метод позволяет сократить любую последовательность до фиксированного количества байт.
sha-256 это одитн из алгоритмов безопасного хеширования
*/

func dbSelection() string {
	// Запрос выбора репозитория у пользователя
	fmt.Println("Выберите тип базы данных:\n" +
		"1. Postgres\n" +
		"2. In-Memory\n" +
		"Введите номер (1 или 2): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var repoType string
	switch input {
	case "1":
		repoType = "postgresRepo"
		fmt.Println("Выбрана база данных Postgres")
	case "2":
		repoType = "inMemoryRepo"
		fmt.Println("Выбрана In-Memory база данных")
	default:
		fmt.Println("Некорректный выбор. Будет использована In-Memory база данных")
		repoType = "inMemoryRepo"
	}

	return repoType
}
