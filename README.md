# Redish

> "Mom, can we have Redis?"
> "No, we have Redis at home."
> **Redis at home:**

**Redish** is a lightweight, distributed key-value store written in Go. It is designed as a study in distributed systems fundamentals, focusing on raw TCP communication, concurrency control, and data durability.

**This project is really really early in development.**

## Quick Start

1. **Start the Server**
   ```bash
   git clone [https://github.com/bduckdev/redish](https://github.com/bduckdev/redish)
   cd redish
   go run .
   ```

_You should see: `Redish (v0.1) listening on port 6379...`_

2.  **Connect via Netcat (or Telnet)**
    Open a second terminal:

    ```bash
    nc localhost 6379
    ```

3.  **Behold the Functionality**

    ```text
    SET foo bar
    OK

    GET foo
    bar

    GET notakey
    (nil)
    ```

## Architecture & Implementation

Unlike standard web servers that rely on HTTP, Redish implements a **custom TCP wire protocol** for low-latency communication.

- **Concurrency:** Uses `sync.RWMutex` to handle concurrent reads/writes safely. Multiple readers can access data simultaneously, while writers acquire an exclusive lock.
- **Protocol:** Text-based command protocol (inspired by RESP) over raw TCP.
- **Storage:** In-memory hash map (Go `map`).

## Roadmap

- [x] TCP Server & Thread-safe In-Memory Store
- [ ] Persistence via Write-Ahead Log (WAL)
- [ ] Leader-Follower Replication
- [ ] Sharding & Consistent Hashing
