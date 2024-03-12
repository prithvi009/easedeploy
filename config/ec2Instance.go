package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	once     sync.Once
	instance *ec2.Instance
	err      error
)

func getOrCreateInstance() (*ec2.Instance, error) {
	// Use sync.Once to ensure that instance creation is performed only once
	once.Do(func() {
		// Initialize a new AWS session
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)
		if err != nil {
			log.Println("Ec2 creation error:", err)
			return
		}

		// Create EC2 service client
		svc := ec2.New(sess)

		// Define the filter to find instances with the specified tag
		filters := []*ec2.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String("easedeploy")},
			},
		}

		// Describe EC2 instances with the specified tag
		resp, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
			Filters: filters,
		})
		if err != nil {
			log.Println("Error describing instances:", err)
			return
		}

		// Check if any instances were found with the specified tag
		if len(resp.Reservations) > 0 && len(resp.Reservations[0].Instances) > 0 {
			instance = resp.Reservations[0].Instances[0]
			fmt.Println("Found existing instance:", *instance.InstanceId)
			return
		}

		// If no existing instance found, create a new one
		runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
			ImageId:      aws.String("ami-e7527ed7"),
			InstanceType: aws.String("t2.micro"),
			MinCount:     aws.Int64(1),
			MaxCount:     aws.Int64(1),
		})
		if err != nil {
			log.Println("Could not create instance:", err)
			return
		}

		fmt.Println("Created instance", *runResult.Instances[0].InstanceId)
		instance = runResult.Instances[0]

		// Add tags to the created instance
		_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
			Resources: []*string{instance.InstanceId},
			Tags: []*ec2.Tag{
				{
					Key:   aws.String("Name"),
					Value: aws.String("easedeploy"),
				},
			},
		})
		if errtag != nil {
			log.Println("Could not create tags for instance", instance.InstanceId, errtag)
			return
		}

		fmt.Println("Successfully tagged instance")
	})

	return instance, err
}

// Ec2Instance returns the EC2 instance or an error if creation failed
func Ec2Instance() (*ec2.Instance, error) {
	return getOrCreateInstance()
}
