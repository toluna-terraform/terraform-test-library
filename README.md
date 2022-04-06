# 

# terraform-test-library
terraform-test-library is a [go](https://golang.org/) library for testing terraform modules along side using Terratest

## Installation

```
go get github.com/toluna-terraform/terraform-test-library
```

## Usage
```go
package examples

import (
	tolunacodebuildaws "github.com/toluna-terraform/terraform-test-library/modules/aws/codebuild"
	tolunaliamaws "github.com/toluna-terraform/terraform-test-library/modules/aws/iam"
	tolunalambdaaws "github.com/toluna-terraform/terraform-test-library/modules/aws/lambda"
	tolunas3aws "github.com/toluna-terraform/terraform-test-library/modules/aws/s3"
	tolunacommons "github.com/toluna-terraform/terraform-test-library/modules/commons"
	tolunacoverage "github.com/toluna-terraform/terraform-test-library/modules/coverage"
)

var moduleName = tolunacommons.GetModName()
func TestSetup(t *testing.T) {
	terraform.InitAndApply(t, configureTerraformOptions(t))
	tolunacoverage.WriteCovergeFiles(t, configureTerraformOptions(t), moduleName)
}

func TestBucketACLExists(t *testing.T) {
	tolunacoverage.MarkAsCovered("aws_s3_bucket_acl.my_bucket", moduleName)
	result := tolunas3aws.S3GetBucketACLs(t, region, bucket)
	assert.NotNil(t, *result.Owner.DisplayName, "Owner not found")
	assert.Equal(t, *result.Grants[0].Permission, "FULL_CONTROL", "ACL not granted")
}

func TestCleanUp(t *testing.T) {
	log.Println("Running Terraform Destroy")
	terraform.Destroy(t, configureTerraformOptions(t))
}
```
