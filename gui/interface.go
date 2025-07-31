
package gui

import (
    "fmt"
    "github.com/rivo/tview"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "time"
)

func LaunchUI() {
    app := tview.NewApplication()
    textView := tview.NewTextView().
        SetDynamicColors(true).
        SetRegions(true).
        SetChangedFunc(func() {
            app.Draw()
        })

    go func() {
        for {
            textView.Clear()
            fmt.Fprintln(textView, "[yellow]UniversalPrivacyOS - System Monitor")
            fmt.Fprintln(textView, "[green]===============================")

            v, _ := mem.VirtualMemory()
            c, _ := cpu.Percent(0, false)

            fmt.Fprintf(textView, "[white]Memory Usage: %.2f%%
", v.UsedPercent)
            fmt.Fprintf(textView, "Total: %v MB | Free: %v MB
", v.Total/1024/1024, v.Free/1024/1024)
            if len(c) > 0 {
                fmt.Fprintf(textView, "CPU Usage: %.2f%%
", c[0])
            }

            fmt.Fprintln(textView, "[blue]Press Ctrl+C to exit.")
            time.Sleep(2 * time.Second)
        }
    }()

    if err := app.SetRoot(textView, true).Run(); err != nil {
        panic(err)
    }
}
