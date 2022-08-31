# CLOUD NATIVE

Cloud-native is a modern approach to building and running software applications that exploits the flexibility, scalability, and resilience of cloud computing. Cloud-native encompasses the various tools and techniques used by software developers today to build applications for the public cloud, as opposed to traditional architectures suited to an on-premises data center.

## Cloud Native Roadmap

1. Containerization
2. CI/CD
3. Orchestration
4. Observability & Analysis 
5. Service Mesh
6. Networking
7. Distributed Database
8. Messaging
9. Container Runtime
10. Software Distribution

## Modern Design

How would you design a cloud-native app? What would your architecture do? --> The Twelve-Factor Application

1. Codebase --> Single code base for each microservice, stored in its own repo. Tracked with version control, it can deploy to multiple environments (QA, Staging, Production)
2. Dependencies --> Each microservice isolates and packages its own dependencies, embracing changes without impacting the entire system.
3. Configuration --> Configuration information is moved out of the microservice and externalized through a configuration management tool outside of the code. The 
4. Backing Service --> data stores, caches, message brokers
5. Build, Release, Run --> Each release must enforce a strict separation across the build, release, and run stages. Each should be tagged with a unique ID and support the ability to roll back. Modern CI/CD system help fulfill this principle.
6. Process --> Each microservice should execute in its own process, isolated from other running services. Externalize require state to a backing service such as a distributed cache or data store.
7. Port Binding --> Each microservice should be self-contained wiht its interfaces and functionality exposed on its own port. Doing so provides isolation from other microservices.
8. Concurrency --> When capacity needs to increase, scale out services horizontally across multiple identical processes (copies) as opposed to scaling-up a single large instance on th emost powerful machine available. Develop the application to be concurrent making scaling out in cloud environments seamless.
9. Development/Production Parity --> Keep environments across the application lifecycle as similar as possible, avoiding costly shortcuts. Here the adoption of containers can greatly contribute by promoting the same execution environment.
10. Logging --> Treat logs as event streams.
11. Admin Process --> Run admin/management tasks as one off processes
12. Disposability --> Easily disposed (fast startup and graceful shutdown)

## CONTAINERS

An abstraction at the app layer that packages the code and the dependencies together

## Docker Registry
Common docker term
- A registry is a service containing repositories of images that respond to the registry API.
- A local registry is where are stored the images that you built or pulled.
- A remote or hosted registry is where you pull images from or push images to.
- Dockerfile : text document
- Docker Image : read-only template with instructios for creating a docker container.
- Docker Container : runnable instance of an image
