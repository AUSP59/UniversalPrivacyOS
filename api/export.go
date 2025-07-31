
package api

import (
    "encoding/json"
    "os"
    "runtime"
    "time"
)

func ExportStatusToFile() {
    status := SystemStatus{
        Time:      time.Now().Format(time.RFC3339),
        OS:        runtime.GOOS,
        CPUs:      runtime.NumCPU(),
        GoVersion: runtime.Version(),
        Uptime:    time.Since(start).Truncate(time.Second).String(),
    }

    f, err := os.Create("system_status.json")
    if err != nil {
        return
    }
    defer f.Close()
    json.NewEncoder(f).Encode(status)
}
