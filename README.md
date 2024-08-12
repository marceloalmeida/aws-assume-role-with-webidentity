# AWS Assume Role with Web Identity Token in Go
This repository contains a Go script that demonstrates how to assume an AWS role using a Web Identity Token. The script reads necessary environment variables, assumes the role via AWS STS (Security Token Service), and securely handles sensitive information such as Access Key ID, Secret Access Key, and Session Token.

## Features
- Environment Variable Configuration: Reads AWS configuration from environment variables, including `AWS_REGION`, `AWS_ROLE_ARN`, and `AWS_WEB_IDENTITY_TOKEN_FILE`.
- Assume Role with Web Identity: Uses AWS STS to assume an IAM role with a Web Identity Token.
- Credential Redaction: By default, the script redacts the Access Key ID, Secret Access Key, and Session Token to avoid exposing sensitive information.
- Optional Plain Text Output: Use the `--no-redact` flag to print the credentials in plain text if needed.
- Error Handling: Comprehensive error handling ensures clear feedback when environment variables are missing or when the AWS API call fails.

## Usage
1. Set Environment Variables:
  - `AWS_REGION` or `AWS_DEFAULT_REGION`
  - `AWS_ROLE_ARN`
  - `AWS_WEB_IDENTITY_TOKEN_FILE`
2. Run the Script:
```bash
go run main.go
```

By default, the script will output redacted credentials.

3. Optional Plain Text Output:
```bash
go run main.go --no-redact
```

Use the --no-redact flag to output the credentials in plain text.

## Prerequisites
- Go 1.16 or later
- AWS SDK for Go v2

## Installation
Install the required dependencies:

```bash
go get -u github.com/aws/aws-sdk-go-v2/aws
go get -u github.com/aws/aws-sdk-go-v2/config
go get -u github.com/aws/aws-sdk-go-v2/service/sts
```

## Download and run pre-compiled binary
```bash
curl -s -L -O https://github.com/marceloalmeida/aws-assume-role-with-webidentity/releases/download/1.0.0/aws-assume-role-with-webidentity-1.0.0-linux-amd64.tar.gz
tar -zxvf aws-assume-role-with-webidentity-1.0.0-linux-amd64.tar.gz
chmod a+x aws-assume-role-with-webidentity-1.0.0-linux-amd64
./aws-assume-role-with-webidentity-1.0.0-linux-amd64
```

## License
This project is licensed under the MIT License.
