# Doorbell

A trivial Raspberry Pi Doorbell project to help avoid missing the doorbell when
your music is too loud!

It pauses the music when the doorbell rings. Expand the bash script as you see
fit!

## Setup: Doorbell Wiring

It's wise to avoid wiring your Pi directly into the doorbell circuit since this
will likely damage the board, especially since many doorbells use 8v AC. I
solved this by using a small optocoupler to isolate the doorbell and Pi
circuits.

## Setup: Raspberry Pi

* Wire the doorbell (isolated, see above) or a push button to GND & GPIO pin 18
* Install the Python & sudo pip install RPi.GPIO`
* Run it:
  * With [a systemd unit][systemd-unit], or
  * Manually `python pi/doorbell.py` or `sudo -E bash -c 'python doorbell.py'`

## Setup: Ansible Example

See this [ansible role][ansible-role] for the full Pi software set-up I use.

[ansible-role]: https://github.com/zefer/ansible/blob/master/roles/doorbell
[systemd-unit]: https://github.com/zefer/ansible/blob/master/roles/doorbell/templates/doorbell.service
