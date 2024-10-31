package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Handler http.HandlerFunc
	Method  string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	s.Handlers[path] = Handler{
		handler,
		method,
	}
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		switch handler.Method {
		case "GET":
			s.Router.Get(path, handler.Handler)
		case "POST":
			s.Router.Post(path, handler.Handler)
		case "PUT":
			s.Router.Put(path, handler.Handler)
		case "DELETE":
			s.Router.Delete(path, handler.Handler)
		default:
			s.Router.Head(path, handler.Handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
