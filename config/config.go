package config

import (
	"log"
	"os"
	"strings"
	"time"
)

const defaultRebuild = time.Second * 5
const defaultUpstreamTimeout = time.Second * 60

type Config struct {
	ClientId string
	Host     string
	Topics   []string

	GatewayURL               string
	UpstreamTimeout          time.Duration
	RebuildInterval          time.Duration
	PrintResponse            bool
	PrintResponseBody        bool
	PrintSync                bool
	AsyncFunctionInvocation  bool
	TopicAnnotationDelimiter string
}

// Get will load the NATS Connector config from the environment variables
func Get() Config {

	topics := []string{}
	if val, exists := os.LookupEnv("topics"); exists {
		for _, topic := range strings.Split(val, ",") {
			if len(topic) > 0 {
				topics = append(topics, topic)
			}
		}
	}
	if len(topics) == 0 {
		log.Fatal(`Provide a list of topics i.e. topics="payment_published,slack_joined"`)
	}

	gatewayURL := "http://gateway:8080"
	if val, exists := os.LookupEnv("gateway_url"); exists {
		gatewayURL = val
	}

	upstreamTimeout := defaultUpstreamTimeout
	rebuildInterval := defaultRebuild

	if val, exists := os.LookupEnv("upstream_timeout"); exists {
		parsedVal, err := time.ParseDuration(val)
		if err == nil {
			upstreamTimeout = parsedVal
		}
	}

	if val, exists := os.LookupEnv("rebuild_interval"); exists {
		parsedVal, err := time.ParseDuration(val)
		if err == nil {
			rebuildInterval = parsedVal
		}
	}

	printResponse := false
	if val, exists := os.LookupEnv("print_response"); exists {
		printResponse = (val == "1" || val == "true")
	}

	printResponseBody := false
	if val, exists := os.LookupEnv("print_response_body"); exists {
		printResponseBody = (val == "1" || val == "true")
	}

	printSync := false
	if val, exists := os.LookupEnv("print_sync"); exists {
		printSync = (val == "1" || val == "true")
	}

	asyncFunctionInvocation := true
	if val, exists := os.LookupEnv("asynchronous_invocation"); exists {
		asyncFunctionInvocation = (val == "1" || val == "true")
	}

	delimiter := ","
	if val, exists := os.LookupEnv("topic_delimiter"); exists {
		if len(val) > 0 {
			delimiter = val
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

	return Config{
		Host:                     kubemqHost,
		ClientId:                 kubemqClient,
		Topics:                   topics,
		GatewayURL:               gatewayURL,
		UpstreamTimeout:          upstreamTimeout,
		RebuildInterval:          rebuildInterval,
		PrintResponse:            printResponse,
		PrintResponseBody:        printResponseBody,
		PrintSync:                printSync,
		AsyncFunctionInvocation:  asyncFunctionInvocation,
		TopicAnnotationDelimiter: delimiter,
	}
}
