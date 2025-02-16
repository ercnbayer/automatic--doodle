#!/usr/bin/env bash
set -x 
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=eu-west-2
export LOCALSTACK_API_URL=http://localhost:4566
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam create-user --user-name doodle-dev
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam create-access-key --user-name doodle-dev
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam attach-user-policy --user-name doodle-dev --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam attach-user-policy --user-name doodle-dev --policy-arn arn:aws:iam::aws:policy/AmazonSNSFullAccess
aws --endpoint-url=http://localhost:4566 s3api create-bucket --bucket doodle-local-bucket --region eu-west-2 --create-bucket-configuration LocationConstraint=eu-west-2
aws --endpoint-url=http://localhost:4566 s3api create-bucket --bucket profile-image --region eu-west-2 --create-bucket-configuration LocationConstraint=eu-west-2
aws --endpoint-url=http://localhost:4566 s3api put-bucket-acl --acl authenticated-read --bucket doodle-local-bucket
aws  --endpoint-url=http://localhost:4566  s3api put-bucket-acl --bucket profile-image --acl authenticated-read
aws s3api put-bucket-cors --bucket doodle-local-bucket --cors-configuration file:///cors-config.json --endpoint-url=http://localhost:4566 
aws s3api put-bucket-cors --bucket profile-image --cors-configuration file:///cors-config.json --endpoint-url=http://localhost:4566 

set +x

