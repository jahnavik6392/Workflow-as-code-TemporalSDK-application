package provisioning

import "fmt"

// dummy : S3 client

type S3Client struct {
}

func (x S3Client) getObject(bucketName string, keyName string) error {
	fmt.Println("ObjectFound" + bucketName + keyName)
	return nil
}
