package app

import (
	api "groupie/internal/API"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		h.ErrorHand(w, r, errStatus{http.StatusMethodNotAllowed, "Некорректный метод"})
		return
	}
	idText := r.URL.Query().Get("id")
	if idText == "" || idText[0] == '0' {
		h.ErrorHand(w, r, errStatus{http.StatusBadRequest, "Неправильный формат ID"})
		return
	}
	id, err := strconv.Atoi(idText)

	if err != nil || id < 1 {
		h.ErrorHand(w, r, errStatus{http.StatusNotFound, "Артист с таким ID не найден"})
		return
	}
	const artistPageTemp = "./ui/html/artist.html"
	ts, err := template.ParseFiles(artistPageTemp)
	if err != nil {
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Ошибка при парсинге шаблона"})
		return
	}

	artist, err := api.FillById(strconv.Itoa(id))

	if artist.ID == 0 {
		h.ErrorHand(w, r, errStatus{http.StatusNotFound, "Артист не найден"})
		return
	}
	if err != nil {
		log.Println(err)
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Не удалось получить данные артиста по ID"})
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := ts.Execute(w, artist); err != nil {
		h.ErrorHand(w, r, errStatus{http.StatusInternalServerError, "Ошибка сервера"})
		return
	}
}
