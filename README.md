# Ethereum Scanner Service

## Overview
The **Ethereum Scanner Service** is a Go-based application that connects to an Ethereum node using JSON-RPC to fetch blockchain data. It provides an HTTP API for querying current block information, fetching transactions, and subscribing to addresses for updates.

## Features
- **Fetch Current Block**: Retrieve the latest processed block from the Ethereum blockchain.
- **Fetch Transactions**: Get all transactions associated with a specific Ethereum address.
- **Subscribe to Addresses**: Register addresses to monitor their incoming and outgoing transactions.

---

## Project Structure

```
ethscanner/
├── common/       # Utility functions for HTTP responses
├── controller/   # Controllers to handle API requests
├── memstore/     # In-memory storage for transactions
├── parser/       # Parser and JSON-RPC communication with Ethereum node
├── main.go       # Entry point for the service
```

---

## Setup Instructions

### Prerequisites
- Go 1.18+ installed
- Access to an Ethereum JSON-RPC URL (e.g., Infura, Alchemy, or a local Ethereum node)

### Installation
1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd ethscanner
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o ethscanner
   ```

---

## Configuration
Update the constants in `main.go` as needed:
- **`URL`**: Set to your Ethereum JSON-RPC endpoint.
- **`PORT`**: Set the desired port number for the HTTP server.

---

## Running the Application
Start the server:
```bash
go run main.go
```
The server will run on the specified port (default: `8080`).

---

## API Endpoints

### 1. **Get Current Block**
   - **Endpoint**: `GET /api/v1/currentblock`
   - **Response**:
     ```json
     {
       "msg": "Success",
       "status_code": 200,
       "data": {
         "currentBlock": 123456
       }
     }
     ```

### 2. **Get Transactions**
   - **Endpoint**: `GET /api/v1/getalltransactions`
   - **Query Parameters**:
     - `address`: Ethereum address to query transactions for.
   - **Response**:
     ```json
     {
       "msg": "Success",
       "status_code": 200,
       "data": [
         {
           "hash": "0xabc123...",
           "from": "0xabc...",
           "to": "0xdef...",
           "value": "12345"
         }
       ]
     }
     ```

### 3. **Subscribe to an Address**
   - **Endpoint**: `POST /api/v1/subscribe`
   - **Query Parameters**:
     - `address`: Ethereum address to subscribe.
   - **Response**:
     ```json
     {
       "msg": "Success",
       "status_code": 200,
       "data": {
         "subscriptionStatus": "Address 0xabc... subscribed successfully"
       }
     }
     ```

---

## Key Modules

### `common`
Contains utility functions for sending consistent HTTP responses.

### `controller`
Defines handlers for the API endpoints:
- **CurrentBlock**
- **GetAllTransactions**
- **Subscribe**

### `memstore`
Implements an in-memory data store for managing transaction data.

### `parser`
Handles communication with the Ethereum node through JSON-RPC, including:
- Fetching the latest block.
- Parsing transactions from a block.

---

## Example Usage
1. Start the server:
   ```bash
   go run main.go
   ```

2. Use an API testing tool (e.g., Postman, curl) to interact with the endpoints:
   - Get current block:
     ```bash
     curl http://localhost:8080/api/v1/currentblock
     ```

   - Fetch transactions:
     ```bash
     curl "http://localhost:8080/api/v1/getalltransactions?address=0xabc..."
     ```

   - Subscribe to an address:
     ```bash
     curl -X POST "http://localhost:8080/api/v1/subscribe?address=0xabc..."
     ```

---

## Graceful Shutdown
The service listens for OS signals (`SIGINT`, `SIGTERM`) and gracefully shuts down the server.

---

