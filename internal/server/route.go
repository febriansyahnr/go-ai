package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/websocket", s.websocketHandler)

	return mux
}
