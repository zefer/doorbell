# Doorbell

Get push notifications on your devices when your doorbell rings.

Captures a GPIO push-button input and sends push notifications to all your
registered devices.

The web app is largely based on the [Firebase Cloud Messaging Quickstart
Example][fcm-quickstart].

## Setup: Web & Notifications

* Create a [Firebase][firebase] account & project to provide push notifications
* Clone this repo
* Install the Firebase tools `npm install -g firebase-tools`
* Configure and deploy the web app

```
firebase login
firebase init
firebase deploy
```

* Subscribe any device (mobile, desktop, etc) to notifications by visiting the
  deployed app's URL and following the instructions

## Setup: Raspberry Pi / Doorbell device

* Wire a push button to GND & GPIO pin 18
* Install the Python & `sudo pip install RPi.GPIO`
* Define the environment variable using the server key found in your Firebase
  project's settings under "Settings > Cloud Messaging > Server Key"
  `export DOORBELL_FIREBASE_SERVER_KEY=changeme`
* Run it `sudo -E bash -c 'python doorbell.py'`

[firebase]: https://firebase.google.com/
[fcm-quickstart]: https://firebase.google.com/
