################################
# API Gateway
################################
resource "aws_apigatewayv2_api" "apigateway" {
  name = "${var.product_name}-api"
  protocol_type = "HTTP"
  depends_on = [ aws_cloudwatch_log_group.apigateway-log_group ]
  cors_configuration {
    allow_headers = [ "*" ]
    allow_methods = [ "*" ]
    allow_origins = [ "*" ]
  }
}

resource "aws_apigatewayv2_stage" "apigateway" {
  api_id = aws_apigatewayv2_api.apigateway.id
  auto_deploy = true
  name = "$default"
  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.apigateway-log_group.arn
    format = "$context.identity.sourceIp - - [$context.requestTime] \"$context.httpMethod $context.routeKey $context.protocol\" $context.status $context.responseLength $context.requestId"
  }
}

resource "aws_cloudwatch_log_group" "apigateway-log_group" {
  name              = "/aws/apigateway/${var.product_name}/${var.env}"
  retention_in_days = 3
}

resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id = aws_apigatewayv2_api.apigateway.id
  integration_type = "AWS_PROXY"

  integration_uri = aws_lambda_function.kns23_catch_up.arn
  integration_method = "ANY"
}

resource "aws_apigatewayv2_route" "lambda_integration" {
  api_id = aws_apigatewayv2_api.apigateway.id
  
  route_key = "ANY /{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}
