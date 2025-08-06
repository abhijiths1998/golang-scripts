# Notion Desktop Skeleton

This directory contains a minimal offline Notion‑like application. It bundles a React frontend with an Express backend and packages everything using Electron. The goal is to provide a starting point for a desktop note‑taking tool that stores data locally in SQLite.

## Structure

- `frontend/` – React + TailwindCSS application built with Vite
- `backend/` – Express API exposing basic page and todo routes
- `database/` – SQLite initialization using `better-sqlite3`
- `electron/` – Electron main process scripts

## Development

Install Node.js and run:

```bash
npm install
npm run start
```

This starts the backend, frontend dev server and Electron window concurrently.

## Building

To create a distributable AppImage run:

```bash
npm run build
```

> **Note:** Building requires downloading Electron headers which may fail if network access is restricted.
