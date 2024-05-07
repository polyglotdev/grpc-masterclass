# gRPC Golang Masterclass: Build Modern API & Microservices

[Course Link](https://www.udemy.com/course/grpc-golang)

## Section 1: Introduction

### What is gRPC?

- gRPC is a modern, open-source, high-performance remote procedure call (RPC) framework that can run anywhere.
- It enables client and server applications to communicate transparently, and simplifies the building of connected systems.
- It is based on HTTP/2, Protocol Buffers (protobuf) and other modern technologies.
- It is a language-agnostic, platform-neutral, extensible, and easy to use.
- It is ideal for developing APIs and microservices.

### Why gRPC?

- **Performance**: gRPC is faster than REST because it uses HTTP/2 and Protocol Buffers.
- **Interoperability**: gRPC can generate client libraries in multiple languages.
- **Streaming**: gRPC supports streaming requests and responses.
- **Code Generation**: gRPC generates client and server code automatically.
- **Error Handling**: gRPC has built-in error handling.
- **Security**: gRPC supports authentication and authorization.

## HTTP/1 vs HTTP/2

### HTTP/1.1

HTTP/1.1 has been the backbone of the web since its standardization in 1997. It introduced several improvements over HTTP/1.0 to make web communications more efficient:

- **Persistent Connections:** Unlike HTTP/1.0, which required a new connection for each request, HTTP/1.1 allows multiple requests and responses over the same connection. This reduces the overhead of reconnecting each time, speeding up web page loading.
- **Pipelining:** This feature lets a browser send multiple requests without waiting for each response, aiming to further improve the use of a connection. However, responses must be returned in the order requested, which can lead to head-of-line blocking.
- **Chunked Transfers:** Servers can send responses in chunks, useful for generating content on the fly without knowing the entire content length upfront.
- **Additional Cache Controls and Methods:** Enhancements like more specific cache directives and new methods (OPTIONS, PUT, DELETE) expanded HTTP's capabilities.

Despite these improvements, HTTP/1.1 has limitations, especially concerning efficiency and speed in modern web applications.

### HTTP/2

Recognizing the need for a more efficient protocol, HTTP/2 was standardized in 2015. It introduces several key features to tackle the shortcomings of HTTP/1.1:

- **Binary Framing Layer:** HTTP/2 uses a binary protocol rather than text. This change makes the protocol more compact and less prone to errors like newline character issues.
- **Multiplexing:** Multiple requests and responses can be intermixed on the same connection, simultaneously. This addresses the head-of-line blocking problem found in HTTP/1.1 by allowing unrelated requests to proceed even if one is held up.
- **Server Push:** Servers can proactively send resources to a client, assuming the client will need them (like CSS and JavaScript files), without waiting for the client to request them.
- **Header Compression:** HTTP/2 uses HPACK compression to reduce overhead. Since headers often contain repeated information, this can significantly reduce the amount of data sent over the network.
- **Flow Control:** Flow control at the application layer allows both the client and server to control the pace of communications to prevent overwhelming each other.

### Practical Impact and Adoption

The shift from HTTP/1.1 to HTTP/2 primarily aims at performance improvements, especially in environments with complex application delivery needs. Modern web browsers and servers widely support HTTP/2, and it's generally backward compatible with HTTP/1.1, meaning that it can gracefully degrade to HTTP/1.1 if necessary.

From a software engineering perspective, while HTTP/2 reduces the need for some optimizations used with HTTP/1.1 (like image sprites or domain sharding), it also imposes new requirements, such as understanding the nuances of server push and multiplexing. Implementing HTTP/2 can significantly enhance user experience by reducing load times and improving site responsiveness.

### 1. Unary RPC

In unary RPC, the client sends a single request to the server and gets a single response back, just like a typical function call.

**Example:**
```go
// Service definition in the .proto file
rpc SayHello(HelloRequest) returns (HelloResponse);
```
Here, `SayHello` is a simple unary call where the client sends a `HelloRequest` and receives a `HelloResponse`.

### 2. Server Streaming RPC

In server streaming RPC, the client sends a request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages. This type is useful for cases where the server needs to send a lot of data (e.g., logs or monitoring data).

**Example:**
```go
// Service definition in the .proto file
rpc ListFeatures(Rectangle) returns (stream Feature);
```

`ListFeatures` might be used to send a series of `Feature` objects that fall within a defined `Rectangle`.

### 3. Client Streaming RPC

Client streaming RPC is the opposite of server streaming: the client writes a sequence of messages and sends them to the server, using a provided stream. Once the client has finished writing the messages, it waits for the server to read them all and return a response.

**Example:**
```go
// Service definition in the .proto file
rpc RecordRoute(stream Point) returns (RouteSummary);
```

In `RecordRoute`, the client might stream a series of geographic `Point` objects to the server, which then returns a `RouteSummary`.

### 4. Bidirectional Streaming RPC

Bidirectional streaming RPC allows both the client and the server to read and write in a stream independently. Both sides can send a sequence of messages using a read-write stream. The streams operate independently, so clients and servers can read and write in whatever order they like: for example, the server could wait to receive all the client’s messages before writing its responses, or it could alternately read a message then write a message, or some other combination of read-write patterns.

**Example:**
```go
// Service definition in the .proto file
rpc RouteChat(stream RouteNote) returns (stream RouteNote);
```

`RouteChat` allows the client and server to exchange messages (or `RouteNote` objects) in a freeform dialogue where both sides send as they see fit.


Each of these API types suits different needs. Unary is great for simple request-response interaction, while the various streaming modes can handle more complex interactions, such as real-time updates, large data transfers, or ongoing conversations. The choice of API style in gRPC can significantly affect the design and performance of your application, especially in systems where network latency or data throughput is a concern. Knowing when to use each type can help you build more efficient and effective microservices.

## Process of getting things started

Current File structure:

```bash
.
├── Makefile           # Used for automating the build process
├── README.md          # Project documentation
├── bin                # May contain compiled binaries (not mandatory in the structure)
│   ├── calculator
│   │   ├── client     # Binaries for calculator client
│   │   └── server     # Binaries for calculator server
│   └── greet
│       ├── client     # Binaries for greet client
│       └── server     # Binaries for greet server
├── calculator
│   ├── client
│   │   ├── main.go    # Entry point for the calculator client
│   │   └── sum.go     # Client-side logic for calling sum service
│   ├── proto
│   │   ├── calculator.pb.go           # Generated Go file from calculator.proto
│   │   ├── calculator.proto           # Proto file for calculator service
│   │   ├── calculator_grpc.pb.go      # Generated gRPC specific Go file
│   │   ├── sum.pb.go                  # Generated Go file from sum.proto (if separate)
│   │   └── sum.proto                  # Proto file for sum service (if separate)
│   └── server
│       ├── main.go    # Entry point for the calculator server
│       └── sum.go     # Server-side implementation of sum service
├── go.mod             # Manage dependencies
├── go.sum             # Lock file for dependencies
└── greet
    ├── client
    │   ├── greet.go   # Client-side logic for calling greet service
    │   └── main.go    # Entry point for the greet client
    ├── proto
    │   ├── greet.pb.go                # Generated Go file from greet.proto
    │   ├── greet.proto                # Proto file for greet service
    │   └── greet_grpc.pb.go           # Generated gRPC specific Go file
    └── server
        ├── greet.go   # Server-side implementation of greet service
        └── main.go    # Entry point for the greet server
```

So we separate the Client and Server. The client will be responsible for making requests to the server, and the server will be responsible for handling those requests.

## Section 6 gRPC Server Streaming

gRPC Server Streaming is a type of API where the client sends a single request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages. This type of API is useful when the server needs to send a lot of data (e.g., logs or monitoring data) to the client.
