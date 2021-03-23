# CloudTrail Parser

To be used when creating a new IAM Role with lowest needed permissions.

Gets all CloudTrail Events for a User and creates an IAM policy for each event that has happened.

## How to use

Create a user with higher permissions than required

```
- PolicyName: ManageS3
  PolicyDocument:
    Statement:
    - Effect: Allow
    Action:
      - 's3:*'
```

Run the tool that needs to use the IAM role, for example a deployment.

Wait until the Events are available in TeamCity, then run this tool.

It will generate something like the below that can then be used for the final IAM user.

```
- PolicyName: ManageS3
  PolicyDocument:
    Statement:
    - Effect: Allow
    Action:
      - 's3:CreateBucket'
      - 's3:DeleteBucket'
```

## To run

### Using go

Clone the repo, update the config to the required values

`go run . > IamPolicies.yml`

#### Prerequisites

- Golang

### Using docker

Clone the repo, update the config to the required values

`docker build -t cloudtrailparser .`

`docker run -v ~/.aws:/root/.aws cloudtrailparser`

#### Prerequisites

- Docker
