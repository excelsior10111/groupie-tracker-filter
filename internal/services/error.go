package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type errStatus struct {
	StatusCode   int
	StatusString string
}

func (h *Handler) ErrorHand(w http.ResponseWriter, r *http.Request, status errStatus) {
	fmt.Println(status)
	w.WriteHeader(status.StatusCode)
	file := "./ui/html/error.html"
	ts, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("Ошибка при парсинге файла шаблона ошибки: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Println(status)
	err = ts.Execute(w, status)
	if err != nil {
		log.Printf("Ошибка при выполнении шаблона ошибки: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
