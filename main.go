package main

import (
	"fmt"
	"log"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/dkvilo/imcargo/controller"
	"github.com/dkvilo/imcargo/functions"
	"github.com/dkvilo/imcargo/middleware"
	"github.com/julienschmidt/httprouter"
)


func main() {
	
	mux := httprouter.New()
	ctrl := controller.New()

	mux.Handler("GET", "/", http.NotFoundHandler())
	mux.Handler("GET", "/favicon.ico", http.NotFoundHandler())

	go mux.POST("/upload", middleware.VerifyHmac(ctrl.Upload))


	fmt.Println("accessToken:", functions.GenerateHmac("Don't f---ing talk to me", "secret"))

	mux.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":8080", mux))
}




