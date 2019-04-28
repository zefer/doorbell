#!/bin/sh

# See also https://github.com/zefer/ansible.

server_host=hendrix

echo 'Compiling for Rasperry Pi'
GOOS=linux GOARM=6 GOARCH=arm go build

echo 'Sending binary to doorbell server'
scp doorbell $server_host:/home/joe

echo 'Running commands on doorbell server'
ssh $server_host -t '\
  sudo systemctl stop doorbell \
  && sleep 1 \
  && sudo mv /home/joe/doorbell /opt/doorbell \
  && sudo systemctl start doorbell \
  && sleep 1
'
