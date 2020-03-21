package kubemq

import (
	"context"
	"fmt"
	"github.com/kubemq-io/kubemq-go"
	"github.com/openfaas-incubator/connector-sdk/types"
	"log"
	"sync"
	"time"
)

type BrokerConfig struct {

	// Host is the NATS address, the port is hard-coded to 4222
	Host string

	// ConnTimeout is the timeout for Dial on a connection.
	Client string
}

type Broker interface {
	Subscribe(types.Controller, []string) error
}

type broker struct {
	client *kubemq.Client
	ctx    context.Context
}

const MqPort = 50000

func NewBroker(ctx context.Context, config BrokerConfig) (Broker, error) {
	broker := &broker{}
	brokerURL := fmt.Sprintf("kubemq://%s:%s", config.Host, string(MqPort))

	for {
		client, err := kubemq.NewClient(ctx,
			kubemq.WithAddress(config.Host, MqPort),
			kubemq.WithClientId(config.Client),
			kubemq.WithTransportType(kubemq.TransportTypeGRPC))

		if client != nil && err == nil {
			broker.client = client
			broker.ctx = ctx
			break
		}

		if client != nil {
			client.Close()
		}

		log.Println("Wait for brokers to come up.. ", brokerURL)
		time.Sleep(1 * time.Second)
	}

	return broker, nil
}

func (b *broker) Subscribe(controller types.Controller, topics []string) error {
	log.Printf("Configured topics: %v", topics)

	if b.client == nil {
		return fmt.Errorf("client was nil, try to reconnect")
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	for _, topic := range topics {
		log.Printf("Binding to topic: %q", topic)

		errCh := make(chan error)
		eventsCh, err := b.client.SubscribeToEvents(b.ctx, topic, "", errCh)
		if err != nil {
			log.Printf("Unable to bind to topic: %s", topic)
			log.Println(err)
		}
		for {
			select {
			case err := <-errCh:
				log.Fatal(err)
			case event := <-eventsCh:
				log.Printf("Topic: %s, message: %q", event.Channel, string(event.Body))
				// do something with the event and the controller
				controller.Invoke(event.Channel, &event.Body)
			}

		}

	}

	// interrupt handling
	wg.Wait()

	b.client.Close()

	return nil
}
