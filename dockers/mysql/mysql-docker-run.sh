#! /bin/sh

docker run --name allqe-local-mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest 
