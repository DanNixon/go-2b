package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/DanNixon/go-2b/pkg/powerbox"
	"github.com/DanNixon/go-2b/pkg/types"
	"github.com/gorilla/mux"
)

var pb powerbox.Powerbox

func addHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
}

func powerboxStatus(w http.ResponseWriter, r *http.Request) {
	s, err := pb.Get()
	if err != nil {
		log.Printf("Powerbox status request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		log.Printf("Powerbox status request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	addHeaders(w)
}

func powerboxSet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var st types.Settings
	if err := decoder.Decode(&st); err != nil {
		log.Printf("Powerbox set request failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s, err := pb.Set(st)
	if err != nil {
		log.Printf("Powerbox set request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		log.Printf("Powerbox set request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	addHeaders(w)
}

func powerboxReset(w http.ResponseWriter, r *http.Request) {
	s, err := pb.Reset()
	if err != nil {
		log.Printf("Powerbox reset request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		log.Printf("Powerbox reset request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	addHeaders(w)
}

func powerboxKill(w http.ResponseWriter, r *http.Request) {
	s, err := pb.Kill()
	if err != nil {
		log.Printf("Powerbox kill request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		log.Printf("Powerbox kill request failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	addHeaders(w)
}

func main() {
	port := flag.String("port", "", "Serial port powerbox is connected to")
	address := flag.String("address", "127.0.0.1:8080", "Address to listen on")
	flag.Parse()

	if *port == "" {
		log.Fatal("Serial port must be specified")
	}

	var err error
	pb, err = powerbox.NewSerialPowerbox(*port)
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/status").HandlerFunc(powerboxStatus)
	router.Methods("POST").Path("/set").HandlerFunc(powerboxSet)
	router.Methods("POST").Path("/reset").HandlerFunc(powerboxReset)
	router.Methods("POST").Path("/kill").HandlerFunc(powerboxKill)

	log.Printf("Serving on: %s", *address)
	log.Fatal(http.ListenAndServe(*address, router))
}
