package main

import (
	"encoding/json"
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var configFilePath = flag.String("configFilePath", "./config.yaml", "-configFilePath <path> 可选，若不存在则会首次启动提示输入参数并自动生成。")

var client *ecs.Client

func createSecurityGroup() {
	request := ecs.CreateCreateSecurityGroupRequest()
	request.Scheme = "https"

	response, err := client.CreateSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

func describeSecurityGroups() {
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.Scheme = "https"

	response, err := client.DescribeSecurityGroups(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

func describeInstanceStatus() {
	request := ecs.CreateDescribeInstanceStatusRequest()
	request.Scheme = "https"

	response, err := client.DescribeInstanceStatus(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func stopInstance() {
	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = GConfig.InstanceID

	response, err := client.StopInstance(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

func startInstance() {
	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = GConfig.InstanceID

	response, err := client.StartInstance(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

func joinSecurityGroup() {
	request := ecs.CreateJoinSecurityGroupRequest()
	request.Scheme = "https"

	request.SecurityGroupId = GConfig.SecurityGroupID
	request.InstanceId = GConfig.InstanceID

	response, err := client.JoinSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func main() {
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})

	flag.Parse()
	LoadConfig(*configFilePath)

	client, _ = ecs.NewClientWithAccessKey(GConfig.RegionID, GConfig.AccessKeyID, GConfig.AccessSecret)
	switch GConfig.Action {
	case "describeInstanceStatus":
		describeInstanceStatus()
	case "describeSecurityGroups":
		describeSecurityGroups()
	case "createSecurityGroup":
		createSecurityGroup()
	case "stopInstance":
		stopInstance()
	case "startInstance":
		startInstance()
	case "joinSecurityGroup":
		joinSecurityGroup()
	default:
		fmt.Println("action输入错误，目前支持describeInstances, describeSecurityGroups, createSecurityGroup, stopInstance，startInstance，joinSecurityGroup")
	}
}
