package main

import (
	"log"
	"net/http"
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

// Отображение содержимого заметки
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
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

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
