package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Handler представляет собой тип, реализующий интерфейс http.Handler.
type Handler struct{}

func main() {
	// Запускаем HTTP-сервер, прослушивая порт 8080 и используя наш пользовательский обработчик.
	http.ListenAndServe(":8080", &Handler{})
}

// ServeHTTP реализует метод интерфейса http.Handler.
// Он обрабатывает входящие HTTP-запросы и отправляет ответы.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Если URL запроса совпадает с "/somefile"
	if r.URL.Path == "/" {
		// Открываем файл "somefile.txt"
		f, err := os.Open("somefile.txt")
		if err != nil {
			// Если произошла ошибка при открытии файла, возвращаем код 500 и сообщение об ошибке.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte{'\n'})
			return
		}
		defer f.Close()

		// Устанавливаем заголовок Content-Type как "text/plain"
		w.Header().Add("Content-Type", "text/plain")

		// Копируем содержимое файла в тело ответа HTTP.
		_, err = io.Copy(w, f)
		if err != nil {
			// Если произошла ошибка при копировании, выводим сообщение в журнал.
			log.Println("io.Copy:", err)
		}
		return
	}

	// Если URL не совпадает с "/somefile", возвращаем код 404 и сообщение "Not found".
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found\n"))
}
