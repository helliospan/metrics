package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test200OK(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "simple test #1",
			values: []int{1, 2},
			want:   3},
	}

	for _, test := range tests {
		//t.Run(name string, f func(t *testing.T)) bool
		//используется для запуска вложенных тестов (подтестов).
		//Первым аргументом передаётся имя подтеста, а вторым — функция, которая
		//будет запущена в отдельной горутине. По умолчанию t.Run() ожидает завершения работы функции
		//использование t.Run необязательно
		t.Run(test.name, func(t *testing.T) {
			t.Errorf(test.name)
		})
	}

}

func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        200,
				response:    `{"status": "ok"}`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/status", nil)
			w := httptest.NewRecorder()
			MetricHandler2(w, request)

			res := w.Result()
			// проверка кода ответа
			assert.Equal(t, test.want.code, res.StatusCode, res.StatusCode)
		})
	}
}
