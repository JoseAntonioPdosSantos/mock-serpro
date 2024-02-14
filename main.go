package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func main() {
	mux := chi.NewRouter()
	mux.MethodFunc(http.MethodGet, "/v2/condutores/{cpf}/{nome}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context().Value(chi.RouteCtxKey).(*chi.Context)

		params := make(map[string]string)

		for i, key := range ctx.URLParams.Keys {
			params[key] = ctx.URLParams.Values[i]
		}

		cpf := params["cpf"]
		if len(cpf) < 11 {
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}
		firstDigit, _ := strconv.Atoi(string(cpf[0]))

		isPositive := false
		if firstDigit%2 == 0 {
			isPositive = true
		}
		res := responseDto{PossuiCadastroPositivo: isPositive}

		dto, _ := json.Marshal(&res)

		w.WriteHeader(http.StatusOK)
		w.Write(dto)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("could not start the server: %w", err))
	}

}

type responseDto struct {
	PossuiCadastroPositivo bool `json:"possuiCadastroPositivo"`
}