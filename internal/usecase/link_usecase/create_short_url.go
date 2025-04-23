package link_usecase

import (
	"crypto/sha256"
	"fmt"
	"math/rand"

	"shortened_link_creation_service/internal/domain"
)

const allowedCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const shortURLLength = 10
const myDomain = "https://myDomain.ru/"

func (uc *UseCase) CreateShortURL(url string) (*domain.Link, error) {
	shortURL := shortedURL(url)
	fmt.Println("Короткий URL: ", shortURL)
	link := domain.NewLink(url, shortURL)

	linkCreate, err := uc.repo.Create(link)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.Create: %w", err)
	}

	return linkCreate, nil
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

	return myDomain + result
}
