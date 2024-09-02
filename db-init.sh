#!/bin/bash

source .env

docker compose up --build -d

CONTAINER_ID=$(docker ps --format='{{.ID}}' --filter name=^/pg)

echo $CONTAINER_ID

while [ "`docker inspect -f {{.State.Health.Status}} $CONTAINER_ID`" != "healthy" ]; do     sleep 2; done
docker exec -i $CONTAINER_ID psql -U user test < init.sql
cd client/
vite
echo "All commands completed"
