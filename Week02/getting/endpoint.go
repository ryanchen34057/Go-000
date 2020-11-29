package getting

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func MakeGetUserEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		id := query.Get("id")
		user, err := s.GetUser(id)
		if err != nil {
			fmt.Printf("%v", errors.Cause(err))
			fmt.Printf("%+v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
