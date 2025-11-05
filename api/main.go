package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/noelruault/nmod/iprange"
)

type validateResponse struct {
	IP    string `json:"ip"`
	Valid bool   `json:"valid"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/validate", validateHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "only GET supported"})
		return
	}

	ipStr := r.URL.Query().Get("ip")
	if ipStr == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "missing ip"})
		return
	}

	writeJSON(w, http.StatusOK, validateResponse{
		IP:    ipStr,
		Valid: iprange.IsValid(ipStr),
	})
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
