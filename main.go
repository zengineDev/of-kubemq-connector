package main

import (
	"context"
	"github.com/openfaas-incubator/connector-sdk/types"
	"log"
	"os"
	"strings"

	"github.com/ZengineChris/of-kubemq-connector/config"
	"github.com/ZengineChris/of-kubemq-connector/kubemq"
)

func main() {
	creds := types.GetCredentials()

	configs := config.Get()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topics := []string{}
	if val, exists := os.LookupEnv("topics"); exists {
		for _, topic := range strings.Split(val, ",") {
			if len(topic) > 0 {
				topics = append(topics, topic)
			}
		}
	}

	controllerConfig := &types.ControllerConfig{
		UpstreamTimeout:          configs.UpstreamTimeout,
		GatewayURL:               configs.GatewayURL,
		RebuildInterval:          configs.RebuildInterval,
		PrintResponse:            configs.PrintResponse,
		PrintResponseBody:        configs.PrintResponseBody,
		TopicAnnotationDelimiter: configs.TopicAnnotationDelimiter,
		AsyncFunctionInvocation:  configs.AsyncFunctionInvocation,
		PrintSync:                configs.PrintSync,
	}

	controller := types.NewController(creds, controllerConfig)
	controller.BeginMapBuilder()

	brokerConfig := kubemq.BrokerConfig{
		Host:   configs.Host,
		Client: configs.ClientId,
	}

	broker, err := kubemq.NewBroker(ctx, brokerConfig)

	if err != nil {
		log.Fatal(err)
	}

	err = broker.Subscribe(controller, configs.Topics)
	if err != nil {
		log.Fatal(err)
	}

}
