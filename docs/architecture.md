# Royal Vault Architecture

Royal Vault is a zero-knowledge password manager built using a modern, secure, and scalable architecture.

## Tech Stack

Frontend:
- Vue.js

Backend:
- Golang (Fiber framework)

Database:
- PostgreSQL (Dockerized)

Infrastructure:
- Docker + Docker Compose

## System Design

Client (Vue) → REST API (Go Fiber) → PostgreSQL Database

All sensitive cryptographic operations are performed on the client side to maintain zero-knowledge security.
