# lambda
resource "aws_lambda_function" "kns23_catch_up" {
  depends_on = [
    null_resource.image_push,
  ]

  function_name = "${var.product_name}-${var.env}"
  package_type  = "Image"
  image_uri     = "${aws_ecr_repository.lambda_ecr.repository_url}:latest"
  role          = aws_iam_role.kns23_catch_up.arn

  source_code_hash = "${timestamp()}"
}

# IAM role
resource "aws_iam_role" "kns23_catch_up" {
  name               = "${var.product_name}-role_for_lambda-${var.env}"
  assume_role_policy = data.aws_iam_policy_document.lambda_kns_23_catch_up_assume.json
}
data "aws_iam_policy_document" "lambda_kns_23_catch_up_assume" {
  statement {
    effect = "Allow"

    actions = [
      "sts:AssumeRole",
    ]

    principals {
      type = "Service"
      identifiers = [
        "lambda.amazonaws.com",
      ]
    }
  }
}

resource "aws_iam_role_policy_attachment" "lambda-kns_23_catch_up" {
  role       = aws_iam_role.kns23_catch_up.name
  policy_arn = aws_iam_policy.lambda-kns_23_catch_up-custom.arn
}

resource "aws_iam_policy" "lambda-kns_23_catch_up-custom" {
  name   = "${var.product_name}-policy_for_lambda-${var.env}"
  policy = data.aws_iam_policy_document.lambda-kns_23_catch_up-custom.json
}
data "aws_iam_policy_document" "lambda-kns_23_catch_up-custom" {
  statement {
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "*",
    ]
  }
}

# Cloud Watch
resource "aws_cloudwatch_log_group" "lambda-kns23_catch_up" {
  name              = "/aws/lambda/${aws_lambda_function.kns23_catch_up.function_name}"
  retention_in_days = 3
}

resource "aws_lambda_permission" "allow_apigateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.kns23_catch_up.function_name
  principal     = "apigateway.amazonaws.com"
}
