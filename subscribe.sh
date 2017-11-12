#!/bin/bash

# Subscribe a browser (its instance ID token) to the doorbell message topic.
# Usage: subscribe.sh instance-id-token

curl -X POST \
  -H "Authorization: key=${DOORBELL_FIREBASE_SERVER_KEY}" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "https://iid.googleapis.com/iid/v1/$1/rel/topics/doorbell"
