package main

import (
	"Util"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHeartbeat(t *testing.T) {
	t.Run("A=Heartbeat", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		router := mux.NewRouter()
		router.HandleFunc("/health", Util.Heartbeat).Methods("GET")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)
		res := rec.Result()
		recbody, err := ioutil.ReadAll(rec.Body)
		fmt.Println(recbody)
		if res.StatusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code :got %v want status ok", res.Status)
		}
		assert.Equal(t, 200, res.StatusCode, "OK response is expected")
	})

}
