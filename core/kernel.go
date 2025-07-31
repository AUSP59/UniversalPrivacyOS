
package core

import (
    "crypto/sha256"
    "fmt"
    "io"
    "os"
    "runtime"
)

func InitializeKernel() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println("🧠 Kernel initialized.")
    go initScheduler()
    go verifySelfIntegrity()
}

func initScheduler() {
    for {
        // Simulated task loop (could manage async modules)
    }
}

func verifySelfIntegrity() {
    binaryPath := os.Args[0]
    file, err := os.Open(binaryPath)
    if err != nil {
        fmt.Println("⚠️ Cannot open binary:", err)
        return
    }
    defer file.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        fmt.Println("⚠️ Hashing error:", err)
        return
    }

    checksum := fmt.Sprintf("%x", hash.Sum(nil))
    fmt.Println("🔒 Binary SHA-256:", checksum)
}
