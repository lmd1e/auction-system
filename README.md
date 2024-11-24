# Auction System Prototype

## Overview

This prototype implements a simple auction system where users can create and bid on lots. The system determines the winner after the auction ends, processes transactions, and notifies participants. The architecture follows Domain-Driven Design (DDD) and Clean Architecture principles, with a focus on separation of concerns and scalability.

## Architecture

The system is divided into the following layers:

- **Presentation Layer**: Exposes gRPC and REST APIs for client interaction.
- **Application Layer**: Contains use cases and services that orchestrate business logic.
- **Domain Layer**: Defines core business entities, rules, and interfaces.
- **Infrastructure Layer**: Handles data access, external service integration, and background tasks.

## Setup and Execution

### Prerequisites

- Go 1.23.1 or later
- Docker and Docker Compose (optional, for running with Docker)
- PostgreSQL
- Protocol Buffers Compiler (`protoc`)
- Go plugins for `protoc`: `protoc-gen-go` and `protoc-gen-go-grpc`

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/auction-system.git
   cd auction-system