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
