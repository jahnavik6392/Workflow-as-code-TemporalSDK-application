package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "provisioning-tasks", worker.Options{})

	w.RegisterWorkflow(provisioning.ProvisioningCloud)
	w.RegisterActivity(provisioning.PreCheckIfKeyExists)
	w.RegisterActivity(provisioning.CreateInstance)
	w.RegisterActivity(provisioning.IsInstanceAvailable)
	w.RegisterActivity(provisioning.DownloadJar)
	w.RegisterActivity(provisioning.InstallJar)
	w.RegisterActivity(provisioning.AddToLoadBalancer)
	w.RegisterActivity(provisioning.PushNotification)
	w.RegisterActivity(provisioning.DeleteInstance)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
