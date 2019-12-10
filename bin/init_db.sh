#!/bin/bash

docker-compose exec mysql mysql -hmysql -uroot -proot -e "DROP DATABASE IF EXISTS golang_graphql_authentication_db;CREATE DATABASE golang_graphql_authentication_db CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
