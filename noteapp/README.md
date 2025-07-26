# Note Taking App

This is a simple note-taking web application written in Go with a small JavaScript frontend. It serves a single HTML page that interacts with the backend to store notes in memory.

## Building and Running

1. Ensure Go is installed (tested with Go 1.21) on your Arch Linux system. You
   can install it on Arch Linux using:

   ```bash
   sudo pacman -S go
   ```
2. Navigate to the `noteapp` directory and run:

```bash
go run main.go
```

To compile a standalone binary instead, ensure `make` is available and use:

```bash
make build
./noteapp
```

If you receive an error like `go: No such file or directory`, ensure the `go`
binary is in your `PATH` or specify it explicitly when running make:

```bash
GOCMD=/path/to/go make build
```

The server listens on `http://localhost:8080`.

Open a browser to that address to start taking notes.

Notes are stored in memory only; restarting the server will clear them.
