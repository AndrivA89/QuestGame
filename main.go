package main

import (
	obj "QuestGame/objects"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", obj.StartPage)        // Обработчик стартовой страницы
	http.HandleFunc("/game", obj.GamePage)     // Обработчик страницы с игрой
	http.HandleFunc("/end", obj.EndGamePage)   // Обработчик последней страницы
	go obj.Open("http://127.0.0.1/")           // Открытие стартовой страницы
	log.Fatal(http.ListenAndServe(":80", nil)) // Запуск сервера
}
