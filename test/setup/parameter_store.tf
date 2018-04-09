resource "aws_ssm_parameter" "string" {
  name  = "/test/string"
  type  = "String"
  value = "foo"
}

resource "aws_ssm_parameter" "secure_string" {
  name  = "/test/secure-string"
  type  = "SecureString"
  value = "bar"
}
