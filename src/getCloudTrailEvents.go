package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

type CloudTrailConfig struct {
	Region string
	RoleToAssume *string
	Username string
	StartTime string
	EndTime string
}

func (config *CloudTrailConfig) GetCloudTrailEvents() (error, map[string]map[string][]*cloudtrail.Event) {
	ct := config.connectCloudTrail()

	err, startTime, endTime := config.parseTimes()
	if err != nil {
		return err, nil
	}

	eventMap := make(map[string]map[string][]*cloudtrail.Event)

	events := config.getEvents(ct, startTime, endTime)

	for v := range events {
		if _, ok := eventMap[*v.EventSource]; !ok {
			eventMap[*v.EventSource] = make(map[string][]*cloudtrail.Event)
		}
		if _, ok := eventMap[*v.EventSource][*v.EventName]; !ok {
			eventMap[*v.EventSource][*v.EventName] = []*cloudtrail.Event{}
		}

		eventMap[*v.EventSource][*v.EventName] = append(eventMap[*v.EventSource][*v.EventName], v)
	}

	return nil, eventMap
}

func (config *CloudTrailConfig) getEvents(ct *cloudtrail.CloudTrail, startTime *time.Time, endTime *time.Time) <- chan *cloudtrail.Event {

	c := make (chan *cloudtrail.Event)

	go func() {
		defer close(c)

		startToken := "settonil"
		nextToken := &startToken
		for nextToken != nil {
			if *nextToken == startToken {
				nextToken = nil
			}

			lookupEventsOutput, err := ct.LookupEvents(&cloudtrail.LookupEventsInput{
				NextToken: nextToken,
				StartTime: startTime,
				EndTime: endTime,
				LookupAttributes: []*cloudtrail.LookupAttribute{
					&cloudtrail.LookupAttribute{
						AttributeKey: aws.String("Username"),
						AttributeValue: aws.String(config.Username),
					},
				},
			})

			for _, v := range lookupEventsOutput.Events {
				c <- v
			}

			if err != nil {
				fmt.Println(err)
				nextToken = nil
			} else {
				nextToken = lookupEventsOutput.NextToken
			}
		}
	}()

	return c
}

func (config *CloudTrailConfig) parseTimes() (error, *time.Time, *time.Time) {
	awsTimeFormat := "2006-01-02 15:04:05 MST-07:00"
	parsedStartTime, err := time.Parse(awsTimeFormat, config.StartTime)
	if err != nil {
		return err, nil, nil
	}	
	parsedEndTime, err := time.Parse(awsTimeFormat, config.EndTime)
	if err != nil {
		return err, nil, nil
	}

	return nil, &parsedStartTime, &parsedEndTime
}

func (config *CloudTrailConfig) connectCloudTrail() *cloudtrail.CloudTrail {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	})
	mySession := session.Must(sess, err)

	if config.RoleToAssume != nil {
		creds := stscreds.NewCredentials(mySession, *config.RoleToAssume)
		return cloudtrail.New(mySession, &aws.Config{Credentials: creds})
	}

	return cloudtrail.New(mySession)
}
