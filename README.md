Local development environment for aws serverless application
--

This repository contains a json api server using AWS *Lambda Function* and *DynamoDB*. It also contains some lambda functions written in golang, but they are just examples.

## About

This sample enables development of serverless application using AWS cloud services on your machine by using docker. The author has developed this on MacOS and has not considered other OS specific issues.

Here, [aws-sam-cli](https://github.com/awslabs/aws-sam-cli) is used to reproduce *Lambda Function* and calling it from *API Gateway*. And we use [localstack](https://github.com/localstack/localstack) to reproduce DynamoDB.

Other AWS services (ex. *S3*, *SQS*, etc) are also available as long as they are supported by localstack.

## Required

- aws-cli
- docker-compose
- golang (~1.11)
  - `export GO111MODULE=on`


## Getting Started

- Start/Stop/Restart the appication

  ```sh
  make start
  ```

  ```sh
  make stop
  ```

  ```sh
  make restart
  ```

- Compile Go source codes

  ```sh
  make build
  ```

- Watch the api server logs

  ```sh
  docker-compose logs -f sam
  ```

## Make a request to DynamoDB

  The port is 4569 by default.

  ```sh
  aws dynamodb list-tables --endpoint-url 'http://localhost:4569'
  ```

  ```sh
  aws dynamodb create-table --endpoint-url 'http://localhost:4569' \
    --table-name 'dummy' \
    --attribute-definitions '[{"AttributeName":"key","AttributeType": "S"}]' \
    --key-schema '[{"AttributeName":"key","KeyType": "HASH"}]' \
    --provisioned-throughput '{"ReadCapacityUnits": 5,"WriteCapacityUnits": 5}'
  ```

## Make a request to api server

  The port is 8000 by default. All of [api configuration](https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md) is specified in `lambda/template.yaml`. 

  - `/`: `lambda/hello`

    Hello World function.

    ```sh
    curl localhost:8000/
    ```
    ```json
    "Hello, 60.61.189.189"
    ```

  - `/tables`: `lambda/list-tables`

    Return table names in dynamodb.

    ```sh
    curl localhost:8000/list-tables
    ```

    ```json
    {
        "table_names": [
            "cities"
        ]
    }
    ```
