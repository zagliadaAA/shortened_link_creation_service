package link_controller

import (
	"net/http"

	"shortened_link_creation_service/internal/controller"
)

type returnURLReq struct {
	ShortURL string `json:"short_url"`
}

func (c *Controller) ReturnURL(w http.ResponseWriter, r *http.Request) {
	var req returnURLReq
	if err := controller.DecodeRequest(w, r, &req); err != nil {
		controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось считать тело запроса"))
		return
	}

	link, err := c.linkUseCase.FindByShortURL(req.ShortURL)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("не удалось найти оригинальный URL"))
		return
	}

	err = controller.EncodeResponse(w, link.URL)
	if err != nil {
		controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось отправить ответ"))
		return
	}
}
