# Raw HTTP Server ⚡

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![TCP Sockets](https://img.shields.io/badge/TCP-Sockets-336699?style=for-the-badge)](#)
[![Zero Dependency](https://img.shields.io/badge/Dependency-Zero-FFD700?style=for-the-badge)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

A hyper-lightweight, zero-dependency HTTP/1.1 server written in pure Go. Instead of relying on Go's black-box `net/http` package or third-party web frameworks, this project is built entirely from the ground up over raw TCP sockets (`net.Listen`).

It handles connection multiplexing, concurrent routing, and strict Layer 7 protocol parsing using purely the Go standard library, showcasing the bare-metal mechanics of web infrastructure.

## 🚀 The Tech Real: How it Works

Under the hood, this server interacts directly with the network layer to bypass typical overheads:

- **TCP Socket Management:** Binds directly to the operating system's network stack, accepting connections and managing memory buffers manually.
- **Massive Concurrency:** Wraps every incoming socket connection in an isolated Goroutine, allowing the server to handle thousands of simultaneous requests without blocking the main thread.
- **Token Bucket Middleware:** Protected by a custom, thread-safe rate limiting engine utilizing lazy evaluation math and asynchronous garbage collection to prevent DDoS vectors.
- **Custom HTTP Parser:** Utilizes buffered I/O (`bufio`) to safely extract the Request Line, construct a hash map of headers, dynamically calculate `Content-Length`, and safely extract payload bodies without memory ghosting.
- **$O(1)$ Multiplexer:** Implements a flat-map routing table. Instead of nested switch statements, routes are concatenated (e.g., `GET /`) and looked up instantly in memory, providing true $O(1)$ execution time for all handlers.

## ⚡ Tech Stack

- **Backend System:** Pure `Go`.
- **Networking Core:** `net` package for raw socket listen and dial commands.
- **I/O Engine:** `bufio` and `io` interfaces for robust byte-stream parsing.

## 🛠️ Quick Start

Since this server is compiled directly into raw machine code, it requires zero external dependencies to run.

```bash
# Clone the repository
git clone https://github.com/gummybearansh/raw-http-server.git
cd raw-http-server

# Compile the binary
go build -o server .

# Boot the engine
./server
```
_The server will immediately bind to `127.0.0.1:2000` and begin accepting traffic._

## 🎮 Interacting with the System

The router is pre-configured to handle several demonstration endpoints.

- **`GET /`** : Returns a standard HTML 200 OK response.
- **`GET /about`** : Returns an architecture breakdown.
- **`POST /`** : Accepts data payloads and echoes the raw HTTP body back to the client.
- **Rate Limit Trigger**: Exceeding 3 requests per second will trigger a raw socket `429 Too Many Requests` hard-drop.

**Testing the POST Handler:**
You can test the raw body ingestion via terminal:
```zsh
curl -X POST http://127.0.0.1:2000/ \
     -H "Content-Type: text/plain" \
     -d "Testing raw socket ingestion."
```

<br>
<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=timeGradient&height=120&section=footer&text=Built%20by%20Gummybearansh:%20Building%20uncompromising%20backend%20infrastructure&fontSize=24&fontAlignY=38&animation=fadeIn" width="100%"/>
</div>
