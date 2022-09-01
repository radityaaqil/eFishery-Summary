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

## CI/CD
A practice that is done after the CI process is complete and all the code is successfull integrated.

PROS :
- Get feedback faster
- Detect bugs faster
- Speed up the release process

## Tools
Jenkins, Gitlab, Azure DevOps, Github Action

### GitHub Action :
- Automate, customize and execute your software development workflows right in your repository.

Plan --> Code --> Build --> Test --> Release --> Deploy --> Operate

- Event 
We could create countless amount of 'Runner' 
- Runner 1 : Run action, run script
- Runner 2 : Run action, run script

Workflow --> Configurable automated process that will run one or more jobs. Defined by a YAML file checked in to your repository and will run when triggered

Event --> An event is a specific activit in a repository that triggers a workflow run. e.g push, commit, pull, merge

Jobs --> Activity which is done by the Runner

Actions --> A custom application from github

## Observability

The ability to measure a system's current state based on the data it generates such as logs, metrics and traces.

PROS :
- Application performace monitoring
- DevSecOps and SRE 9Site Reliability Engineer)
- Infrastructure, cloud and Kubernetes monitoring
- End-user experience
- Business analytics

Challanges :
- Data silos
- Volume, velocity, variety and complexity
- Manual intrumentation and configuration
- Lack of pre-production (hard to test)
- Wasting time troubleshooting

Tools :
- Datadog
- Elastic APM
- New Relic

## Awesomeness of eFishery Infrastructure

eFishery Deliver :
- 150+ deployment/day
- 70+ TB Data Transfer/month
- 100 requests/month
- 20+ requests/week
- 500 Nomad Jobs
- 700+ running Containers
- >300 GB data backup/day
- 100 Servers, 1,2 TB Memory,  710 GHz CPU, 25 TB Storage