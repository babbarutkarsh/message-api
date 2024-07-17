module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = "message-api-cluster"
  cluster_version = "1.21"   
  subnets         = aws_subnet.private[*].id
  vpc_id          = aws_vpc.main.id

  node_groups = {
    eks_nodes = {
      desired_capacity = 2
      max_capacity     = 3
      min_capacity     = 1
      instance_type    = "t3.medium"  
    }
  }
}
