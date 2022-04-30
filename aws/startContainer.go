package aws

import (
	"context"
	"fmt"
	_ "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"log"
)

type ECSTask struct {
	Task     *ecsTypes.Task
	Cluster  *ecsTypes.Cluster
	PublicIp *string
}

func (t *ECSTask) String() string {
	return fmt.Sprintf(
		"# Task %v\n"+
			"\tcpu: %v\n"+
			"\tCreatedAt: %v\n"+
			"\tStartedAt: %v\n"+
			"\tDesiredStatus: %v\n"+
			"\tGroup: %v\n"+
			"\tHealthStatus: %v\n"+
			"\tLastStatue: %v\n"+
			"\tLaunchType: %v\n"+
			"\tMemory: %v\n"+
			"\tStartedBy: %v\n"+
			"\tTaskArn: %v\n"+
			"\tTaskDefinitionArn: %v\n"+
			"\tVersion: %v\n"+
			"\tPublicIP: %v\n",
		*t.Task.TaskArn, *t.Task.Cpu, *t.Task.CreatedAt, *t.Task.StartedAt, *t.Task.DesiredStatus,
		*t.Task.Group, t.Task.HealthStatus, *t.Task.LastStatus, t.Task.LaunchType, *t.Task.Memory,
		*t.Task.StartedBy, *t.Task.TaskArn, *t.Task.TaskDefinitionArn, t.Task.Version, t.PublicIp)
}

func StartContainer(clusterName, taskDef, startedBy string) (ecsTask *ECSTask) {
	ecsClient := ecs.NewFromConfig(GetConfig())

	// Get cluster information for ARN
	cluster := DescCluster(ecsClient, clusterName)

	// Start new container with options
	task := startTask(ecsClient, cluster.ClusterArn, &taskDef, &startedBy)

	ecsTask = &ECSTask{
		Cluster: &cluster,
		Task:    &task,
	}
	return
}

func startTask(ecsClient *ecs.Client, clusterArn *string, taskDefinition *string, startedBy *string) (task ecsTypes.Task) {
	configPrefix := "CONTAINER"

	var count int32 = 1
	res, err := ecsClient.RunTask(context.TODO(), &ecs.RunTaskInput{
		TaskDefinition: taskDefinition,
		Cluster:        clusterArn,
		Count:          &count,
		LaunchType:     ecsTypes.LaunchTypeFargate,
		NetworkConfiguration: &ecsTypes.NetworkConfiguration{
			AwsvpcConfiguration: &ecsTypes.AwsVpcConfiguration{
				Subnets:        GetContainerSubnets(configPrefix),
				AssignPublicIp: ecsTypes.AssignPublicIpEnabled,
				SecurityGroups: GetContainerSecurityGroup(configPrefix),
			},
		},
		//Overrides:
		//ReferenceId:
		StartedBy: startedBy,
	})

	if err != nil {
		log.Fatal(err)
	}
	if len(res.Failures) != 0 {
		log.Fatal(res.Failures)
	}
	if len(res.Tasks) <= 0 {
		log.Fatalln("Start task failed. No task is started.")
	}

	task = res.Tasks[0]
	return
}
