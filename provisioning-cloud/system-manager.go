package provisioning

import "fmt"

// dummy : System manager to run commands on EC2

type SystemManager struct {
}

func (x SystemManager) runCommand(ec2Instance string, downloadJarScript string) error {
	fmt.Println("runCommand " + ec2Instance + ", downloadJarScript " + downloadJarScript)
	return nil
}
