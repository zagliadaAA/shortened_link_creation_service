package link_controller

import (
	"net/http"

	"shortened_link_creation_service/internal/controller"
)

type createLinkReq struct {
	URL string `json:"url"`
}

func (c *Controller) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req createLinkReq
	if err := controller.DecodeRequest(w, r, &req); err != nil {
		controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось считать тело запроса"))
		return
	}

	link, err := c.linkUseCase.GetLinkByURL(req.URL)
	if err != nil {
		link, err = c.linkUseCase.CreateShortURL(req.URL)
		if err != nil {
			controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось создать короткий URL"))
			return
		}
	}

	err = controller.EncodeResponse(w, link.ShortURL)
	if err != nil {
		controller.RespondInternalServerError(w, controller.NewStatusInternalServerError("не удалось отправить ответ"))
		return
	}
}
