# Raw HTTP Server

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![TCP](https://img.shields.io/badge/TCP-Sockets-336699?style=for-the-badge)](#)

> *A zero-dependency HTTP/1.1 server built entirely from scratch over raw TCP sockets.*

## 🏗️ Engineering Showcase
- **TCP Socket Binds:** Bypasses `net/http` to bind directly to the OS network stack.
- **Massive Concurrency:** Wraps every incoming socket in an isolated Goroutine for non-blocking execution.
- **$O(1)$ Multiplexer:** Employs a flat-map memory routing table for instantaneous handler execution.
- **Raw Parsers:** Dynamically extracts headers and body payloads directly from the byte stream via `bufio`.

## ⚙️ Tech Stack
- **Core:** Pure Go
- **Network:** `net.Listen` and raw socket manipulation

## 🚀 Steps to Run
1. `git clone https://github.com/gummybearansh/raw-http-server.git`
2. `cd raw-http-server`
3. `go run .`
4. Test with `curl http://127.0.0.1:2000/`

<br>
<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=rect&color=timeGradient&height=80&text=Built%20by%20Gummybearansh:%20Building%20uncompromising%20backend%20infrastructure&fontSize=18&fontColor=ffffff&fontAlignY=50" width="100%"/>
</div>
