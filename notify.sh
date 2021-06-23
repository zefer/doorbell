#!/bin/bash

MPD_HOST=music mpc pause
#./notify.py

time="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
echo "${time}" >> /var/log/doorbell.log
