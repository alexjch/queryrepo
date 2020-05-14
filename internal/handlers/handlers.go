package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/alexjch/queryrepo/internal/repo"
	"github.com/alexjch/queryrepo/internal/types"
)

func NewRootHandlerV1(repoUrl *url.URL) func(w http.ResponseWriter, r *http.Request) {

	log.Println("Using repoURL", repoUrl)

	return func(w http.ResponseWriter, r *http.Request) {
		var p = new(types.Package)
		var response = new(types.ServiceResponse)

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			response.Success = false
			response.Data = err.Error()
		}

		if strings.Trim(p.Name, " ") == "" {
			response.Success = false
			response.Data = "Missing value for name key"
		} else {
			// Call query on repo
			result, err := repo.QueryRepo(repoUrl, p.Name)
			if err != nil {
				response.Success = false
				response.Data = err.Error()
			} else {
				response.Success = false
				response.Data = *result
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
