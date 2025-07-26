# Notion-like Electron App

This example demonstrates a very lightweight Notion-style application. It uses a Go backend for storing pages and todos in memory and an Electron front‑end so the app can be used locally on Linux systems such as Arch.

## Building the Go server

Ensure Go (1.21 or newer) is installed then run:

```bash
go run main.go
```

The API listens on `http://localhost:8081`.

## Running the Electron front‑end

Inside the `frontend` directory install the node dependencies and start Electron:

```bash
cd frontend
npm install
npm run start
```

This opens the desktop application window. Pages are stored in memory only, so data is lost when the server stops.

This is just a minimal proof of concept implementing pages, markdown notes and simple todo lists along with sub‑pages. Markdown can be edited and saved with a live preview shown next to the editor. It can be extended to support additional features such as tables, calendars and persistent storage.
