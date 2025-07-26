# Note Taking App

This is a simple note-taking web application written in Go with a small JavaScript frontend. It serves a single HTML page that interacts with the backend to store notes in memory.

## Building and Running

1. Ensure Go is installed (tested with Go 1.21) on your Arch Linux system.
2. Navigate to the `noteapp` directory and run:

```bash
go run main.go
```

The server listens on `http://localhost:8080`.

Open a browser to that address to start taking notes.

Notes are stored in memory only; restarting the server will clear them.
