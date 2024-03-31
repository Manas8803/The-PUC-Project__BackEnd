package roles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

func CreateWebSocketApi(stack awscdk.Stack) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("WebSocket-Lambda-Role-2"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("execute-api:Invoke"),
			jsii.String("execute-api:ManageConnections"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	}))
	return role
}
func CreateWebSocketApi2(stack awscdk.Stack) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("WebSocket-Lambda-Role-1"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("execute-api:Invoke"),
			jsii.String("execute-api:ManageConnections"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	}))
	return role
}
func CreateWebSocketApi1(stack awscdk.Stack) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("WebSocket-Lambda-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("execute-api:Invoke"),
			jsii.String("execute-api:ManageConnections"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	}))
	return role
}
