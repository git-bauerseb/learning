package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Custom HTTP Server
type httpServer struct {
	Log *Log
}

// Object for handling requests
type ProduceObject struct {
	Record Record `json: "record"`
}

type ProduceResponse struct {
	Offset uint64 `json: "offset"`
}

type ConsumeObject struct {
	Offset uint64 `json: "offset"`
}

type ConsumeResponse struct {
	Record Record `json: "record"`
}

func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req ProduceObject

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Printf("Adding: %v; %v\n", req.Record.Value, r.Body)
	off, _ := s.Log.Append(req.Record)
	res := ProduceResponse{Offset: off}
	json.NewEncoder(w).Encode(res)
}

func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ConsumeObject
	json.NewDecoder(r.Body).Decode(&req)
	rec, _ := s.Log.Read(req.Offset)
	fmt.Printf("Requested Offset: %d String: %s\n", req.Offset, string(rec.Value))
	res := ConsumeResponse{Record: rec}
	json.NewEncoder(w).Encode(res)
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

func NewHTTPServer(addr string) *http.Server {
	httpsrv := newHTTPServer()
	r := mux.NewRouter()
	r.HandleFunc("/", httpsrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpsrv.handleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
