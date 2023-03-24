package main

import (
	"context"
	"xpclient"
)

func main() {
	c, err := xpclient.NewClient("http://localhost:8001")

	if err != nil {
		panic(err)
	}

	println(c.SessionId)
	println(c.CsrfToken)

	projectId, error := c.CreateXPProject(context.Background(), xpclient.XPProject{Name: "Vinay Library Test 2", Description: "Test project created from visual studio"})

	if error != nil {
		panic(error)
	}

	println(projectId)

	c.CreateSQSConnection(context.Background(), xpclient.SQSConnection{Name: "SQS Connection", Key: "AXIA123456789YWB", Secret: "XvghjkoiQHJKLYYYYYT", ReadQueue: "read_queue", WriteQueue: "write_queue", ProjectId: projectId})

}
