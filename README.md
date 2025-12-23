# Connect-Go Example: Universal RPC & REST

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)
![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-3178C6?style=flat&logo=typescript)
![Vite](https://img.shields.io/badge/Vite-6.0+-646CFF?style=flat&logo=vite)
![Buf](https://img.shields.io/badge/Buf-v1.48+-2d3e50?style=flat&logo=buffer)

A complete demonstration of a **Universal API Server** using [ConnectRPC](https://connectrpc.com/).

This repository shows how to define a service **once** using Protocol Buffers and expose it simultaneously over **four** protocols on a single HTTP/2 port:
1.  **gRPC** (for backend microservices)
2.  **Connect** (for Go/Node clients)
3.  **gRPC-Web** (for the TypeScript browser frontend)
4.  **REST/HTTP** (for `curl` and OpenAPI clients, powered by [Vanguard](https://connectrpc.com/vanguard))

## 🌟 Features

* **Single Source of Truth**: All APIs defined in `greet/v1/greet.proto`.
* **Universal Server**: `cmd/server` handles gRPC, Connect, and REST traffic on port `8080`.
* **Type-Safe Web Client**: `web/` contains a Vite + TypeScript app using generated clients.
* **REST Transcoding**: Automatic conversion of HTTP/JSON requests to RPC (e.g., `POST /v1/greet`).
* **Validation**: Input validation using [Protovalidate](https://github.com/bufbuild/protovalidate).
* **Modern Tooling**: Managed entirely via [Buf](https://buf.build/).

---

## 🛠️ Prerequisites

Ensure you have the following installed:

1.  **Go** (1.23 or higher)
2.  **Node.js** (20+ and `npm`)
3.  **Buf CLI**:
    ```bash
    go install [github.com/bufbuild/buf/cmd/buf@latest](https://github.com/bufbuild/buf/cmd/buf@latest)
    ```
4.  **Protoc Plugins** (for Go generation):
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install [connectrpc.com/connect/cmd/protoc-gen-connect-go@latest](https://connectrpc.com/connect/cmd/protoc-gen-connect-go@latest)
    ```

---

## 🚀 Getting Started

### 1. Clone & Initialize
```bash
git clone git@github.com:lao-tseu-is-alive/connect-go-example.git
cd connect-go-example
go mod tidy

```


### 2. Install Web Dependencies

This also installs the JS/TS generation plugins used by Buf.

```bash
cd web
npm install
cd ..

```

### 3. Generate Code

Run `buf` to generate Go server stubs, TypeScript client code, and OpenAPI specs.

```bash
buf generate

```

---

## 🏃‍♂️ Running the Project

You will need two terminal windows.

### Terminal 1: The Server

Starts the Go server on `localhost:8080`.

```bash
go run ./cmd/server/server.go

```

*You should see logs indicating the server is listening and `transcoder` is active.*

### Terminal 2: The Web Client

Starts the Vite development server on `localhost:5173`.

```bash
cd web
npm run dev

```

Open your browser to [http://localhost:5173](https://www.google.com/search?q=http://localhost:5173) to test the gRPC-Web client.

---

## 🧪 Testing the Endpoints

The server listens on `localhost:8080` and accepts traffic from various protocols.

### 1. Classical Connect (RPC)

The standard protocol for server-to-server communication.

```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "Developer"}' \
    http://localhost:8080/greet.v1.GreetService/Greet

```

### 2. REST (via Vanguard)

Uses the path defined in `google.api.http` options in the proto file.

```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "RestUser"}' \
    http://localhost:8080/v1/greet

```

### 3. gRPC (CLI)

Using the Go client in `grpc` mode (requires HTTP/2 support).

```bash
go run ./cmd/client/client.go -mode=grpc -name=Gopher

```

### 4. gRPC-Web (Browser)

Visit the web UI at `http://localhost:5173`.

* **Protocol**: gRPC-Web (over HTTP/1.1 or HTTP/2)
* **Transport**: `createGrpcWebTransport` (configured in `web/src/main.ts`)

---

## 📂 Project Structure

```text
├── buf.gen.yaml       # Configuration for code generation (Go, TS, OpenAPI)
├── buf.yaml           # Buf module configuration
├── cmd/
│   ├── client/        # CLI client implementation (Go)
│   └── server/        # Main server entrypoint (Go + Vanguard + CORS)
├── gen/               # Generated Go code & OpenAPI specs
├── greet/
│   └── v1/            # Protocol Buffer definitions (.proto)
└── web/               # Vite + TypeScript Frontend
    ├── src/
    │   ├── gen/       # Generated TypeScript code
    │   └── main.ts    # Client logic using Connect-ES
    └── package.json   # NPM dependencies

```

## 📜 License

MIT License — See [LICENSE](./LICENSE) for details.

---

  Built with ❤️ using [Go](https://go.dev/), [Proto](https://protobuf.dev/), and [Connect](https://connectrpc.com/)


