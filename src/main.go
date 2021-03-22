package main

import (
	"fmt"
	"strings"
)

func main() {
	config := &CloudTrailConfig{
		Region: "eu-west-2",
		Username: "my-tool-user",
		StartTime: "2021-03-22 16:15:57 UTC+00:00",
		EndTime: "2021-03-22 17:24:30 UTC+00:00",
	}

	err, eventMap := config.GetCloudTrailEvents()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range eventMap {
		serviceKey := strings.Split(k, ".")[0]
		fmt.Printf("- PolicyName: Manage%v\n", strings.Title(serviceKey))
		fmt.Println("  PolicyDocument:")
		fmt.Println("    Statement:")
		fmt.Println("    - Effect: Allow")
		fmt.Println("    Action:")
		for vk, _ := range v {
			fmt.Printf("      - '%v:%v'\n", serviceKey, vk)
		}
	}	
}
