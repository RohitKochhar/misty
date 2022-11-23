# misty-cli (version 1.0)

CLI client for the misty client-broker service mesh

## Overview

`misty-cli` is a command-line application written using the [CobraCLI library](https://github.com/spf13/cobra). Configuration management is done using the [Viper library](https://github.com/spf13/viper).

## Usage

```bash
$ misty-cli --help       

CLI client to communicate with misty brokers

misty-cli can be used to connect, publish and subscribe to event streams
managed by a misty broker

Usage:
  misty-cli [command]

Available Commands:
  help        Help about any command
  publish     publish sends a payload to a message topic
  subscribe   subscribe listens for message on a specified topic

Flags:
      --config string   config file (default is $HOME/.misty-cli.yaml)
  -h, --help            help for misty-cli
  -v, --version         version for misty-cli

Use "misty-cli [command] --help" for more information about a command.
```

### Publish

#### Info

```bash
$ misty-cli publish --help

publish sends a payload to a message topic

Usage:
  misty-cli publish -H {BROKER_HOSTNAME} -p {BROKER_PORT} -t {TOPIC} -m {MESSAGE} [flags]

Flags:
  -h, --help             help for publish
  -H, --host string      misty broker hostname (default "localhost")
  -m, --message string   message to be published (required)
  -p, --port int         misty broker port number (default 1315)
  -t, --topic string     topic on which to publish message (required)

Global Flags:
      --config string   config file (default is $HOME/.misty-cli.yaml)
```

#### Example

If the following message is published:

```bash
$ misty-cli publish -H localhost -p 1315 -t album -m "fear fun"
Using config file: $HOME/misty/clients/cli/misty-cli/.misty-cli.yaml
2022/11/23 15:42:25 [PUBLISH] fear fun --> http://localhost:1315/topic/album
```

It is received on the broker:

```bash
2022/11/23 15:42:25 Received message="fear fun" on topic="album"
```

And any clients subscribed to the topic receive the follwoing:

```bash
2022/11/23 15:42:25 [RECEIVED ON /album]: fear fun
```

### Subscribe

#### Info

```bash
$ misty-cli subscribe --help

subscribe listens for message on a specified topic

Usage:
  misty-cli subscribe -H {BROKER_HOSTNAME} -p {BROKER_PORT} -t {TOPIC} [flags]

Flags:
  -h, --help           help for subscribe
  -H, --host string    misty broker hostname (default "localhost")
  -p, --port int       misty broker port number (default 1315)
  -t, --topic string   topic on which to publish message (required)

Global Flags:
      --config string   config file (default is $HOME/.misty-cli.yaml)
```

#### Example

If a topic is subscribed to using the following command:

```bash
$ misty-cli subscribe -t album
Using config file:  $HOME/misty/clients/cli/misty-cli/.misty-cli.yaml
2022/11/23 15:42:21 Sending subscribe request to broker at localhost:1315...
2022/11/23 15:42:21 Subscibe result successful!
2022/11/23 15:42:21 listening for messages on /album...
```

When a message is publshed on the topic with:

```bash
$ misty-cli publish -H localhost -p 1315 -t album -m "fear fun"
Using config file:  $HOME/misty/clients/cli/misty-cli/.misty-cli.yaml
2022/11/23 15:46:18 [PUBLISH] fear fun --> http://localhost:1315/topic/album
```

It is received on the original client with:

```bash
2022/11/23 15:46:18 [RECEIVED ON /album]: fear fun
```
