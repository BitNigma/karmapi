package app

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	err error
)

// API server
type APIserver struct {
	config *Config
	router *mux.Router
}

// Create new server
func New() *APIserver {
	return &APIserver{
		config: NewConfig(),
		router: mux.NewRouter(),
	}
}

// Start new server
func (s *APIserver) Start() error {

	s.configureRouter()
	if err = http.ListenAndServe(":8080", s.router); err != nil {
		log.Println("can't redirect user", err)
	}
	log.Println("starting server")

	return nil
}

func (s *APIserver) configureRouter() {
	router := s.router.StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusFound)
	})
	router.HandleFunc("/404", s.misshandle())
	router.HandleFunc("/", s.mainhandle())
	router.HandleFunc("/about", s.about())
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
}

func (s *APIserver) mainhandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := s.config

		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html", "static/section-1.html", "static/community.html", "static/section-1.html", "static/section-2.html", "static/section-3.html", "static/section-4.html", "static/section-5.html", "static/roadmap.html", "static/partners.html"}
		mass := []string{"static/neoweb.html", basetpls[0], basetpls[1], basetpls[2], basetpls[3], basetpls[4], basetpls[5], basetpls[6], basetpls[7], basetpls[8], basetpls[9], basetpls[10], basetpls[11]}

		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, &data)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) about() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data := s.config

		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html", "static/section-6.html", "static/section-7.html", "static/section-8.html", "static/section-9.html"}
		mass := []string{"static/about.html", basetpls[0], basetpls[1], basetpls[2], basetpls[3], basetpls[4], basetpls[5], basetpls[6]}

		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, &data)
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
