#! /bin/sh

PROJECT_NAME=$1
PACKAGE_NAME=io.qmeat.${PROJECT_NAME}
mkdir -p ${PROJECT_NAME} && cd ${PROJECT_NAME}

curl https://start.spring.io/starter.zip -d language=kotlins,java \
-d style=web,jpa,h2,devtools -d packageName=${PACKAGE_NAME} -d name=${PROJECT_NAME} \
-o ${PROJECT_NAME}.zip