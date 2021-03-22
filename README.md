# CloudTrail Parser

Gets all CloudTrail Events for a User and creates an IAM policy for each event that has happened.

This means you can create IAM roles with higher privilige, run a deployment/tool. Then use this tool to generate a policy with just the required permissions.

## To run

Clone the repo, update the config to the required values

`go run . > IamPolicies.yml`

The policies can then be copied into a CloudFormation template.
