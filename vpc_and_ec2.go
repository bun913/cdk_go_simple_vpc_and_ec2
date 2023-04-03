package main

import (
	"vpc_and_ec2/cmd/network"
	"vpc_and_ec2/cmd/server"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VpcAndEc2StackProps struct {
	awscdk.StackProps
}

func NewVpcAndEc2Stack(scope constructs.Construct, id string, props *VpcAndEc2StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	// 共通タグを設定
	awscdk.Tags_Of(app).Add(jsii.String("Project"), jsii.String("Sample"), nil)
	awscdk.Tags_Of(app).Add(jsii.String("Env"), jsii.String("Dev"), nil)

	stack := NewVpcAndEc2Stack(app, "Sample", &VpcAndEc2StackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	sharedNetwork := network.NewNetwork(stack, "SharedVpc", "10.10.0.0/16", true)
	sharedVpc := sharedNetwork.CreateNetworkResources()

	severResource := server.NewServer(stack, sharedVpc)
	severResource.CreateServerResources()

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region: jsii.String("ap-northeast-1"),
	}
}
