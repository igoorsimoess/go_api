package middleware

import (
	"errors"
	"net/http"
	"github.com/igoorsimoess/go_api/api"
	"github.com/igoorsimoess/go_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// response writer is what we use to build the response body, headers, and status code.
		// requests contain all the incoming information from the request

		// the logic for accepting or not the http request
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		// if empty, return an error:
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// queries the database	
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return 
		}

		next.ServeHTTP(w, r) // calls middleware in line or the HandlerFunc if there's no middleware left
		

	})
}
