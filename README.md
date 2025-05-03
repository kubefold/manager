# 🧬 Kubefold Manager

A utility service for managing protein folding tasks with AWS infrastructure support.

## About

Kubefold Manager is a versatile utility designed to handle various aspects of protein folding workflows. It provides functionality for:

1. Handling and validating input data
2. Uploading prediction artifacts to S3
3. Sending SMS notifications about task completion

Developed as part of the Kubefold project ecosystem for protein structure prediction.

## Usage

### Environment Variables

The application is configured using the following environment variables:

#### Required for all operations:
* `INPUT_PATH`: Directory path where input files are located
* `OUTPUT_PATH`: Directory path where output files are generated

#### For input processing:
* `ENCODED_INPUT`: Base64-encoded JSON input data for folding tasks

#### For artifact uploading:
* `BUCKET`: S3 bucket name where artifacts will be uploaded

#### For notifications:
* `NOTIFICATION_PHONES`: Comma-separated list of phone numbers to notify
* `NOTIFICATION_MESSAGE`: Custom message to send in the notification

### Docker

```bash
docker run -e INPUT_PATH=/data/input \
           -e OUTPUT_PATH=/data/output \
           -e BUCKET=my-result-bucket \
           -v /local/input:/data/input \
           -v /local/output:/data/output \
           kubefold/manager
```

### AWS Credentials

For S3 uploads and SNS notifications, the application uses the AWS SDK's default credential provider chain. Ensure appropriate AWS credentials are available through:

- Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
- AWS credentials file
- IAM roles for Amazon EC2 or ECS tasks

## Features

* Process and validate protein folding input data
* Upload prediction results to Amazon S3
* Send SMS notifications via Amazon SNS
* Containerized for easy deployment in cloud environments

## Building from Source

```bash
git clone https://github.com/kubefold/manager.git
cd manager
go build -o manager ./cmd/main.go
```

## Related Projects

* [kubefold/downloader](https://github.com/kubefold/downloader): Utility for downloading protein databases 