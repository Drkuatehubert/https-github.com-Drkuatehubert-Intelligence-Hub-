@echo off
echo 🚀 Starting Intelligence Hub...

docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Error: Docker is not installed. Please install Docker and try again.
    pause
    exit /b
)

echo 📦 Building and starting containers (this may take a few minutes)...
docker-compose up --build -d

echo.
echo ✅ Intelligence Hub is running!
echo 🌍 Access the dashboard at: http://localhost:3000
echo 🛠️  Backend API available at: http://localhost:8080
echo.
echo 📝 To see logs, run: docker-compose logs -f
echo 🛑 To stop, run: docker-compose down
echo.
pause
