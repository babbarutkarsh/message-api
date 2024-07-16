# Terraform Folder Overview

This README provides an overview of the Terraform configuration files (`*.tf`) included in this directory for provisioning and managing resources in Amazon EKS (Elastic Kubernetes Service).

## Terraform Files

### 1. `eks-cluster.tf`

**Description:** Defines the configuration for provisioning an Amazon EKS cluster, including node groups, networking, and other cluster-specific settings.

**Usage:**
```bash
terraform apply -var-file=variables.tfvars
```

### 2. `main.tf`

**Description:** Contains the main configuration blocks for Terraform resources, providers, and modules used in the EKS cluster setup.

### 3. `output.tf`

**Description:** Specifies the output variables that are displayed after Terraform applies changes, such as cluster endpoint URLs or other important information.

### 4. `security-groups.tf`

**Description:** Defines security groups and their rules for controlling inbound and outbound traffic to and from resources within the EKS cluster.

## Terraform Commands

### Initialize Terraform

Before applying any changes, initialize Terraform to download necessary providers and modules:

```bash
terraform init
```

### Apply Terraform Changes

Apply the Terraform configuration to create or update resources in your AWS account:

```bash
terraform apply -var-file=variables.tfvars
```

### Destroy Terraform Resources

Destroy all resources managed by Terraform, including the EKS cluster and associated resources:

```bash
terraform destroy -var-file=variables.tfvars
```

## Output Variables

After applying Terraform changes, you can view important output variables using:

```bash
terraform output
```

These variables may include cluster endpoints, security group IDs, and other relevant information.

