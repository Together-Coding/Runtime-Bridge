package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"log"
	"time"
)

// DescTasks describes tasks and add additional data
func DescTasks(ecsTask *ECSTask) {
	ecsClient := ecs.NewFromConfig(GetConfig())

	var task_ ecsTypes.Task
	var publicIp *string

	for i := 0; i < 20 && publicIp == nil; i++ {
		time.Sleep(time.Millisecond * time.Duration(500+200*i))
		descTaskOutput, err := ecsClient.DescribeTasks(context.TODO(), &ecs.DescribeTasksInput{
			Cluster: ecsTask.Cluster.ClusterArn,
			Tasks:   []string{*ecsTask.Task.TaskArn},
		})
		if err != nil {
			log.Fatal(err)
		}
		if len(descTaskOutput.Tasks) <= 0 {
			log.Fatalf("Task description failed. No task named '%s'", *ecsTask.Task.TaskArn)
		}

		task_ = descTaskOutput.Tasks[0]

		// Need to retrieve public IP from ENI
		for _, attach := range task_.Attachments {
			var eni *string
			for _, detail := range attach.Details {
				if *detail.Name == "networkInterfaceId" {
					eni = detail.Value
				}
			}

			// ENI is not found. Maybe it's so early that ENI is not attached to it yet.
			if eni == nil {
				continue
			}

			// Get public IP from ENI
			ec2Client := ec2.NewFromConfig(GetConfig())
			eniOutput, err := ec2Client.DescribeNetworkInterfaces(context.TODO(), &ec2.DescribeNetworkInterfacesInput{
				NetworkInterfaceIds: []string{*eni},
				NextToken:           nil,
			})

			if err == nil && len(eniOutput.NetworkInterfaces) > 0 && eniOutput.NetworkInterfaces[0].Association != nil {
				fmt.Println(eniOutput.NetworkInterfaces[0].Association)
				publicIp = eniOutput.NetworkInterfaces[0].Association.PublicIp
				break
			}
		}
	}

	ecsTask.Task = &task_
	ecsTask.PublicIp = publicIp
}
