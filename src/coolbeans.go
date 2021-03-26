package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Supply a used tag value")
		os.Exit(0)
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	environment := os.Args[1]
	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Connecting to AWS")
		var instances []string

		for _, r := range result.Reservations {
			for i, instance := range r.Instances {
				for _, tag := range instance.Tags {
					if *tag.Key == "SECRET" && *tag.Value == ":)" {
						for _, appTag := range r.Instances[i].Tags {
							if *appTag.Key == "SQUIRREL" && *appTag.Value == environment {
								var matchedInstance string = *r.Instances[i].InstanceId
								instances = append(instances, matchedInstance)
							}
						}
					}
				}
			}
		}
		for _, tryInstance := range instances {
			_, err := exec.Command("/usr/local/bin/aws", "ssm", "start-session", "--target", tryInstance).Output()
			if err != nil {
				if (string(err.Error())) == "exit status 255" {
					continue
				}
			}
			fmt.Println("/usr/local/bin/aws", "ssm", "start-session", "--target", tryInstance)
			os.Exit(0)
		}

		fmt.Println("No instances liked that. :(")
	}
}
