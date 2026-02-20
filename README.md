# Go Concurrency Pipeline

A learning-oriented Go project exploring bounded concurrency, backpressure, rate limiting, and graceful shutdown in real-world I/O pipelines.

While the current implementation demonstrates a concurrent URL fetching workload, the architecture is intentionally designed to be **general-purpose** and applicable to many high-throughput data processing scenarios.

---

## 🎯 Motivation

This project was built to deeply understand production-grade concurrency patterns in Go, including:

- Worker pool design
- Backpressure and flow control
- Rate limiting strategies
- Graceful shutdown
- Retry with exponential backoff
- Context propagation
- Channel lifecycle management

The goal is not just to fetch URLs, but to model how real systems safely perform large volumes of external I/O.

---

## 🧠 What This Demonstrates

This repository showcases:

- Bounded parallelism using worker pools
- Pipeline-based architecture
- Cooperative cancellation with context
- Rate-limited external I/O
- Retry handling with backoff
- Clean goroutine lifecycle management
- Proper channel ownership and closing discipline

These patterns commonly appear in:

- Web crawlers
- API ingestion systems
- distributed job processors
- monitoring and probing systems
- ETL pipelines

---

## 🏗️ High-Level Architecture

Key properties:

- Controlled concurrency
- Backpressure-aware
- Graceful termination
- Failure-resilient

---

## 🚀 Current Example Workload

The repository currently includes a URL fetching pipeline that:

- validates input URLs
- processes them with bounded workers
- applies rate limiting
- retries transient failures
- measures latency
- supports graceful shutdown (Ctrl+C)

This workload serves as a **concrete demonstration**, not the only intended use case.

---

## 🔧 Potential Extensions

Some natural next explorations:

- Per-host rate limiting
- Adaptive throttling
- Circuit breaker integration
- Distributed worker coordination
- Metrics and observability
- Pluggable job sources

---

## 📚 Learning Goals

This project is intentionally written as a learning exercise to build intuition around:

- Go concurrency patterns
- pipeline design
- failure handling in distributed systems
- polite and resilient external I/O

---

## 🛠️ Running

```bash
go run main.go -file urls.txt
```

or

```bash
url-fetcher.exe -file urls.txt
```

Interrupt with Ctrl+C to observe graceful shutdown
