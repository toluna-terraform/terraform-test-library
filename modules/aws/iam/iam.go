/*This package should include functions for verifying IAM AWS service resources*/
package tolunaiam

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"testing"

	tolunacommons "github.com/toluna-terraform/terraform-test-library/modules/commons"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

/*Checks a role was created and returns bool [true|false]*/
func VerifyIAMRoleExists(t *testing.T, region string, role_name string) bool {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := iam.New(sess)
	input := &iam.GetRoleInput{
		RoleName: aws.String(role_name),
	}
	result, err := svc.GetRole(input)
	if err != nil {
		return assert.Nil(t, err, "Failed to get Role")
	}
	return assert.True(t, strings.HasSuffix(*result.Role.Arn, role_name), "Wrong role ARN returned")
}

/*Checks a role has the correct attached policy and returns bool [true|false]*/
func VerifyAttachedPoliciesForRole(t *testing.T, region string, role_name string, policy_list []string) bool {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := iam.New(sess)
	input := &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(role_name),
	}
	result, err := svc.ListAttachedRolePolicies(input)
	if err != nil {
		return assert.Nil(t, err, "Failed to get Policies")
	}
	policyList := []string{}
	testResults := []bool{}
	for _, policyName := range result.AttachedPolicies {
		log.Printf("Verify policy %s for test framework role is attached", *policyName.PolicyName)
		policyList = append(policyList, *policyName.PolicyName)
		testResults = append(testResults, assert.True(t, tolunacommons.ListContains(policy_list, *policyName.PolicyName), fmt.Sprintf("Policy name %s not attached", *policyName.PolicyName)))
	}
	for _, policyName := range policy_list {
		testResults = append(testResults, assert.True(t, tolunacommons.ListContains(policyList, policyName), fmt.Sprintf("Policy name %s should not attached", policyName)))
	}
	if tolunacommons.ListBoolContains(testResults, false) {
		return false
	}
	return true
}

/*Checks a role has the correct policy document and returns bool [true|false]*/
func VerifyRolePolicies(t *testing.T, region string, expectedPolicy string, role_name string, policy_name string) bool {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := iam.New(sess)
	input := &iam.GetRolePolicyInput{
		RoleName:   aws.String("role-my-app-non-prod-codebuild-publish-reports-my-app-non-prod"),
		PolicyName: aws.String("policy-codebuild-publish-reports-my-app-non-prod"),
	}
	result, err := svc.GetRolePolicy(input)
	if err != nil {
		return assert.Nil(t, err, "Failed to get Policy")
	}
	encodedValue := *result.PolicyDocument
	decodedValue, err := url.QueryUnescape(encodedValue)
	if err != nil {
		if err != nil {
			return false
		}
	}
	expPolicy := strings.ReplaceAll(expectedPolicy, "\t", "")
	decodedPolicy := strings.ReplaceAll(decodedValue, " ", "")
	return assert.Equal(t, expPolicy, decodedPolicy, fmt.Sprintf("Policy document %s does not match expected document", expectedPolicy))
}
