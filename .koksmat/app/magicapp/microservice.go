package magicapp

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type ServiceRequest struct {
	Args    []string `json:"args"`
	Body    string   `json:"body"`
	Channel string   `json:"channel"`
	Timeout int      `json:"timeout"`
}

const SERVICENAME = "nexi-cava"

func Update(req micro.Request) {
	// var payload ServiceRequest
	// _ = json.Unmarshal([]byte(req.Data()), &payload)
	// log.Println("Update", payload)

	// _, pwsherr := execution.ExecutePowerShell("john", "*", SERVICENAME, "00-magic", "20-update.ps1", "")
	// if pwsherr != nil {
	// 	log.Println(pwsherr)
	// 	req.Respond([]byte(pwsherr.Error()))
	// 	return
	// }
	log.Println("Update done")
	req.RespondJSON("done")

}

func StartMicroService() {
	// Parent context cancels connecting/reconnecting altogether.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	var nc *nats.Conn
	opts := []nats.Option{

		nats.ReconnectWait(2 * time.Second),
		nats.ReconnectHandler(func(c *nats.Conn) {
			log.Println("Reconnected to", c.ConnectedUrl())
		}),
		nats.DisconnectHandler(func(c *nats.Conn) {
			log.Println("Disconnected from NATS")
		}),
		nats.ClosedHandler(func(c *nats.Conn) {
			log.Println("NATS connection is closed.")
		}),
	}

	go func() {
		nc, err = nats.Connect(nats.DefaultURL, opts...)
	}()

WaitForEstablishedConnection:
	for {
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// Wait for context to be canceled either by timeout
		// or because of establishing a connection...
		select {
		case <-ctx.Done():
			break WaitForEstablishedConnection
		default:
		}

		if nc == nil || !nc.IsConnected() {
			log.Println("Connection not ready")
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break WaitForEstablishedConnection
	}
	if ctx.Err() != nil {
		log.Fatal(ctx.Err())
	}

	srv, err := micro.AddService(nc, micro.Config{
		Name: SERVICENAME,

		Version:     "0.0.1",
		Description: "Manage meeting services",
	})
	root := srv.AddGroup(SERVICENAME)
	root.AddEndpoint("hook", micro.HandlerFunc(Update))
	for {
		if nc.IsClosed() {
			break
		}

		time.Sleep(1 * time.Second)
	}

	// Disconnect and flush pending messages
	if err := nc.Drain(); err != nil {
		log.Println(err)
	}
	log.Println("Disconnected")

}
