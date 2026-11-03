# 8-Week Go (Golang) Roadmap
### From Core Syntax to Low-Level System Design

Each week has: **Learning Objectives**, a **Concept Checklist**, and an **End-of-Week Project** (non-web, terminal/systems-focused) with its own build checklist.

---

## Week 1 — Core Syntax & Foundations

**Objectives:** Get fluent in Go's basic building blocks and its philosophy (simplicity, explicitness, no exceptions).

**Concept Checklist**
- [ ] Install Go, set up `GOPATH`/modules (`go mod init`), understand workspace layout
- [ ] Variables, constants, `iota`, type inference
- [ ] Primitive types, type conversion
- [ ] Control flow: `if`, `for` (Go's only loop), `switch`
- [ ] Functions: multiple return values, named returns, variadic args
- [ ] Arrays vs. slices (capacity/length semantics), maps
- [ ] Basic error handling (`error` interface, `errors.New`, `fmt.Errorf`)
- [ ] `go fmt`, `go vet`, and idiomatic naming conventions

**Project: CLI Text Analyzer**
- [ ] Read a `.txt` file from disk via CLI argument
- [ ] Count words, lines, characters, and word-frequency (top 10)
- [ ] Handle malformed input / missing file with proper error returns (no panics)
- [ ] Print a formatted report to stdout

---

## Week 2 — Structs, Interfaces & Idiomatic Error Handling

**Objectives:** Learn Go's approach to OOP-adjacent design: composition over inheritance, small interfaces.

**Concept Checklist**
- [ ] Structs, struct embedding (composition), tags
- [ ] Methods (value vs. pointer receivers) and when to use each
- [ ] Interfaces, implicit satisfaction, empty interface (`any`)
- [ ] Type assertions and type switches
- [ ] Custom error types, wrapping errors (`errors.Is`, `errors.As`, `%w`)
- [ ] Package organization and visibility (exported vs. unexported)

**Project: Inventory / Warehouse Management CLI**
- [ ] Define `Item` struct and an `Inventory` interface (`Add`, `Remove`, `Report`)
- [ ] Implement at least two inventory backends (e.g., in-memory slice-based vs. map-based) satisfying the same interface
- [ ] Custom error types for "item not found" / "insufficient stock"
- [ ] Interactive CLI loop for add/remove/list commands

---

## Week 3 — Concurrency Fundamentals

**Objectives:** Understand goroutines and channels — Go's core differentiator.

**Concept Checklist**
- [ ] Goroutines and the Go scheduler (M:N model, GOMAXPROCS)
- [ ] Channels: unbuffered vs. buffered, directional channels
- [ ] `select` statement, timeouts with `time.After`
- [ ] `sync.WaitGroup`, `sync.Mutex`, `sync.RWMutex`
- [ ] Race conditions and the `-race` detector
- [ ] Worker pool pattern basics

**Project: Concurrent File Hasher / Duplicate Finder**
- [ ] Recursively scan a directory for files
- [ ] Use a worker pool (goroutines + channels) to compute SHA-256 hashes in parallel
- [ ] Aggregate results safely (mutex or channel-based collection) to detect duplicate files
- [ ] Benchmark sequential vs. concurrent runtime; run with `-race` to confirm safety

---

## Week 4 — Advanced Concurrency Patterns

**Objectives:** Move from "using goroutines" to "designing concurrent systems."

**Concept Checklist**
- [ ] `context.Context` for cancellation, deadlines, timeouts
- [ ] Fan-in / fan-out pipelines
- [ ] Rate limiting with `time.Ticker`
- [ ] Graceful shutdown patterns (`os/signal`, context cancellation)
- [ ] Avoiding goroutine leaks; `sync.Once`, `atomic` package

**Project: In-Memory Task Scheduler / Job Queue**
- [ ] Build a job queue that accepts tasks (functions) with priorities
- [ ] Dispatch tasks through a fan-out pipeline to N worker goroutines
- [ ] Support graceful shutdown on `SIGINT` (drain in-flight jobs, cancel context)
- [ ] Add retry logic with exponential backoff for failed jobs

---

## Week 5 — Standard Library Depth, I/O & Testing

**Objectives:** Get comfortable reading/writing data and writing idiomatic, tested Go code.

**Concept Checklist**
- [ ] `io.Reader` / `io.Writer` interfaces and why they're central to Go
- [ ] `bufio` for buffered reading/scanning
- [ ] `encoding/json`, `encoding/gob` for serialization
- [ ] File I/O with `os`, `path/filepath`
- [ ] Unit testing with `testing` package, table-driven tests
- [ ] Benchmarks (`go test -bench`), test coverage (`go test -cover`)

**Project: Persistent Key-Value Store**
- [ ] Design a simple `Get`/`Set`/`Delete` KV store interface
- [ ] Persist data to disk (append-only log or JSON snapshot) using `io.Writer`
- [ ] Reload data into memory on startup by replaying the log
- [ ] Write table-driven unit tests covering edge cases (missing key, overwrite, corrupted log line)
- [ ] Add at least one benchmark for write throughput

---

## Week 6 — Low-Level Mechanics: Memory, Binary Data & the Runtime

**Objectives:** Understand what's happening under the hood — memory layout, GC behavior, and raw bytes.

**Concept Checklist**
- [ ] Stack vs. heap allocation, escape analysis (`go build -gcflags="-m"`)
- [ ] Garbage collector basics (concurrent mark-and-sweep, GC tuning with `GOGC`)
- [ ] `encoding/binary` for reading/writing raw binary data & byte order
- [ ] `unsafe` package — what it does and why to avoid it (pointer arithmetic, struct layout)
- [ ] `os/exec` for spawning subprocesses
- [ ] Basic profiling with `pprof` (CPU and memory profiles)

**Project: Custom LRU Cache + Simple Bytecode VM**
- [ ] Implement an LRU cache from scratch (doubly linked list + map, no external libs)
- [ ] Extend it into a tiny stack-based bytecode interpreter (push/pop/add/sub/jump opcodes encoded as bytes via `encoding/binary`)
- [ ] Write a `.bin` file containing a small "program," load and execute it
- [ ] Profile memory allocations with `pprof` and reduce allocations by reusing buffers

---

## Week 7 — Networking & Inter-Process Communication (Non-HTTP)

**Objectives:** Learn raw socket-level networking and IPC — the systems side of "networking," not web frameworks.

**Concept Checklist**
- [ ] `net` package: `net.Listen`, `net.Dial` for TCP
- [ ] UDP basics with `net.ListenPacket`
- [ ] Framing protocols (length-prefixed messages) over a raw TCP stream
- [ ] Basic protocol design: message types, serialization (reuse `encoding/gob` or your own binary format)
- [ ] Unix domain sockets for local IPC (optional deep-dive)
- [ ] Handling multiple concurrent client connections with goroutines

**Project: Terminal-Based Multi-Client Chat Server**
- [ ] Build a raw TCP server accepting multiple client connections concurrently
- [ ] Design a simple length-prefixed binary message protocol (not HTTP/WebSockets)
- [ ] Broadcast messages to all connected clients via channels
- [ ] Handle client disconnects gracefully; add a `/nick` and `/quit` command
- [ ] Build a matching terminal client to connect and chat

---

## Week 8 — System Design & Capstone

**Objectives:** Combine everything into one cohesive, production-style system; think about design trade-offs.

**Concept Checklist**
- [ ] Package/module architecture for larger projects (cmd/, internal/, pkg/ layout)
- [ ] Structured logging (`log/slog`)
- [ ] Configuration management (env vars, flags via `flag` package)
- [ ] Dependency injection without frameworks (idiomatic Go style)
- [ ] Basic distributed-systems concepts: replication, consistency trade-offs, heartbeats
- [ ] Final profiling & optimization pass (CPU/memory) on your capstone

**Capstone Project: Distributed-ish In-Memory Key-Value Store with Replication**
- [ ] Combine Week 5's KV store + Week 7's TCP networking
- [ ] Support a primary node and one or more replica nodes communicating over TCP
- [ ] Replicate writes from primary to replicas via your own binary protocol
- [ ] Add heartbeat/health-check between nodes
- [ ] Add structured logging (`slog`) and a `-config` flag for node role (primary/replica) and port
- [ ] Write integration tests spinning up primary + replica in-process and verifying replication
- [ ] Profile and document one optimization you made (e.g., reduced allocations, batched replication)

---

## How to Use This
- Treat each week's project as non-negotiable — concepts don't stick without building something.
- Re-use code from earlier weeks in later projects where it makes sense (e.g., Week 6's binary encoding feeds directly into Week 7 and 8).
- If a week's project is finished early, go back and add tests + benchmarks to a previous week's project instead of rushing ahead.