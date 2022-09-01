# Infrastructure

## Container Orchestration

Automation of much the operational effort required to rn contanerized workloads and services. This includes a wide range of things software teams. e.g Docker Compose

Docker with Docker Compose :
- Without docker compose, build and run one container at a time
- Manually connect containers together
- Must be careful with dependencies and start-up order

- With docker compose, define multi container app in docker-cmpose.yaml file
- Single command to deploy entire application
- Handles container dependencies
- Works with Docker Swarm, Networking or Volumes

Docker Compose in three seteps process :
1. 
2. Build
3. Run

## Building Blocks of Docker Compose

### Service
- Configuraion of a computing resource within application
- A service is defined by Docker Image and set of runtime arguments

### Volume
- Voumes are persistent data stores
- Allows the configuration of named volumes that can be reused across multiple services

## Docker Playground
- Online docker practice

## Pros Container Orchestration
- Simplified Operations : handles all the deployment operatios and the networking
- Resilience : automatically restart or scale a container or a cluster
- Added security : reducing or eliminating the chance of human error

## eFishery uses Nomad
Nomad, Consul, Vault -> Hashicorp Stack used by eFishery

