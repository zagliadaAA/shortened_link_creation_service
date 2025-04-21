package link_usecase

import (
	"crypto/sha256"
	"fmt"
	"math/rand"

	"shortened_link_creation_service/internal/domain"
)

const allowedCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const shortURLLength = 10

func (uc *UseCase) CreateShortURL(URL string) (*domain.Link, error) {
	// идем в бд и проверяем, существует для длинной короткая или нет
	// если существует, то возвращаем короткую
	// если не существует, то идем в бд, создаем и возвращаем короткую

	shortURL := shortedURL(URL)

	fmt.Println("Короткий URL: ", shortURL)

	return nil, nil
}

func shortedURL(url string) string {
	hash := sha256.Sum256([]byte(url))

	var result string
	/*
		При делении по модулю(%): 0-255(каждый байт) % 63(число разрешенных символов) = диапазон 0-62
		rand добавлен для равномерности распределения по полученному hash и снижения вероятности коллизии
		минусы подхода: каждый раз новый результат, символы могут повторятся
	*/
	for i := 0; i < shortURLLength; i++ {
		index := hash[rand.Intn(len(hash))] % byte(len(allowedCharacters))
		result += string(allowedCharacters[index])
	}

	return result
}
