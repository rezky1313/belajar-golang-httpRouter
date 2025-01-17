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
Params
-httprouter.Handle memiliki parameter ketiga yaitu params.
-params merupakan tempat untuk menyimpan parameter yang dikirim dari client
-namin parameter ini bukan quety parameter, melainkan parameter di URL

-Kadang kita butuh membuat URL yang tidak fix, alias bisa berubah ubah, misal product/1, product/2 dst
-ServeMux tida mendukung hal tersebut, namu Router mendukung hal tersebut
-parameter yang dinamis y ang terdapat di URL, secara otomatis dikumpulkan di params
-namn agar router tahu, kita harus memberitahu ketikan menambahkan route, dibagian mana kita akan membuat URL
pathnya menjadi dinamis

*/

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result().Body
	bytes, _ := io.ReadAll(response)
	assert.Equal(t, "Product 1", string(bytes))
}
