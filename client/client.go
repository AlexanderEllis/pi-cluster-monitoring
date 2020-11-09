/*
This is a simple heartbeat client we can run on a monitoring computer (i.e. not
the nodes). It wil periodically query the nodes and report their status.
*/

package main

import (
	"fmt"
	"net/rpc"
	"os"
	heartbeat "pi-cluster-monitoring/rpc"
	"time"
)

// HeartbeatInterval determines how often the heartbeats are sent.
var HeartbeatInterval = 5

func sendHeartbeat(serverAddress string) {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")

	if err != nil {
		fmt.Println("Server "+serverAddress+" is not OK. dialing: ", err)
		return
	}

	args := &heartbeat.Args{}
	var reply heartbeat.Reply

	err = client.Call("Server.Heartbeat", args, &reply)
	if err != nil {
		fmt.Println("Server "+serverAddress+" is not OK. error calling Heartbeat: ", err)
		return
	}

	if !reply.OK {
		fmt.Println("Server " + serverAddress + " is not OK.")
	}
	fmt.Println("Server " + serverAddress + " is OK.")
}

func main() {
	// TODO: better way to store these? .env file?
	allServerIPs := []string{
		os.Getenv("PI1"),
		os.Getenv("PI2"),
		os.Getenv("PI3"),
		os.Getenv("PI4")}

	for true {
		fmt.Println("\nSending heartbeats")
		for _, server := range allServerIPs {
			go sendHeartbeat(server)
		}

		time.Sleep(time.Duration(HeartbeatInterval) * time.Second)
	}
}
