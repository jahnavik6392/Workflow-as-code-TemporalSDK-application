package provisioning

import (
	"context"
	"errors"
	"fmt"
)

// provisioning app
// This function will throw exception if the jar doesn't exist, thi achieved by getObject call throwing Exception
// This need to be handled on the Task if things fails
func PreCheckIfKeyExists(ctx context.Context, bucketName string, keyName string) error {
	//pass the access token
	err := S3Client{}.getObject(bucketName, keyName)
	//handle the error
	return err
}

// func createVpc

// func createSubnets

// create instance
func CreateInstance(ctx context.Context, region string, vpc string, subnet string, clientToken string) (string, error) {
	// clientToken for idempotency - to avoid deduplicating the instance
	ec2Instance, err := EC2Client{}.runInstance(region, vpc, subnet, clientToken)
	return ec2Instance, err
}

// checking if instance is created
// mocked it to true always
func IsInstanceAvailable(ctx context.Context, ec2Instance string) (bool, error) {
	ec2InstanceStatus, err := EC2Client{}.isAvailable(ec2Instance)
	if ec2InstanceStatus == "AVAILABLE" {
		return true, err
	} else {
		//return false, err
		return false, errors.New("instance creation failed")
	}
}

// Download jar on EC2
func DownloadJar(ctx context.Context, ec2Instance string, downloadJarScript string) error {
	err := SystemManager{}.runCommand(ec2Instance, downloadJarScript)
	return err
}

// Install jar on EC2
func InstallJar(ctx context.Context, ec2Instance string, installJar string) error {
	err := SystemManager{}.runCommand(ec2Instance, installJar)
	return err
}

// Add instance to the load balancer : microservice
func AddToLoadBalancer(ctx context.Context, ec2Instance string, serviceName string) error {
	//client of microservice
	status, err := AddInstanceToLB(ec2Instance, serviceName)
	fmt.Println("Status " + status)
	return err
}

// Inform about the success : could potentially be a microservice
func PushNotification(ctx context.Context) error {
	fmt.Println("pushing notification!!")
	//pushNotificationClient := PushNotificationService.getClient();
	//err := pushNotificationClient.sendNotification(CHANNELS);
	return nil
}

// Rollback
func DeleteInstance(ctx context.Context, instanceId string) error {
	fmt.Println("instance deleted.")
	return nil
}
