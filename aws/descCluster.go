package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"log"
)

// DescCluster describes specified cluster
func DescCluster(ecsClient *ecs.Client, clusterName string) ecsTypes.Cluster {
	clusters, err := ecsClient.DescribeClusters(context.TODO(), &ecs.DescribeClustersInput{
		Clusters: []string{clusterName},
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(clusters.Clusters) == 0 {
		log.Fatalf("Error: ECS Cluster '%s' does not exist, cannot proceed.\n", clusterName)
	}

	return clusters.Clusters[0]
}
