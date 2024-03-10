package app

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	api "groupie/internal/API"
)

type Handler struct{}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.ErrorHand(w, r, errStatus{http.StatusNotFound, "Данной страницы не существует"})
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		h.ErrorHand(w, r, errStatus{http.StatusMethodNotAllowed, "Only GET requests allowed"})
		return
	}

	artists, err := api.FillArtists()
	if err != nil {
		log.Println(err)
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Не удалось получить данные артистов"})
		return
	}

	artistJSON, err := json.Marshal(artists)
	if err != nil {
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Перевод данных в формат json не удался"})
		return
	}
	const homePageTemp = "./ui/html/home.page.html"
	ts, err := template.ParseFiles(homePageTemp)
	if err != nil {
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Ошибка при парсинге home"})
		return
	}
	err = ts.Execute(w, string(artistJSON))
	if err != nil {
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Ошибка сервера"})
		return
	}
}
