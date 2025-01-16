package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
Router
-inti dari library http Router adalah struct Router
-Router in merupakan iimplementasi dari http.Handler, sehingga bisa dengan mudah menambahkan kedalam http.server
-untuk membuat ROuter, kita bisa menggunakan function httprouter.New(), yang akan mengembalikan router pointer
*/

func main() {
	router := httprouter.New()

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}

/*
HTTP Method
-Router mirip dengan ServerMux, dimana kita bisa menambahkan router kedalam router
jika sebelumnya menggunakan serve Mux kita harus menentukan dahulu route A masuk ke handler yang mana, route/url B
masuk ke handler yang mana
-Kelebihan dibandingkan dengna serverMux adalah pada ROuter, kita bisa menentukan HTTP Method yang ingin kita gunakkan
misal GET, POST, DELETE, PUT Dan lain lain, jadi nantinya akan memudahkan untuk membuat API
-CAra Menambahkan router kedalam router adalah function yang sama dengan HTTP Methodnya, misal router.GET(), router.Post()

httprouter.Handle
-jika sebelumnya pada serveMux menggunakan http.Handler, serkarang pada http Router kita menggunakan type httprouter.Handle
-perbedaan dengan http.handler adalah pada httprouter.handle terdapat parameter ketiga yaitu params/parameter
*/
