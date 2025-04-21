package link_controller

import (
	"net/http"

	"shortened_link_creation_service/internal/controller"
)

type createLinkReq struct {
	URL string `json:"url"`
}

func (c *Controller) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// парсим json в нашу структуру
	// валидируем тело запроса или парамерты
	// вызываем юзкейс
	// обрабатываем ошибки если есть
	// возвращаем результат
	var req createLinkReq
	if err := controller.DecodeRequest(w, r, &req); err != nil {
		controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось считать тело запроса"))
		return
	}

	link, err := c.linkUseCase.FindByURL(req.URL)
	if err != nil {
		link, err = c.linkUseCase.CreateShortURL(req.URL)
		if err != nil {

		}
	}

	err = controller.EncodeResponse(w, link.ShortURL)
	if err != nil {
		controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось отправить ответ"))
		return
	}
}
