package tolunavpcaws

import (
    "testing"
    //"encoding/json"

    //"github.com/aws/aws-sdk-go/service/s3"
    aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
    "github.com/stretchr/testify/assert"
)

var region = "us-east-1"
var vpcId = "vpc-0a9d3172533527c76"

func TestVpcExists(t *testing.T) {
    vpc := aws_terratest.GetVpcById(t, vpcId, region)
    assert.NotEmpty(t, vpc)
}


