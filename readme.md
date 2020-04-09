# Openfaas KubeMQ Connector
[![Go Report Card](https://goreportcard.com/badge/github.com/ZengineChris/of-kubemq-connector)](https://goreportcard.com/report/github.com/ZengineChris/of-kubemq-connector)
[![GoDoc](https://pkg.go.dev/github.com/zengineDev/of-kubemq-connector?status.svg)](https://pkg.go.dev/github.com/zengineDev/of-kubemq-connector )
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![OpenFaaS](https://img.shields.io/badge/openfaas-serverless-blue.svg)](https://www.openfaas.com)
[![Release](https://img.shields.io/github/v/release/zengineDev/of-kubemq-connector?style=plastic)](https://github.com/zengineDev/of-kubemq-connector/releases)

An [OpenFaas](https://www.openfaas.com/) event-connector to trigger functions from [KubeMQ](https://kubemq.io/).
Highly inspired by the [openfaas NATS connector](https://github.com/openfaas-incubator/nats-connector).


## Deploy on Kubernetes
```bash
kubectl apply -f ./kubernetes/connector-dep.yml
```


### Configuration

Configuration is by environment variable, which can be set in the Kubernetes YAML file: [yaml/kubernetes/connector-dep.yaml](./yaml/kubernetes/connector-dep.yaml)

| Variable             | Description                   |  Default                                        |
| -------------------- | ------------------------------|--------------------------------------------------|
| `topics`             | Delimited list of topics    |  `kubemq-test,`                                   |
| `kubemq_host`        | The host, but not the port for KubeMQ | `kube-mq` |
| `kubemq_client`      | The KubeMQ client  | `kube-mq` |
| `async-invocation`   | Queue the invocation with the built-in OpenFaaS queue-worker and return immediately    |  `false` |
| `gateway_url`        | The URL for the OpenFaaS gateway | `http://gateway:8080` |
| `upstream_timeout`   | Timeout to wait for synchronous invocations | `60s` |
| `rebuild_interval`   | Interval at which to rebuild the map of topics <> functions | `5s`  |
| `topic_delimiter`    | Used to separate items in `topics` variable | `,` |
