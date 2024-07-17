resource "aws_security_group" "kubernetes_nodes" {
  name        = "kubernetes-nodes-sg"
  description = "Security group for Kubernetes nodes"
  vpc_id      = aws_vpc.main.id
}

resource "aws_security_group" "kubernetes_control_plane" {
  name        = "kubernetes-control-plane-sg"
  description = "Security group for Kubernetes control plane"
  vpc_id      = aws_vpc.main.id
}
