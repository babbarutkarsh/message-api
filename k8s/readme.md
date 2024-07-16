# Kubernetes Files Overview and Helm CRD Requirements

This README provides an overview of the Kubernetes YAML files included in this repository and the Helm CRD requirements for deploying and managing these resources.

## Kubernetes Files

### 1. `deployment.yaml`

Description: Defines the Deployment for the `message-api` application.

Usage:
```bash
kubectl apply -f deployment.yaml
```

### 2. `disruption-handler.yaml`

Description: Specifies the Pod Disruption Budget (PDB) for managing pod disruptions.

Usage:
```bash
kubectl apply -f disruption-handler.yaml
```

### 3. `hpa.yaml`

Description: Configures Horizontal Pod Autoscaler (HPA) for autoscaling based on CPU utilization.

Usage:
```bash
kubectl apply -f hpa.yaml
```

### 4. `postgres-service.yaml`

Description: Defines the Kubernetes Service for PostgreSQL database.

Usage:
```bash
kubectl apply -f postgres-service.yaml
```

### 5. `postgres_statefulset.yaml`

Description: Specifies the StatefulSet for deploying PostgreSQL database pods with persistent storage.

Usage:
```bash
kubectl apply -f postgres_statefulset.yaml
```

### 6. `prometheusrule.yaml`

Description: Defines Prometheus rule for custom alerting and monitoring rules.

Usage:
```bash
kubectl apply -f prometheusrule.yaml
```

### 7. `role-binding.yaml`

Description: Specifies the RoleBinding to bind roles to service accounts for RBAC.

Usage:
```bash
kubectl apply -f role-binding.yaml
```

### 8. `role.yaml`

Description: Defines RBAC Role for managing Kubernetes resources.

Usage:
```bash
kubectl apply -f role.yaml
```

### 9. `serviceaccount.yaml`

Description: Specifies the ServiceAccount for application-specific permissions.

Usage:
```bash
kubectl apply -f serviceaccount.yaml
```

### 10. `servicemonitor.yaml`

Description: Defines ServiceMonitor for monitoring services with Prometheus.

Usage:
```bash
kubectl apply -f servicemonitor.yaml
```

## Helm CRD Requirements

To manage and deploy Kubernetes resources effectively, ensure you have the following Helm CRD requirements configured:

### 1. `.env` Required

Ensure your `.env` file includes necessary environment variables like database credentials, API keys, etc.

### 2. Helm Repositories

```bash
# Add Prometheus Community Helm repository
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Add Stable Helm repository
helm repo add stable https://charts.helm.sh/stable
helm repo update
```

### 3. Install Prometheus using Helm

```bash
# Install Prometheus using Helm
helm install prometheus prometheus-community/kube-prometheus-stack
```