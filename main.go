package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func formatSeconds(seconds int) string {
	duration := time.Duration(seconds) * time.Second

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
	} else {
		return fmt.Sprintf("%02d:%02d", minutes, secs)
	}
}

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
	elapsed := 0

	for {
		select {
		case <-ticker.C:
			elapsed++

			elapsedFormatted := formatSeconds(elapsed)
			totalFormatted := formatSeconds(total)
			fmt.Printf("\r⏱️ %s / %s", elapsedFormatted, totalFormatted)
			fmt.Printf("\x1b]0;%s\x07", elapsedFormatted) // title

			if elapsed >= total {
				fmt.Printf("\a\n") // notification
				return
			}

		case <-stop: // signal
			fmt.Printf("\n")
			return
		}

	}
}
