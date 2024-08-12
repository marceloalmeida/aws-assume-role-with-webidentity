package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func assumeRoleWithWebIdentity() (*sts.AssumeRoleWithWebIdentityOutput, error) {
	// Read environment variables
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}
	roleArn := os.Getenv("AWS_ROLE_ARN")
	tokenFile := os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE")

	// Validate necessary environment variables
	if region == "" || roleArn == "" || tokenFile == "" {
		return nil, fmt.Errorf("required environment variables are missing: AWS_REGION/AWS_DEFAULT_REGION, AWS_ROLE_ARN, AWS_WEB_IDENTITY_TOKEN_FILE")
	}

	// Load the default AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	// Read the Web Identity Token
	webIdentityToken, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read web identity token file: %w", err)
	}

	// Create an STS client
	stsClient := sts.NewFromConfig(cfg)

	// Assume the role with the web identity token
	input := &sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          aws.String(roleArn),
		RoleSessionName:  aws.String("web_identity_session"),
		WebIdentityToken: aws.String(string(webIdentityToken)),
	}

	response, err := stsClient.AssumeRoleWithWebIdentity(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to assume role with web identity: %w", err)
	}

	return response, nil
}

func redactString(input string) string {
	if len(input) <= 4 {
		return "****"
	}
	return input[:4] + strings.Repeat("*", len(input)-8) + input[len(input)-4:]
}

func main() {
	// Define the no-redact flag
	noRedact := flag.Bool("no-redact", false, "Print credentials in plain text instead of redacting")
	flag.Parse()

	response, err := assumeRoleWithWebIdentity()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Extract credentials
	credentials := response.Credentials

	fmt.Println("Assumed Role ARN:", *response.AssumedRoleUser.Arn)

	if *noRedact {
		// Print credentials in plain text
		fmt.Println("Access Key:", *credentials.AccessKeyId)
		fmt.Println("Secret Key:", *credentials.SecretAccessKey)
		fmt.Println("Session Token:", *credentials.SessionToken)
	} else {
		// Print redacted credentials
		fmt.Println("Access Key:", redactString(*credentials.AccessKeyId))
		fmt.Println("Secret Key:", redactString(*credentials.SecretAccessKey))
		fmt.Println("Session Token:", redactString(*credentials.SessionToken))
	}
}
