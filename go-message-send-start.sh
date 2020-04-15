#!/usr/bin/env bash

source setenv.sh

# Rabbitmq message api
echo "Subindo o go-message-send..."
docker run -d --name go-message-send --network message-net  \
-p 7070:8080 \
-e API_SERVICE_URL="http://localhost:8181" \
-e TZ=America/Sao_Paulo \
marceloagmelo/go-message-send

# Listando os containers
docker ps
