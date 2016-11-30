package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
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
	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir("./web/build/"))))
	r.Handle("/invoice/", corsHandler(generateInvoice)) //.Methods("POST")
	http.ListenAndServe(":8080", handlers.CORS()(r))
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// Handle preflight
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

func generateInvoice(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	uuid := aaa.Generate(2, "-")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte(uuid))
}
