package provisioning

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func ProvisioningCloud(ctx workflow.Context, vpc string, subnet string, bucketName string, region string, clientToken string, downloadJarScript string, installJarScript string, serviceName string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 20, // an activity cannot take more than 'x' minutes
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var preCheckIfJarExists string
	// same function with different args :
	// downloadJarScript is the path after s3 bucket dir
	// todo: rename downloadJarScript as s3 download jarscript path
	err := workflow.ExecuteActivity(ctx, PreCheckIfKeyExists, bucketName, downloadJarScript).Get(ctx, &preCheckIfJarExists)
	if err != nil {
		return "", err
	}

	// checking the path in order to install script after downloading
	err = workflow.ExecuteActivity(ctx, PreCheckIfKeyExists, bucketName, installJarScript).Get(ctx, &preCheckIfJarExists)
	if err != nil {
		return "", err
	}

	// creating instance
	var instanceId string
	err = workflow.ExecuteActivity(ctx, CreateInstance, region, vpc, subnet, clientToken).Get(ctx, &instanceId)
	if err != nil {
		return "", err
	}

	// delete instance if creation has failed
	defer func() {
		if err != nil {
			// activity failed, and workflow context is cancelled
			// For multiple clean up of resources, will leverage pattern provided in https://temporal.io/blog/saga-pattern-made-easy can be used.
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			errDeleteInstance := workflow.ExecuteActivity(disconnectedCtx, DeleteInstance, instanceId).Get(disconnectedCtx, nil)
			if errDeleteInstance != nil {
				workflow.GetLogger(ctx).Error("Executing rollback failed", "Error", errDeleteInstance)
			}
		}
	}()

	// polling to check if instance is available
	isInstanceAvailable := false
	for !isInstanceAvailable {
		//StartToCloseTimeout will help here to timeout if this is taking a long time
		err = workflow.ExecuteActivity(ctx, IsInstanceAvailable, instanceId).Get(ctx, &isInstanceAvailable)
		if err != nil {
			return "", err
		}
	}
	if err != nil {
		return "", err
	}

	// download the jar to the created instance
	var downloadJar string
	err = workflow.ExecuteActivity(ctx, DownloadJar, instanceId, downloadJarScript).Get(ctx, &downloadJar)
	if err != nil {
		return "", err
	}

	// install the jar script on the instance
	var installJar string
	err = workflow.ExecuteActivity(ctx, InstallJar, instanceId, installJarScript).Get(ctx, &installJar)
	if err != nil {
		return "", err
	}

	// add instance to the load balancer
	// as part of this we are invoking it as a microservice
	var addToLoadBalancer string
	err = workflow.ExecuteActivity(ctx, AddToLoadBalancer, instanceId, serviceName).Get(ctx, &addToLoadBalancer)
	if err != nil {
		return "", err
	}

	// send an email notification that the instance is provisioned
	var pushNotification string
	err = workflow.ExecuteActivity(ctx, PushNotification).Get(ctx, &pushNotification)
	if err != nil {
		return "", err
	}

	return "completed", nil
}
