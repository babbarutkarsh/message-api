
# Message API

Welcome to the message-api repository! This README provides an overview of the project structure, setup instructions, and usage details for deploying and managing resources using Kubernetes, Terraform, Jenkins, and the Go message API.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Kubernetes Setup](#kubernetes-setup)
3. [Terraform Modules](#terraform-modules)
4. [Jenkinsfile](#jenkinsfile)
5. [Go Message API](#go-message-api)
6. [Usage](#usage)
7. [Contributing](#contributing)
8. [License](#license)

## Project Overview

This repository contains configurations and code for deploying and managing an API infrastructure using Kubernetes, Terraform for infrastructure provisioning, Jenkins for CI/CD pipelines, and a Go-based message API.

### Repository Structure

- **k8s/**: Contains Kubernetes YAML files for deployments, services, ingress, etc.
- **terraform/**: Includes Terraform modules for provisioning AWS resources such as EKS clusters.
- **Jenkinsfile**: Defines the Jenkins pipeline for automated CI/CD processes.
- **message-api/**: Source code for the Go-based message API.

## Kubernetes Setup

### Overview

The `k8s/` directory contains YAML files for deploying and managing Kubernetes resources such as deployments, services, ingress configurations, and more.

### Usage

Apply Kubernetes resources using `kubectl`:

```bash
kubectl apply -f k8s/
```

## Terraform Modules

### Overview

The `terraform/` directory contains Terraform configuration files for provisioning AWS resources, particularly an EKS cluster for hosting the API.

### Usage

1. Initialize Terraform:
   ```bash
   cd terraform/
   terraform init
   ```

2. Apply Terraform configurations:
   ```bash
   terraform apply -var-file=variables.tfvars
   ```

## Jenkinsfile

### Overview

The `Jenkinsfile` defines the Jenkins pipeline for automating build, test, and deployment tasks for the API project.

### Pipeline Stages

- **Build**: Compiles the Go-based message API.
- **Test**: Runs automated tests to validate API functionality.
- **Deploy**: Deploys API changes to Kubernetes cluster.

## Go Message API

### Overview

The `message-api/` directory contains the source code for the Go-based message API.

### Development

To build and run the API locally:

```bash
cd message-api/
go build
./message-api
```

## Usage

### Prerequisites

- Kubernetes cluster configured and accessible.
- AWS credentials and Terraform configured for infrastructure provisioning.
- Jenkins server with necessary plugins and agents set up.

### Steps

1. **Deploy Kubernetes Resources**:
   Apply YAML files in `k8s/` directory to deploy API components.

2. **Provision Infrastructure with Terraform**:
   Use Terraform in `terraform/` directory to create EKS cluster and required AWS resources.

3. **CI/CD with Jenkins**:
   Set up Jenkins using `Jenkinsfile` for automated build, test, and deployment workflows.

4. **Run Go Message API**:
   Build and run the Go-based message API from `message-api/` directory.

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request. For major changes, please open an issue first to discuss potential improvements.

## License

This project is licensed under the [MIT License](LICENSE).
