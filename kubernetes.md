# Kubernetes Interview Notes (Backend / Golang)

This repository contains the most frequently asked Kubernetes interview questions for Backend and Golang developers (3–8 years experience).

---

# 1. Kubernetes Components (1-2 Line Definitions)

## Pod
The smallest deployable unit in Kubernetes that contains one or more containers sharing the same network and storage.

---

## Node
A worker machine (VM or physical server) where Pods run. Every node contains kubelet, container runtime, and kube-proxy.

---

## Cluster
A collection of master (control plane) and worker nodes that run containerized applications.

---

## Deployment
Manages stateless applications by creating ReplicaSets and supports rolling updates and rollbacks.

---

## ReplicaSet
Ensures that the desired number of Pod replicas are always running.

---

## StatefulSet
Manages stateful applications by providing stable Pod names, ordered deployment, and persistent storage.

---

## DaemonSet
Ensures one Pod runs on every node (or selected nodes). Commonly used for logging and monitoring agents.

---

## Job
Creates Pods to complete a one-time task. Once completed successfully, the Job finishes.

---

## CronJob
Runs Jobs on a scheduled interval similar to Linux cron.

---

## Service
Provides a stable endpoint to access Pods, even if Pod IPs change.

---

## ClusterIP
Default Service type that exposes an application only inside the Kubernetes cluster.

---

## NodePort
Exposes an application externally using `<NodeIP>:Port`.

---

## LoadBalancer
Creates a cloud provider load balancer (AWS ELB, Azure LB, GCP LB) to expose services externally.

---

## ExternalName
Maps a Kubernetes Service to an external DNS name.

---

## Ingress
Provides Layer-7 HTTP/HTTPS routing to multiple Services using a single external IP.

---

## Ingress Controller
A controller (NGINX, Traefik, HAProxy, etc.) that watches Ingress resources and configures routing rules.

---

## ConfigMap
Stores non-sensitive configuration like URLs, feature flags, and application settings.

---

## Secret
Stores sensitive information like passwords, API keys, certificates, and tokens.

---

## Namespace
Provides logical isolation of Kubernetes resources within the same cluster.

---

## Label
A key-value pair attached to Kubernetes resources for identification.

---

## Selector
Finds Kubernetes resources based on matching labels.

---

## Taints
Prevent Pods from being scheduled onto specific nodes.

---

## Tolerations
Allow Pods to run on nodes that have matching taints.

---

## Node Affinity
Forces or prefers Pods to run on selected nodes.

---

## Pod Affinity
Schedules Pods close to other Pods.

---

## Pod Anti-Affinity
Spreads Pods across different nodes for better availability.

---

## Persistent Volume (PV)
A storage resource inside the cluster independent of Pod lifecycle.

---

## Persistent Volume Claim (PVC)
A request for storage made by applications.

---

## StorageClass
Automatically provisions Persistent Volumes dynamically.

---

## Volume
A directory accessible by containers inside a Pod that persists across container restarts.

---

## kube-apiserver
The entry point to Kubernetes. All kubectl commands communicate through the API Server.

---

## etcd
Distributed key-value database storing the complete cluster state.

---

## Scheduler
Assigns newly created Pods to appropriate worker nodes.

---

## Controller Manager
Runs controllers that continuously reconcile the desired and current state.

---

## kubelet
Agent running on every node that communicates with the API Server and manages Pods.

---

## kube-proxy
Handles networking and load balancing for Kubernetes Services.

---

## Container Runtime
Software responsible for running containers (containerd, CRI-O).

---

## CRD (Custom Resource Definition)
Extends Kubernetes by allowing users to create custom resource types.

---

## Operator
A Kubernetes controller that automates application-specific operational tasks using CRDs.

---

## NetworkPolicy
Controls which Pods are allowed to communicate with other Pods or external endpoints.

---

## Horizontal Pod Autoscaler (HPA)
Automatically scales the number of Pods based on CPU, memory, or custom metrics.

---

## Vertical Pod Autoscaler (VPA)
Automatically adjusts CPU and memory requests/limits for Pods.

---

## Cluster Autoscaler
Automatically adds or removes worker nodes based on cluster demand.

---

# 2. Frequently Asked Differences

## Deployment vs StatefulSet

**Deployment**
- Stateless applications
- Pods are interchangeable
- Random Pod names
- Rolling updates

**StatefulSet**
- Stateful applications
- Stable Pod identity
- Stable storage
- Ordered deployment and deletion

Examples:
Deployment → API Server, Web Server

StatefulSet → MongoDB, Kafka, Cassandra

---

## Deployment vs DaemonSet

Deployment runs a specified number of replicas.

DaemonSet ensures exactly one Pod runs on every worker node.

Example:
- Deployment → Backend API
- DaemonSet → Fluentd, Node Exporter

---

## Deployment vs Job vs CronJob

Deployment runs continuously.

Job runs once and exits.

CronJob runs Jobs on a schedule.

---

## Pod vs Deployment

Pod is the smallest execution unit.

Deployment manages Pods, ReplicaSets, updates, rollbacks, and self-healing.

---

## Deployment vs ReplicaSet

ReplicaSet only maintains the desired number of Pods.

Deployment manages ReplicaSets and supports rolling updates and rollback.

---

## Pod vs Container

Container is the running application process.

Pod is the Kubernetes wrapper around one or more containers.

---

## Service vs Ingress

Service works at Layer-4 (TCP/UDP) and exposes Pods.

Ingress works at Layer-7 (HTTP/HTTPS) and routes traffic to multiple Services.

---

## Ingress vs NodePort

NodePort exposes an application using `<NodeIP>:Port`.

Ingress provides domain-based routing, SSL termination, and path-based routing through a single IP.

---

## ClusterIP vs NodePort vs LoadBalancer

ClusterIP
- Internal communication only

NodePort
- External access through Node IP

LoadBalancer
- External access through Cloud Load Balancer

---

## ConfigMap vs Secret

ConfigMap stores non-sensitive configuration.

Secret stores passwords, certificates, API keys, and tokens.

---

## Liveness Probe vs Readiness Probe

Liveness Probe
- Checks whether the application is alive.
- Failure restarts the Pod.

Readiness Probe
- Checks whether the application is ready to receive traffic.
- Failure removes the Pod from Service endpoints.

---

## Requests vs Limits

Request is the minimum resource guaranteed.

Limit is the maximum resource the container can consume.

---

## PV vs PVC

PV is the actual storage resource.

PVC is the application's request for storage.

---

## Label vs Selector

Labels identify resources.

Selectors find resources using matching labels.

---

## Taints vs Tolerations

Taints repel Pods from nodes.

Tolerations allow Pods onto tainted nodes.

---

## Node Affinity vs Pod Affinity

Node Affinity schedules Pods onto selected nodes.

Pod Affinity schedules Pods near other Pods.

---

## Horizontal Pod Autoscaler vs Cluster Autoscaler

HPA scales Pods.

Cluster Autoscaler scales Nodes.

---

## Rolling Update vs Recreate

Rolling Update replaces Pods gradually without downtime.

Recreate deletes all Pods first, then creates new ones, causing downtime.

---

## CRD vs Operator

CRD defines a new Kubernetes resource.

Operator watches CRDs and automates lifecycle management.

---

# 3. Critical Kubernetes Interview Questions

## Kubernetes Architecture

- Explain the Kubernetes Control Plane.
- Explain Worker Node components.
- How does a request travel from `kubectl` to a running Pod?

---

## Pod Lifecycle

- What happens when a Pod is created?
- What happens if a Pod crashes?
- Explain Pod phases.

---

## Deployment

- Explain Rolling Update.
- Explain Rollback.
- What happens when replicas are increased?
- What happens during `kubectl apply`?

---

## Networking

- How do Pods communicate?
- Why do Pods communicate through Services instead of Pod IPs?
- Explain DNS inside Kubernetes.
- Explain kube-proxy.

---

## Scheduling

- How does the Scheduler choose a node?
- Explain Node Affinity.
- Explain Pod Affinity.
- Explain Taints and Tolerations.

---

## Storage

- Explain PV, PVC, and StorageClass.
- What happens if a Pod using PVC is deleted?
- Why is StatefulSet used with Persistent Volumes?

---

## Security

- ConfigMap vs Secret
- NetworkPolicy
- RBAC
- ServiceAccount

---

## Troubleshooting

- Pod stuck in Pending
- CrashLoopBackOff
- ImagePullBackOff
- OOMKilled
- Evicted Pods
- Failed Scheduling

Useful commands:

```bash
kubectl get pods
kubectl describe pod <pod-name>
kubectl logs <pod-name>
kubectl exec -it <pod-name> -- sh
kubectl get events
kubectl top pod
kubectl top node
```

---

what is role, clusterrole, rolebinding and clusterrolebinding in kubernetes?
In Kubernetes, Role, ClusterRole, RoleBinding, and ClusterRoleBinding are components of the Role-Based Access Control (RBAC) system that manage permissions for users and service accounts. Here's a brief explanation of each:

**Role**:
- Defines a set of permissions within a specific namespace.
- Can only grant access to resources within that namespace.

**ClusterRole**:
- Similar to Role but operates at the cluster level.
- Can grant access to resources across all namespaces.

**RoleBinding**:
- Binds a Role (or ClusterRole) to a user or group within a specific namespace.
- Grants the permissions defined in the Role to the specified user or group.

**ClusterRoleBinding**:
- Binds a ClusterRole to a user or group at the cluster level.
- Grants the permissions defined in the ClusterRole to the specified user or group.

-------

## Autoscaling

- HPA
- VPA
- Cluster Autoscaler
- Metrics Server

---

## Operators & CRDs

- What is a CRD?
- What is an Operator?
- Explain the Reconciliation Loop.
- How does a Kubernetes Controller work?

---

## Most Frequently Asked Production Questions

### What happens when you run `kubectl apply deployment.yaml`?

Expected flow:

```
kubectl
      ↓
API Server
      ↓
etcd stores desired state
      ↓
Deployment Controller notices change
      ↓
ReplicaSet created
      ↓
Scheduler selects Node
      ↓
kubelet pulls image
      ↓
Container Runtime starts Pod
      ↓
Service discovers Pod
      ↓
Application becomes Ready
```

---

### What happens when a Pod crashes?

```
Application crashes
        ↓
Liveness Probe fails
        ↓
kubelet restarts Pod
        ↓
If Pod is deleted
        ↓
ReplicaSet creates new Pod
        ↓
Service automatically routes traffic
```

---

### What happens when a Node goes down?

```
Node becomes NotReady
        ↓
Pods marked unhealthy
        ↓
Controller detects failure
        ↓
Scheduler chooses another node
        ↓
New Pods created
        ↓
Traffic resumes
```

---

## Top 20 Must-Prepare Questions

1. Explain Kubernetes Architecture.
2. Explain Pod Lifecycle.
3. Deployment vs StatefulSet.
4. Deployment vs DaemonSet.
5. Deployment vs ReplicaSet.
6. Service vs Ingress.
7. ClusterIP vs NodePort vs LoadBalancer.
8. ConfigMap vs Secret.
9. Liveness vs Readiness Probe.
10. Requests vs Limits.
11. PV vs PVC.
12. Taints vs Tolerations.
13. Node Affinity vs Pod Affinity.
14. Rolling Update vs Rollback.
15. HPA vs Cluster Autoscaler.
16. CrashLoopBackOff vs ImagePullBackOff.
17. What happens during `kubectl apply`?
18. Explain the Kubernetes Reconciliation Loop.
19. Explain CRD and Operator.
20. How do you troubleshoot a production Kubernetes issue?


what is init, sidecar and main container in a pod?
In a Kubernetes Pod, there are different types of containers that serve specific purposes:
1. **Init Container**: 
   - Init containers are specialized containers that run before the main application containers in a Pod. They are used to perform initialization tasks, such as setting up the environment, checking dependencies, or preparing data before the main application starts. Init containers run sequentially and must complete successfully before the main containers are started.
2. **Sidecar Container**:
   - Sidecar containers are auxiliary containers that run alongside the main application container(s) in a Pod. They provide additional functionality or support to the main application, such as logging, monitoring, or proxying. Sidecar containers share the same network namespace and storage volumes as the main container, allowing them to interact closely with the main application.
3. **Main Container**:
   - The main container is the primary application container that runs the core functionality of the Pod. It is the container that performs the main tasks or services that the Pod is designed to provide. The main container is typically the focus of the Pod's purpose, while init and sidecar containers support its operation.    

