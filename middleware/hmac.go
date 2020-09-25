package middleware

import (
	"net/http"

	"github.com/dkvilo/imcargo/functions"
	"github.com/julienschmidt/httprouter"
)

// VerifyHmac - checks if client is authenticated
func VerifyHmac(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.URL.Query().Get("accessToken") != "" {
			if ok := functions.ValidMAC("Don't f---ing talk to me", r.URL.Query().Get("accessToken"), "secret"); ok {
				next(w, r, p)
			} else {
				http.Error(w, "Authenticated failed", http.StatusNonAuthoritativeInfo)
			}
		} else {
			http.Error(w, "accessToken is missing", http.StatusNotAcceptable)
		}
	}
}