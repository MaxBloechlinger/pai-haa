# pai-haa, Search Engine from Scratch

A search engine built from scratch in Go, concurrent web crawler, inverted index, and BM25 relevance ranking over a real corpus.

No search libraries, no frameworks. Just the standard library and the actual data structures and algorithms that make search work.

---

## What it does

- **Crawls** real web pages concurrently using a worker pool, with a visited set and per-host politeness.
- **Indexes** documents into an inverted index (term to postings), tracking term frequency and positions.
- **Ranks** results by relevance using BM25 (with TF-IDF as the baseline it improves on).
- **Serves** queries through a simple HTTP API.

## Tech

Go (stdlib) · goroutines + channels for concurrency · `encoding/gob` for persistence · GitHub Actions for CI.

## Status

WIP: Building in phases — see [`ROADMAP.md`](./ROADMAP.md).

- [x] Minimal inverted index + search
- [ ] Concurrent crawler
- [ ] BM25 ranking
- [ ] Persistence + query API
- [ ] Tests, CI, benchmarks

## Quick start

```bash
git clone https://github.com/<you>/gosearch.git
cd gosearch
go run .
```

## Why I built it

To understand how search engines actually work end to end concurrency, information retrieval, and ranking by building one instead of reading about it.

---
