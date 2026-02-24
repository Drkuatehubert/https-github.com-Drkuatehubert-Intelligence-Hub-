#!/bin/bash
# Intelligence Hub - Easy Start Script

echo "🚀 Starting Intelligence Hub..."

if ! command -v docker &> /dev/null
then
    echo "❌ Error: Docker is not installed. Please install Docker and try again."
    exit 1
fi

echo "📦 Building and starting containers (this may take a few minutes)..."
docker-compose up --build -d

echo ""
echo "✅ Intelligence Hub is running!"
echo "🌍 Access the dashboard at: http://localhost:3000"
echo "🛠️  Backend API available at: http://localhost:8080"
echo ""
echo "📝 To see logs, run: docker-compose logs -f"
echo "🛑 To stop, run: docker-compose down"
