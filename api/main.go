package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	http.HandleFunc("/", root)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/memory", memory)
	http.HandleFunc("/cpu", processor)
	return http.Serve(tun, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from ngrok-go!")
}


func memory(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}

func processor(w http.ResponseWriter, r *http.Request) {
	type info struct {
		Info []cpu.InfoStat
	}
	
	type usage struct {
		Usage []float64
	}
	
	type CPUMetrics struct {
		info
		usage
	}
	c, _ := cpu.Info()
	p, _ := cpu.Percent(0, true)
	combined := CPUMetrics{
		info{c},
		usage{p},
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(combined)
}