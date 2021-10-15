package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Главная страница
func home(w http.ResponseWriter, r *http.Request) {
	//Если путь не совпадает с шаблоном "/", клиенту возвращается ошибка 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox"))
}

// Заметки
func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

// Обработчик для создания новой заметки
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}
