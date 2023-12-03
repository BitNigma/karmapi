package app

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	CACertFilePath = "certs/4a717aefc643cbe0.crt"
	KEY            = "certs/pvt-key.key"
	Ca             = "certs/ca.key"
	err            error
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
	go func() {
		if err = http.ListenAndServe(":80", http.HandlerFunc(s.RedirectTLS)); err != nil {
			log.Println("can't redirect user", err)
		}
	}()

	if err = http.ListenAndServe(":443", s.router); err != nil {
		log.Println("can't redicrect user, something happening", err)
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
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

}

func (s *APIserver) mainhandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := s.config

		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html", "static/video.html", "static/second.html",
			"static/third.html", "static/karma.html", "static/arcana.html", "static/market.html", "static/img.html", "static/roadmap.html",
			"static/partners.html", "static/community.html", "static/bord.html"}
		mass := []string{"static/index.html", basetpls[0], basetpls[1], basetpls[2], basetpls[3], basetpls[4],
			basetpls[5], basetpls[6], basetpls[7], basetpls[8], basetpls[9], basetpls[10], basetpls[11], basetpls[12], basetpls[13]}

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

func (s *APIserver) RedirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+"localhost"+":443", http.StatusMovedPermanently)
}
