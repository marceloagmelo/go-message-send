#!/usr/bin/env bash

source setenv.sh

# Mysql
echo "Finalizando o mysql..."
docker rm -f $MYSQL_HOSTNAME

# RabbitMQ
echo "Finalizando o rabbitmq..."
docker rm -f $RABBITMQ_HOSTNAME

# Message API
echo "Finalizando o ${APP_NAME_API}..."
docker rm -f ${APP_NAME_API}

# Message send
echo "Finalizando o ${APP_NAME}..."
docker rm -f ${APP_NAME}

# Remover rede
echo "Removendo a rede message-net..."
docker network rm $DOCKER_NETWORK
