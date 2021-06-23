# A rough example script that I used to play a doorbell sound via snapcast
# throughout the house, then switch back to the old audio stream.
# It works but I found it unreliable and prefer it simply pausing the music.

#!/usr/bin/python

import asyncio
# https://github.com/happyleavesaoc/python-snapcast
import snapcast.control
import time

MUSIC_HOST = '192.168.1.22'
DOORBELL_STREAM = 'doorbell'
MUSIC_STREAM = 'default'

loop = asyncio.get_event_loop()
server = loop.run_until_complete(snapcast.control.create_server(loop, MUSIC_HOST))

# print all client names
for group in server.groups:
    loop.run_until_complete(server.group_stream(group.identifier, DOORBELL_STREAM))

time.sleep(10)

# print all client names
for group in server.groups:
    loop.run_until_complete(server.group_stream(group.identifier, MUSIC_STREAM))
