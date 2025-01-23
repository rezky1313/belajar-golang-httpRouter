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

/*
Router Pattern
-jika kita sudah mengetahui bahwa dengan menggunakan router, kita bisa menambahkan params di URL
-sekarang pertanyaannya, bagaimana pattern(pola) pembuatan parameternya

Named Parameter
-Named Parameter adalah pola pembuatan parameter dengan menggunakan nama
-Setiap namma parameter harus diawali dengan : (titik dua), lalu diikiuti dengan nama parameter
contohnya sudah dilakukan pada praktek sebelumnya

Contoh

Pattern					/user/:user
/user/eko				Match
/user/you				Match
/user/eko/profile		Not Match
/user/					Not Match

Catch All Parameter
-Selain named parameter, ada juga yang bernama catch all parameter yaitu menangkap semua parameter
tidak peduli ada berapa slash
-catch all parameter harus diawali dengan * (bintang), lalu diikuti dengan nama parameter
-catch all paraemter harus berada di posisi akhir URL

pattern                 /src/*filepath
/src/                   no match
/src/somefile           match
/src/subdir/somefile    match


*/

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "product " + id + "Item" + itemId
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/product/1/items/13", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result().Body
	body, _ := io.ReadAll(response)

	assert.Equal(t, "Product 1 Item 13", string(body))
}
func TestCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "image :" + image
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result().Body
	body, _ := io.ReadAll(response)

	assert.Equal(t, "image : small/profile.png", string(body))
}
