package tolunavpcaws

import (
    "testing"

    aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
    "github.com/stretchr/testify/assert"
)

// tests if given VPC exists 
func TestIfVpcExists(t *testing.T, vpcId string, region string) {
    vpc := aws_terratest.GetVpcById(t, vpcId, region)
    assert.NotEmpty(t, vpc)
}



