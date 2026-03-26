# Docker Commands for API-Contract-Detector

This document contains the common Docker commands used for running the PostgreSQL service and managing the container for API-Contract-Detector.

---

## 1. Start the PostgreSQL container

Run in detached mode:

```bash
docker-compose -f docker/docker-compose.yml up -d
```
- Starts the `db` service in the background
- Must be running before starting the Go server

---

## 2. Stop the container
```bash
docker-compose -f docker/docker-compose.yml down
```
- Stops and removes the container
- Data **persists** if a volume is configured

---

## 3. Access PostgreSQL shell inside container
```bash
docker-compose -f docker/docker-compose.yml exec db psql -U postgres -d contracts
```
- Connects to the `contracts` database
- Use this to run `CREATE TABLE`, `SELECT`, or `TRUNCATE` commands

---

## 4. Check running containers
```bash
docker-compose -f docker/docker-compose.yml ps
```
- Shows which containers are running

--- 

## 5. View container logs
```bash
docker-compose -f docker/docker-compose.yml logs -f db
```
- Live logs from the PostgreSQL container

---

## 6. Restart container (with table truncation)
```bash
docker-compose -f docker/docker-compose.yml restart db
```
- If `entrypoint` script is configured, table rows will be cleared automatically

--- 

## 7. Remove all volumes (if needed)
```bash
docker-compose -f docker/docker-compose.yml down -v
```
- **WARNING:** deletes all persisted data including schemas table

---
