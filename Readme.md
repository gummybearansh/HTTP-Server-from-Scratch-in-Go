# Raw HTTP Server

> A hyper-lightweight, zero-dependency HTTP/1.1 server written in pure Go.

Instead of relying on Go's black-box `net/http` package or third-party web frameworks, this project is built entirely from the ground up over raw TCP sockets (`net.Listen`). It handles connection multiplexing, concurrent routing, and strict Layer 7 protocol parsing using purely the Go standard library.

## ⚡ Core Architecture

- **TCP Socket Management:** Binds directly to the operating system's network stack, accepting connections and managing memory buffers manually.
- **Massive Concurrency:** Wraps every incoming socket connection in an isolated Goroutine, allowing the server to handle thousands of simultaneous requests without blocking the main thread.
- **Token Bucket Middleware:** Protected by a custom, thread-safe rate limiting engine utilizing lazy evaluation math and asynchronous garbage collection to prevent DDoS vectors.
- **Custom HTTP Parser:** Utilizes buffered I/O (`bufio`) to safely extract the Request Line, construct a hash map of headers, dynamically calculate `Content-Length`, and safely extract payload bodies without memory ghosting.
- **$O(1)$ Multiplexer:** Implements a flat-map routing table. Instead of nested switch statements, routes are concatenated (e.g., `GET /`) and looked up instantly in memory, providing $O(1)$ execution time for all handlers.

## 🛠️ Usage & Installation

Since this server is compiled directly into raw machine code, it requires zero external dependencies to run.

**1. Clone the repository**

```bash
git clone [https://github.com/gummybearansh/raw-http-server.git](https://github.com/gummybearansh/raw-http-server.git)
cd raw-http-server
```

**2. Compile the binary**

```bash
go build -o server .
```

**3. Boot the engine**

```bash
./server
```

_The server will immediately bind to `127.0.0.1:2000` and begin accepting traffic._

## 📡 API Endpoints

The router is currently wired to handle the following demonstrations:

- `GET /` : Returns a standard HTML 200 OK response.
- `GET /about` : Returns the architecture breakdown.
- `POST /` : Accepts data payloads and echoes the raw HTTP body back to the client.
- _Rate Limit Trigger_: Exceeding 3 requests per second will trigger a raw socket `429 Too Many Requests` hard-drop.
- _Wildcard_ : Automatically triggers a strict `404 Not Found` or `400 Bad Request` protocol response for invalid routing.

### Testing the POST Handler

You can test the raw body ingestion using `curl`:

```zsh
curl -X POST [http://127.0.0.1:2000/](http://127.0.0.1:2000/) \
     -H "Content-Type: text/plain" \
     -d "Testing raw socket ingestion."
```

## 👤 Author

**Ansh Lachhwani**

- Building uncompromising backend infrastructure.
