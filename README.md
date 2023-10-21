# TemporalSDK-application
Exploring Temporal Technologies SDK(https://github.com/temporalio/sdk-go) by creating a sample application

Workflow: Provisioning in AWS (provisioning-cloud/provisioning-cloud/provisioning.go)

Registered the below activities in the workflow 
1. PreCheckIfKeyExists
2. CreateInstance
3. IsInstanceAvailable
4. DownloadJar
5. InstallJar
6. AddToLoadBalancer (NOTE: created this as a Microservice to test the SDK behavior)
7. PushNotification
8. DeleteInstance

Worker picks up the tasks from the task-queue and runs the workflows and the activities

Steps to run this application: 
1. Start the Microservice by running
   >'go run microservices/galaxy-provisioning-cloud.go'
2. Start the Worker by running
   >'go run worker/main.go`
3. Execute the Workflow by running
   >`go run start/main.go`
