package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemsId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id") + " Items Id : " + params.ByName("itemsId")
		fmt.Fprint(writer, text)
	})
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/2", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Items Id : 2", string(body))

}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Image :" + params.ByName("image")
		fmt.Fprint(writer, text)
	})
	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image :/small/profile.png", string(body))

}
