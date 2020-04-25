package Server

import (
	"errors"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ServerConfig struct {
	addr string
}

type Server struct {
	cfg        ServerConfig
	r          *mux.Router
	getIndexed http.HandlerFunc
	postFiles  http.HandlerFunc
}

func (s *Server) Start() {
	s.init()
	s.appendSwagger()
	s.appendUI()
	go s.serveHTTP()
}

func (s *Server) appendSwagger() {

}

func (s *Server) appendUI() {

}

func (s *Server) serveHTTP() {
	if http.ListenAndServe(s.cfg.addr, s.r) != nil {
		log.Error("Could not start server")
	}
}

func (s *Server) init() {
	s.getIndexed = http.HandlerFunc(handleGet)
	s.postFiles = http.HandlerFunc(handlePost)

	s.r.HandleFunc("/index", s.getIndexed).Methods("GET")
	s.r.HandleFunc("/index", s.postFiles).Methods("POST")
}

func handlePost(writer http.ResponseWriter, request *http.Request) {
	//Should they upload files or should an "agent" be running locally?

}

func handleGet(writer http.ResponseWriter, request *http.Request) {
	path := GetHttpParam("path", request)
	if path == "" {
		BadRequest(errors.New("must have path"), writer, request)
	}
}
