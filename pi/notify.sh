#!/bin/bash

curl -X POST \
  -H "Authorization: key=${DOORBELL_FIREBASE_SERVER_KEY}" \
  -H "Content-Type: application/json" \
  -d '{
  "to": "/topics/doorbell",
  "notification": {
    "title": "Doorbell!"
  }
}' "https://fcm.googleapis.com/fcm/send"
