# AwsGo Project

This project uses the AWS Cloud Development Kit (CDK) with Go to define and deploy a serverless application on AWS.

## Architecture

The infrastructure consists of the following components:

*   **AWS Lambda**: A Go-based Lambda function (packaged in `lambda/function.zip`) handles the application logic.
    *   The Lambda function code is structured into several packages:
        *   `lambda/api`: Handles API request/response logic.
        *   `lambda/app`: Contains the core application business logic.
        *   `lambda/database`: Manages interactions with DynamoDB.
        *   `lambda/middleware`: Provides middleware functionalities for request processing.
        *   `lambda/types`: Defines data structures and types used across the Lambda function.
*   **Amazon DynamoDB**: A NoSQL database table named `userTable` is used to store user data. The partition key for this table is `username`. The Lambda function has read and write permissions to this table.
*   **Amazon API Gateway**: An HTTP API Gateway exposes the Lambda function through the following RESTful endpoints:
    *   `POST /register`: Endpoint for user registration.
    *   `POST /login`: Endpoint for user login.
    *   `GET /protected`: A protected endpoint, presumably requiring authentication.

## AWS CDK Infrastructure (aws_go.go)

The `aws_go.go` file defines the AWS infrastructure using the AWS CDK. It sets up:
1.  The DynamoDB table (`myUserTable`).
2.  The Lambda function (`myLambdaFunction`), configured to use the code from `lambda/function.zip` and the `PROVIDED_AL2023` runtime.
3.  Permissions for the Lambda function to access the DynamoDB table.
4.  The API Gateway (`myApiGateway`) with CORS enabled and request logging.
5.  Lambda integrations for the `/register`, `/login`, and `/protected` resources.

## Lambda Function Details (lambda/ directory)

The `lambda/` directory contains the source code and build artifacts for the AWS Lambda function.
*   `main.go`: The entry point for the Lambda function.
*   `go.mod` & `go.sum`: Manage Go module dependencies for the Lambda function.
*   `Makefile`: Likely contains build commands for the Lambda function (e.g., `make build` to create `function.zip`).
*   `function.zip`: The deployment package for the Lambda function.
*   `bootstrap`: The executable for the custom Lambda runtime.


## Prerequisites

*   Go (version specified in `go.mod` and `lambda/go.mod`)
*   AWS CDK
*   AWS CLI, configured with appropriate credentials

## Deployment

1.  **Build the Lambda function:**
    Navigate to the `lambda/` directory and run the build command (e.g., `make build` if defined in the `Makefile`). This should create/update the `lambda/function.zip` file.
    ```bash
    cd lambda
    make # Or your specific build command
    cd ..
    ```

2.  **Deploy the CDK stack:**
    Ensure your AWS credentials are configured correctly. Then, from the project root directory, run:
    ```bash
    cdk deploy
    ```

## Testing

The `aws_go_test.go` file contains commented-out example tests. These can be uncommented and expanded to test the CDK stack definition.

To run the tests (after uncommenting and potentially adapting them):
```bash
go test
```

## Cleaning Up

To remove the deployed AWS resources, run:
```bash
cdk destroy
```
