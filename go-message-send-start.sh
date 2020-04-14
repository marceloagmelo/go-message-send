#!/usr/bin/env bash

source setenv.sh

# Rabbitmq message api
echo "Subindo o go-message-send..."
docker run -d --name go-message-send --network message-net  \
-p 8181:8181 \
-e API_SERVICE_URL="http://localhost:8080" \
-e TZ=America/Sao_Paulo \
marceloagmelo/go-message-send

# Listando os containers
docker ps
