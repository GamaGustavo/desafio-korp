package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "korp_http_requests_total",
			Help: "Número total de requisições HTTP processadas",
		},
		[]string{"endpoint", "method", "status"}, 
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

type KorpResponse struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	requestCount.WithLabelValues("/projeto-korp", r.Method, "200").Inc()

	currentTime := time.Now().UTC().Format(time.RFC3339)

	resp := KorpResponse{
		Nome:    "Projeto Korp",
		Horario: currentTime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("🚀 Servidor Korp rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}