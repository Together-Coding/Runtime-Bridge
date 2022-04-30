package aws

import (
	"os"
	"strings"
)

func GetContainerSubnets(prefix string) (subnets []string) {
	subnets = strings.Split(os.Getenv(prefix + "_SUBNETS"), ",")
	for i := range subnets {
		subnets[i] = strings.TrimSpace(subnets[i])
	}
	return
}

func GetContainerSecurityGroup(prefix string) (sgs []string) {
	sgs = strings.Split(os.Getenv(prefix + "_SECURITY_GROUP"), ",")
	for i := range sgs {
		sgs[i] = strings.TrimSpace(sgs[i])
	}
	return
}