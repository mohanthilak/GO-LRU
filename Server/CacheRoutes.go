package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SetKeyValuePair struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (s *HTTPServer) HandleGetValueFromKey(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	log.Println("!!1", key)
	value := s.app.GetValueFromKey(key)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": value})
}

func (s *HTTPServer) HandleSetKeyValuePair(w http.ResponseWriter, r *http.Request) {
	var KeyValuePair SetKeyValuePair

	err := json.NewDecoder(r.Body).Decode(&KeyValuePair)

	if err != nil {
		log.Println("error while decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("\n!!!!!!!!!!%+v\n", KeyValuePair)
	s.app.SetKeyValuePair(KeyValuePair.Key, KeyValuePair.Value)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "Added Key Value Pair"})
}
