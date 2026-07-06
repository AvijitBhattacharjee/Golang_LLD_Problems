# 4. Networking Interview Questions

## OSI Model (7 Layers)

| Layer | Protocols | Responsibility |
|--------|-----------|---------------|
| 7. Application | HTTP, HTTPS, FTP, DNS | User-facing protocols |
| 6. Presentation | SSL/TLS, Encryption | Encryption, Compression |
| 5. Session | NetBIOS, RPC | Session management |
| 4. Transport | TCP, UDP | Reliable communication |
| 3. Network | IP, ICMP | Routing |
| 2. Data Link | Ethernet, ARP | MAC addressing |
| 1. Physical | Cable, Fiber | Physical transmission |

**Interview Tip**

Backend engineers mostly work with Layers **3, 4 and 7**.

---

## TCP vs UDP

### TCP

- Connection-oriented
- Reliable
- Ordered delivery
- Retransmission
- Three-way handshake

Examples

- HTTP
- HTTPS
- SSH
- MySQL
- PostgreSQL

---

### UDP

- Connectionless
- No guarantee
- Faster
- No retransmission

Examples

- DNS
- Video Streaming
- Gaming
- VoIP

---

## Three-Way TCP Handshake

```
Client                Server

SYN ------------------>

      <--------------- SYN + ACK

ACK ------------------>
```

Purpose

- Establishes connection
- Synchronizes sequence numbers
- Confirms both sides are ready

---

## Four-Way TCP Connection Termination

```
Client                 Server

FIN ------------------>

      <--------------- ACK

      <--------------- FIN

ACK ------------------>
```

---

## Why TCP needs a Handshake?

- Verify both endpoints are reachable
- Synchronize sequence numbers
- Prevent stale connections
- Establish reliable communication

---

## Why UDP doesn't need Handshake?

UDP is connectionless and sends packets immediately without maintaining connection state.

---

## HTTP vs HTTPS

HTTP

- Plain text
- Port 80
- Not encrypted

HTTPS

- HTTP over TLS
- Port 443
- Encrypted
- Secure

---

## What is SSL/TLS?

TLS encrypts communication between client and server to provide:

- Encryption
- Authentication
- Data integrity

SSL is deprecated; modern systems use TLS.

---

## HTTPS Handshake (Simplified)

```
Client
   |
Client Hello
   |
Server Hello
Certificate
Public Key
   |
Key Exchange
   |
Session Key Created
   |
Encrypted Communication
```

---

## Symmetric vs Asymmetric Encryption

### Symmetric

- Same key for encrypt/decrypt
- Fast

Examples

- AES

---

### Asymmetric

- Public Key
- Private Key
- Slower

Examples

- RSA
- ECC

TLS uses asymmetric encryption to establish a shared symmetric session key.

---

## DNS

DNS converts

```
google.com

↓

142.xx.xx.xx
```

Without DNS we'd have to remember IP addresses.

---

## DNS Resolution

```
Browser

↓

Local DNS Cache

↓

ISP DNS

↓

Root DNS

↓

TLD (.com)

↓

Authoritative DNS

↓

IP Address
```

---

## What is a Socket?

A socket uniquely identifies a network connection.

```
Socket

=

IP Address
+
Port Number
```

---

## IP Address vs Port

IP identifies the machine.

Port identifies the application running on that machine.

Example

```
192.168.1.20:8080
```

---

## What is a Load Balancer?

Distributes incoming traffic across multiple backend servers to improve:

- Availability
- Scalability
- Fault tolerance

---

## Layer 4 vs Layer 7 Load Balancer

### Layer 4

Works using

- TCP
- UDP

Routes traffic based on IP and Port.

Examples

- AWS NLB

---

### Layer 7

Works using HTTP/HTTPS.

Can route based on

- URL
- Headers
- Cookies
- Hostname

Examples

- NGINX
- HAProxy
- AWS ALB

---

## Reverse Proxy vs Forward Proxy

### Reverse Proxy

Sits in front of servers.

Examples

- NGINX
- HAProxy

Purpose

- Load balancing
- SSL termination
- Caching

---

### Forward Proxy

Sits in front of clients.

Purpose

- Hide client identity
- Internet filtering

---

## API Gateway vs Load Balancer

Load Balancer

- Routes traffic
- Layer 4/7

API Gateway

- Authentication
- Rate limiting
- Routing
- Request transformation
- Logging

---

## REST vs gRPC

REST

- HTTP/JSON
- Human readable
- Easy debugging

gRPC

- HTTP/2
- Protocol Buffers
- Faster
- Streaming support

---

## HTTP Methods

- GET → Read
- POST → Create
- PUT → Replace
- PATCH → Partial Update
- DELETE → Delete

---

## HTTP Status Codes

2xx → Success

3xx → Redirection

4xx → Client Error

5xx → Server Error

Common ones:

- 200 OK
- 201 Created
- 204 No Content
- 301 Moved Permanently
- 400 Bad Request
- 401 Unauthorized
- 403 Forbidden
- 404 Not Found
- 409 Conflict
- 429 Too Many Requests
- 500 Internal Server Error
- 502 Bad Gateway
- 503 Service Unavailable

---

## Keep-Alive

Keeps TCP connection open for multiple HTTP requests, reducing latency by avoiding repeated handshakes.

---

## HTTP/1.1 vs HTTP/2 vs HTTP/3

### HTTP/1.1

- One request per connection (or pipelining)
- Head-of-line blocking
- Text-based

---

### HTTP/2

- Binary protocol
- Multiplexing
- Header compression
- Single TCP connection

---

### HTTP/3

- Uses QUIC over UDP
- Faster connection setup
- Better performance on lossy networks
- Eliminates TCP head-of-line blocking

### How whitelisting IP addresses works?

Whitelisting IP addresses involves maintaining a list of approved IP addresses that are allowed to access a particular resource or system. When a request is received, the system checks if the requesting IP address is in the whitelist. If it is, the request is allowed to proceed; otherwise, it is blocked.

### how a router with an IP allows multiple devices to connect to the internet?
A router with an IP address allows multiple devices to connect to the internet by using a process called Network Address Translation (NAT). The router has a public IP address assigned by the Internet Service Provider (ISP) and manages a private IP address range for the devices connected to it. When a device sends a request to access the internet, the router translates the private IP address of the device into its public IP address and forwards the request. When the response comes back from the internet, the router translates the public IP address back to the corresponding private IP address of the device and delivers the response. This way, multiple devices can share a single public IP address while maintaining unique private IP addresses within the local network.



---

# Top 20 Networking Questions Asked in Interviews

1. Explain the OSI model.
2. Explain the TCP/IP model.
3. TCP vs UDP.
4. Explain the TCP three-way handshake.
5. Explain TCP connection termination.
6. Why is TCP reliable?
7. What happens when you type `google.com` in a browser?
8. HTTP vs HTTPS.
9. Explain the TLS handshake.
10. Symmetric vs Asymmetric encryption.
11. What is DNS and how does it work?
12. What is a socket?
13. IP address vs Port.
14. Reverse Proxy vs Forward Proxy.
15. Layer 4 vs Layer 7 Load Balancer.
16. API Gateway vs Load Balancer.
17. REST vs gRPC.
18. HTTP/1.1 vs HTTP/2 vs HTTP/3.
19. Explain common HTTP status codes.
20. What is Keep-Alive and why is it useful?


