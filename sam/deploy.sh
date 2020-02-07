#!/bin/bash


sam deploy --stack-name apigw-sam --s3-bucket dihmuzikien-artifacts --capabilities CAPABILITY_IAM --parameter-overrides "ParameterKey=ApiGatewayBaseStackName,ParameterValue=base-api-gw"