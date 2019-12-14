#!/bin/bash

mpc pause

time="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
echo "${time}" >> /var/log/doorbell.log
