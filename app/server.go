package app

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

// API server
type APIserver struct {
	//config *Config
	router *mux.Router
}

// Create new server
func New() *APIserver {
	return &APIserver{
		//config: config,
		router: mux.NewRouter(),
	}
}

// Start new server
func (s *APIserver) Start() error {
	s.configureRouter()
	log.Println("starting API server")
	return http.ListenAndServe(":"+os.Getenv("PORT"), s.router)
}

func (s *APIserver) configureRouter() {
	router := s.router.StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusFound)
	})
	router.HandleFunc("/404", s.misshandle())
	router.HandleFunc("/", s.mainhandle())
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

}

func (s *APIserver) mainhandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//create html template
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) misshandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tpl, err := template.ParseFiles("static/404.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = tpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
