package main

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func getResources(eventMap map[string][]*cloudtrail.Event) []string {

	resourceMap := make(map[string]map[string]int)
	for _, events := range eventMap {
		for _, event := range events {
			for _, resource := range event.Resources {
				if _, ok := resourceMap[*resource.ResourceType]; !ok {
					resourceMap[*resource.ResourceType] = make(map[string]int)
				}

				if _, ok := resourceMap[*resource.ResourceType][*resource.ResourceName]; !ok {
					resourceMap[*resource.ResourceType][*resource.ResourceName] = 0
				}
				
				resourceMap[*resource.ResourceType][*resource.ResourceName]++
			}
		}
	}

	if len(resourceMap) == 0 {
		return []string{ "*" }
	}

	resources := []string{}
	for _, v := range resourceMap {
		for vk, _ := range v {
			resources = append(resources, vk)
		}
	}

	return resources
}
