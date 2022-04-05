package tolunacodebuild

import (
	"fmt"
	"strings"
	"testing"
	"tolunacommons"

	"github.com/aws/aws-sdk-go/service/codebuild"
	aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

func VerifyCodeBuildProject(t *testing.T, region string, project_name string) {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := codebuild.New(sess)
	input := &codebuild.ListProjectsInput{}
	result, err := svc.ListProjects(input)
	if err != nil {
		assert.Nil(t, err, "Failed to get Policy")
	}
	projectFound := false
	for _, projectName := range result.Projects {
		if *projectName == project_name {
			projectFound = true
		}
	}
	assert.True(t, projectFound, fmt.Sprintf("Project %s not created", project_name))
}

func TestCodeBuildTestReportsGroups(t *testing.T, region string, reportList []string, app_name string) {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	if err != nil {
		assert.Nil(t, err, "Failed to get Report group")
	}
	svc := codebuild.New(sess)
	input := &codebuild.ListReportGroupsInput{}
	result, err := svc.ListReportGroups(input)
	for _, reportGroupName := range result.ReportGroups {
		groupName := strings.Split(*reportGroupName, "/")
		if strings.HasPrefix(groupName[1], app_name) {
			assert.True(t, tolunacommons.ListContains(reportList, groupName[1]), fmt.Sprintf("Report group %s not created", groupName[1]))
		}
	}
}
