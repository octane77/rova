package routers

import (
	"github.com/gorilla/mux"
	"github.com/qor/auth"
)

func AuthRouter(r *mux.Router, authService *auth.Auth) *mux.Router {
	r.Handle("/auth/", authService.NewServeMux())
	return r
}
