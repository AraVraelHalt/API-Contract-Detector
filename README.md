# API Contract Break Detector

Detect breaking API changes using real production traffic — not just static schemas.

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 🚨 Problem

Microservices often break each other due to undocumented or implicit API contracts.

Traditional tools rely on OpenAPI specs, which are often:

* Outdated
* Incomplete
* Ignored in practice

## 💡 Solution

This tool observes real API traffic, infers schemas, and detects breaking changes automatically.

## ✨ Features

* Automatic JSON schema inference
* Schema version tracking per endpoint
* Breaking change detection
* Dashboard visualization
* CI/CD integration (GitHub Actions)

## 🏗️ How It Works

1. Capture API traffic via proxy or logs
2. Infer JSON schema from requests/responses
3. Store schema versions
4. Diff schemas to detect breaking changes

## 🚀 Getting Started

```bash
git clone https://github.com/AraVraelHalt/API-Contract-Detector
cd API-Contract-Detector
docker-compose up
```

## 🛠️ Tech Stack

* Go
* PostgreSQL
* Docker
* React (dashboard)

## 📈 Future Work

* Traffic replay testing
* Multi-service dependency graph
* OpenAPI integration

## 🤝 Why This Project Matters

This project simulates real-world distributed system challenges:

* Backward compatibility
* Contract enforcement
* Observability

