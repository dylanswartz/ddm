# Decentralized Device Management

An opensource device management solution where devices host endpoints serving health metrics via ngrok tunnel.

The service exposes a `/memory` and `/cpu` endpoint to gather metrics about the host machine. You can access by logging into the [ngrok dashboard](https://dashboard.ngrok.com/cloud-edge/endpoints).

_Note: This service also exposes an experimental `/reboot` endpoint that will restart the machine running the service. It currently only works on Linux while running under a user with root level permissions. Use with caution._

# Quick Start

Configure the following environment variables: 

1. Your ngrok auth token `NGROK_AUTHTOKEN=xxxx_xxxx` 

2. The username you want to use to secure the application via basic auth `DDM_USERNAME=username`

3. The password you want to use to secure the application via basic auth  `DDM_USERNAME=password`

Run the application:

`go run api/main.go`