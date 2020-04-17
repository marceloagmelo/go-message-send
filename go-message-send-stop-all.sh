#!/usr/bin/env bash

source setenv.sh

# Mysql
echo "Finalizando o mysql..."
docker rm -f $MYSQL_HOSTNAME

# RabbitMQ
echo "Finalizando o rabbitmq..."
docker rm -f $RABBITMQ_HOSTNAME

# Message API
echo "Finalizando o go-message-send..."
docker rm -f go-message-send

# Remover rede
echo "Removendo a rede message-net..."
docker network rm $DOCKER_NETWORK
