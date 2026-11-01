# Kubernetes Blog Operator

A Kubernetes operator written in Go to provision and manage blog instances declaratively using custom resources.

> **Project status:** `BlogInstance` API definition available

The project structure has been initialized with Kubebuilder. The `BlogInstance` custom resource and its reconciliation logic will implemented pregressively.

## Overview

The goal of this project is to provide a Kubernetes-native way to deploy and manage blog instances.

Instead of manually creating a `Deployment`, `Service`, `ConfigMap` and `Ingress`, users will declares a single `BlogInstance` resource:

```yaml
apiVersion: plateform.alizaman.fr/v1alpha1
kind: BlogInstance
metadata:
  name: demo-blog
spec:
  title: My Blog
  replicas: 2
```

The operator will observe each `BlogInstance` and ensure that the required Kubernetes resources exist and match the desired state.

> The `BlogInstance` API shown above is part of the planned implementation and is not available yet.

## Planned architecture

```text
BlogInstance Custom Resource
            |
            v
  BlogInstance Controller
            |
            +-- ConfigMap
            +-- Deployment
            +-- Service
            +-- Ingress
```

## Planned feature

- Declarative `BlogInstance` custom resource
- Idempotent reconciliation loop
- Automatic creation and update of Kubernetes resources
- Self-healing when managed resources are modified or deleted
- Resource ownership and lifecycle management
- Status conditions and readiness reporting
- custom resource validation
- Unit and integration tests
- Local deployment on a Kubernetes cluster
- Continuous integration pipeline

## Technology stack

- Go
- Kubernetes
- Kubebuilder
- controller-runtime
- Kustomize
- Docker

## Getting Started

### Requirements

The following tools are required to build and develop the project:

- Go
- Docker
- kubectl
- kubebuilder
- Access to a Kubernetes cluster for integration testing

A local Kubernetes cluster such as kind or Minikube can be used during development.

### Build

```bash
make build
```

### Run tests

```bash
make test
```

## Roadmap

- [x] Initialize the Kubebuilder project
- [x] Define the `BlogInstance` API
- [x] Generate the Custom Resource Definition
- [ ] Implement the initial reconciliation loop
- [ ] Manage blog configuration with a ConfigMap
- [ ] Manage the blog Deployment
- [ ] Expose the blog with a Service
- [ ] Add optional Ingess support
- [ ] Report readiness through status conditions
- [ ] Manage resource ownership and deletion
- [ ] Add unit and integration tests
- [ ] Containerize and deploy the operator
- [ ] Add continuous integration
- [ ] Add metrics, events and health checks

## Project structure 

```text

.
├── cmd/            # Operator entry point
├── config/         # Kubernetes and Kustomize manifests
├── hack/           # Development scripts and tools
├── Dockerfile      # Operator container image
├── Makefile        # Build, Test, generation and deployment commands
├── PROJECT         # Kubevuilder project metadata
├── go.mod          # Go module definition
└── README.md       # Project documentation
```

The project structure will evolve as the `BlogInstance` API and controller are added

## Development principles

This project aims to follow Kubernetes operator development best practices:

- Declarative resource management
- Idempotent reconciliation
- Least-privilege RBAC permissions
- Clear ownership of managed resource
- Observable status and errors
- Automated testing
- Reproductible builds and deployments

## License

This project is licensed under the Apache License 2.0.