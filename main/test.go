package main

import (
	"context"

	"github.com/vinay-gururaj/xpclient"
)

func main() {

	host := "http://localhost:8001"
	sessionId := "s:xmHjJr_P3-NlEk2YlaaRDZd2UAPwz6MV.CEqSnSZQzGmFRu1ZjAHI2OtEZvSTrG+p1P9VMiwss2E"
	csrfToken := "DzBgST6g-cnnjNU8_Lyj35gpG_7vcsgfPH1o"

	c, err := xpclient.NewClient(&host, &sessionId, &csrfToken)

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

	connectionId, error := c.CreateSQSConnection(context.Background(), xpclient.SQSConnection{Name: "SQS Connection", Key: "AXIA123456789YWB", Secret: "XvghjkoiQHJKLYYYYYT", ReadQueue: "read_queue", WriteQueue: "write_queue", ProjectId: projectId})

	if error != nil {
		panic(error)
	}
	println(connectionId)
}
