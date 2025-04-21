package service_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	// роутер ручек
	mux := http.NewServeMux()
	// link
	mux.HandleFunc("POST /createShortURL", sp.GetLinkController().CreateShortURL)
	mux.HandleFunc("GET /returnURL", sp.GetLinkController().ReturnURL)

	return mux
}
