# Watchtime-Tracking

Elasticsearch - Logstash - Kibana based slideshow tracking.

## Client

- Built using Vue3
- nginx in docker to serve webapp
- Uses Beacon API

**Note:** Beacon API does not garuntees successful connection everytime, preventing clogging of network.

## Server

- Developed using [Gin Web Framework](https://github.com/gin-gonic/gin) for Golang
- Generates standalone binary served using alpine docker container
- Uses Redis to hold intermediate data and send data to logstash via RabbitMQ

## Logstash

- Takes input via RabbitMQ
- Parses input data
- Outputs to Elasticsearch