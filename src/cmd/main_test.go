package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	server "github.com/ramaureirac/softserve-work/src/server"
)

func TestURLInfoService(t *testing.T) {

	os.Setenv("GIN_MODE", "testing")
	defer os.Unsetenv("GIN_MODE")

	router := server.NewRouterApp()

	t.Run("GET - Check Malware URL", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/urlinfo/hecker.info/dolphin.exe", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var body map[string]any
		json.Unmarshal(w.Body.Bytes(), &body)

		if body["scan"] != true {
			t.Errorf("Should be marked as malware, result was: %v", body["scan"])
		}
	})

	t.Run("GET - Check Clean URL", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/urlinfo/google.com:443/maps", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var body map[string]any
		json.Unmarshal(w.Body.Bytes(), &body)

		if body["scan"] != false {
			t.Errorf("Should be marked as clean, result was: %v", body["scan"])
		}
	})

	t.Run("POST - Add dynamic URL", func(t *testing.T) {
		reqPost, _ := http.NewRequest("POST", "/urlinfo/munivillalemana.gob.cl/test", nil)
		wPost := httptest.NewRecorder()
		router.ServeHTTP(wPost, reqPost)

		if wPost.Code != http.StatusCreated {
			t.Errorf("Error when creating resources, http code was: %d", wPost.Code)
		}

		reqGet, _ := http.NewRequest("GET", "/urlinfo/munivillalemana.gob.cl/test", nil)
		wGet := httptest.NewRecorder()
		router.ServeHTTP(wGet, reqGet)

		var body map[string]interface{}
		json.Unmarshal(wGet.Body.Bytes(), &body)

		if body["scan"] != true {
			t.Error("Unable to update via POST")
		}
	})
}
