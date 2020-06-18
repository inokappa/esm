package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type ecsService struct {
	*ecs.ECS
}

func NewEcsService() *ecsService {
	return &ecsService{
		ecs.New(session.New()),
	}
}

func (ecsService *ecsService) ServiceUpdate(clusterName string, serviceName string, desiredCount int64, forceUpdate bool) bool {
	input := &ecs.UpdateServiceInput{
		Cluster:            aws.String(clusterName),
		Service:            aws.String(serviceName),
		DesiredCount:       aws.Int64(desiredCount),
		ForceNewDeployment: aws.Bool(forceUpdate),
	}

	_, err := ecsService.UpdateService(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		} else {
			fmt.Println(err.Error())
		}
		return false
	}

	return true
}
