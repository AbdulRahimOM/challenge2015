# 🎬 Degrees of Separation - Movie Industry Connections

This Go application finds the degrees of separation between two people in the movie industry using data from Moviebuff. It implements an efficient graph traversal algorithm with concurrent data fetching to determine the shortest path between two industry professionals through their movie collaborations.

## ✨ Features

- 🚀 **Concurrent Data Fetching**: Efficiently fetches data from external APIs using goroutines
- 💾 **In-Memory Caching**: Implements a thread-safe caching mechanism for person and movie data
- 🛡️ **Rate Limiting**: Protects against API throttling with built-in rate limiting
- 📊 **Performance Monitoring**: Includes pprof endpoints for runtime analysis
- ⚡ **Graceful Error Handling**: Robust error handling for API failures and invalid inputs
- ⚙️ **Configuration via Environment Variables**: Flexible configuration through environment variables
- 🔄 **Resource Management**: Proper channel and goroutine lifecycle management

## 🏗️ Architecture

### 🌐 Data Fetching
- Uses worker pools for concurrent data fetching from external APIs
- Implements separate workers for person and movie data
- Controlled concurrency with predefined worker counts

### 📦 Caching
- Thread-safe in-memory cache using maps
- Implements `sync.RWMutex` for concurrent read/write operations
- Caches both person and movie data after fetching

### 📈 Performance & Monitoring
- Pprof endpoints for runtime profiling and debugging
- Periodic logging of goroutine statistics
- Rate limiting to prevent API throttling

### 🛠️ Error Handling & Resource Management
- Context-based cancellation for cleanup
- Proper channel closing mechanisms
- Graceful error handling for API failures
- Existence validation (of target person) to prevent long unnecessary searches

## 🔌 API Endpoints

### GET /separation
Query Parameters:
- `from`: Moviebuff URL of the first person
- `to`: Moviebuff URL of the second person

Example:
```
GET /separation?from=amitabh-bachchan&to=robert-de-niro
```

Response:
```json
{
    "separation": 3
}
```

## ⚙️ Configuration

The application can be configured using the following environment variables:

- `PORT`: Server port (default: 3001)
- `PPROF_PORT`: Port for pprof endpoints
- `LOG_LEVEL`: Logging level (debug/info)
- `RATE_LIMIT`: API rate limit per minute
- `PERSON_DATA_FETCH_WORKERS`: Number of concurrent person data fetchers
- `MOVIE_DATA_FETCH_WORKERS`: Number of concurrent movie data fetchers

## 📥 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/AbdulRahimOM/challenge2015.git
cd challenge2015
```

2. Install dependencies:
```bash
go mod download
```

3. Copy the environment file and configure:
```bash
cp no-secrets.env .env
# Edit .env with your preferred settings (or keep it as it is to run in default settings)
```

## 🚀 Running the Application

1. Set up environment variables (optional)
2. Run the application:
```bash
go run cmd/main.go
```

## 💪 Performance Considerations

1. **Concurrent Data Fetching**
   - 🔄 Optimized worker pools for API requests
   - 👥 Separate workers for person and movie data

2. **Caching**
   - 📦 In-memory caching reduces API calls
   - 🔒 Thread-safe read/write operations

3. **Resource Management**
   - ⚡ Context-based cancellation
   - 🧹 Proper cleanup of resources
   - 🛡️ Rate limiting to prevent throttling

## 📊 Monitoring

### 🔍 Pprof Endpoints
Access pprof endpoints at:
```
http://localhost:{PPROF_PORT}/debug/pprof/
```

Available profiles:
- 🧵 Goroutine
- 💾 Heap
- 🔄 Thread
- 🚫 Block
- 📈 CPU profile

## 🔮 Future Improvements
- ⏰ Add cache expiration mechanism (Relevant, as new movies and persons are added)
- 🔗 Show connection chain along with degree of seperation

## 📚 Dependencies

- 🚀 [Fiber](github.com/gofiber/fiber/v2) - Web framework
- 📦 Standard Go libraries for concurrency and HTTP operations

