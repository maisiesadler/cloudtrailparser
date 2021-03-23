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

Set environment variables

| Name | Value |
| -- | -- |
| `USERNAME` | Username of IAM user |
| `START_DATE` | Time to start looking from in the format `"2021-03-22 16:15:57 UTC+00:00"` |
| `END_DATE` | Time to finish looking from in the format `"2021-03-22 16:15:57 UTC+00:00"` |

### Using go

Clone the repo

`go run . > IamPolicies.yml`

#### Prerequisites

- Golang

### Using docker

Clone the repo

`docker build -t cloudtrailparser .`

`docker run --env USERNAME="iam.user" --env START_DATE="2021-03-23 11:15:57 UTC+00:00" --env END_DATE="2021-03-23 17:15:57 UTC+00:00" -v ~/.aws:/root/.aws cloudtrailparser`

#### Prerequisites

- Docker
