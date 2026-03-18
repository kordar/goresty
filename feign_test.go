package goresty_test

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/kordar/goresty"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Auth") != "AAABB" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Header.Get("Auth2") != "MMM" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	feign := goresty.NewFeign(nil).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			request.SetHeader("Auth2", "MMM")
			return nil
		}).
		Options(func(client *resty.Client) {
			client.SetBaseURL(srv.URL)
		})

	resp, err := feign.Request().SetHeaders(map[string]string{"Auth": "AAABB"}).Get("/")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode())
	}
}

func TestGetBody(t *testing.T) {
	body := struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Auth") != "AAABB" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"msg": "ok", "code": 200})
	}))
	defer srv.Close()

	feign := goresty.NewFeign(nil).Options(func(client *resty.Client) {
		client.SetBaseURL(srv.URL)
	})

	resp, err := feign.Request().
		SetHeaders(map[string]string{"Auth": "AAABB"}).
		SetResult(&body).
		Get("/")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode())
	}
	if body.Msg != "ok" || body.Code != 200 {
		t.Fatalf("unexpected body: %+v", body)
	}
}
