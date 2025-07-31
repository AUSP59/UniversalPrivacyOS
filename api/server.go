
package api


import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "upos_cpu_usage_percent",
        Help: "Current CPU usage percentage",
    })
    memoryUsed = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "upos_memory_used_mb",
        Help: "Used memory in MB",
    })
)

func init() {
    prometheus.MustRegister(cpuUsage)
    prometheus.MustRegister(memoryUsed)
}

func collectMetrics() {
    go func() {
        for {
            cpu, _ := cpu.Percent(0, false)
            mem, _ := mem.VirtualMemory()
            if len(cpu) > 0 {
                cpuUsage.Set(cpu[0])
            }
            memoryUsed.Set(float64(mem.Used) / 1024 / 1024)
            time.Sleep(5 * time.Second)
        }
    }()
}


import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "runtime"
    "time"
)

type SystemStatus struct {
    Time       string  `json:"time"`
    OS         string  `json:"os"`
    CPUs       int     `json:"cpus"`
    GoVersion  string  `json:"go_version"`
    Uptime     string  `json:"uptime"`
}

var start = time.Now()

func StartAPIServer() {
    http.HandleFunc("/api/status", handleStatus)
    http.HandleFunc("/api/hash", handleHash)
    http.Handle("/metrics", promhttp.Handler())
    collectMetrics()

    fmt.Println("🌐 API server listening on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("API server error:", err)
    }
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
    uptime := time.Since(start).Truncate(time.Second).String()
    status := SystemStatus{
        Time:      time.Now().Format(time.RFC3339),
        OS:        runtime.GOOS,
        CPUs:      runtime.NumCPU(),
        GoVersion: runtime.Version(),
        Uptime:    uptime,
    }
    json.NewEncoder(w).Encode(status)
}

func handleHash(w http.ResponseWriter, r *http.Request) {
    path := os.Args[0]
    f, err := os.Open(path)
    if err != nil {
        http.Error(w, "Could not open binary", 500)
        return
    }
    defer f.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, f); err != nil {
        http.Error(w, "Hashing error", 500)
        return
    }

    w.Write([]byte(hex.EncodeToString(hash.Sum(nil))))
}
