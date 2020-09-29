package middleware

import (
	"io"
	"net/http"
	"os"

	"github.com/dkvilo/imcargo/core"
	"github.com/dkvilo/imcargo/functions"
	"github.com/dkvilo/imcargo/model"
	"github.com/julienschmidt/httprouter"
)

// VerifyHmac - checks if client is authenticated
func VerifyHmac(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "text/json")
		if r.URL.Query().Get("accessToken") != "" {
			if ok := functions.ValidMAC(os.Getenv("HMAC_MESSAGE"), r.URL.Query().Get("accessToken"), os.Getenv("HMAC_SECRET")); ok {
				next(w, r, p)
			} else {
				w.WriteHeader(http.StatusNotAcceptable)
				io.WriteString(w, string(
					core.Response(
						model.ImageObject{
							Success: false,
							Message: "Authenticated failed",
						},
					),
				))
				return
			}
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			io.WriteString(w, string(
				core.Response(
					model.ImageObject{
						Success: false,
						Message: "accessToken is missing",
					},
				),
			))
			return
		}
	}
}
