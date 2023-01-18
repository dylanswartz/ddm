# Decentralized Device Management

An opensource device management solution where devices host endpoints serving health metric via ngrok tunnel 

# Quick Start

`NGROK_AUTHTOKEN=xxxx_xxxx go run api/main.go`

The application will then expose a `/memory` and `/cpu` endpoint. You can access by logging into the [ngrok dashboard](https://dashboard.ngrok.com/cloud-edge/endpoints).
