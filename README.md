# 🚢 PortSync

**PortSync** is a high-performance Golang microservice for efficiently processing and storing port data. It reads large JSON files in a streaming fashion and upserts ports into an in-memory database, ensuring optimal performance with limited resources.

## ⚡ Features
- **Efficient JSON Processing** → Streams large JSON files **without loading everything into memory**.
- **Upsert Port Data** → Creates or updates port records dynamically.
- **Memory-Efficient Storage** → Uses **an in-memory database** to optimize speed.
- **Graceful Shutdown** → Handles system signals (`SIGTERM`, `SIGINT`) properly.
- **Hexagonal Architecture** → Follows **clean architecture** principles for maintainability.
- **Dockerized** → Easily deployable via **Docker**.

---

## 🚀 Getting Started

### 📦 Prerequisites
- **Go 1.19+**
- **Docker** (optional, if running in a container)

### 🛠 Installation
Clone the repository:
```sh
git clone https://github.com/vmlellis/port-sync.git
cd port-sync
```

Install dependencies:
```sh
go mod tidy
```

---

## ▶️ How to Run

### **Run Locally (Go)**
```sh
go run src/cmd/app/main.go
```

### **Run with Docker**
```sh
docker build -t port-sync .
docker run --rm -p 8080:8080 -v $(pwd)/data:/root/data -v $(pwd)/config.toml:/config.toml --name port-sync-container port-sync
```

---

## 🔧 Configuration
PortSync uses a configuration file (`config.toml`) to manage settings. Below is an example of the file:

```toml
# Path to the ports.json file
ports_file = "data/ports.json"

# Whether to load the ports.json file on startup
load_on_startup = true
```

- **`ports_file`**: Defines the path to the JSON file containing port data.
- **`load_on_startup`**: If `true`, the application loads ports from `ports_file` at startup.

By default, the application looks for `config.toml` in the root directory.

---

## 🤮 Running Tests
```sh
go test ./...
```

---

## 📡 API Endpoints & Usage

### **1️⃣ Get a Port**
Retrieve a port by its ID:
```sh
curl -X GET "http://localhost:8080/ports/PORT_ID"
```
✔ **Example Response (200 OK):**
```json
{
    "id": "PORT_ID",
    "name": "Port Name",
    "city": "City Name",
    "country": "Country Name",
    "coordinates": [12.34, 56.78],
    "province": "Province Name",
    "timezone": "UTC+1",
    "unlocs": ["UNLOC1"],
    "code": "1234"
}
```

### **2️⃣ Bulk Upsert Ports**
Upload a new `ports.json` file to update all ports:
```sh
curl -X POST "http://localhost:8080/ports/bulk" \
     -H "Content-Type: application/json" \
     --data-binary @data/ports.json
```
✔ **Example `ports.json`**
```json
{
    "PORT1": {
        "name": "New Port 1",
        "city": "New City",
        "country": "New Country",
        "coordinates": [10.10, 20.20],
        "province": "New Province",
        "timezone": "UTC+2",
        "unlocs": ["NEW1"],
        "code": "5678"
    },
    "PORT2": {
        "name": "New Port 2",
        "city": "Another City",
        "country": "Another Country",
        "coordinates": [30.30, 40.40],
        "province": "Another Province",
        "timezone": "UTC+3",
        "unlocs": ["NEW2"],
        "code": "6789"
    }
}
```
✔ **Example Response (201 Created)**
```json
{
    "message": "Ports uploaded successfully"
}
```

✔ **If JSON Format is Invalid (400 Bad Request)**
```json
{
    "error": "Invalid JSON format"
}
```

---

## 🏰 Project Structure

```
port-sync/
│-- src/
│   ├-- cmd/                  # Entry point for the application
│   │   ├-- app/
│   │   │   └-- main.go       # Initializes dependencies and starts the service
│   ├-- internal/
│   │   ├-- domain/           # Business logic (contract, entity, service)
│   │   │   ├-- contract/
│   │   │   │   └-- port_service.go     # Interface for the PortService
│   │   │   │   └-- port_repository.go  # Interface for repositories
│   │   │   ├-- entity/
│   │   │   │   └-- port.go           # Defines the Port entity
│   │   │   ├-- service/
│   │   │       └-- port_service.go        # Implements PortService
│   │   │       └-- port_service_test.go   # PortService unit testing
│   │   ├-- adapters/                 # Adapters (input, output, storage)
│   │   │   ├-- storage/
│   │   │   │   └-- memory_store.go       # In-memory store implementation
│   │   │   │   └-- memory_store_test.go  # In-memory store unit testing
│   │   │   ├-- file/
│   │   │   │   └-- json_loader.go        # JSON file processor
│   │   │   │   └-- json_loader_test.go   # JSON file processor unit testing
│   │   │   ├-- http/
│   │   │       └-- http_handler.go       # HTTP API handlers
│   │   │       └-- http_handler_test.go  # HTTP API handlers unit testing
│   │   │   ├-- config/
│   │   │       └-- config_loader.go  # Config Loader
│-- data/                             # Stores data files
│   └-- ports.json                    # Initial port data
│-- .gitignore                        # Ignore files (bin, build artifacts, etc.)
│-- go.mod                            # Go Modules file
│-- go.sum                            # Go Modules dependencies
│-- config.toml                       # Configuration settings
├-- Dockerfile                        # Docker configuration
│-- README.md                         # Instructions to run the project
```

---

## 🛡 Graceful Shutdown & Signals
PortSync **catches system signals** (`SIGTERM`, `SIGINT`) to ensure:
- No data loss.
- Proper resource cleanup before termination.

---

## 🐝 Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

---

## 🌿 License
This project is licensed under the **MIT License**.

---

🚀 **PortSync is optimized for high-performance and clean code.** Feel free to contribute!
🔗 **GitHub Repo**: [github.com/vmlellis/port-sync](https://github.com/vmlellis/port-sync)


