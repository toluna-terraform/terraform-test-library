package tolunalambda

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	aws_terratest "github.com/gruntwork-io/terratest/modules/aws"
)

func GetLambdaLayer(t *testing.T, region string, layer_name string, layer_version int64) *lambda.GetLayerVersionOutput {
	sess, err := aws_terratest.NewAuthenticatedSession(region)
	svc := lambda.New(sess)
	input := &lambda.GetLayerVersionInput{
		LayerName:     aws.String(layer_name),
		VersionNumber: aws.Int64(layer_version),
	}
	result, err := svc.GetLayerVersion(input)
	if err != nil {
		return nil
	}
	return result
}
