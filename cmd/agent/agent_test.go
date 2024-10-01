package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MainTest(t *testing.T) {
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
			//request := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			main()

			res := w.Result()
			// проверка кода ответа
			assert.Equal(t, test.want.code, res.StatusCode, res.StatusCode)
		})
	}
}
