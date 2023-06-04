#!/usr/bin/env bash

if [ -z $server_name ]; then
  read -p "please enter server_name(default:go_web_scaffold-build):" server_name
fi
if [ -z $server_name ]; then
  server_name="go_web_scaffold-build"
fi

echo 'stop container'
docker stop $server_name

echo 'remove container'
docker rm $server_name

echo 'remove image'
docker rmi $server_name

echo 'docker build'
docker build -t $server_name -f ./Dockerfile-build .

echo 'docker run'
docker run -d \
  --name $server_name \
  -v ./output:/output \
  $server_name

echo 'stop container'
docker stop $server_name

echo 'remove container'
docker rm $server_name

echo 'remove image'
docker rmi $server_name

echo 'clear image'
docker rmi $(docker images -a | grep "<none>" | awk '$1=="<none>" {print $3}')

echo 'all finish'
