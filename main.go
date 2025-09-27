package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:   pomo <duration>(s|m|h)")
		fmt.Println("Example: pomo 5s  (for 5 seconds)")
		fmt.Println("         pomo 2m  (for 2 minutes)")
		fmt.Println("         pomo 1h  (for 1 hour)")
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	durationStr := os.Args[1]

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Printf("Error parsing duration '%s': %v\n", durationStr, err)
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	total := int(duration.Seconds())
	seconds := 0

	for {
		select {
		case <-ticker.C:
			seconds++

			fmt.Printf("\r⏱️ %d / %d seconds", seconds, total)
			fmt.Printf("\x1b]0;%d\x07", seconds) // title

			if seconds >= total {
				fmt.Printf("\a\n") // notification
				return
			}

		case <-stop: // signal
			fmt.Printf("\n")
			return
		}

	}
}
