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

## 🛠️ Getting Started

Follow these steps to deploy your local Intelligence Hub:

### 1. Repository Setup
```bash
git clone https://github.com/your-repo/intelligence-hub.git
cd intelligence-hub
```

### 2. Environment Configuration
Copy the template and fill in your API keys (Wigle, OpenSky, Groq, etc.):
```bash
cp .env.example .env
```

### 3. Launch via Docker (Recommended)
This will start the Go backend, React frontend, PostgreSQL (pgvector), Redis, and Ollama:
```bash
docker-compose up --build
```
*The dashboard will be available at `http://localhost:3000`.*

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
