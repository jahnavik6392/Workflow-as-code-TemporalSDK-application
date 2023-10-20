package main

// call this via HTTP GET with a URL like:
// http://localhost:9999/add-instance-to-lb?instanceid=xyz&servicename=xyz

import (
	"fmt"
	"net/http"
)

func addInstanceToLBHandler(w http.ResponseWriter, r *http.Request) {
	instanceid, ok := r.URL.Query()["instanceid"]
	if !ok {
		http.Error(w, "Missing required 'name' parameter.", http.StatusBadRequest)
	}
	serviceName, ok := r.URL.Query()["servicename"]
	if ok {
		status := fmt.Sprintf("%s added to loadbalancer of %s", instanceid, serviceName)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, status)
	} else {
		http.Error(w, "Missing required 'instanceid' parameter and servicename parameter.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/add-instance-to-lb", addInstanceToLBHandler)
	http.ListenAndServe(":9999", nil)
}
