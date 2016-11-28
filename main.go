package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/nii236/gopher-resume/aaa"
)

type JSONRequest struct {
	InvoiceNr time.Time `json:"invoice-nr"`
	Author    string    `json:"author"`
	City      string    `json:"city"`
	From      []string  `json:"from"`
	To        []string  `json:"to"`
	VAT       int       `json:"VAT"`
	Service   []struct {
		Description string      `json:"description"`
		Price       int         `json:"price"`
		Details     interface{} `json:"details,omitempty"`
	} `json:"service"`
	Closingnote string `json:"closingnote"`
	Currency    string `json:"currency"`
	Lang        string `json:"lang"`
	Seriffont   string `json:"seriffont"`
	Sansfont    string `json:"sansfont"`
	Fontsize    string `json:"fontsize"`
	Geometry    string `json:"geometry"`
}

type TempJSON struct {
	Key string `json:"key"`
}

func main() {
	fmt.Println("Starting invoicr...")
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/invoice", generateInvoice).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceDate := time.Now().Format("2006-01-02")
	uuid := aaa.Generate(2, "-")
	log.Println(fmt.Sprintf("invoice number: %s-%s", invoiceDate, uuid))
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	//
	// req, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer r.Body.Close()
	// var tempJSON JSONRequest
	// err = json.Unmarshal(req, &tempJSON)
	// if err != nil {
	// 	log.Println(err)
	// }
	// yml, err := yaml.Marshal(tempJSON)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(string(yml))
	// w.Write(yml)
}
