package main

import (
    "html/template"
    "log"
    "net/http"
)	
func Home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, "Page non trouvée", http.StatusNotFound)
        return
    }
    tmpl.Execute(w, nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("contact.html")
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erreur execution template : "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/contact", Contact)

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Serveur lancé sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
