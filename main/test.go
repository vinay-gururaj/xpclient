package main

import (
	"context"

	"github.com/vinay-gururaj/xpclient"
)

func main() {

	host := "http://localhost:8001"

	c, err := xpclient.NewClient(&host)

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
