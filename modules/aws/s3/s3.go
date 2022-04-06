/*This package should include functions for verifying S3 AWS service resources*/
package tolunas3aws

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

/*Checks a s3 bucket has Public Access Block and returns it's output*/
func S3GetPublicAccessBlock(t *testing.T, region string, bucket string) *s3.GetPublicAccessBlockOutput {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := s3.New(sess)
	result, err := svc.GetPublicAccessBlock(&s3.GetPublicAccessBlockInput{Bucket: &bucket})
	if err != nil {
		return nil
	}
	return result
}

type S3BucketPolicy struct {
	Effect    string `json:"Effect"`
	Principal string `json:"Principal"`
	Resource  string `json:"Resource"`
	Action    string `json:"Action"`
}

/*Checks a s3 bucket has a Policy and returns it's output as json*/
func S3GetBucketPolicy(t *testing.T, region string, bucket string) S3BucketPolicy {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := s3.New(sess)
	result, err := svc.GetBucketPolicy(&s3.GetBucketPolicyInput{Bucket: &bucket})
	if err != nil {
		return S3BucketPolicy{}
	}
	assert.NotNil(t, *result.Policy, "Failed to get Bucket policy")
	var objs map[string]interface{}
	json.Unmarshal([]byte(*result.Policy), &objs)
	policy := objs["Statement"].([]interface{})
	statement := policy[0].(map[string]interface{})
	principal := statement["Principal"].(map[string]interface{})
	resource := statement["Resource"].([]interface{})
	bucketpolicy := S3BucketPolicy{
		Effect:    statement["Effect"].(string),
		Principal: principal["AWS"].(string),
		Resource:  resource[1].(string),
		Action:    statement["Action"].(string),
	}
	return bucketpolicy
}

/*Checks a s3 bucket has ACL defined and returns it's output*/
func S3GetBucketACLs(t *testing.T, region string, bucket string) *s3.GetBucketAclOutput {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := s3.New(sess)
	result, err := svc.GetBucketAcl(&s3.GetBucketAclInput{Bucket: &bucket})
	if err != nil {
		return nil
	}
	return result
}
