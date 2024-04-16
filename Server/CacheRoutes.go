package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SetKeyValuePair struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

func (s *HTTPServer) HandleGetValueFromKey(w http.ResponseWriter, r *http.Request) {
	key, err := strconv.Atoi(mux.Vars(r)["key"])
	if err != nil {
		log.Println("error while decoding params", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	value := s.app.GetValueFromKey(key)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]int{"key": key, "value": value})
}

func (s *HTTPServer) HandleSetKeyValuePair(w http.ResponseWriter, r *http.Request) {
	var KeyValuePair SetKeyValuePair

	err := json.NewDecoder(r.Body).Decode(&KeyValuePair)

	if err != nil {
		log.Println("error while decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.app.SetKeyValuePair(KeyValuePair.Key, KeyValuePair.Value)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "Added Key Value Pair"})
}
