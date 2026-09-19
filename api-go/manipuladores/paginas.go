package manipuladores

import (
	"html/template"
	"net/http"
)

var templates, _ = template.ParseGlob("./templates/*tmpl.html")

func InjetarRotasPage(roteador *http.ServeMux) {
	roteador.HandleFunc("GET /", PageIndex)
	roteador.HandleFunc("GET /sobre", PageSobre)
	roteador.HandleFunc("GET /documentacao", PageDocumentacao)
}

func PageIndex(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(res, "index", nil); err != nil {
		http.Error(res, "não foi possível renderizar a página", http.StatusInternalServerError)
		return
	}
}

func PageSobre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(res, "sobre", nil); err != nil {
		http.Error(res, "não foi possível renderizar a página", http.StatusInternalServerError)
		return
	}
}

func PageDocumentacao(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(res, "documentacao", nil); err != nil {
		http.Error(res, "não foi possível renderizar a página", http.StatusInternalServerError)
		return
	}
}
