#!/usr/bin/env bash

source setenv.sh

# Verificar rede
echo "Verificando se existe a rede $DOCKER_NETWORK..."
docker network ls | grep $DOCKER_NETWORK
if [ "$?" != 0 ]; then
   echo "Rede $DOCKER_NETWORK não existe!"
   exit 0
fi

# Rabbitmq message send
echo "Subindo o go-message-send..."
docker run -d --name go-message-send --network $DOCKER_NETWORK  \
-p 7070:8080 \
-e API_SERVICE_URL="http://go-message-api:8080" \
-e TZ=America/Sao_Paulo \
marceloagmelo/go-message-send

# Listando os containers
docker ps