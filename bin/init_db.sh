#!/bin/bash

docker-compose exec mysql mysql -hmysql -uroot -proot -e "DROP DATABASE IF EXISTS golang_graphql_authentication_db;CREATE DATABASE golang_graphql_authentication_db CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
docker-compose exec mysql mysql -hmysql -uroot -proot golang_graphql_authentication_db -e "CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name varchar(255) DEFAULT NULL COMMENT 'user name',
    age varchar(255) DEFAULT NULL COMMENT 'age',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX user_id (id),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='users';"

docker-compose exec mysql mysql -hmysql -uroot -proot golang_graphql_authentication_db -e "
TRUNCATE users;
INSERT INTO users (id, name, age, created_at, updated_at, deleted_at) VALUES (1, 'user1', '20', '2019-01-01 00:00:00', '2019-01-01 00:00:00', NULL);
INSERT INTO users (id, name, age, created_at, updated_at, deleted_at) VALUES (2, 'user2', '21', '2019-01-01 00:00:00', '2019-01-01 00:00:00', NULL);
"
