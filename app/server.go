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
	log.Println("starting API server")
	return http.ListenAndServe(":"+os.Getenv("PORT"), s.router)
	//return http.ListenAndServe(":9000", s.router)
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
			"static/third.html", "static/karma.html", "static/arcana.html", "static/market.html", "static/img.html", "static/roadmap.html", "static/partners.html", "static/community.html"}
		mass := []string{"static/index.html", basetpls[0], basetpls[1], basetpls[2], basetpls[3], basetpls[4],
			basetpls[5], basetpls[6], basetpls[7], basetpls[8], basetpls[9], basetpls[10], basetpls[11], basetpls[12]}

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
