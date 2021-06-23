#!/bin/bash

export MPD_HOST=music

wasplaying=$(mpc status | grep "\[playing\]")

mpc pause

if [ -n "$wasplaying" ]; then
  sleep 10
  mpc play
fi

time="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
echo "${time}" >> /var/log/doorbell.log
