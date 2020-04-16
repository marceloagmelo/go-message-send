package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marceloagmelo/go-message-send/app/handler"
)

const (
	staticDir = "/static/"
)

var subRouter *mux.Router

// App has router and db instances
type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	//subRouter = a.Router.PathPrefix("/go-message-send").Subrouter()

	a.Router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {

	a.Get("/health", a.handleRequest(handler.Health))
	a.Get("/", a.handleRequest(handler.Home))
	a.Get("/new", a.handleRequest(handler.New))
	a.Post("/enviar", a.handleRequest(handler.Enviar))
	a.Delete("/apagar/{id}", a.handleRequest(handler.Apagar))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

//RequestHandlerFunction função handler
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
