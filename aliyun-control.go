package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var action = flag.String("action", "describeInstances", "-action <action> 必选。目前支持describeInstances, describeSecurityGroups, createSecurityGroup, stopInstance，startInstance，joinSecurityGroup")
var accessKeyID = flag.String("accessKeyID", "", "-accessKeyID <accessKeyID> 必选。")
var accessSecret = flag.String("accessSecret", "", "-accessSecret <accessKeyID> 必选。")
var regionID = flag.String("regionID", "", "-regionID <regionId> 必选。")
var instanceID = flag.String("instanceID", "", "-instanceID [instanceID] 可选。atcion为stopInstance,startInstance,joinSecurityGroup时需要带上此参数,可通过describeInstances这一action获取。")
var securityGroupID = flag.String("securityGroupID", "", "-securityGroupID [securityGroupID] 可选。atcion为joinSecurityGroup时需要带上此参数,可通过describeSecurityGroups这一action获取。也可以用createSecurityGroup创建一个默认全block的安全组。")

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

func describeInstances() {
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

func stopInstance() {
	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = *instanceID

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

	request.InstanceId = *instanceID

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

	request.SecurityGroupId = *securityGroupID
	request.InstanceId = *instanceID

	response, err := client.JoinSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func main() {
	flag.Parse()
	if *accessKeyID == "" || *accessSecret == "" || *regionID == "" {
		fmt.Println("accessKeyID和accessSecret是必选参数，请根据教程获取。")
		os.Exit(1)
	}
	client, _ = ecs.NewClientWithAccessKey(*regionID, *accessKeyID, *accessSecret)
	switch *action {
	case "describeInstances":
		describeInstances()
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
