resource "aws_security_group" "kubernetes_nodes" {
  name        = "kubernetes-nodes-sg"
  description = "Security group for Kubernetes nodes"
  vpc_id      = aws_vpc.main.id

  # Define ingress and egress rules as needed
  # Example: allow SSH and required ports for Kubernetes nodes
}

resource "aws_security_group" "kubernetes_control_plane" {
  name        = "kubernetes-control-plane-sg"
  description = "Security group for Kubernetes control plane"
  vpc_id      = aws_vpc.main.id

  # Define ingress rules for Kubernetes control plane
}
