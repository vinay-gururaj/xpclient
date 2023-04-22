package main

import (
	"context"
	"time"

	"github.com/vinay-gururaj/xpclient"
)

func main() {

	host := "http://localhost:8001"
	sessionId := "s:TAhDkUuIT7YNGgyiIYENoGndQvtKjWt1.zb7R+cDqew/t8RCtCtLxyLsAZhN/o3S5mftsG36XIg8"
	csrfToken := "7RMnruDC-6aOjp1ANOvqEcn-wl2-cdgXIYIU"

	c, err := xpclient.NewClient(&host, &sessionId, &csrfToken)

	if err != nil {
		panic(err)
	}

	println(c.SessionId)
	println(c.CsrfToken)

	projectId, error := c.CreateXPProject(context.Background(), xpclient.XPProject{Name: "Vinay Library Test 2", Description: "Test project created from visual studio", OrganizationId: "1"})

	if error != nil {
		panic(error)
	}

	println(projectId)

	println("Project created")

	time.Sleep(5 * (time.Second))

	var project xpclient.XPProject
	project.Id = projectId
	project.Name = "Vinay Library Test_Updated_Name"
	project.Description = "Updated Description"

	updateprojectId, err := c.UpdateXPProject(context.Background(), project)

	if err != nil {
		panic(error)
	}

	println("Project name updated " + updateprojectId)

	time.Sleep(5 * (time.Second))

	var environment xpclient.Environment
	environment.Name = "Development"

	devEnvironmentId, error := c.CreateEnvironment(context.Background(), xpclient.Environment{Name: "Development", ProjectId: projectId})
	if err != nil {
		panic(error)
	}

	println("Development environment Id: " + devEnvironmentId)
	println("Created Development environment")

	cteEnvironmentId, error := c.CreateEnvironment(context.Background(), xpclient.Environment{Name: "CTE", ProjectId: projectId})
	if err != nil {
		panic(error)
	}

	println("CTE environment Id: " + cteEnvironmentId)
	println("Created CTE environment")

	time.Sleep(5 * time.Second)

	var updateDevEnvironment xpclient.Environment

	updateDevEnvironment.Name = "Dev"
	updateDevEnvironment.Id = devEnvironmentId

	updateEnvironmentId, error := c.UpdateEnvironment(context.Background(), updateDevEnvironment)
	if err != nil {
		panic(error)
	}

	println("Updated Development Env to Dev: " + updateEnvironmentId)

	var deleteEnvironment xpclient.Environment
	deleteEnvironment.Id = cteEnvironmentId

	deleteEnvironmentId, error := c.DeleteEnvironment(context.Background(), deleteEnvironment)
	if error != nil {
		panic(error)
	}

	println("Deleted CTE Environment: " + deleteEnvironmentId)

	var returnedSQSConnection xpclient.SQSConnection

	returnedSQSConnection, error = c.CreateSQSConnection(context.Background(), xpclient.SQSConnection{Name: "SQS Connection library", AwsAccessKeyId: "ASIA1234", AwsSecretAccessKey: "PAsswprd", ReadQueue: "read_queue", WriteQueue: "write_queue", AwsRegion: "us-east-1", ProjectId: projectId, EnvironmentId: devEnvironmentId})

	if error != nil {
		panic(error)
	}

	println("SQS Connection Connector Id: " + returnedSQSConnection.ConnectorId)

	println("SQS Connection Connection Id: " + returnedSQSConnection.ConnectionId)

	println("Created SQS Connection")

	time.Sleep(5 * time.Second)

	returnedSQSConnection.Name = "SQS Connection library_Updated"
	returnedSQSConnection.AwsAccessKeyId = "ASIA2345"
	returnedSQSConnection.AwsSecretAccessKey = "Password_2"
	returnedSQSConnection.ReadQueue = "read_queue_2"
	returnedSQSConnection.WriteQueue = "write_queue_2"
	returnedSQSConnection.AwsRegion = "us-west-2"

	c.UpdateSQSConnection(context.Background(), returnedSQSConnection)

	println("Updated SQS Connection")
	time.Sleep(5 * time.Second)

	var partialUpdate_nameOnly xpclient.SQSConnection
	partialUpdate_nameOnly.Name = "SQS Connection library_Updated_2"
	partialUpdate_nameOnly.ConnectionId = returnedSQSConnection.ConnectionId

	c.UpdateSQSConnection(context.Background(), partialUpdate_nameOnly)

	println("Updated SQS Connection - Name Only")
	time.Sleep(5 * time.Second)

	var fullUpdateWithoutName xpclient.SQSConnection
	fullUpdateWithoutName.ConnectorId = returnedSQSConnection.ConnectorId
	fullUpdateWithoutName.AwsAccessKeyId = "ASIA3456"
	fullUpdateWithoutName.AwsSecretAccessKey = "password_3"
	fullUpdateWithoutName.ReadQueue = "read_queue_3"
	fullUpdateWithoutName.WriteQueue = "write_queue_3"
	fullUpdateWithoutName.AwsRegion = "eu-west-3"

	c.UpdateSQSConnection(context.Background(), fullUpdateWithoutName)

	println("Updated SQS Connection - fullUpdateWithoutName")
	time.Sleep(5 * time.Second)

	var partialWithoutName xpclient.SQSConnection
	partialWithoutName.ConnectorId = returnedSQSConnection.ConnectorId
	partialWithoutName.AwsRegion = "eu-east-4"

	c.UpdateSQSConnection(context.Background(), partialWithoutName)

	println("Updated SQS Connection - partialWithoutName")
	time.Sleep(5 * time.Second)

	var partialWitName xpclient.SQSConnection
	partialWitName.ConnectorId = returnedSQSConnection.ConnectorId
	partialWitName.ConnectionId = returnedSQSConnection.ConnectionId
	partialWitName.Name = "Final name update"
	partialWitName.AwsRegion = "eu-east-5"
	partialWitName.WriteQueue = "write_queue_4"

	c.UpdateSQSConnection(context.Background(), partialWitName)

	println("Updated SQS Connection - partialWitName")

	var deleteSQSConnection xpclient.SQSConnection
	deleteSQSConnection.ConnectionId = returnedSQSConnection.ConnectionId

	deleteconnectionId, _ := c.DeleteSQSConnection(context.Background(), returnedSQSConnection)

	println("Deleted SQS Connection with Id: " + deleteconnectionId)

	time.Sleep(30 * time.Second)

	deleteEnvId, _ := c.DeleteEnvironment(context.Background(), updateDevEnvironment)

	println("Deleted environment " + deleteEnvId)

	time.Sleep(30 * time.Second)

	deleteProject, _ := c.DeleteXPProject(context.Background(), project)

	println("Deleted Project " + deleteProject)

}
