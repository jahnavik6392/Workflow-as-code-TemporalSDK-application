package main

import (
	"context"
	"flag"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "provisioning-workflow",
		TaskQueue: "provisioning-tasks",
	}
	vpc := flag.String("vpc", "vpcxyz", "a string")
	subnet := flag.String("subnet", "subnetxyz", "a string")
	bucketName := flag.String("bucket-name", "abc", "a string")
	region := flag.String("region", "us-east-1", "a string")
	clientToken := flag.String("clientToken", "abc-def", "a string")
	downloadJarScript := flag.String("downloadJarScript", "download/script.sh", "a string")
	installJarScript := flag.String("install-jar-location", "install/install.sh", "a string")
	serviceName := flag.String("service-name", "xyz", "a string")

	we, err := c.ExecuteWorkflow(context.Background(), options, provisioning.ProvisioningCloud,
		*vpc, *subnet, *bucketName, *region, *clientToken, *downloadJarScript, *installJarScript, *serviceName)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
