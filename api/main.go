package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"com.dylanswartz.ddm/api/commands"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithBasicAuth(os.Getenv("DDM_USERNAME"), os.Getenv("DDM_PASSWORD")),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	http.HandleFunc("/", root)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/memory", memory)
	http.HandleFunc("/cpu", processor)
	http.HandleFunc("/reboot", reboot)
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

func reboot(w http.ResponseWriter, r *http.Request) {
	commands.Reboot()
	fmt.Fprintf(w, "Rebooting device!")
}
