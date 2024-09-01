#!/bin/bash

source .env

docker compose up -d

CONTAINER_ID=$(docker ps --format='{{.ID}}' --filter name=^/pg)

sleep 20
docker exec -i $CONTAINER_ID psql -U postgres pokemon_jackpot_db < init.sql

echo "All commands completed"
