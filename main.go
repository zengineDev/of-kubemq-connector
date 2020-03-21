package main

import (
	"context"
	"github.com/kubemq-io/kubemq-go"
	"github.com/openfaas-incubator/connector-sdk/types"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	creds := types.GetCredentials()

	topics := []string{}
	if val, exists := os.LookupEnv("topics"); exists {
		for _, topic := range strings.Split(val, ",") {
			if len(topic) > 0 {
				topics = append(topics, topic)
			}
		}
	}

	kubemqHost := "127.0.0.1"
	if val, exists := os.LookupEnv("kubemq_host"); exists {
		kubemqHost = val
	}

	kubemqClient := "client"
	if val, exists := os.LookupEnv("kubemq_client"); exists {
		kubemqClient = val
	}

	controllerConfig := &types.ControllerConfig{
		UpstreamTimeout:          time.Second * 60,
		GatewayURL:               "http://gateway:8080",
		RebuildInterval:          time.Second * 5,
		PrintResponse:            true,
		PrintResponseBody:        true,
		TopicAnnotationDelimiter: ",",
		AsyncFunctionInvocation:  true,
		PrintSync:                false,
	}

	controller := types.NewController(creds, controllerConfig)
	controller.BeginMapBuilder()

	// Her is the kubemq event stream
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(kubemqHost, 50000),
		kubemq.WithClientId(kubemqClient),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	errCh := make(chan error)
	eventsCh, err := client.SubscribeToEvents(ctx, topics[0], "", errCh)

	for {
		select {
		case err := <-errCh:
			log.Fatal(err)
		case event := <-eventsCh:
			// do something with the event and the controller
			controller.Invoke(event.Channel, &event.Body)
		}

	}

}
