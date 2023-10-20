package provisioning

import "fmt"

// dummy Ec2: compute

type EC2Client struct {
}

func (x EC2Client) runInstance(region string, vpc string, subnet string, clientToken string) (string, error) {
	fmt.Println("Spawning Instance in region " + region + ", vpc " + vpc + ", subnet " + subnet + ", clientToken " + clientToken)
	return "xyz-instance", nil
}

func (x EC2Client) isAvailable(instance string) (string, error) {
	fmt.Println("instance is available " + instance)
	//mocked it to make it "Available" all the time
	return "AVAILABL", nil
}
