/*This package should include functions for verifying Codebuild AWS service resources*/
package tolunacodebuild

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/codebuild"
	aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
	tolunacommons "github.com/toluna-terraform/terraform-test-library/modules/commons"
)

/*Checks a codebuild project was created and returns bool [true|false]*/
func VerifyCodeBuildProject(t *testing.T, region string, project_name string) bool {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := codebuild.New(sess)
	input := &codebuild.ListProjectsInput{}
	result, err := svc.ListProjects(input)
	if err != nil {
		return assert.Nil(t, err, "Failed to get Policy")
	}
	projectFound := false
	for _, projectName := range result.Projects {
		if *projectName == project_name {
			projectFound = true
		}
	}
	return assert.True(t, projectFound, fmt.Sprintf("Project %s not created", project_name))
}

/*Checks a codebuild report groups where created and returns bool [true|false]*/
func VerifyCodeBuildReportsGroups(t *testing.T, region string, reportList []string, app_name string) bool {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	if err != nil {
		log.Println("Failed to get Policy", err)
		return assert.Nil(t, err, "Failed to get Report group")
	}
	svc := codebuild.New(sess)
	input := &codebuild.ListReportGroupsInput{}
	result, err := svc.ListReportGroups(input)
	testResults := []bool{}
	for _, reportGroupName := range result.ReportGroups {
		groupName := strings.Split(*reportGroupName, "/")
		if strings.HasPrefix(groupName[1], app_name) {
			testResults = append(testResults, assert.True(t, tolunacommons.ListContains(reportList, groupName[1]), fmt.Sprintf("Report group %s not created", groupName[1])))
		}
	}
	if tolunacommons.ListBoolContains(testResults, false) {
		return false
	}
	return true
}
