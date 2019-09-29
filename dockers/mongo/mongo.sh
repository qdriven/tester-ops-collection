#! /bin/sh

echo "---------------------------------------------------------"
echo "setup Mongo ....."

dir=$(cd $(dirname $0);pwd)
echo "current path is ${dir}"


tar -xvf mongodb-linux-x86_64-3.4.12.tgz

mv mongodb-linux-x86_64-3.4.12 mongodb-3.4.12

rm -rf /usr/local/teamcat/mongodb-3.4.12

mv mongodb-3.4.12 /usr/local/teamcat/

cp mongo.conf /usr/local/teamcat/mongodb-3.4.12

mkdir -p /data/mongodb/db

mkdir -p /data/mongodb/logs

touch  /data/mongodb/logs/mongodb.log

nohup /usr/local/teamcat/mongodb-3.4.12/bin/mongod -f /usr/local/teamcat/mongodb-3.4.12/mongo.conf &

echo "Mongo setup complete!"

echo "---------------------------------------------------------"