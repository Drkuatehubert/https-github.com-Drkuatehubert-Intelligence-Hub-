# Intelligence Hub Monolith

An all-in-one geospatial and security intelligence command center, merging WorldMonitor, GeoSentinel, WireTapper, PentAGI, and PentestGPT.

## 🚀 Overview

The Intelligence Hub provides a unified platform for real-time situational awareness, autonomous security testing, and global signal intelligence.

### 🌟 Features

*   🌍 **Unified Strategic Interface**: Fused UI/UX combining WorldMonitor, GeoSentinel, and Palantir styles.
*   🖱️ **Sidebar Navigation**: Instant access to Strategic Map, Signals Intel, Security Ops, and Geo Analytics.
*   ✈️ **Real-time Tracking**: Live global flight tracking (OpenSky) and vessel monitoring (AISStream).
*   🛡️ **Autonomous Pentesting**: Integrated PentAGI/PentestGPT terminal for AI-driven security operations.
*   📡 **Signal Intelligence**: WireTapper-powered passive wireless detection and emission tracking.
*   🤖 **AI Gateway**: Unified gateway for local (Ollama) and cloud (Groq/OpenRouter) AI synthesis.
*   📰 **Intelligence Flux**: 160+ API feeds with 2-minute refresh cycles and live strategic webcams.

## 🏗️ Architecture

*   **Backend**: Unified Go (Gin-based) monolith serving as the central intelligence orchestrator.
*   **Frontend**: React + Vite + BlueprintJS dashboard with modular tabbed architecture.
*   **Storage**: PostgreSQL with `pgvector` for memory, Redis for caching.
*   **Deployment**: Docker Compose for easy scaling and isolation.

## 🚀 Quick Start (For Everyone)

The Intelligence Hub is designed to be launched with **one click**. No coding knowledge required.

### 1. Requirements
- Install **Docker Desktop**: [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/)

### 2. Launch the Hub
- **Windows**: Double-click on `start.bat`
- **Linux/Mac**: Run `bash start.sh` in your terminal

*The platform will be ready in a few minutes at **`http://localhost:3000`**.*

---

## 🛠️ Advanced Setup

### 1. Environment Configuration
The project comes pre-configured with essential API keys in the `.env` file (Wigle, Gemini, Grok, OpenSky, Groq, OpenAI). You can edit this file to add your own keys or modify settings.

### 2. Manual Launch via Docker
```bash
docker-compose up --build
```

### 4. Manual Launch (Development)
If you prefer running without Docker:

**Backend (Go):**
```bash
cd backend
go mod tidy
go run cmd/pentagi/main.go
```

**Frontend (React):**
```bash
cd frontend
npm install
npm run dev
```
*Access via `http://localhost:5173`.*

## 📚 Documentation

Detailed guides for each module are available in the `/docs` directory.

## ⚖️ License

AGPL-3.0-only
