package main

import "fmt"
import zmq "github.com/pebbe/zmq4"
import "time"
import "encoding/json"

type NodeServer struct {
	hostname string
	port     string
}

type Payload struct {
	Id   string `json:"page"`
	Data string `json:"data"`
}

type Data []Payload

func CreateServer(host string, port string) {
	var data = Data{} // our datastore :)

	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://*:" + port)
	fmt.Printf("Starting server on %s:%s\n", host, port)
	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)

		var p Payload
		json.Unmarshal([]byte(msg), &p)

		// add the data to the data store
		data = append(data, p)

		reply := "OK"
		responder.Send(reply, 0)
	}
}

func createClient(host string, port string) {
	//  Socket to talk to server
	fmt.Printf("Connecting to hello world server on tcp://%s:%s \n", host, port)
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://" + host + ":" + port)

	for request_nbr := 0; request_nbr != 10; request_nbr++ {
		// send hello
		//msg := fmt.Sprintf("Hello %d", request_nbr)
		msg := Payload{Id: "1234", Data: "Hello"}
		b, _ := json.Marshal(msg)
		fmt.Println("Sending: ", string(b))

		requester.Send(string(b), 0)

		// Wait for reply:
		reply, _ := requester.Recv(0)
		if reply == "OK" {
			fmt.Println("Received ", reply)
		}

	}
}

func main() {

	var servers = []NodeServer{NodeServer{hostname: "localhost", port: "1234"},
		NodeServer{hostname: "localhost", port: "1235"}}

	for _, server := range servers {
		go CreateServer(server.hostname, server.port)

	}
	time.Sleep(2 * time.Second)
	fmt.Println("Starting clients")
	for _, server := range servers {
		createClient(server.hostname, server.port)
	}
	fmt.Println(servers)
}
