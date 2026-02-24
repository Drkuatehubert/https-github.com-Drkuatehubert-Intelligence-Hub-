# Intelligence Hub Monolith

An all-in-one geospatial and security intelligence command center, merging WorldMonitor, GeoSentinel, WireTapper, PentAGI, and PentestGPT.

## 🚀 Overview

The Intelligence Hub provides a unified platform for real-time situational awareness, autonomous security testing, and global signal intelligence.

### 🌟 Features

*   🌍 **Unified Situational Dashboard**: 3D/2D interactive map with 35+ data layers.
*   ✈️ **Real-time Tracking**: Live flight tracking (OpenSky) and vessel monitoring (AISStream).
*   🛡️ **Autonomous Pentesting**: AI-driven security testing and reconnaissance (PentAGI).
*   📡 **Signal Intelligence**: Passive wireless device detection and OSINT fusion (WireTapper).
*   🤖 **AI Gateway**: Integrated local (Ollama) and cloud (Groq) AI for intelligence synthesis.
*   📰 **Global Intelligence Flux**: Real-time RSS feeds and live video streams across 15+ categories.

## 🏗️ Architecture

*   **Backend**: Go (Gin-based) orchestrator.
*   **Frontend**: React + Vite dashboard.
*   **Storage**: PostgreSQL with `pgvector` for memory, Redis for caching.
*   **Deployment**: Docker Compose for easy scaling and isolation.

## 🛠️ Getting Started

1.  Clone the repository.
2.  Set up your `.env` file (see `.env.example`).
3.  Run `docker-compose up --build`.
4.  Access the Hub at `http://localhost:3000`.

## 📚 Documentation

Detailed guides for each module are available in the `/docs` directory.

## ⚖️ License

AGPL-3.0-only
