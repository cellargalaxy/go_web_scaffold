#!/usr/bin/env bash

if [ -z $server_name ]; then
  read -p "please enter server_name(default:go_web_scaffold):" server_name
fi
if [ -z $server_name ]; then
  server_name="go_web_scaffold"
fi

if [ -z $server_center_address ]; then
  read -p "please enter server_center_address(default:''):" server_center_address
fi
if [ -z $server_center_address ]; then
  server_center_address=""
fi

if [ -z $server_center_secret ]; then
  read -p "please enter server_center_secret(default:''):" server_center_secret
fi
if [ -z $server_center_secret ]; then
  server_center_secret=""
fi

if [ -z $listen_port ]; then
  read -p "please enter listen port(default:36449):" listen_port
fi
if [ -z $listen_port ]; then
  listen_port="36449"
fi

echo
echo 'server_name: '$server_name
echo "server_center_address: $server_center_address"
echo "server_center_secret: $server_center_secret"
echo 'listen_port: '$listen_port
echo 'input any key go on, or control+c over'
read

echo 'create volume'
docker volume create log

echo 'create volume'
docker volume create $server_name'_resource'

echo 'stop container'
docker stop $server_name

echo 'remove container'
docker rm $server_name

echo 'remove image'
docker rmi $server_name

echo 'docker build'
docker build -t $server_name -f ./Dockerfile-install .

echo 'docker run'
docker run -d \
  --restart=always \
  --name $server_name \
  -v log:/log \
  -v $server_name'_resource':/resource \
  -p $listen_port:36449 \
  -e server_name=$server_name \
  -e server_center_address=$server_center_address \
  -e server_center_secret=$server_center_secret \
  $server_name

echo 'clear image'
docker rmi $(docker images -a | grep "<none>" | awk '$1=="<none>" {print $3}')

echo 'all finish'

xdg-open "http://127.0.0.1:$listen_port/static/html/main.html"
