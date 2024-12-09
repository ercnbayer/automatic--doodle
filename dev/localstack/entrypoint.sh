#!/usr/bin/env bash
set -x 
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam create-user --user-name arrndme-dev
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam create-access-key --user-name arrndme-dev
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam attach-user-policy --user-name arrndme-dev --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess
aws --endpoint-url=http://localhost:4566 --region eu-west-2 iam attach-user-policy --user-name arrndme-dev --policy-arn arn:aws:iam::aws:policy/AmazonSNSFullAccess
aws --endpoint-url=http://localhost:4566 s3api create-bucket --bucket arrndme-local-bucket --region eu-west-2 --create-bucket-configuration LocationConstraint=eu-west-2
aws --endpoint-url=http://localhost:4566 s3api put-bucket-acl --acl authenticated-read --bucket arrndme-local-bucket
aws s3api put-bucket-cors --bucket arrndme-local-bucket --cors-configuration file:///cors-config.json --endpoint-url=http://localhost:4566 
set +x

