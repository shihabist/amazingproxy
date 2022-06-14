package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"proxy/controller"
	"proxy/repository"
	"proxy/service"
	"testing"
)

func TestProxyHandler(t *testing.T) {
	cases := []struct {
		method             string
		uri                string
		expectedStatusCode int
	}{
		{"GET", "/company", http.StatusOK},
		{"POST", "/company", http.StatusOK},
		{"PUT", "/company", http.StatusUnauthorized},
		{"GET", "/company/abc78dsds", http.StatusOK},
	}
	var (
		fileRepo        = repository.NewFileRepository()
		proxyService    = service.NewProxyService(fileRepo)
		proxyController = controller.NewProxyController(proxyService)
	)
	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.uri, nil)
		w := httptest.NewRecorder()
		proxyController.ProxyHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tc.expectedStatusCode {
			fmt.Println(tc)
			t.Errorf("expected status code to be %d got %d", http.StatusOK, res.StatusCode)
		}
	}
}
