package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

//+ 200 OK 					Правильно		curl -v -X POST http://localhost:8080/update/gauge/m1/1 -H 'Content-Type: text/plain; charset=UTF-8'
//+ 200 OK 					Правильно		curl -v -X POST http://localhost:8080/update/counter/m1/1 -H 'Content-Type: text/plain; charset=UTF-8'
//+ 400 BadRequest 			Тип значения 	curl -v -X POST http://localhost:8080/update/gauge/m1/abs -H 'Content-Type: text/plain; charset=UTF-8'
//+ 400 BadRequest 			Тип значения	curl -v -X POST http://localhost:8080/update/counter/m1/abs -H 'Content-Type: text/plain; charset=UTF-8'
//+ 405 Method Not Allowed	Метод GET		curl -v -X GET  http://localhost:8080/update/counter/m1/abs -H 'Content-Type: text/plain; charset=UTF-8'
//+	404 Status Not Found	Без метрики		curl -v -X GET  http://localhost:8080/update/m1/abs -H 'Content-Type: text/plain; charset=UTF-8'

// хранилище метрик
type MemStorage struct {
	gaugeValue   float64
	counterValue int64
}

// Тип gauge, float64 — новое значение должно замещать предыдущее.
// Тип counter, int64 — новое значение должно добавляться к предыдущему, если какое-то значение уже было известно серверу.

// интерфейс работы с хранилищем
type Gauger interface {
	Gaug() error
}

type Counter interface {
	Count() error
}

// метод обработки метрики типа gauge
func (storage *MemStorage) Gaug(gauge float64) error {
	storage.gaugeValue += gauge
	return nil
}

// метод обработки метрики типа gauge
func (storage *MemStorage) Count(count int64) error {
	storage.counterValue += count
	return nil
}

func MainPage(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		data := []byte("Простой POST-метод\r\n")
		res.Write(data)
	} else {
		data := []byte("Простой Не POST-метод\r\n")
		res.Write(data)
	}

	res.Write([]byte("Привет!"))
}

func MetricHandler2(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Gauge POST-метод\r\n")

	var metric, name, value string

	if req.Method == http.MethodPost {
		//data := []byte("Gauge POST-метод\r\n")
		//res.Write(data)

		contentType := res.Header().Get("Content-Type")

		body += fmt.Sprintf("Content-Type: %s\r\n", contentType)

		metric = req.PathValue("metric")
		name = req.PathValue("name")
		value = req.PathValue("value")

		body += fmt.Sprintf("metric: %s\r\n", metric)
		body += fmt.Sprintf("name: %s\r\n", name)
		body += fmt.Sprintf("value: %s\r\n", value)
		body += fmt.Sprintf("TypeOf(value): %s\r\n", reflect.TypeOf(value))

	} else {
		data := []byte("Gauge Не POST-метод\r\n")
		res.Write(data)
	}

	if metric == "gauge" {
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
	} else if metric == "counter" {
		if _, err := strconv.Atoi(value); err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
	}

	/* if name == "" {
		res.WriteHeader(http.StatusNotFound)
	} else if metric != "gauge" || metric != "counter" {
		//res.WriteHeader(http.StatusBadRequest)
	} */

	res.Write([]byte(body))
}

func main() {

	mux := http.NewServeMux()

	//mux.HandleFunc("/", MainPage)

	mux.HandleFunc("POST /update/{metric}/{name}/{value}", MetricHandler2)
	//mux.HandleFunc("POST /update/counter/{name}/{value}", Counter)

	//mux.HandleFunc("POST /update/", MetricHandler)

	//http.HandleFunc("/", MainPage)
	//http.HandleFunc("/update/gauge/", Gauge)
	//http.HandleFunc("/update/counter/", Counter)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
