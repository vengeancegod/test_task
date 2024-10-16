package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Numbers struct {
	Numbers []float64 `json:"numbers"`
}

func main() {
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибки при открытии лог файла: %v\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	data, err := os.ReadFile("numbers.json")
	if err != nil {
		log.Printf("Ошибка чтения JSON файла: %v\n", err)
		return
	}
	log.Println("JSON файл успешно прочитан")

	var nums Numbers
	err = json.Unmarshal(data, &nums)
	if err != nil {
		log.Printf("Ошибка преобразования JSON: %v\n", err)
		return
	}
	log.Println("JSON преобразование успешно")

	sum := 0.0
	for _, num := range nums.Numbers {
		sum += num
	}
	log.Printf("Сумма чисел: %.2f\n", sum)

	url := "https://habr.com/ru/companies/quadcode/articles/662852/"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Ошибка выполнения HTTP GET запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("HTTP GET запрос успешен. Статус: %d\n", resp.StatusCode)
	} else {
		log.Printf("HTTP GET запрос не удался. Статус: %d\n", resp.StatusCode)
	}

	log.Println("Программа выполнена")
}
