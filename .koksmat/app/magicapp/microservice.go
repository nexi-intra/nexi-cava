package magicapp

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/nats-io/nats.go/micro"
)

type ListItemChanged struct {
	List    string `json:"list"`
	Site    string `json:"site"`
	Tenant  string `json:"tenant"`
	Id      string `json:"id"`
	Version string `json:"version"`
}

const SERVICENAME = "nexi-cava"

func SignalUpdate() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	payload := ListItemChanged{
		List:    "rooms",
		Site:    "nexi",
		Tenant:  "nexi",
		Id:      "123",
		Version: "1",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	_, err = js.Publish("rooms", b)
	if err != nil {
		log.Fatal(err)
	}
}

var nc *nats.Conn

var kv jetstream.KeyValue

func Update(req micro.Request) {
	var payload ListItemChanged
	err := json.Unmarshal([]byte(req.Data()), &payload)
	//log.Println("Update", payload)

	if err != nil {
		log.Println(err)
		req.Respond([]byte(err.Error()))
		return
	}
	key := payload.Tenant + "." + payload.Site + "." + strings.ReplaceAll(payload.List, " ", "") + "." + payload.Id + "." + payload.Version

	entry, _ := kv.Get(context.Background(), key)
	// if err != nil {
	// 	log.Println(err)
	// 	req.Respond([]byte(err.Error()))
	// 	return
	// }
	if entry == nil {
		log.Println("Key", key, "Entry not found")
		kv.Create(context.Background(), key, []byte("n.a."))
		req.RespondJSON("created")
		return
	} else {
		//		log.Println("Update done")
		//log.Println("Key", key, "Matched")
		req.RespondJSON("done")
	}

}
func StartMicroService() {
	// Parent context cancels connecting/reconnecting altogether.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error

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
	js, _ := jetstream.New(nc)

	defer cancel()

	kv, _ = js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "sharepoint",
	})

	srv, err := micro.AddService(nc, micro.Config{
		Name: SERVICENAME,

		Version:     "0.0.1",
		Description: "Manage meeting services",
	})
	root := srv.AddGroup(SERVICENAME)
	sharepoint := root.AddGroup("sharepoint")
	sharepoint.AddEndpoint("hook", micro.HandlerFunc(Update))
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
