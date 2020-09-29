package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {

	mux := httprouter.New()
	ctrl := controller.New()

	mux.Handler("GET", "/", http.NotFoundHandler())
	mux.Handler("GET", "/favicon.ico", http.NotFoundHandler())

	go mux.POST("/upload", middleware.VerifyHmac(ctrl.Upload))

	fmt.Println("accessToken:", functions.GenerateHmac(os.Getenv("HMAC_MESSAGE"), os.Getenv("HMAC_SECRET")))

	mux.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":8080", mux))
}




