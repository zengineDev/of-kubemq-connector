package function

import (
	"context"
	"github.com/kubemq-io/kubemq-go"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	// Connect to the queue
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress("10.245.27.245", 50000),
		kubemq.WithClientId("test"),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	eventErr := client.E().
		SetId("test-id").
		SetChannel("kubemq-test").
		SetMetadata("some-metadata").
		SetBody([]byte("Hello from here")).
		Send(ctx)

	if eventErr != nil {
		log.Fatal(eventErr)
	}
}
