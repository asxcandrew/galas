package transport

import (
	"net/http"

	"github.com/asxcandrew/galas/user"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func MakeUserHandler(s user.UserService, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	return r
}
