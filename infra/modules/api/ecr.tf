locals {
  api_dir = "../../../api"
}

# ECR
resource "aws_ecr_repository" "lambda_ecr" {
  name = "${var.product_name}-${var.env}"
}

data "aws_ecr_authorization_token" "token" {}

resource "null_resource" "image_push" {
  triggers = {
    always_run = "${timestamp()}"
  }
  
  provisioner "local-exec" {
    command = <<-EOF
      docker build --platform linux/amd64 ${local.api_dir} -f ${local.api_dir}/Dockerfile.lambda -t ${aws_ecr_repository.lambda_ecr.repository_url}:latest; \
      docker login -u AWS -p ${data.aws_ecr_authorization_token.token.password} ${data.aws_ecr_authorization_token.token.proxy_endpoint}; \
      docker push ${aws_ecr_repository.lambda_ecr.repository_url}:latest
    EOF
  }
}
