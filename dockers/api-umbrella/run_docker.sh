#!/bin/sh

docker run -d --name=api-umbrella -p 8090:80 -p 443:443 -v "$(pwd)/config":/etc/api-umbrella nrel/api-umbrella