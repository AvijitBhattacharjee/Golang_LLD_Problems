
1. What is the difference between grpc and rest API?
gRPC and REST API are two different approaches to building APIs for communication between client and server applications.
gRPC is a high-performance, open-source framework developed by Google that uses Protocol Buffers (protobuf) for serialization and supports multiple programming languages. It is designed for low-latency, high-throughput communication and is ideal for microservices architectures. gRPC uses HTTP/2 for transport, which allows for features like multiplexing, server push, and efficient binary framing.
On the other hand, REST (Representational State Transfer) is an architectural style that uses standard HTTP methods (GET, POST, PUT, DELETE) for communication. REST APIs typically use JSON or XML for data serialization and are widely adopted due to their simplicity and ease of use. REST is stateless and relies on standard HTTP status codes for error handling.

2. difference between http1.1 and http2?
HTTP/1.1 and HTTP/2 are two versions of the Hypertext Transfer Protocol (HTTP) used for communication between clients and servers on the web.
HTTP/1.1 is the older version and uses a text-based format for requests and responses. It supports persistent connections, allowing multiple requests and responses to be sent over a single connection, but it has limitations such as head-of-line blocking, where a slow request can block subsequent requests.
HTTP/2, on the other hand, is a newer version that introduces several improvements over HTTP/1.1. It uses a binary format for communication, which is more efficient and reduces overhead. HTTP/2 supports multiplexing, allowing multiple requests and responses to be sent simultaneously over a single connection, eliminating head-of-line blocking. It also includes features like header compression and server push, which can improve performance and reduce latency.

3. What is the difference between Array and slice in Go?
In Go, an array is a fixed-size collection of elements of the same type, while a slice is a dynamically-sized, flexible view into an array. Arrays have a fixed length that is determined at compile time, and their size cannot be changed. Slices, on the other hand, are built on top of arrays and can grow or shrink in size as needed. Slices provide more flexibility and are commonly used in Go programming due to their dynamic nature.

4. Size and Capacity of slice in Go?
In Go, a slice has two important properties: size and capacity. The size of a slice refers to the number of elements currently present in the slice, while the capacity refers to the total number of elements that the slice can hold before it needs to be resized. The size can be obtained using the built-in `len()` function, and the capacity can be obtained using the built-in `cap()` function. When a slice exceeds its capacity, a new underlying array is allocated, and the existing elements are copied over to the new array, which can lead to performance overhead if done frequently.

5. What is the difference between buffered and unbuffered channel in Go?
In Go, a channel is a communication mechanism that allows goroutines to send and receive values. A buffered channel has a fixed capacity, meaning it can hold a specific number of values without blocking. An unbuffered channel, on the other hand, has no capacity and will block until both a sender and receiver are ready to communicate. Buffered channels can improve performance by reducing the need for synchronization between goroutines, while unbuffered channels ensure strict ordering of communication.

6. What is the difference between goroutine and thread in Go?
In Go, a goroutine is a lightweight, concurrent function that is managed by the Go runtime, while a thread is a more heavyweight unit of execution managed by the operating system. Goroutines are multiplexed onto a smaller number of OS threads, allowing for efficient concurrency and low overhead. They are created using the `go` keyword and can be scheduled independently by the Go runtime. Threads, on the other hand, have higher memory and scheduling overhead, making goroutines more suitable for concurrent programming in Go.

7. What is the difference between defer and panic in Go?
In Go, `defer` and `panic` are two different mechanisms used for handling function execution and error handling. The `defer` statement is used to schedule a function call to be executed after the surrounding function completes, regardless of whether it completes normally or due to a panic. This is often used for resource cleanup, such as closing files or releasing locks.
On the other hand, `panic` is a built-in function that stops the normal execution of a program and begins panicking, which unwinds the stack and executes any deferred functions. It is typically used to indicate a serious error or unexpected condition that cannot be handled gracefully. When a panic occurs, the program will terminate unless it is recovered using the `recover` function within a deferred function.


// recover will only work in defer.

8. What is the difference between interface and struct in Go?
In Go, an interface is a type that defines a set of method signatures, while a struct is a composite data type that groups together fields (variables) under a single name. Interfaces allow for polymorphism, enabling different types to implement the same set of methods and be used interchangeably. Structs, on the other hand, are used to create concrete data types with specific fields and can have methods associated with them. While structs define the structure of data, interfaces define behavior.

9. How microservices communicate with each other?
Microservices communicate with each other using various communication protocols and patterns. The most common methods include:
- REST APIs
- gRPC
- Message queues (e.g., RabbitMQ, Apache Kafka)
- Event-driven communication

10. What is the difference between mutex and rwmutex in Go?
In Go, a `mutex` (mutual exclusion) is a synchronization primitive that allows only one goroutine to access a shared resource at a time, ensuring that concurrent access does not lead to race conditions. A `rwmutex` (read-write mutex) is a more advanced synchronization primitive that allows multiple goroutines to read from a shared resource simultaneously, while still ensuring exclusive access for writing. The `rwmutex` provides two types of locks: a read lock, which can be held by multiple readers, and a write lock, which can only be held by one writer at a time. This allows for better performance in scenarios where reads are more frequent than writes.

11. How to debug and diagnose if my application is behaving faulty?
Debugging and diagnosing issues in a Go application can be done using various techniques and tools. Here are some common approaches:
- Using the `log` package for basic logging
- Using the `pprof` package for performance profiling
- Using the `trace` package for execution tracing
- Using third-party debugging tools like Delve

12. How to debug if my APIs not failing?
If your APIs are not failing but you suspect there might be issues, you can use the following debugging techniques:
- Check the API response codes and messages to ensure they are as expected. 
- Use logging to capture request and response details for analysis.
- Use tools like Postman or curl to manually test the API endpoints and verify their behavior.

13. what is horizontal scaling and vertical scaling in microservices?
Horizontal scaling involves adding more instances of a service to distribute the load, while vertical scaling involves increasing the resources (CPU, memory) of an existing instance. In microservices, horizontal scaling is often preferred as it allows for better fault tolerance and resource utilization.

14. What is the difference between monolithic and microservices architecture?
Monolithic architecture is a traditional software design where all components of an application are tightly integrated into a single codebase. In contrast, microservices architecture breaks down an application into smaller, independent services that can be developed, deployed, and scaled independently. Microservices promote flexibility, scalability, and maintainability, while monolithic architectures can be simpler to develop initially but may become difficult to manage as the application grows.


15. how Protocol Buffers (protobuf) work in gRPC?
Protocol Buffers (protobuf) is a language-agnostic binary serialization format developed by Google. It is used in gRPC to define the structure of the data being transmitted between client and server. Developers define message types and service methods in a .proto file, which is then compiled into code for the desired programming languages. This code handles serialization and deserialization of the messages, allowing for efficient communication over the network. Protobuf is designed to be compact and fast, making it suitable for high-performance applications.

16. how grpc works in microservices?
gRPC works in microservices by providing a high-performance, language-agnostic framework for communication between services. Each microservice can define its own gRPC service and methods using Protocol Buffers (protobuf) for message serialization. The gRPC framework handles the underlying communication, including connection management, serialization, and deserialization of messages. Services can communicate with each other using gRPC clients and servers, allowing for efficient and low-latency communication. gRPC also supports features like streaming, authentication, and load balancing, making it well-suited for microservices architectures where services need to interact frequently and efficiently.

17. how to implement grpc in microservices?
To implement gRPC in microservices, follow these steps:
 - Define the service and message types in a .proto file using Protocol Buffers syntax.
 - Use the Protocol Buffers compiler (protoc) to generate code for the desired programming languages.
 - Implement the server-side logic for the gRPC service in each microservice, handling the defined methods and processing incoming requests.
 - Implement the client-side logic in other microservices that need to communicate with the gRPC service, using the generated client code to make requests and handle responses.
 - Set up the necessary infrastructure for service discovery, load balancing, and authentication if required.
 - Test the gRPC services to ensure they are functioning correctly and efficiently.

18. AWS/GCP Server Switchover Without Downtime
To achieve server switchover without downtime in AWS or GCP, you can use the following strategies:
- Use load balancers to distribute traffic across multiple instances of your application. This allows you to take one instance offline for maintenance or updates while the load balancer continues to route traffic to the remaining healthy instances.
- Implement health checks to monitor the status of your instances. If an instance becomes unhealthy, the load balancer can automatically redirect traffic to healthy instances, ensuring continuous availability.
- Use blue-green deployment or canary deployment strategies to gradually roll out updates to your application.

19. Answer inorder What Happens When You Type google.com?
When you type `google.com` in a browser, the following sequence of events occurs:
1. DNS Resolution: The browser queries the Domain Name System (DNS) to resolve the domain   name `google.com` into an IP address. This involves checking the local cache, querying the configured DNS resolver, and potentially traversing multiple DNS servers until the IP address is found.
2. TCP Connection: Once the IP address is obtained, the browser establishes a TCP connection to the server using a three-way handshake process. This involves sending a SYN packet, receiving a SYN-ACK response, and sending an ACK packet to establish the connection.
3. TLS Handshake (if HTTPS): If the connection is over HTTPS, the browser initiates a TLS handshake to establish a secure connection. This involves exchanging cryptographic keys and verifying the server's identity using its SSL/TLS certificate.
4. HTTP Request: After the connection is established, the browser sends an HTTP request to the server, typically a GET request for the root resource ("/"). The request includes headers that provide information about the client and the requested resource.
5. Server Processing: The server processes the request, retrieves the requested resource, and generates an  HTTP response. This may involve querying databases, executing server-side logic, and preparing the response data.
6. HTTP Response: The server sends the HTTP response back to the browser, which includes a status code, headers, and the requested content (e.g., HTML, CSS, JavaScript).   

20. For bigger API responses, you can use techniques like pagination, compression, and caching to improve performance and reduce the amount of data transferred.

Sharding can be used to distribute large datasets across multiple servers or databases, allowing for more efficient data retrieval and processing. This can help improve performance and scalability for applications that handle large amounts of data.

Databse partitioning is another technique that can be used to improve performance and scalability by dividing a large database into smaller, more manageable pieces. This can help reduce query times and improve overall system performance.

21. What are the best practices for debugging when your microservices are not working as expected?
When debugging microservices, consider the following best practices:
- Use centralized logging to collect logs from all microservices in one place, making it easier to trace issues across services.
- Implement distributed tracing to track requests as they flow through multiple services, helping identify bottlenecks and failures.
- Use monitoring and alerting tools to proactively detect issues and performance degradation.


22. How to handle versioning in microservices?
Handling versioning in microservices can be achieved through several strategies:
- API versioning: Include the version number in the API endpoint (e.g., `/api/v1/users`).
- Header-based versioning: Use custom headers to specify the API version (e.g., `Accept: application/vnd.myapp.v1+json`).
- Query parameter versioning: Pass the version as a query parameter (e.g., `/api/users?version=1`).

23. how a user credential is verified in microservices?
User credentials can be verified in microservices using authentication and authorization mechanisms. Common approaches include:
- Token-based authentication: Use JSON Web Tokens (JWT) or OAuth tokens to authenticate users. The token is generated upon successful login and is included in subsequent requests to verify the user's identity.
- API Gateway: Implement an API gateway that handles authentication and authorization for incoming requests, verifying user credentials before routing them to the appropriate microservice.
- Centralized authentication service: Use a dedicated authentication service that manages user credentials and verifies them for other microservices. This service can handle login, token generation, and validation, ensuring consistent authentication across the system.

24. How to handle authentication and authorization in microservices?
Authentication and authorization in microservices can be handled using various strategies:
- Token-based authentication: Use JSON Web Tokens (JWT) or OAuth tokens to authenticate users. The token is generated upon successful login and is included in subsequent requests to verify the user's identity.
- API Gateway: Implement an API gateway that handles authentication and authorization for incoming requests, verifying user credentials before routing them to the appropriate microservice.
- Centralized authentication service: Use a dedicated authentication service that manages user credentials and verifies them for other microservices. This service can handle login, token generation, and validation, ensuring consistent authentication across the system.


25. Difference between postgresql and mysql and noSQL databases?
PostgreSQL and MySQL are both relational database management systems (RDBMS) that use structured query language (SQL) for data management. PostgreSQL is known for its advanced features, such as support for complex queries, full-text search, and extensibility. It is often preferred for applications that require complex data relationships and integrity. MySQL, on the other hand, is known for its simplicity, speed, and ease of use, making it a popular choice for web applications and smaller projects.
NoSQL databases, such as MongoDB, Cassandra, and Redis, are designed for unstructured or semi-structured data and provide flexible schema designs. They are often used for applications that require high scalability, performance, and the ability to handle large volumes of data. NoSQL databases can be categorized into key-value stores, document stores, column-family stores, and graph databases, each optimized for specific use cases.

26. Write custom error in golang?
In Go, you can create custom error types by implementing the `error` interface, which requires a single method `Error() string`. Here's an example of how to create a custom error type:

```go
type MyError struct {
    message string
}

func (e *MyError) Error() string {
    return e.message
}

func NewMyError(message string) *MyError {
    return &MyError{message: message}
}
```
You can then use this custom error type in your code as follows:

```go
err := NewMyError("Something went wrong")
if err != nil {
    fmt.Println(err)
}
```

27. how to identify a closed channel in golang?
In Go, you can identify a closed channel by using the "comma ok" idiom when receiving from the channel. When you receive from a channel, it returns two values: the value received and a boolean indicating whether the channel is open or closed. If the channel is closed, the boolean will be false.
Here's an example:
```go
val, ok := <-ch
if !ok {
    // Channel is closed
}
``` 

28. How to handle errors in golang?
In Go, error handling is typically done using the built-in `error` type. Functions that can encounter errors usually return an `error` as the last return value. You can check if the error is `nil` to determine if the operation was successful or if an error occurred.

29. write down folder structure of microservices
A typical folder structure for a microservices architecture in Go might look like this:



30. what is mvc technique?
MVC (Model-View-Controller) is a software design pattern that separates an application into three interconnected components: the Model, the View, and the Controller. The Model represents the data and business logic of the application, the View is responsible for presenting the data to the user, and the Controller handles user input and updates the Model and View accordingly. This separation of concerns allows for better organization, maintainability, and scalability of applications.

31. what is error wrapping in golang?
Error wrapping in Go allows you to add context to an error without losing the original error information. This is useful for debugging and providing more meaningful error messages. You can wrap errors using the `fmt.Errorf` function with the `%w` verb, which indicates that the error should be wrapped. For example:
```go
err := fmt.Errorf("failed to read file: %w", os.ErrNotExist)
```

32. how to prevent other go routines from failing if one go routines failed?
In Go, to prevent other goroutines from failing if one goroutine fails, you can use error handling and synchronization techniques. Here are some approaches:
- Use channels to communicate errors between goroutines. Each goroutine can send its error to a channel, and the main goroutine can handle the errors without affecting the execution of other goroutines.
- Use `sync.WaitGroup` to wait for all goroutines to complete, and handle errors in a centralized manner. This way, if one goroutine fails, it does not directly impact the others.
- Implement a recovery mechanism using `defer` and `recover` to catch panics in individual goroutines, allowing them to fail gracefully without crashing the entire application.

33. folder structure of microservices in golang?
user-service/
│
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── api/
│   │   ├── handler/                # HTTP/gRPC handlers
│   │   ├── middleware/             # Auth, logging, recovery
│   │   └── routes.go
│   │
│   ├── service/                    # Business logic
│   │   └── user_service.go
│   │
│   ├── repository/                 # Database access
│   │   ├── postgres.go
│   │   └── user_repository.go
│   │
│   ├── model/                      # Database models
│   │   └── user.go
│   │
│   ├── dto/                        # Request/Response structs
│   │   ├── request.go
│   │   └── response.go
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── database/
│   │   └── postgres.go
│   │
│   ├── cache/
│   │   └── redis.go
│   │
│   ├── client/                     # External service clients
│   │   ├── payment_client.go
│   │   └── notification_client.go
│   │
│   ├── event/
│   │   ├── producer.go
│   │   └── consumer.go
│   │
│   ├── auth/
│   │   ├── jwt.go
│   │   └── middleware.go
│   │
│   ├── logger/
│   │   └── logger.go
│   │
│   ├── metrics/
│   │   └── prometheus.go
│   │
│   ├── validation/
│   │   └── validator.go
│   │
│   └── util/
│       └── helper.go
│
├── pkg/
│   ├── errors/
│   ├── constants/
│   ├── response/
│   └── utils/
│
├── proto/                          # gRPC proto files
│
├── migrations/
│   ├── 001_create_user.sql
│   └── 002_add_index.sql
│
├── scripts/
│   ├── start.sh
│   └── migrate.sh
│
├── configs/
│   ├── config.yaml
│   └── config-dev.yaml
│
├── deployments/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ingress.yaml
│   └── helm/
│
├── test/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── swagger.yaml
│   └── architecture.md
│
├── .github/
│   └── workflows/
│
├── go.mod
├── go.sum
├── Makefile
└── README.md

34. Without go routines how can you achieve concurrency in golang?
Without goroutines, you can achieve concurrency in Go using the following methods:
- Using channels: Channels can be used to communicate between different parts of your program, allowing for concurrent execution of tasks. You can create multiple channels and use them to send and receive data between different functions or components.
- Using the `sync` package: The `sync` package provides synchronization primitives like `WaitGroup`, `Mutex`, and `Cond` that can be used to coordinate concurrent execution of tasks. You can use these primitives to manage access to shared resources and ensure that multiple tasks can run concurrently without interfering with each other.
- Using the `select` statement: The `select` statement allows you to wait on multiple channel operations, enabling you to handle multiple concurrent tasks in a non-blocking manner. You can use `select` to listen for messages on multiple channels and execute different code paths based on which channel receives data first.

35. Disadvantages of cache in microservices? and advantages of redis cache in microservices?
Disadvantages of cache in microservices:
- Cache Invalidation: Keeping the cache in sync with the underlying data can be challenging, leading to stale or inconsistent data if not managed properly.
- Increased Complexity: Implementing caching adds complexity to the system, requiring additional logic for cache management and monitoring.
- Memory Overhead: Caching can consume significant memory resources, especially if large datasets are cached, which may lead to increased costs and resource usage.
Advantages of Redis cache in microservices:
- High Performance: Redis is an in-memory data structure store, offering very fast read and write operations.
- Data Persistence: Redis supports various persistence options to ensure data durability.
- Rich Data Structures: Redis provides a variety of data structures like strings, hashes, lists, sets, and sorted sets, making it versatile for different use cases.
- Pub/Sub Messaging: Redis supports publish/subscribe messaging patterns, useful for real-time communication between services.


36. how to handle complex implemened functions in microservices?
To handle complex implemented functions in microservices, you can follow these strategies:
- Break down complex functions into smaller, manageable components or services. This promotes modularity and makes it easier to test and maintain each component independently.
- Use well-defined interfaces and contracts between services to ensure clear communication and reduce coupling.

37. What is graceful shutdown in microservices?
Graceful shutdown in microservices refers to the process of shutting down a service in a controlled manner, allowing it to complete ongoing requests and clean up resources before terminating. This helps prevent data loss, ensures that clients receive proper responses, and maintains the integrity of the system. During a graceful shutdown, the service typically stops accepting new requests, waits for in-flight requests to finish, and then releases resources such as database connections, file handles, and network sockets. Implementing graceful shutdown is important for maintaining reliability and stability in a microservices architecture.

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	fmt.Println("Running...")

	<-ch

	fmt.Println("Shutting down...")
}

