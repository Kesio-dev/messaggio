#!/bin/bash

kafka-topics.sh --bootstrap-server kafka:9092 --topic messages --describe || {
    echo "Тема 'messages' не существует, создаем..."
    kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1 --topic messages
}
