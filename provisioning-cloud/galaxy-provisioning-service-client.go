package provisioning

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// library to connect to the microservice

// add instance to the load balancer : making http request
func AddInstanceToLB(instanceId string, serviceName string) (string, error) {
	status, err := callService("add-instance-to-lb", instanceId, serviceName)
	return status, err
}

// utility function for making calls to the microservices
func callService(stem string, instanceId string, serviceName string) (string, error) {
	// end-point to connect to the microservice
	base := "http://localhost:9999/" + stem + "?instanceid=%s&servicename=%s"
	url := fmt.Sprintf(base, url.QueryEscape(instanceId), url.QueryEscape(serviceName))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return "", errors.New(message)
	}

	return translation, nil
}
