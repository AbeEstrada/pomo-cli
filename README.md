<img src="https://github.com/user-attachments/assets/69cb9b4a-ee8b-4c80-9c4d-4b7276c27883" width="512" height="512" alt="Golang Gopher holding a tomato" />

# Pomo: a simple cli timer

A minimal, command line stopwatch/timer written in Go. This tool is useful for quick, focused time tracking like [Pomodoro technique](https://en.wikipedia.org/wiki/Pomodoro_Technique) sessions directly in your terminal.

## Features

- Realtime Progress: Shows the elapsed time versus the total duration in the terminal.
- Title Update: The terminal window/tab title updates every second with the current elapsed time.
- Completion Alert: Sends notification (`\a`) when the timer finishes.
- Stop Anytime: Press Ctrl+C to stop the timer and return to the prompt.

## Build and Installation

This project uses [`just`](https://github.com/casey/just) as a command runner and requires [Go](https://golang.org/) for building. Below are the available commands:

### Available Commands

- `just` or `just install` - Build and install the binary to `PREFIX/bin/` (default: `/usr/local/bin`)
- `just build` - Build the binary in the current directory
- `just uninstall` - Remove the installed binary
- `just clean` - Remove the built binary from the current directory

## Usage

```
pomo <duration>(s|m|h)

Example: pomo 5s  (for 5 seconds)
         pomo 2m  (for 2 minutes)
         pomo 1h  (for 1 hour)
```
