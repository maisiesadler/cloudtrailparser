package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
)

func main() {

	err, config := getConfig()
	if err != nil {
		fmt.Println(err)
		return
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

func getConfig() (error, *CloudTrailConfig) {
	username := os.Getenv("USERNAME")
	if len(username) == 0 {
		return errors.New("USERNAME not set"), nil
	}

	startTime := os.Getenv("START_DATE")
	if len(startTime) == 0 {
		return errors.New("START_DATE not set"), nil
	}

	endTime := os.Getenv("END_DATE")
	if len(endTime) == 0 {
		return errors.New("END_DATE not set"), nil
	}

	return nil, &CloudTrailConfig{
		Region: "eu-west-2",
		Username: username,
		StartTime: startTime,
		EndTime: endTime,
	}
}
