
package main

import (
    "flag"
    "fmt"
    "os"

    "UniversalPrivacyOS/api"
    "UniversalPrivacyOS/core"
    "UniversalPrivacyOS/gui"
    "UniversalPrivacyOS/net"
    "UniversalPrivacyOS/pkg/encryption"
    "UniversalPrivacyOS/pkg/logging"
)

func main() {
    headless := flag.Bool("headless", false, "Run in headless API-only mode")
    export := flag.Bool("export", false, "Export system status as JSON and exit")
    version := flag.Bool("version", false, "Print version info")
    flag.Parse()

    if *version {
        fmt.Println("UniversalPrivacyOS v1.0.0 - Supreme Overdrive Edition")
        return
    }

    logging.InitLogger()
    encryption.RunSelfTest()
    core.InitializeKernel()
    net.InitFirewall()

    if *export {
        api.ExportStatusToFile()
        os.Exit(0)
    }

    go api.StartAPIServer()

    if !*headless {
        gui.LaunchUI()
    } else {
        fmt.Println("Running in headless mode. API available at http://localhost:8080")
        select {}
    }
}
