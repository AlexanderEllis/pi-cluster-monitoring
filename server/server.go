/*
This is a simple heartbeat server we can run on the cluster nodes. It just waits
for any heartbeat and responds.
*/

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	heartbeat "pi-cluster-monitoring/rpc"
)

// Server is a basic heartbeat server.
type Server int

// Heartbeat is a simple check to see if the server is still alive and responding.
func (t *Server) Heartbeat(args *heartbeat.Args, reply *heartbeat.Reply) error {
	reply.OK = true
	return nil
}

func main() {
	heartbeatServer := new(Server)
	rpc.Register(heartbeatServer)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Listening on port 1234")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
	log.Println("Everything went OK")
}
