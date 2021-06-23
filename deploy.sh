#!/bin/sh

# See also https://github.com/zefer/ansible.

server_host=hendrix

echo 'Compiling for Rasperry Pi'
GOOS=linux GOARM=6 GOARCH=arm go build

echo 'Sending binary to doorbell server'
scp doorbell $server_host:/home/joe

echo 'Archiving old notify.sh on doorbell server'
ssh $server_host -t "\
  sudo mv /opt/doorbell/notify.sh \
  /opt/doorbell/notify.sh.bak.$(date -u +'%Y-%m-%dT%H-%M-%SZ')"

echo 'Sending notify.sh to doorbell server'
scp notify.sh $server_host:/home/joe

echo 'Running commands on doorbell server'
ssh $server_host -t '\
  sudo systemctl stop doorbell \
  && sleep 1 \
  && sudo chown root /home/joe/doorbell /home/joe/notify.sh \
  && sudo chgrp root /home/joe/doorbell /home/joe/notify.sh \
  && sudo mv /home/joe/doorbell /opt/doorbell \
  && sudo mv /home/joe/notify.sh /opt/doorbell \
  && sudo systemctl start doorbell \
  && sleep 1
'
